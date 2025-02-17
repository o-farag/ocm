package gc

import (
	"context"
	"strings"
	"time"

	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	rbacv1listers "k8s.io/client-go/listers/rbac/v1"
	"k8s.io/client-go/metadata"
	"k8s.io/klog/v2"

	clientset "open-cluster-management.io/api/client/cluster/clientset/versioned"
	informerv1 "open-cluster-management.io/api/client/cluster/informers/externalversions/cluster/v1"
	clusterv1listers "open-cluster-management.io/api/client/cluster/listers/cluster/v1"
	worklister "open-cluster-management.io/api/client/work/listers/work/v1"
	clusterv1 "open-cluster-management.io/api/cluster/v1"
	"open-cluster-management.io/sdk-go/pkg/patcher"

	"open-cluster-management.io/ocm/pkg/common/queue"
	"open-cluster-management.io/ocm/pkg/registration/register"
)

type gcReconcileOp int

const (
	gcReconcileRequeue gcReconcileOp = iota
	gcReconcileStop
	gcReconcileContinue
)

var (
	addonGvr = schema.GroupVersionResource{Group: "addon.open-cluster-management.io",
		Version: "v1alpha1", Resource: "managedclusteraddons"}
	workGvr = schema.GroupVersionResource{Group: "work.open-cluster-management.io",
		Version: "v1", Resource: "manifestworks"}
)

// gcReconciler is an interface for reconcile cleanup logic after cluster is deleted.
// clusterName is from the queueKey, cluster may be nil.
type gcReconciler interface {
	reconcile(ctx context.Context, cluster *clusterv1.ManagedCluster, clusterNamespace string) (gcReconcileOp, error)
}

type GCController struct {
	clusterLister  clusterv1listers.ManagedClusterLister
	clusterPatcher patcher.Patcher[*clusterv1.ManagedCluster, clusterv1.ManagedClusterSpec, clusterv1.ManagedClusterStatus]
	gcReconcilers  []gcReconciler
}

// NewGCController ensures the related resources are cleaned up after cluster is deleted
func NewGCController(
	clusterRoleLister rbacv1listers.ClusterRoleLister,
	roleBindingLister rbacv1listers.RoleBindingLister,
	clusterInformer informerv1.ManagedClusterInformer,
	manifestWorkLister worklister.ManifestWorkLister,
	clusterClient clientset.Interface,
	kubeClient kubernetes.Interface,
	metadataClient metadata.Interface,
	approver register.Approver,
	eventRecorder events.Recorder,
	gcResourceList []string,
	resourceCleanupFeatureGateEnable bool,
) factory.Controller {
	clusterPatcher := patcher.NewPatcher[
		*clusterv1.ManagedCluster, clusterv1.ManagedClusterSpec, clusterv1.ManagedClusterStatus](
		clusterClient.ClusterV1().ManagedClusters())

	controller := &GCController{
		clusterLister:  clusterInformer.Lister(),
		clusterPatcher: clusterPatcher,
		gcReconcilers:  []gcReconciler{},
	}

	// do not clean resources if featureGate is disabled or no gc resource list for backwards compatible.
	if resourceCleanupFeatureGateEnable && len(gcResourceList) != 0 {
		gcResources := []schema.GroupVersionResource{}
		for _, gcResource := range gcResourceList {
			subStrings := strings.Split(gcResource, "/")
			if len(subStrings) != 3 {
				klog.Errorf("invalid gc-resource-list flag: %v", gcResources)
				continue
			}
			gcResources = append(gcResources, schema.GroupVersionResource{
				Group: subStrings[0], Version: subStrings[1], Resource: subStrings[2]})
		}
		controller.gcReconcilers = append(controller.gcReconcilers,
			newGCResourcesController(metadataClient, gcResources, eventRecorder))
	}

	controller.gcReconcilers = append(controller.gcReconcilers,
		newGCClusterRbacController(kubeClient, clusterPatcher, clusterRoleLister, roleBindingLister,
			manifestWorkLister, approver, eventRecorder, resourceCleanupFeatureGateEnable))

	return factory.New().
		WithInformersQueueKeysFunc(queue.QueueKeyByMetaName, clusterInformer.Informer()).
		WithSync(controller.sync).ToController("GCController", eventRecorder)
}

// gc controller is watching cluster and to do these jobs:
//  1. add a cleanup finalizer to managedCluster if the cluster is not deleting.
//  2. clean up all rolebinding and resources in the cluster ns after the cluster is deleted.
func (r *GCController) sync(ctx context.Context, controllerContext factory.SyncContext) error {
	clusterName := controllerContext.QueueKey()
	if clusterName == "" || clusterName == factory.DefaultQueueKey {
		return nil
	}

	// cluster could be nil, that means the cluster is gone but the gc is not finished.
	cluster, err := r.clusterLister.Get(clusterName)
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}

	var copyCluster *clusterv1.ManagedCluster
	if cluster != nil {
		if cluster.DeletionTimestamp.IsZero() {
			_, err = r.clusterPatcher.AddFinalizer(ctx, cluster, clusterv1.ManagedClusterFinalizer)
			return err
		}

		copyCluster = cluster.DeepCopy()
	}

	var errs []error
	var requeue bool
	for _, reconciler := range r.gcReconcilers {
		op, err := reconciler.reconcile(ctx, copyCluster, clusterName)
		if err != nil {
			errs = append(errs, err)
		}
		if op == gcReconcileRequeue {
			requeue = true
		}
		if op == gcReconcileStop {
			break
		}
	}

	if requeue {
		controllerContext.Queue().AddAfter(clusterName, 5*time.Second)
	}

	if cluster == nil {
		return utilerrors.NewAggregate(errs)
	}

	// update cluster condition
	if len(errs) != 0 {
		applyErrors := utilerrors.NewAggregate(errs)
		meta.SetStatusCondition(&copyCluster.Status.Conditions, metav1.Condition{
			Type:    clusterv1.ManagedClusterConditionDeleting,
			Status:  metav1.ConditionFalse,
			Reason:  clusterv1.ConditionDeletingReasonResourceError,
			Message: applyErrors.Error(),
		})
	}

	if _, err = r.clusterPatcher.PatchStatus(ctx, cluster, copyCluster.Status, cluster.Status); err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 || requeue {
		return utilerrors.NewAggregate(errs)
	}

	return r.clusterPatcher.RemoveFinalizer(ctx, cluster, clusterv1.ManagedClusterFinalizer)
}

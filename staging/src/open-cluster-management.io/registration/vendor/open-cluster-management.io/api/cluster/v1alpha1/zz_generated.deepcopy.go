//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScore) DeepCopyInto(out *AddOnPlacementScore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScore.
func (in *AddOnPlacementScore) DeepCopy() *AddOnPlacementScore {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnPlacementScore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScoreItem) DeepCopyInto(out *AddOnPlacementScoreItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScoreItem.
func (in *AddOnPlacementScoreItem) DeepCopy() *AddOnPlacementScoreItem {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScoreItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScoreList) DeepCopyInto(out *AddOnPlacementScoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddOnPlacementScore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScoreList.
func (in *AddOnPlacementScoreList) DeepCopy() *AddOnPlacementScoreList {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnPlacementScoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScoreStatus) DeepCopyInto(out *AddOnPlacementScoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scores != nil {
		in, out := &in.Scores, &out.Scores
		*out = make([]AddOnPlacementScoreItem, len(*in))
		copy(*out, *in)
	}
	if in.ValidUntil != nil {
		in, out := &in.ValidUntil, &out.ValidUntil
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScoreStatus.
func (in *AddOnPlacementScoreStatus) DeepCopy() *AddOnPlacementScoreStatus {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaim) DeepCopyInto(out *ClusterClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaim.
func (in *ClusterClaim) DeepCopy() *ClusterClaim {
	if in == nil {
		return nil
	}
	out := new(ClusterClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaimList) DeepCopyInto(out *ClusterClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaimList.
func (in *ClusterClaimList) DeepCopy() *ClusterClaimList {
	if in == nil {
		return nil
	}
	out := new(ClusterClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaimSpec) DeepCopyInto(out *ClusterClaimSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaimSpec.
func (in *ClusterClaimSpec) DeepCopy() *ClusterClaimSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterClaimSpec)
	in.DeepCopyInto(out)
	return out
}
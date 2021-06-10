// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberAccountAssociation) DeepCopyInto(out *MemberAccountAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberAccountAssociation.
func (in *MemberAccountAssociation) DeepCopy() *MemberAccountAssociation {
	if in == nil {
		return nil
	}
	out := new(MemberAccountAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemberAccountAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberAccountAssociationList) DeepCopyInto(out *MemberAccountAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MemberAccountAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberAccountAssociationList.
func (in *MemberAccountAssociationList) DeepCopy() *MemberAccountAssociationList {
	if in == nil {
		return nil
	}
	out := new(MemberAccountAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemberAccountAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberAccountAssociationSpec) DeepCopyInto(out *MemberAccountAssociationSpec) {
	*out = *in
	in.MemberAccountAssociationSpec2.DeepCopyInto(&out.MemberAccountAssociationSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberAccountAssociationSpec.
func (in *MemberAccountAssociationSpec) DeepCopy() *MemberAccountAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(MemberAccountAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberAccountAssociationSpec2) DeepCopyInto(out *MemberAccountAssociationSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.MemberAccountID != nil {
		in, out := &in.MemberAccountID, &out.MemberAccountID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberAccountAssociationSpec2.
func (in *MemberAccountAssociationSpec2) DeepCopy() *MemberAccountAssociationSpec2 {
	if in == nil {
		return nil
	}
	out := new(MemberAccountAssociationSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberAccountAssociationStatus) DeepCopyInto(out *MemberAccountAssociationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberAccountAssociationStatus.
func (in *MemberAccountAssociationStatus) DeepCopy() *MemberAccountAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(MemberAccountAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAssociation) DeepCopyInto(out *S3BucketAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAssociation.
func (in *S3BucketAssociation) DeepCopy() *S3BucketAssociation {
	if in == nil {
		return nil
	}
	out := new(S3BucketAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAssociationList) DeepCopyInto(out *S3BucketAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3BucketAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAssociationList.
func (in *S3BucketAssociationList) DeepCopy() *S3BucketAssociationList {
	if in == nil {
		return nil
	}
	out := new(S3BucketAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAssociationSpec) DeepCopyInto(out *S3BucketAssociationSpec) {
	*out = *in
	in.S3BucketAssociationSpec2.DeepCopyInto(&out.S3BucketAssociationSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAssociationSpec.
func (in *S3BucketAssociationSpec) DeepCopy() *S3BucketAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAssociationSpec2) DeepCopyInto(out *S3BucketAssociationSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.ClassificationType != nil {
		in, out := &in.ClassificationType, &out.ClassificationType
		*out = new(S3BucketAssociationSpecClassificationType)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberAccountID != nil {
		in, out := &in.MemberAccountID, &out.MemberAccountID
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAssociationSpec2.
func (in *S3BucketAssociationSpec2) DeepCopy() *S3BucketAssociationSpec2 {
	if in == nil {
		return nil
	}
	out := new(S3BucketAssociationSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAssociationSpecClassificationType) DeepCopyInto(out *S3BucketAssociationSpecClassificationType) {
	*out = *in
	if in.Continuous != nil {
		in, out := &in.Continuous, &out.Continuous
		*out = new(string)
		**out = **in
	}
	if in.OneTime != nil {
		in, out := &in.OneTime, &out.OneTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAssociationSpecClassificationType.
func (in *S3BucketAssociationSpecClassificationType) DeepCopy() *S3BucketAssociationSpecClassificationType {
	if in == nil {
		return nil
	}
	out := new(S3BucketAssociationSpecClassificationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAssociationStatus) DeepCopyInto(out *S3BucketAssociationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAssociationStatus.
func (in *S3BucketAssociationStatus) DeepCopy() *S3BucketAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(S3BucketAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

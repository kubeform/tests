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
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecret) DeepCopyInto(out *ManagerSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecret.
func (in *ManagerSecret) DeepCopy() *ManagerSecret {
	if in == nil {
		return nil
	}
	out := new(ManagerSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamBinding) DeepCopyInto(out *ManagerSecretIamBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamBinding.
func (in *ManagerSecretIamBinding) DeepCopy() *ManagerSecretIamBinding {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretIamBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamBindingList) DeepCopyInto(out *ManagerSecretIamBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerSecretIamBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamBindingList.
func (in *ManagerSecretIamBindingList) DeepCopy() *ManagerSecretIamBindingList {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretIamBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamBindingSpec) DeepCopyInto(out *ManagerSecretIamBindingSpec) {
	*out = *in
	in.ManagerSecretIamBindingSpec2.DeepCopyInto(&out.ManagerSecretIamBindingSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamBindingSpec.
func (in *ManagerSecretIamBindingSpec) DeepCopy() *ManagerSecretIamBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamBindingSpec2) DeepCopyInto(out *ManagerSecretIamBindingSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ManagerSecretIamBindingSpecCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamBindingSpec2.
func (in *ManagerSecretIamBindingSpec2) DeepCopy() *ManagerSecretIamBindingSpec2 {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamBindingSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamBindingSpecCondition) DeepCopyInto(out *ManagerSecretIamBindingSpecCondition) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamBindingSpecCondition.
func (in *ManagerSecretIamBindingSpecCondition) DeepCopy() *ManagerSecretIamBindingSpecCondition {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamBindingSpecCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamBindingStatus) DeepCopyInto(out *ManagerSecretIamBindingStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamBindingStatus.
func (in *ManagerSecretIamBindingStatus) DeepCopy() *ManagerSecretIamBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamMember) DeepCopyInto(out *ManagerSecretIamMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamMember.
func (in *ManagerSecretIamMember) DeepCopy() *ManagerSecretIamMember {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretIamMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamMemberList) DeepCopyInto(out *ManagerSecretIamMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerSecretIamMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamMemberList.
func (in *ManagerSecretIamMemberList) DeepCopy() *ManagerSecretIamMemberList {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretIamMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamMemberSpec) DeepCopyInto(out *ManagerSecretIamMemberSpec) {
	*out = *in
	in.ManagerSecretIamMemberSpec2.DeepCopyInto(&out.ManagerSecretIamMemberSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamMemberSpec.
func (in *ManagerSecretIamMemberSpec) DeepCopy() *ManagerSecretIamMemberSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamMemberSpec2) DeepCopyInto(out *ManagerSecretIamMemberSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ManagerSecretIamMemberSpecCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamMemberSpec2.
func (in *ManagerSecretIamMemberSpec2) DeepCopy() *ManagerSecretIamMemberSpec2 {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamMemberSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamMemberSpecCondition) DeepCopyInto(out *ManagerSecretIamMemberSpecCondition) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamMemberSpecCondition.
func (in *ManagerSecretIamMemberSpecCondition) DeepCopy() *ManagerSecretIamMemberSpecCondition {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamMemberSpecCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamMemberStatus) DeepCopyInto(out *ManagerSecretIamMemberStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamMemberStatus.
func (in *ManagerSecretIamMemberStatus) DeepCopy() *ManagerSecretIamMemberStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamPolicy) DeepCopyInto(out *ManagerSecretIamPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamPolicy.
func (in *ManagerSecretIamPolicy) DeepCopy() *ManagerSecretIamPolicy {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretIamPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamPolicyList) DeepCopyInto(out *ManagerSecretIamPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerSecretIamPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamPolicyList.
func (in *ManagerSecretIamPolicyList) DeepCopy() *ManagerSecretIamPolicyList {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretIamPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamPolicySpec) DeepCopyInto(out *ManagerSecretIamPolicySpec) {
	*out = *in
	in.ManagerSecretIamPolicySpec2.DeepCopyInto(&out.ManagerSecretIamPolicySpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamPolicySpec.
func (in *ManagerSecretIamPolicySpec) DeepCopy() *ManagerSecretIamPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamPolicySpec2) DeepCopyInto(out *ManagerSecretIamPolicySpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.PolicyData != nil {
		in, out := &in.PolicyData, &out.PolicyData
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamPolicySpec2.
func (in *ManagerSecretIamPolicySpec2) DeepCopy() *ManagerSecretIamPolicySpec2 {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamPolicySpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretIamPolicyStatus) DeepCopyInto(out *ManagerSecretIamPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretIamPolicyStatus.
func (in *ManagerSecretIamPolicyStatus) DeepCopy() *ManagerSecretIamPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretIamPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretList) DeepCopyInto(out *ManagerSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretList.
func (in *ManagerSecretList) DeepCopy() *ManagerSecretList {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretSpec) DeepCopyInto(out *ManagerSecretSpec) {
	*out = *in
	in.ManagerSecretSpec2.DeepCopyInto(&out.ManagerSecretSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretSpec.
func (in *ManagerSecretSpec) DeepCopy() *ManagerSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretSpec2) DeepCopyInto(out *ManagerSecretSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Replication != nil {
		in, out := &in.Replication, &out.Replication
		*out = new(ManagerSecretSpecReplication)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretSpec2.
func (in *ManagerSecretSpec2) DeepCopy() *ManagerSecretSpec2 {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretSpecReplication) DeepCopyInto(out *ManagerSecretSpecReplication) {
	*out = *in
	if in.Automatic != nil {
		in, out := &in.Automatic, &out.Automatic
		*out = new(bool)
		**out = **in
	}
	if in.UserManaged != nil {
		in, out := &in.UserManaged, &out.UserManaged
		*out = new(ManagerSecretSpecReplicationUserManaged)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretSpecReplication.
func (in *ManagerSecretSpecReplication) DeepCopy() *ManagerSecretSpecReplication {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretSpecReplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretSpecReplicationUserManaged) DeepCopyInto(out *ManagerSecretSpecReplicationUserManaged) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]ManagerSecretSpecReplicationUserManagedReplicas, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretSpecReplicationUserManaged.
func (in *ManagerSecretSpecReplicationUserManaged) DeepCopy() *ManagerSecretSpecReplicationUserManaged {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretSpecReplicationUserManaged)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretSpecReplicationUserManagedReplicas) DeepCopyInto(out *ManagerSecretSpecReplicationUserManagedReplicas) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretSpecReplicationUserManagedReplicas.
func (in *ManagerSecretSpecReplicationUserManagedReplicas) DeepCopy() *ManagerSecretSpecReplicationUserManagedReplicas {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretSpecReplicationUserManagedReplicas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretStatus) DeepCopyInto(out *ManagerSecretStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretStatus.
func (in *ManagerSecretStatus) DeepCopy() *ManagerSecretStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretVersion) DeepCopyInto(out *ManagerSecretVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretVersion.
func (in *ManagerSecretVersion) DeepCopy() *ManagerSecretVersion {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretVersionList) DeepCopyInto(out *ManagerSecretVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerSecretVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretVersionList.
func (in *ManagerSecretVersionList) DeepCopy() *ManagerSecretVersionList {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerSecretVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretVersionSpec) DeepCopyInto(out *ManagerSecretVersionSpec) {
	*out = *in
	in.ManagerSecretVersionSpec2.DeepCopyInto(&out.ManagerSecretVersionSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretVersionSpec.
func (in *ManagerSecretVersionSpec) DeepCopy() *ManagerSecretVersionSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretVersionSpec2) DeepCopyInto(out *ManagerSecretVersionSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DestroyTime != nil {
		in, out := &in.DestroyTime, &out.DestroyTime
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.SecretData != nil {
		in, out := &in.SecretData, &out.SecretData
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretVersionSpec2.
func (in *ManagerSecretVersionSpec2) DeepCopy() *ManagerSecretVersionSpec2 {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretVersionSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSecretVersionStatus) DeepCopyInto(out *ManagerSecretVersionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSecretVersionStatus.
func (in *ManagerSecretVersionStatus) DeepCopy() *ManagerSecretVersionStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerSecretVersionStatus)
	in.DeepCopyInto(out)
	return out
}
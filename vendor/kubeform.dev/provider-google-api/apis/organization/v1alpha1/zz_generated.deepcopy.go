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

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettings) DeepCopyInto(out *AccessApprovalSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettings.
func (in *AccessApprovalSettings) DeepCopy() *AccessApprovalSettings {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessApprovalSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsList) DeepCopyInto(out *AccessApprovalSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessApprovalSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsList.
func (in *AccessApprovalSettingsList) DeepCopy() *AccessApprovalSettingsList {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessApprovalSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsSpec) DeepCopyInto(out *AccessApprovalSettingsSpec) {
	*out = *in
	in.AccessApprovalSettingsSpec2.DeepCopyInto(&out.AccessApprovalSettingsSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsSpec.
func (in *AccessApprovalSettingsSpec) DeepCopy() *AccessApprovalSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsSpec2) DeepCopyInto(out *AccessApprovalSettingsSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.EnrolledAncestor != nil {
		in, out := &in.EnrolledAncestor, &out.EnrolledAncestor
		*out = new(bool)
		**out = **in
	}
	if in.EnrolledServices != nil {
		in, out := &in.EnrolledServices, &out.EnrolledServices
		*out = make([]AccessApprovalSettingsSpecEnrolledServices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotificationEmails != nil {
		in, out := &in.NotificationEmails, &out.NotificationEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsSpec2.
func (in *AccessApprovalSettingsSpec2) DeepCopy() *AccessApprovalSettingsSpec2 {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsSpecEnrolledServices) DeepCopyInto(out *AccessApprovalSettingsSpecEnrolledServices) {
	*out = *in
	if in.CloudProduct != nil {
		in, out := &in.CloudProduct, &out.CloudProduct
		*out = new(string)
		**out = **in
	}
	if in.EnrollmentLevel != nil {
		in, out := &in.EnrollmentLevel, &out.EnrollmentLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsSpecEnrolledServices.
func (in *AccessApprovalSettingsSpecEnrolledServices) DeepCopy() *AccessApprovalSettingsSpecEnrolledServices {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsSpecEnrolledServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsStatus) DeepCopyInto(out *AccessApprovalSettingsStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsStatus.
func (in *AccessApprovalSettingsStatus) DeepCopy() *AccessApprovalSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfig) DeepCopyInto(out *IamAuditConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfig.
func (in *IamAuditConfig) DeepCopy() *IamAuditConfig {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamAuditConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigList) DeepCopyInto(out *IamAuditConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamAuditConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigList.
func (in *IamAuditConfigList) DeepCopy() *IamAuditConfigList {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamAuditConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigSpec) DeepCopyInto(out *IamAuditConfigSpec) {
	*out = *in
	in.IamAuditConfigSpec2.DeepCopyInto(&out.IamAuditConfigSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigSpec.
func (in *IamAuditConfigSpec) DeepCopy() *IamAuditConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigSpec2) DeepCopyInto(out *IamAuditConfigSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.AuditLogConfig != nil {
		in, out := &in.AuditLogConfig, &out.AuditLogConfig
		*out = make([]IamAuditConfigSpecAuditLogConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigSpec2.
func (in *IamAuditConfigSpec2) DeepCopy() *IamAuditConfigSpec2 {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigSpecAuditLogConfig) DeepCopyInto(out *IamAuditConfigSpecAuditLogConfig) {
	*out = *in
	if in.ExemptedMembers != nil {
		in, out := &in.ExemptedMembers, &out.ExemptedMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogType != nil {
		in, out := &in.LogType, &out.LogType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigSpecAuditLogConfig.
func (in *IamAuditConfigSpecAuditLogConfig) DeepCopy() *IamAuditConfigSpecAuditLogConfig {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigSpecAuditLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigStatus) DeepCopyInto(out *IamAuditConfigStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigStatus.
func (in *IamAuditConfigStatus) DeepCopy() *IamAuditConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBinding) DeepCopyInto(out *IamBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBinding.
func (in *IamBinding) DeepCopy() *IamBinding {
	if in == nil {
		return nil
	}
	out := new(IamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingList) DeepCopyInto(out *IamBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingList.
func (in *IamBindingList) DeepCopy() *IamBindingList {
	if in == nil {
		return nil
	}
	out := new(IamBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingSpec) DeepCopyInto(out *IamBindingSpec) {
	*out = *in
	in.IamBindingSpec2.DeepCopyInto(&out.IamBindingSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingSpec.
func (in *IamBindingSpec) DeepCopy() *IamBindingSpec {
	if in == nil {
		return nil
	}
	out := new(IamBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingSpec2) DeepCopyInto(out *IamBindingSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(IamBindingSpecCondition)
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
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingSpec2.
func (in *IamBindingSpec2) DeepCopy() *IamBindingSpec2 {
	if in == nil {
		return nil
	}
	out := new(IamBindingSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingSpecCondition) DeepCopyInto(out *IamBindingSpecCondition) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingSpecCondition.
func (in *IamBindingSpecCondition) DeepCopy() *IamBindingSpecCondition {
	if in == nil {
		return nil
	}
	out := new(IamBindingSpecCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingStatus) DeepCopyInto(out *IamBindingStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingStatus.
func (in *IamBindingStatus) DeepCopy() *IamBindingStatus {
	if in == nil {
		return nil
	}
	out := new(IamBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamCustomRole) DeepCopyInto(out *IamCustomRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamCustomRole.
func (in *IamCustomRole) DeepCopy() *IamCustomRole {
	if in == nil {
		return nil
	}
	out := new(IamCustomRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamCustomRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamCustomRoleList) DeepCopyInto(out *IamCustomRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamCustomRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamCustomRoleList.
func (in *IamCustomRoleList) DeepCopy() *IamCustomRoleList {
	if in == nil {
		return nil
	}
	out := new(IamCustomRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamCustomRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamCustomRoleSpec) DeepCopyInto(out *IamCustomRoleSpec) {
	*out = *in
	in.IamCustomRoleSpec2.DeepCopyInto(&out.IamCustomRoleSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamCustomRoleSpec.
func (in *IamCustomRoleSpec) DeepCopy() *IamCustomRoleSpec {
	if in == nil {
		return nil
	}
	out := new(IamCustomRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamCustomRoleSpec2) DeepCopyInto(out *IamCustomRoleSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Deleted != nil {
		in, out := &in.Deleted, &out.Deleted
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.Stage != nil {
		in, out := &in.Stage, &out.Stage
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamCustomRoleSpec2.
func (in *IamCustomRoleSpec2) DeepCopy() *IamCustomRoleSpec2 {
	if in == nil {
		return nil
	}
	out := new(IamCustomRoleSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamCustomRoleStatus) DeepCopyInto(out *IamCustomRoleStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamCustomRoleStatus.
func (in *IamCustomRoleStatus) DeepCopy() *IamCustomRoleStatus {
	if in == nil {
		return nil
	}
	out := new(IamCustomRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMember) DeepCopyInto(out *IamMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMember.
func (in *IamMember) DeepCopy() *IamMember {
	if in == nil {
		return nil
	}
	out := new(IamMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberList) DeepCopyInto(out *IamMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberList.
func (in *IamMemberList) DeepCopy() *IamMemberList {
	if in == nil {
		return nil
	}
	out := new(IamMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberSpec) DeepCopyInto(out *IamMemberSpec) {
	*out = *in
	in.IamMemberSpec2.DeepCopyInto(&out.IamMemberSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberSpec.
func (in *IamMemberSpec) DeepCopy() *IamMemberSpec {
	if in == nil {
		return nil
	}
	out := new(IamMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberSpec2) DeepCopyInto(out *IamMemberSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(IamMemberSpecCondition)
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
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberSpec2.
func (in *IamMemberSpec2) DeepCopy() *IamMemberSpec2 {
	if in == nil {
		return nil
	}
	out := new(IamMemberSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberSpecCondition) DeepCopyInto(out *IamMemberSpecCondition) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberSpecCondition.
func (in *IamMemberSpecCondition) DeepCopy() *IamMemberSpecCondition {
	if in == nil {
		return nil
	}
	out := new(IamMemberSpecCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberStatus) DeepCopyInto(out *IamMemberStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberStatus.
func (in *IamMemberStatus) DeepCopy() *IamMemberStatus {
	if in == nil {
		return nil
	}
	out := new(IamMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicy) DeepCopyInto(out *IamPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicy.
func (in *IamPolicy) DeepCopy() *IamPolicy {
	if in == nil {
		return nil
	}
	out := new(IamPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicyList) DeepCopyInto(out *IamPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicyList.
func (in *IamPolicyList) DeepCopy() *IamPolicyList {
	if in == nil {
		return nil
	}
	out := new(IamPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicySpec) DeepCopyInto(out *IamPolicySpec) {
	*out = *in
	in.IamPolicySpec2.DeepCopyInto(&out.IamPolicySpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicySpec.
func (in *IamPolicySpec) DeepCopy() *IamPolicySpec {
	if in == nil {
		return nil
	}
	out := new(IamPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicySpec2) DeepCopyInto(out *IamPolicySpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.PolicyData != nil {
		in, out := &in.PolicyData, &out.PolicyData
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicySpec2.
func (in *IamPolicySpec2) DeepCopy() *IamPolicySpec2 {
	if in == nil {
		return nil
	}
	out := new(IamPolicySpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicyStatus) DeepCopyInto(out *IamPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicyStatus.
func (in *IamPolicyStatus) DeepCopy() *IamPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(IamPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	in.PolicySpec2.DeepCopyInto(&out.PolicySpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec2) DeepCopyInto(out *PolicySpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.BooleanPolicy != nil {
		in, out := &in.BooleanPolicy, &out.BooleanPolicy
		*out = new(PolicySpecBooleanPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Constraint != nil {
		in, out := &in.Constraint, &out.Constraint
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ListPolicy != nil {
		in, out := &in.ListPolicy, &out.ListPolicy
		*out = new(PolicySpecListPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.RestorePolicy != nil {
		in, out := &in.RestorePolicy, &out.RestorePolicy
		*out = new(PolicySpecRestorePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec2.
func (in *PolicySpec2) DeepCopy() *PolicySpec2 {
	if in == nil {
		return nil
	}
	out := new(PolicySpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecBooleanPolicy) DeepCopyInto(out *PolicySpecBooleanPolicy) {
	*out = *in
	if in.Enforced != nil {
		in, out := &in.Enforced, &out.Enforced
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecBooleanPolicy.
func (in *PolicySpecBooleanPolicy) DeepCopy() *PolicySpecBooleanPolicy {
	if in == nil {
		return nil
	}
	out := new(PolicySpecBooleanPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecListPolicy) DeepCopyInto(out *PolicySpecListPolicy) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = new(PolicySpecListPolicyAllow)
		(*in).DeepCopyInto(*out)
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = new(PolicySpecListPolicyDeny)
		(*in).DeepCopyInto(*out)
	}
	if in.InheritFromParent != nil {
		in, out := &in.InheritFromParent, &out.InheritFromParent
		*out = new(bool)
		**out = **in
	}
	if in.SuggestedValue != nil {
		in, out := &in.SuggestedValue, &out.SuggestedValue
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecListPolicy.
func (in *PolicySpecListPolicy) DeepCopy() *PolicySpecListPolicy {
	if in == nil {
		return nil
	}
	out := new(PolicySpecListPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecListPolicyAllow) DeepCopyInto(out *PolicySpecListPolicyAllow) {
	*out = *in
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = new(bool)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecListPolicyAllow.
func (in *PolicySpecListPolicyAllow) DeepCopy() *PolicySpecListPolicyAllow {
	if in == nil {
		return nil
	}
	out := new(PolicySpecListPolicyAllow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecListPolicyDeny) DeepCopyInto(out *PolicySpecListPolicyDeny) {
	*out = *in
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = new(bool)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecListPolicyDeny.
func (in *PolicySpecListPolicyDeny) DeepCopy() *PolicySpecListPolicyDeny {
	if in == nil {
		return nil
	}
	out := new(PolicySpecListPolicyDeny)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecRestorePolicy) DeepCopyInto(out *PolicySpecRestorePolicy) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecRestorePolicy.
func (in *PolicySpecRestorePolicy) DeepCopy() *PolicySpecRestorePolicy {
	if in == nil {
		return nil
	}
	out := new(PolicySpecRestorePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

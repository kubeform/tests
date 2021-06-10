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
func (in *Fleet) DeepCopyInto(out *Fleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet.
func (in *Fleet) DeepCopy() *Fleet {
	if in == nil {
		return nil
	}
	out := new(Fleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetList) DeepCopyInto(out *FleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetList.
func (in *FleetList) DeepCopy() *FleetList {
	if in == nil {
		return nil
	}
	out := new(FleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpec) DeepCopyInto(out *FleetSpec) {
	*out = *in
	in.FleetSpec2.DeepCopyInto(&out.FleetSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpec.
func (in *FleetSpec) DeepCopy() *FleetSpec {
	if in == nil {
		return nil
	}
	out := new(FleetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpec2) DeepCopyInto(out *FleetSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AuditStreamArn != nil {
		in, out := &in.AuditStreamArn, &out.AuditStreamArn
		*out = new(string)
		**out = **in
	}
	if in.CompanyCode != nil {
		in, out := &in.CompanyCode, &out.CompanyCode
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.DeviceCaCertificate != nil {
		in, out := &in.DeviceCaCertificate, &out.DeviceCaCertificate
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(FleetSpecIdentityProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(FleetSpecNetwork)
		(*in).DeepCopyInto(*out)
	}
	if in.OptimizeForEndUserLocation != nil {
		in, out := &in.OptimizeForEndUserLocation, &out.OptimizeForEndUserLocation
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpec2.
func (in *FleetSpec2) DeepCopy() *FleetSpec2 {
	if in == nil {
		return nil
	}
	out := new(FleetSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpecIdentityProvider) DeepCopyInto(out *FleetSpecIdentityProvider) {
	*out = *in
	if in.SamlMetadata != nil {
		in, out := &in.SamlMetadata, &out.SamlMetadata
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpecIdentityProvider.
func (in *FleetSpecIdentityProvider) DeepCopy() *FleetSpecIdentityProvider {
	if in == nil {
		return nil
	}
	out := new(FleetSpecIdentityProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpecNetwork) DeepCopyInto(out *FleetSpecNetwork) {
	*out = *in
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpecNetwork.
func (in *FleetSpecNetwork) DeepCopy() *FleetSpecNetwork {
	if in == nil {
		return nil
	}
	out := new(FleetSpecNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetStatus) DeepCopyInto(out *FleetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetStatus.
func (in *FleetStatus) DeepCopy() *FleetStatus {
	if in == nil {
		return nil
	}
	out := new(FleetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteCertificateAuthorityAssociation) DeepCopyInto(out *WebsiteCertificateAuthorityAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteCertificateAuthorityAssociation.
func (in *WebsiteCertificateAuthorityAssociation) DeepCopy() *WebsiteCertificateAuthorityAssociation {
	if in == nil {
		return nil
	}
	out := new(WebsiteCertificateAuthorityAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebsiteCertificateAuthorityAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteCertificateAuthorityAssociationList) DeepCopyInto(out *WebsiteCertificateAuthorityAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebsiteCertificateAuthorityAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteCertificateAuthorityAssociationList.
func (in *WebsiteCertificateAuthorityAssociationList) DeepCopy() *WebsiteCertificateAuthorityAssociationList {
	if in == nil {
		return nil
	}
	out := new(WebsiteCertificateAuthorityAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebsiteCertificateAuthorityAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteCertificateAuthorityAssociationSpec) DeepCopyInto(out *WebsiteCertificateAuthorityAssociationSpec) {
	*out = *in
	in.WebsiteCertificateAuthorityAssociationSpec2.DeepCopyInto(&out.WebsiteCertificateAuthorityAssociationSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteCertificateAuthorityAssociationSpec.
func (in *WebsiteCertificateAuthorityAssociationSpec) DeepCopy() *WebsiteCertificateAuthorityAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(WebsiteCertificateAuthorityAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteCertificateAuthorityAssociationSpec2) DeepCopyInto(out *WebsiteCertificateAuthorityAssociationSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FleetArn != nil {
		in, out := &in.FleetArn, &out.FleetArn
		*out = new(string)
		**out = **in
	}
	if in.WebsiteCaID != nil {
		in, out := &in.WebsiteCaID, &out.WebsiteCaID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteCertificateAuthorityAssociationSpec2.
func (in *WebsiteCertificateAuthorityAssociationSpec2) DeepCopy() *WebsiteCertificateAuthorityAssociationSpec2 {
	if in == nil {
		return nil
	}
	out := new(WebsiteCertificateAuthorityAssociationSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteCertificateAuthorityAssociationStatus) DeepCopyInto(out *WebsiteCertificateAuthorityAssociationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteCertificateAuthorityAssociationStatus.
func (in *WebsiteCertificateAuthorityAssociationStatus) DeepCopy() *WebsiteCertificateAuthorityAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(WebsiteCertificateAuthorityAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

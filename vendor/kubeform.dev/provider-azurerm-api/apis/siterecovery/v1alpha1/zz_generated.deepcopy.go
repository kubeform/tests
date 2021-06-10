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
func (in *Fabric) DeepCopyInto(out *Fabric) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fabric.
func (in *Fabric) DeepCopy() *Fabric {
	if in == nil {
		return nil
	}
	out := new(Fabric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fabric) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricList) DeepCopyInto(out *FabricList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fabric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricList.
func (in *FabricList) DeepCopy() *FabricList {
	if in == nil {
		return nil
	}
	out := new(FabricList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricSpec) DeepCopyInto(out *FabricSpec) {
	*out = *in
	in.FabricSpec2.DeepCopyInto(&out.FabricSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricSpec.
func (in *FabricSpec) DeepCopy() *FabricSpec {
	if in == nil {
		return nil
	}
	out := new(FabricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricSpec2) DeepCopyInto(out *FabricSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricSpec2.
func (in *FabricSpec2) DeepCopy() *FabricSpec2 {
	if in == nil {
		return nil
	}
	out := new(FabricSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricStatus) DeepCopyInto(out *FabricStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricStatus.
func (in *FabricStatus) DeepCopy() *FabricStatus {
	if in == nil {
		return nil
	}
	out := new(FabricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkMapping) DeepCopyInto(out *NetworkMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkMapping.
func (in *NetworkMapping) DeepCopy() *NetworkMapping {
	if in == nil {
		return nil
	}
	out := new(NetworkMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkMappingList) DeepCopyInto(out *NetworkMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkMappingList.
func (in *NetworkMappingList) DeepCopy() *NetworkMappingList {
	if in == nil {
		return nil
	}
	out := new(NetworkMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkMappingSpec) DeepCopyInto(out *NetworkMappingSpec) {
	*out = *in
	in.NetworkMappingSpec2.DeepCopyInto(&out.NetworkMappingSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkMappingSpec.
func (in *NetworkMappingSpec) DeepCopy() *NetworkMappingSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkMappingSpec2) DeepCopyInto(out *NetworkMappingSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SourceNetworkID != nil {
		in, out := &in.SourceNetworkID, &out.SourceNetworkID
		*out = new(string)
		**out = **in
	}
	if in.SourceRecoveryFabricName != nil {
		in, out := &in.SourceRecoveryFabricName, &out.SourceRecoveryFabricName
		*out = new(string)
		**out = **in
	}
	if in.TargetNetworkID != nil {
		in, out := &in.TargetNetworkID, &out.TargetNetworkID
		*out = new(string)
		**out = **in
	}
	if in.TargetRecoveryFabricName != nil {
		in, out := &in.TargetRecoveryFabricName, &out.TargetRecoveryFabricName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkMappingSpec2.
func (in *NetworkMappingSpec2) DeepCopy() *NetworkMappingSpec2 {
	if in == nil {
		return nil
	}
	out := new(NetworkMappingSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkMappingStatus) DeepCopyInto(out *NetworkMappingStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkMappingStatus.
func (in *NetworkMappingStatus) DeepCopy() *NetworkMappingStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainer) DeepCopyInto(out *ProtectionContainer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainer.
func (in *ProtectionContainer) DeepCopy() *ProtectionContainer {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectionContainer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerList) DeepCopyInto(out *ProtectionContainerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProtectionContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerList.
func (in *ProtectionContainerList) DeepCopy() *ProtectionContainerList {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectionContainerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerMapping) DeepCopyInto(out *ProtectionContainerMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerMapping.
func (in *ProtectionContainerMapping) DeepCopy() *ProtectionContainerMapping {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectionContainerMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerMappingList) DeepCopyInto(out *ProtectionContainerMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProtectionContainerMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerMappingList.
func (in *ProtectionContainerMappingList) DeepCopy() *ProtectionContainerMappingList {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectionContainerMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerMappingSpec) DeepCopyInto(out *ProtectionContainerMappingSpec) {
	*out = *in
	in.ProtectionContainerMappingSpec2.DeepCopyInto(&out.ProtectionContainerMappingSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerMappingSpec.
func (in *ProtectionContainerMappingSpec) DeepCopy() *ProtectionContainerMappingSpec {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerMappingSpec2) DeepCopyInto(out *ProtectionContainerMappingSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryFabricName != nil {
		in, out := &in.RecoveryFabricName, &out.RecoveryFabricName
		*out = new(string)
		**out = **in
	}
	if in.RecoveryReplicationPolicyID != nil {
		in, out := &in.RecoveryReplicationPolicyID, &out.RecoveryReplicationPolicyID
		*out = new(string)
		**out = **in
	}
	if in.RecoverySourceProtectionContainerName != nil {
		in, out := &in.RecoverySourceProtectionContainerName, &out.RecoverySourceProtectionContainerName
		*out = new(string)
		**out = **in
	}
	if in.RecoveryTargetProtectionContainerID != nil {
		in, out := &in.RecoveryTargetProtectionContainerID, &out.RecoveryTargetProtectionContainerID
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerMappingSpec2.
func (in *ProtectionContainerMappingSpec2) DeepCopy() *ProtectionContainerMappingSpec2 {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerMappingSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerMappingStatus) DeepCopyInto(out *ProtectionContainerMappingStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerMappingStatus.
func (in *ProtectionContainerMappingStatus) DeepCopy() *ProtectionContainerMappingStatus {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerSpec) DeepCopyInto(out *ProtectionContainerSpec) {
	*out = *in
	in.ProtectionContainerSpec2.DeepCopyInto(&out.ProtectionContainerSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerSpec.
func (in *ProtectionContainerSpec) DeepCopy() *ProtectionContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerSpec2) DeepCopyInto(out *ProtectionContainerSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryFabricName != nil {
		in, out := &in.RecoveryFabricName, &out.RecoveryFabricName
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerSpec2.
func (in *ProtectionContainerSpec2) DeepCopy() *ProtectionContainerSpec2 {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionContainerStatus) DeepCopyInto(out *ProtectionContainerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionContainerStatus.
func (in *ProtectionContainerStatus) DeepCopy() *ProtectionContainerStatus {
	if in == nil {
		return nil
	}
	out := new(ProtectionContainerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedVm) DeepCopyInto(out *ReplicatedVm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedVm.
func (in *ReplicatedVm) DeepCopy() *ReplicatedVm {
	if in == nil {
		return nil
	}
	out := new(ReplicatedVm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicatedVm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedVmList) DeepCopyInto(out *ReplicatedVmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicatedVm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedVmList.
func (in *ReplicatedVmList) DeepCopy() *ReplicatedVmList {
	if in == nil {
		return nil
	}
	out := new(ReplicatedVmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicatedVmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedVmSpec) DeepCopyInto(out *ReplicatedVmSpec) {
	*out = *in
	in.ReplicatedVmSpec2.DeepCopyInto(&out.ReplicatedVmSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedVmSpec.
func (in *ReplicatedVmSpec) DeepCopy() *ReplicatedVmSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicatedVmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedVmSpec2) DeepCopyInto(out *ReplicatedVmSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = make([]ReplicatedVmSpecManagedDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]ReplicatedVmSpecNetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RecoveryReplicationPolicyID != nil {
		in, out := &in.RecoveryReplicationPolicyID, &out.RecoveryReplicationPolicyID
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SourceRecoveryFabricName != nil {
		in, out := &in.SourceRecoveryFabricName, &out.SourceRecoveryFabricName
		*out = new(string)
		**out = **in
	}
	if in.SourceRecoveryProtectionContainerName != nil {
		in, out := &in.SourceRecoveryProtectionContainerName, &out.SourceRecoveryProtectionContainerName
		*out = new(string)
		**out = **in
	}
	if in.SourceVmID != nil {
		in, out := &in.SourceVmID, &out.SourceVmID
		*out = new(string)
		**out = **in
	}
	if in.TargetAvailabilitySetID != nil {
		in, out := &in.TargetAvailabilitySetID, &out.TargetAvailabilitySetID
		*out = new(string)
		**out = **in
	}
	if in.TargetNetworkID != nil {
		in, out := &in.TargetNetworkID, &out.TargetNetworkID
		*out = new(string)
		**out = **in
	}
	if in.TargetRecoveryFabricID != nil {
		in, out := &in.TargetRecoveryFabricID, &out.TargetRecoveryFabricID
		*out = new(string)
		**out = **in
	}
	if in.TargetRecoveryProtectionContainerID != nil {
		in, out := &in.TargetRecoveryProtectionContainerID, &out.TargetRecoveryProtectionContainerID
		*out = new(string)
		**out = **in
	}
	if in.TargetResourceGroupID != nil {
		in, out := &in.TargetResourceGroupID, &out.TargetResourceGroupID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedVmSpec2.
func (in *ReplicatedVmSpec2) DeepCopy() *ReplicatedVmSpec2 {
	if in == nil {
		return nil
	}
	out := new(ReplicatedVmSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedVmSpecManagedDisk) DeepCopyInto(out *ReplicatedVmSpecManagedDisk) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(string)
		**out = **in
	}
	if in.StagingStorageAccountID != nil {
		in, out := &in.StagingStorageAccountID, &out.StagingStorageAccountID
		*out = new(string)
		**out = **in
	}
	if in.TargetDiskType != nil {
		in, out := &in.TargetDiskType, &out.TargetDiskType
		*out = new(string)
		**out = **in
	}
	if in.TargetReplicaDiskType != nil {
		in, out := &in.TargetReplicaDiskType, &out.TargetReplicaDiskType
		*out = new(string)
		**out = **in
	}
	if in.TargetResourceGroupID != nil {
		in, out := &in.TargetResourceGroupID, &out.TargetResourceGroupID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedVmSpecManagedDisk.
func (in *ReplicatedVmSpecManagedDisk) DeepCopy() *ReplicatedVmSpecManagedDisk {
	if in == nil {
		return nil
	}
	out := new(ReplicatedVmSpecManagedDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedVmSpecNetworkInterface) DeepCopyInto(out *ReplicatedVmSpecNetworkInterface) {
	*out = *in
	if in.RecoveryPublicIPAddressID != nil {
		in, out := &in.RecoveryPublicIPAddressID, &out.RecoveryPublicIPAddressID
		*out = new(string)
		**out = **in
	}
	if in.SourceNetworkInterfaceID != nil {
		in, out := &in.SourceNetworkInterfaceID, &out.SourceNetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.TargetStaticIP != nil {
		in, out := &in.TargetStaticIP, &out.TargetStaticIP
		*out = new(string)
		**out = **in
	}
	if in.TargetSubnetName != nil {
		in, out := &in.TargetSubnetName, &out.TargetSubnetName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedVmSpecNetworkInterface.
func (in *ReplicatedVmSpecNetworkInterface) DeepCopy() *ReplicatedVmSpecNetworkInterface {
	if in == nil {
		return nil
	}
	out := new(ReplicatedVmSpecNetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedVmStatus) DeepCopyInto(out *ReplicatedVmStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedVmStatus.
func (in *ReplicatedVmStatus) DeepCopy() *ReplicatedVmStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicatedVmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationPolicy) DeepCopyInto(out *ReplicationPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationPolicy.
func (in *ReplicationPolicy) DeepCopy() *ReplicationPolicy {
	if in == nil {
		return nil
	}
	out := new(ReplicationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationPolicyList) DeepCopyInto(out *ReplicationPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationPolicyList.
func (in *ReplicationPolicyList) DeepCopy() *ReplicationPolicyList {
	if in == nil {
		return nil
	}
	out := new(ReplicationPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationPolicySpec) DeepCopyInto(out *ReplicationPolicySpec) {
	*out = *in
	in.ReplicationPolicySpec2.DeepCopyInto(&out.ReplicationPolicySpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationPolicySpec.
func (in *ReplicationPolicySpec) DeepCopy() *ReplicationPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationPolicySpec2) DeepCopyInto(out *ReplicationPolicySpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.ApplicationConsistentSnapshotFrequencyInMinutes != nil {
		in, out := &in.ApplicationConsistentSnapshotFrequencyInMinutes, &out.ApplicationConsistentSnapshotFrequencyInMinutes
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryPointRetentionInMinutes != nil {
		in, out := &in.RecoveryPointRetentionInMinutes, &out.RecoveryPointRetentionInMinutes
		*out = new(int64)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationPolicySpec2.
func (in *ReplicationPolicySpec2) DeepCopy() *ReplicationPolicySpec2 {
	if in == nil {
		return nil
	}
	out := new(ReplicationPolicySpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationPolicyStatus) DeepCopyInto(out *ReplicationPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationPolicyStatus.
func (in *ReplicationPolicyStatus) DeepCopy() *ReplicationPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

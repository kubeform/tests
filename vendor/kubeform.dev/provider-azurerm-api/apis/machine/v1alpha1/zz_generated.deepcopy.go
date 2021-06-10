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
func (in *LearningInferenceCluster) DeepCopyInto(out *LearningInferenceCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningInferenceCluster.
func (in *LearningInferenceCluster) DeepCopy() *LearningInferenceCluster {
	if in == nil {
		return nil
	}
	out := new(LearningInferenceCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LearningInferenceCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningInferenceClusterList) DeepCopyInto(out *LearningInferenceClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LearningInferenceCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningInferenceClusterList.
func (in *LearningInferenceClusterList) DeepCopy() *LearningInferenceClusterList {
	if in == nil {
		return nil
	}
	out := new(LearningInferenceClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LearningInferenceClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningInferenceClusterSpec) DeepCopyInto(out *LearningInferenceClusterSpec) {
	*out = *in
	in.LearningInferenceClusterSpec2.DeepCopyInto(&out.LearningInferenceClusterSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningInferenceClusterSpec.
func (in *LearningInferenceClusterSpec) DeepCopy() *LearningInferenceClusterSpec {
	if in == nil {
		return nil
	}
	out := new(LearningInferenceClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningInferenceClusterSpec2) DeepCopyInto(out *LearningInferenceClusterSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.ClusterPurpose != nil {
		in, out := &in.ClusterPurpose, &out.ClusterPurpose
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KubernetesClusterID != nil {
		in, out := &in.KubernetesClusterID, &out.KubernetesClusterID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MachineLearningWorkspaceID != nil {
		in, out := &in.MachineLearningWorkspaceID, &out.MachineLearningWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Ssl != nil {
		in, out := &in.Ssl, &out.Ssl
		*out = new(LearningInferenceClusterSpecSsl)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningInferenceClusterSpec2.
func (in *LearningInferenceClusterSpec2) DeepCopy() *LearningInferenceClusterSpec2 {
	if in == nil {
		return nil
	}
	out := new(LearningInferenceClusterSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningInferenceClusterSpecSsl) DeepCopyInto(out *LearningInferenceClusterSpecSsl) {
	*out = *in
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = new(string)
		**out = **in
	}
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningInferenceClusterSpecSsl.
func (in *LearningInferenceClusterSpecSsl) DeepCopy() *LearningInferenceClusterSpecSsl {
	if in == nil {
		return nil
	}
	out := new(LearningInferenceClusterSpecSsl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningInferenceClusterStatus) DeepCopyInto(out *LearningInferenceClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningInferenceClusterStatus.
func (in *LearningInferenceClusterStatus) DeepCopy() *LearningInferenceClusterStatus {
	if in == nil {
		return nil
	}
	out := new(LearningInferenceClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningWorkspace) DeepCopyInto(out *LearningWorkspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningWorkspace.
func (in *LearningWorkspace) DeepCopy() *LearningWorkspace {
	if in == nil {
		return nil
	}
	out := new(LearningWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LearningWorkspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningWorkspaceList) DeepCopyInto(out *LearningWorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LearningWorkspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningWorkspaceList.
func (in *LearningWorkspaceList) DeepCopy() *LearningWorkspaceList {
	if in == nil {
		return nil
	}
	out := new(LearningWorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LearningWorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningWorkspaceSpec) DeepCopyInto(out *LearningWorkspaceSpec) {
	*out = *in
	in.LearningWorkspaceSpec2.DeepCopyInto(&out.LearningWorkspaceSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningWorkspaceSpec.
func (in *LearningWorkspaceSpec) DeepCopy() *LearningWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(LearningWorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningWorkspaceSpec2) DeepCopyInto(out *LearningWorkspaceSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.ApplicationInsightsID != nil {
		in, out := &in.ApplicationInsightsID, &out.ApplicationInsightsID
		*out = new(string)
		**out = **in
	}
	if in.ContainerRegistryID != nil {
		in, out := &in.ContainerRegistryID, &out.ContainerRegistryID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FriendlyName != nil {
		in, out := &in.FriendlyName, &out.FriendlyName
		*out = new(string)
		**out = **in
	}
	if in.HighBusinessImpact != nil {
		in, out := &in.HighBusinessImpact, &out.HighBusinessImpact
		*out = new(bool)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(LearningWorkspaceSpecIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultID != nil {
		in, out := &in.KeyVaultID, &out.KeyVaultID
		*out = new(string)
		**out = **in
	}
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
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningWorkspaceSpec2.
func (in *LearningWorkspaceSpec2) DeepCopy() *LearningWorkspaceSpec2 {
	if in == nil {
		return nil
	}
	out := new(LearningWorkspaceSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningWorkspaceSpecIdentity) DeepCopyInto(out *LearningWorkspaceSpecIdentity) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningWorkspaceSpecIdentity.
func (in *LearningWorkspaceSpecIdentity) DeepCopy() *LearningWorkspaceSpecIdentity {
	if in == nil {
		return nil
	}
	out := new(LearningWorkspaceSpecIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LearningWorkspaceStatus) DeepCopyInto(out *LearningWorkspaceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LearningWorkspaceStatus.
func (in *LearningWorkspaceStatus) DeepCopy() *LearningWorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(LearningWorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}

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
func (in *ServiceEnvironment) DeepCopyInto(out *ServiceEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceEnvironment.
func (in *ServiceEnvironment) DeepCopy() *ServiceEnvironment {
	if in == nil {
		return nil
	}
	out := new(ServiceEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceEnvironmentList) DeepCopyInto(out *ServiceEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceEnvironmentList.
func (in *ServiceEnvironmentList) DeepCopy() *ServiceEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(ServiceEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceEnvironmentSpec) DeepCopyInto(out *ServiceEnvironmentSpec) {
	*out = *in
	in.ServiceEnvironmentSpec2.DeepCopyInto(&out.ServiceEnvironmentSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceEnvironmentSpec.
func (in *ServiceEnvironmentSpec) DeepCopy() *ServiceEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceEnvironmentSpec2) DeepCopyInto(out *ServiceEnvironmentSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.AccessEndpointType != nil {
		in, out := &in.AccessEndpointType, &out.AccessEndpointType
		*out = new(string)
		**out = **in
	}
	if in.ConnectorEndpointIPAddresses != nil {
		in, out := &in.ConnectorEndpointIPAddresses, &out.ConnectorEndpointIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConnectorOutboundIPAddresses != nil {
		in, out := &in.ConnectorOutboundIPAddresses, &out.ConnectorOutboundIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.VirtualNetworkSubnetIDS != nil {
		in, out := &in.VirtualNetworkSubnetIDS, &out.VirtualNetworkSubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkflowEndpointIPAddresses != nil {
		in, out := &in.WorkflowEndpointIPAddresses, &out.WorkflowEndpointIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkflowOutboundIPAddresses != nil {
		in, out := &in.WorkflowOutboundIPAddresses, &out.WorkflowOutboundIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceEnvironmentSpec2.
func (in *ServiceEnvironmentSpec2) DeepCopy() *ServiceEnvironmentSpec2 {
	if in == nil {
		return nil
	}
	out := new(ServiceEnvironmentSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceEnvironmentStatus) DeepCopyInto(out *ServiceEnvironmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceEnvironmentStatus.
func (in *ServiceEnvironmentStatus) DeepCopy() *ServiceEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

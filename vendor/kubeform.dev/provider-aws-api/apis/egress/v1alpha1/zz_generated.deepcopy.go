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
func (in *OnlyInternetGateway) DeepCopyInto(out *OnlyInternetGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlyInternetGateway.
func (in *OnlyInternetGateway) DeepCopy() *OnlyInternetGateway {
	if in == nil {
		return nil
	}
	out := new(OnlyInternetGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnlyInternetGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlyInternetGatewayList) DeepCopyInto(out *OnlyInternetGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OnlyInternetGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlyInternetGatewayList.
func (in *OnlyInternetGatewayList) DeepCopy() *OnlyInternetGatewayList {
	if in == nil {
		return nil
	}
	out := new(OnlyInternetGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnlyInternetGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlyInternetGatewaySpec) DeepCopyInto(out *OnlyInternetGatewaySpec) {
	*out = *in
	in.OnlyInternetGatewaySpec2.DeepCopyInto(&out.OnlyInternetGatewaySpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlyInternetGatewaySpec.
func (in *OnlyInternetGatewaySpec) DeepCopy() *OnlyInternetGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(OnlyInternetGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlyInternetGatewaySpec2) DeepCopyInto(out *OnlyInternetGatewaySpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
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
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlyInternetGatewaySpec2.
func (in *OnlyInternetGatewaySpec2) DeepCopy() *OnlyInternetGatewaySpec2 {
	if in == nil {
		return nil
	}
	out := new(OnlyInternetGatewaySpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlyInternetGatewayStatus) DeepCopyInto(out *OnlyInternetGatewayStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlyInternetGatewayStatus.
func (in *OnlyInternetGatewayStatus) DeepCopy() *OnlyInternetGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(OnlyInternetGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

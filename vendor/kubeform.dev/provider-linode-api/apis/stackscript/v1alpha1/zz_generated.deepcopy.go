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
func (in *Stackscript) DeepCopyInto(out *Stackscript) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stackscript.
func (in *Stackscript) DeepCopy() *Stackscript {
	if in == nil {
		return nil
	}
	out := new(Stackscript)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stackscript) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackscriptList) DeepCopyInto(out *StackscriptList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stackscript, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackscriptList.
func (in *StackscriptList) DeepCopy() *StackscriptList {
	if in == nil {
		return nil
	}
	out := new(StackscriptList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackscriptList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackscriptSpec) DeepCopyInto(out *StackscriptSpec) {
	*out = *in
	in.StackscriptSpec2.DeepCopyInto(&out.StackscriptSpec2)
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(StackscriptSpec2)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackscriptSpec.
func (in *StackscriptSpec) DeepCopy() *StackscriptSpec {
	if in == nil {
		return nil
	}
	out := new(StackscriptSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackscriptSpec2) DeepCopyInto(out *StackscriptSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.DeploymentsActive != nil {
		in, out := &in.DeploymentsActive, &out.DeploymentsActive
		*out = new(int64)
		**out = **in
	}
	if in.DeploymentsTotal != nil {
		in, out := &in.DeploymentsTotal, &out.DeploymentsTotal
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.RevNote != nil {
		in, out := &in.RevNote, &out.RevNote
		*out = new(string)
		**out = **in
	}
	if in.Script != nil {
		in, out := &in.Script, &out.Script
		*out = new(string)
		**out = **in
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
	if in.UserDefinedFields != nil {
		in, out := &in.UserDefinedFields, &out.UserDefinedFields
		*out = make([]StackscriptSpecUserDefinedFields, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserGravatarID != nil {
		in, out := &in.UserGravatarID, &out.UserGravatarID
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackscriptSpec2.
func (in *StackscriptSpec2) DeepCopy() *StackscriptSpec2 {
	if in == nil {
		return nil
	}
	out := new(StackscriptSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackscriptSpecUserDefinedFields) DeepCopyInto(out *StackscriptSpecUserDefinedFields) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(string)
		**out = **in
	}
	if in.Example != nil {
		in, out := &in.Example, &out.Example
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.ManyOf != nil {
		in, out := &in.ManyOf, &out.ManyOf
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackscriptSpecUserDefinedFields.
func (in *StackscriptSpecUserDefinedFields) DeepCopy() *StackscriptSpecUserDefinedFields {
	if in == nil {
		return nil
	}
	out := new(StackscriptSpecUserDefinedFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackscriptStatus) DeepCopyInto(out *StackscriptStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackscriptStatus.
func (in *StackscriptStatus) DeepCopy() *StackscriptStatus {
	if in == nil {
		return nil
	}
	out := new(StackscriptStatus)
	in.DeepCopyInto(out)
	return out
}
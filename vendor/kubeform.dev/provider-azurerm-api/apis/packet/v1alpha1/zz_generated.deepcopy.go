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
func (in *Capture) DeepCopyInto(out *Capture) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capture.
func (in *Capture) DeepCopy() *Capture {
	if in == nil {
		return nil
	}
	out := new(Capture)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Capture) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureList) DeepCopyInto(out *CaptureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Capture, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureList.
func (in *CaptureList) DeepCopy() *CaptureList {
	if in == nil {
		return nil
	}
	out := new(CaptureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CaptureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureSpec) DeepCopyInto(out *CaptureSpec) {
	*out = *in
	in.CaptureSpec2.DeepCopyInto(&out.CaptureSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureSpec.
func (in *CaptureSpec) DeepCopy() *CaptureSpec {
	if in == nil {
		return nil
	}
	out := new(CaptureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureSpec2) DeepCopyInto(out *CaptureSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]CaptureSpecFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaximumBytesPerPacket != nil {
		in, out := &in.MaximumBytesPerPacket, &out.MaximumBytesPerPacket
		*out = new(int64)
		**out = **in
	}
	if in.MaximumBytesPerSession != nil {
		in, out := &in.MaximumBytesPerSession, &out.MaximumBytesPerSession
		*out = new(int64)
		**out = **in
	}
	if in.MaximumCaptureDuration != nil {
		in, out := &in.MaximumCaptureDuration, &out.MaximumCaptureDuration
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkWatcherName != nil {
		in, out := &in.NetworkWatcherName, &out.NetworkWatcherName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.StorageLocation != nil {
		in, out := &in.StorageLocation, &out.StorageLocation
		*out = new(CaptureSpecStorageLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetResourceID != nil {
		in, out := &in.TargetResourceID, &out.TargetResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureSpec2.
func (in *CaptureSpec2) DeepCopy() *CaptureSpec2 {
	if in == nil {
		return nil
	}
	out := new(CaptureSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureSpecFilter) DeepCopyInto(out *CaptureSpecFilter) {
	*out = *in
	if in.LocalIPAddress != nil {
		in, out := &in.LocalIPAddress, &out.LocalIPAddress
		*out = new(string)
		**out = **in
	}
	if in.LocalPort != nil {
		in, out := &in.LocalPort, &out.LocalPort
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.RemoteIPAddress != nil {
		in, out := &in.RemoteIPAddress, &out.RemoteIPAddress
		*out = new(string)
		**out = **in
	}
	if in.RemotePort != nil {
		in, out := &in.RemotePort, &out.RemotePort
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureSpecFilter.
func (in *CaptureSpecFilter) DeepCopy() *CaptureSpecFilter {
	if in == nil {
		return nil
	}
	out := new(CaptureSpecFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureSpecStorageLocation) DeepCopyInto(out *CaptureSpecStorageLocation) {
	*out = *in
	if in.FilePath != nil {
		in, out := &in.FilePath, &out.FilePath
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
	if in.StoragePath != nil {
		in, out := &in.StoragePath, &out.StoragePath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureSpecStorageLocation.
func (in *CaptureSpecStorageLocation) DeepCopy() *CaptureSpecStorageLocation {
	if in == nil {
		return nil
	}
	out := new(CaptureSpecStorageLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureStatus) DeepCopyInto(out *CaptureStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureStatus.
func (in *CaptureStatus) DeepCopy() *CaptureStatus {
	if in == nil {
		return nil
	}
	out := new(CaptureStatus)
	in.DeepCopyInto(out)
	return out
}
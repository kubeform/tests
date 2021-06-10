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
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfiguration) DeepCopyInto(out *BucketLifecycleConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfiguration.
func (in *BucketLifecycleConfiguration) DeepCopy() *BucketLifecycleConfiguration {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketLifecycleConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationList) DeepCopyInto(out *BucketLifecycleConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketLifecycleConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationList.
func (in *BucketLifecycleConfigurationList) DeepCopy() *BucketLifecycleConfigurationList {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketLifecycleConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationSpec) DeepCopyInto(out *BucketLifecycleConfigurationSpec) {
	*out = *in
	in.BucketLifecycleConfigurationSpec2.DeepCopyInto(&out.BucketLifecycleConfigurationSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationSpec.
func (in *BucketLifecycleConfigurationSpec) DeepCopy() *BucketLifecycleConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationSpec2) DeepCopyInto(out *BucketLifecycleConfigurationSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]BucketLifecycleConfigurationSpecRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationSpec2.
func (in *BucketLifecycleConfigurationSpec2) DeepCopy() *BucketLifecycleConfigurationSpec2 {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationSpecRule) DeepCopyInto(out *BucketLifecycleConfigurationSpecRule) {
	*out = *in
	if in.AbortIncompleteMultipartUpload != nil {
		in, out := &in.AbortIncompleteMultipartUpload, &out.AbortIncompleteMultipartUpload
		*out = new(BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload)
		(*in).DeepCopyInto(*out)
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = new(BucketLifecycleConfigurationSpecRuleExpiration)
		(*in).DeepCopyInto(*out)
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(BucketLifecycleConfigurationSpecRuleFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationSpecRule.
func (in *BucketLifecycleConfigurationSpecRule) DeepCopy() *BucketLifecycleConfigurationSpecRule {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationSpecRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload) DeepCopyInto(out *BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload) {
	*out = *in
	if in.DaysAfterInitiation != nil {
		in, out := &in.DaysAfterInitiation, &out.DaysAfterInitiation
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload.
func (in *BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload) DeepCopy() *BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationSpecRuleAbortIncompleteMultipartUpload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationSpecRuleExpiration) DeepCopyInto(out *BucketLifecycleConfigurationSpecRuleExpiration) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.ExpiredObjectDeleteMarker != nil {
		in, out := &in.ExpiredObjectDeleteMarker, &out.ExpiredObjectDeleteMarker
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationSpecRuleExpiration.
func (in *BucketLifecycleConfigurationSpecRuleExpiration) DeepCopy() *BucketLifecycleConfigurationSpecRuleExpiration {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationSpecRuleExpiration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationSpecRuleFilter) DeepCopyInto(out *BucketLifecycleConfigurationSpecRuleFilter) {
	*out = *in
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationSpecRuleFilter.
func (in *BucketLifecycleConfigurationSpecRuleFilter) DeepCopy() *BucketLifecycleConfigurationSpecRuleFilter {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationSpecRuleFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLifecycleConfigurationStatus) DeepCopyInto(out *BucketLifecycleConfigurationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLifecycleConfigurationStatus.
func (in *BucketLifecycleConfigurationStatus) DeepCopy() *BucketLifecycleConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(BucketLifecycleConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicy) DeepCopyInto(out *BucketPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicy.
func (in *BucketPolicy) DeepCopy() *BucketPolicy {
	if in == nil {
		return nil
	}
	out := new(BucketPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyList) DeepCopyInto(out *BucketPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyList.
func (in *BucketPolicyList) DeepCopy() *BucketPolicyList {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicySpec) DeepCopyInto(out *BucketPolicySpec) {
	*out = *in
	in.BucketPolicySpec2.DeepCopyInto(&out.BucketPolicySpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicySpec.
func (in *BucketPolicySpec) DeepCopy() *BucketPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BucketPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicySpec2) DeepCopyInto(out *BucketPolicySpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicySpec2.
func (in *BucketPolicySpec2) DeepCopy() *BucketPolicySpec2 {
	if in == nil {
		return nil
	}
	out := new(BucketPolicySpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyStatus) DeepCopyInto(out *BucketPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyStatus.
func (in *BucketPolicyStatus) DeepCopy() *BucketPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.BucketSpec2.DeepCopyInto(&out.BucketSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec2) DeepCopyInto(out *BucketSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.OutpostID != nil {
		in, out := &in.OutpostID, &out.OutpostID
		*out = new(string)
		**out = **in
	}
	if in.PublicAccessBlockEnabled != nil {
		in, out := &in.PublicAccessBlockEnabled, &out.PublicAccessBlockEnabled
		*out = new(bool)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec2.
func (in *BucketSpec2) DeepCopy() *BucketSpec2 {
	if in == nil {
		return nil
	}
	out := new(BucketSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

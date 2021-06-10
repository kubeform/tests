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
func (in *BudgetResourceGroup) DeepCopyInto(out *BudgetResourceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroup.
func (in *BudgetResourceGroup) DeepCopy() *BudgetResourceGroup {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetResourceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupList) DeepCopyInto(out *BudgetResourceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BudgetResourceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupList.
func (in *BudgetResourceGroupList) DeepCopy() *BudgetResourceGroupList {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetResourceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpec) DeepCopyInto(out *BudgetResourceGroupSpec) {
	*out = *in
	in.BudgetResourceGroupSpec2.DeepCopyInto(&out.BudgetResourceGroupSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpec.
func (in *BudgetResourceGroupSpec) DeepCopy() *BudgetResourceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpec2) DeepCopyInto(out *BudgetResourceGroupSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(float64)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(BudgetResourceGroupSpecFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]BudgetResourceGroupSpecNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
	if in.TimePeriod != nil {
		in, out := &in.TimePeriod, &out.TimePeriod
		*out = new(BudgetResourceGroupSpecTimePeriod)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpec2.
func (in *BudgetResourceGroupSpec2) DeepCopy() *BudgetResourceGroupSpec2 {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecFilter) DeepCopyInto(out *BudgetResourceGroupSpecFilter) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]BudgetResourceGroupSpecFilterDimension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(BudgetResourceGroupSpecFilterNot)
		(*in).DeepCopyInto(*out)
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]BudgetResourceGroupSpecFilterTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecFilter.
func (in *BudgetResourceGroupSpecFilter) DeepCopy() *BudgetResourceGroupSpecFilter {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecFilterDimension) DeepCopyInto(out *BudgetResourceGroupSpecFilterDimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecFilterDimension.
func (in *BudgetResourceGroupSpecFilterDimension) DeepCopy() *BudgetResourceGroupSpecFilterDimension {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecFilterDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecFilterNot) DeepCopyInto(out *BudgetResourceGroupSpecFilterNot) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = new(BudgetResourceGroupSpecFilterNotDimension)
		(*in).DeepCopyInto(*out)
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(BudgetResourceGroupSpecFilterNotTag)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecFilterNot.
func (in *BudgetResourceGroupSpecFilterNot) DeepCopy() *BudgetResourceGroupSpecFilterNot {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecFilterNot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecFilterNotDimension) DeepCopyInto(out *BudgetResourceGroupSpecFilterNotDimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecFilterNotDimension.
func (in *BudgetResourceGroupSpecFilterNotDimension) DeepCopy() *BudgetResourceGroupSpecFilterNotDimension {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecFilterNotDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecFilterNotTag) DeepCopyInto(out *BudgetResourceGroupSpecFilterNotTag) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecFilterNotTag.
func (in *BudgetResourceGroupSpecFilterNotTag) DeepCopy() *BudgetResourceGroupSpecFilterNotTag {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecFilterNotTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecFilterTag) DeepCopyInto(out *BudgetResourceGroupSpecFilterTag) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecFilterTag.
func (in *BudgetResourceGroupSpecFilterTag) DeepCopy() *BudgetResourceGroupSpecFilterTag {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecFilterTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecNotification) DeepCopyInto(out *BudgetResourceGroupSpecNotification) {
	*out = *in
	if in.ContactEmails != nil {
		in, out := &in.ContactEmails, &out.ContactEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContactGroups != nil {
		in, out := &in.ContactGroups, &out.ContactGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContactRoles != nil {
		in, out := &in.ContactRoles, &out.ContactRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecNotification.
func (in *BudgetResourceGroupSpecNotification) DeepCopy() *BudgetResourceGroupSpecNotification {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpecTimePeriod) DeepCopyInto(out *BudgetResourceGroupSpecTimePeriod) {
	*out = *in
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(string)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpecTimePeriod.
func (in *BudgetResourceGroupSpecTimePeriod) DeepCopy() *BudgetResourceGroupSpecTimePeriod {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpecTimePeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupStatus) DeepCopyInto(out *BudgetResourceGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupStatus.
func (in *BudgetResourceGroupStatus) DeepCopy() *BudgetResourceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscription) DeepCopyInto(out *BudgetSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscription.
func (in *BudgetSubscription) DeepCopy() *BudgetSubscription {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionList) DeepCopyInto(out *BudgetSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BudgetSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionList.
func (in *BudgetSubscriptionList) DeepCopy() *BudgetSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpec) DeepCopyInto(out *BudgetSubscriptionSpec) {
	*out = *in
	in.BudgetSubscriptionSpec2.DeepCopyInto(&out.BudgetSubscriptionSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpec.
func (in *BudgetSubscriptionSpec) DeepCopy() *BudgetSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpec2) DeepCopyInto(out *BudgetSubscriptionSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(float64)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(BudgetSubscriptionSpecFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]BudgetSubscriptionSpecNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
	if in.TimePeriod != nil {
		in, out := &in.TimePeriod, &out.TimePeriod
		*out = new(BudgetSubscriptionSpecTimePeriod)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpec2.
func (in *BudgetSubscriptionSpec2) DeepCopy() *BudgetSubscriptionSpec2 {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecFilter) DeepCopyInto(out *BudgetSubscriptionSpecFilter) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]BudgetSubscriptionSpecFilterDimension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(BudgetSubscriptionSpecFilterNot)
		(*in).DeepCopyInto(*out)
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]BudgetSubscriptionSpecFilterTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecFilter.
func (in *BudgetSubscriptionSpecFilter) DeepCopy() *BudgetSubscriptionSpecFilter {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecFilterDimension) DeepCopyInto(out *BudgetSubscriptionSpecFilterDimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecFilterDimension.
func (in *BudgetSubscriptionSpecFilterDimension) DeepCopy() *BudgetSubscriptionSpecFilterDimension {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecFilterDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecFilterNot) DeepCopyInto(out *BudgetSubscriptionSpecFilterNot) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = new(BudgetSubscriptionSpecFilterNotDimension)
		(*in).DeepCopyInto(*out)
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(BudgetSubscriptionSpecFilterNotTag)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecFilterNot.
func (in *BudgetSubscriptionSpecFilterNot) DeepCopy() *BudgetSubscriptionSpecFilterNot {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecFilterNot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecFilterNotDimension) DeepCopyInto(out *BudgetSubscriptionSpecFilterNotDimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecFilterNotDimension.
func (in *BudgetSubscriptionSpecFilterNotDimension) DeepCopy() *BudgetSubscriptionSpecFilterNotDimension {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecFilterNotDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecFilterNotTag) DeepCopyInto(out *BudgetSubscriptionSpecFilterNotTag) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecFilterNotTag.
func (in *BudgetSubscriptionSpecFilterNotTag) DeepCopy() *BudgetSubscriptionSpecFilterNotTag {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecFilterNotTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecFilterTag) DeepCopyInto(out *BudgetSubscriptionSpecFilterTag) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecFilterTag.
func (in *BudgetSubscriptionSpecFilterTag) DeepCopy() *BudgetSubscriptionSpecFilterTag {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecFilterTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecNotification) DeepCopyInto(out *BudgetSubscriptionSpecNotification) {
	*out = *in
	if in.ContactEmails != nil {
		in, out := &in.ContactEmails, &out.ContactEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContactGroups != nil {
		in, out := &in.ContactGroups, &out.ContactGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContactRoles != nil {
		in, out := &in.ContactRoles, &out.ContactRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecNotification.
func (in *BudgetSubscriptionSpecNotification) DeepCopy() *BudgetSubscriptionSpecNotification {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpecTimePeriod) DeepCopyInto(out *BudgetSubscriptionSpecTimePeriod) {
	*out = *in
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(string)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpecTimePeriod.
func (in *BudgetSubscriptionSpecTimePeriod) DeepCopy() *BudgetSubscriptionSpecTimePeriod {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpecTimePeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionStatus) DeepCopyInto(out *BudgetSubscriptionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionStatus.
func (in *BudgetSubscriptionStatus) DeepCopy() *BudgetSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

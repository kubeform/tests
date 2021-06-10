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
func (in *BillingBudget) DeepCopyInto(out *BillingBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudget.
func (in *BillingBudget) DeepCopy() *BillingBudget {
	if in == nil {
		return nil
	}
	out := new(BillingBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetList) DeepCopyInto(out *BillingBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BillingBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetList.
func (in *BillingBudgetList) DeepCopy() *BillingBudgetList {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetSpec) DeepCopyInto(out *BillingBudgetSpec) {
	*out = *in
	in.BillingBudgetSpec2.DeepCopyInto(&out.BillingBudgetSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetSpec.
func (in *BillingBudgetSpec) DeepCopy() *BillingBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetSpec2) DeepCopyInto(out *BillingBudgetSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.AllUpdatesRule != nil {
		in, out := &in.AllUpdatesRule, &out.AllUpdatesRule
		*out = new(BillingBudgetSpecAllUpdatesRule)
		(*in).DeepCopyInto(*out)
	}
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(BillingBudgetSpecAmount)
		(*in).DeepCopyInto(*out)
	}
	if in.BillingAccount != nil {
		in, out := &in.BillingAccount, &out.BillingAccount
		*out = new(string)
		**out = **in
	}
	if in.BudgetFilter != nil {
		in, out := &in.BudgetFilter, &out.BudgetFilter
		*out = new(BillingBudgetSpecBudgetFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ThresholdRules != nil {
		in, out := &in.ThresholdRules, &out.ThresholdRules
		*out = make([]BillingBudgetSpecThresholdRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetSpec2.
func (in *BillingBudgetSpec2) DeepCopy() *BillingBudgetSpec2 {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetSpecAllUpdatesRule) DeepCopyInto(out *BillingBudgetSpecAllUpdatesRule) {
	*out = *in
	if in.DisableDefaultIamRecipients != nil {
		in, out := &in.DisableDefaultIamRecipients, &out.DisableDefaultIamRecipients
		*out = new(bool)
		**out = **in
	}
	if in.MonitoringNotificationChannels != nil {
		in, out := &in.MonitoringNotificationChannels, &out.MonitoringNotificationChannels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PubsubTopic != nil {
		in, out := &in.PubsubTopic, &out.PubsubTopic
		*out = new(string)
		**out = **in
	}
	if in.SchemaVersion != nil {
		in, out := &in.SchemaVersion, &out.SchemaVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetSpecAllUpdatesRule.
func (in *BillingBudgetSpecAllUpdatesRule) DeepCopy() *BillingBudgetSpecAllUpdatesRule {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetSpecAllUpdatesRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetSpecAmount) DeepCopyInto(out *BillingBudgetSpecAmount) {
	*out = *in
	if in.LastPeriodAmount != nil {
		in, out := &in.LastPeriodAmount, &out.LastPeriodAmount
		*out = new(bool)
		**out = **in
	}
	if in.SpecifiedAmount != nil {
		in, out := &in.SpecifiedAmount, &out.SpecifiedAmount
		*out = new(BillingBudgetSpecAmountSpecifiedAmount)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetSpecAmount.
func (in *BillingBudgetSpecAmount) DeepCopy() *BillingBudgetSpecAmount {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetSpecAmount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetSpecAmountSpecifiedAmount) DeepCopyInto(out *BillingBudgetSpecAmountSpecifiedAmount) {
	*out = *in
	if in.CurrencyCode != nil {
		in, out := &in.CurrencyCode, &out.CurrencyCode
		*out = new(string)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(int64)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetSpecAmountSpecifiedAmount.
func (in *BillingBudgetSpecAmountSpecifiedAmount) DeepCopy() *BillingBudgetSpecAmountSpecifiedAmount {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetSpecAmountSpecifiedAmount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetSpecBudgetFilter) DeepCopyInto(out *BillingBudgetSpecBudgetFilter) {
	*out = *in
	if in.CreditTypes != nil {
		in, out := &in.CreditTypes, &out.CreditTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreditTypesTreatment != nil {
		in, out := &in.CreditTypesTreatment, &out.CreditTypesTreatment
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Projects != nil {
		in, out := &in.Projects, &out.Projects
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subaccounts != nil {
		in, out := &in.Subaccounts, &out.Subaccounts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetSpecBudgetFilter.
func (in *BillingBudgetSpecBudgetFilter) DeepCopy() *BillingBudgetSpecBudgetFilter {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetSpecBudgetFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetSpecThresholdRules) DeepCopyInto(out *BillingBudgetSpecThresholdRules) {
	*out = *in
	if in.SpendBasis != nil {
		in, out := &in.SpendBasis, &out.SpendBasis
		*out = new(string)
		**out = **in
	}
	if in.ThresholdPercent != nil {
		in, out := &in.ThresholdPercent, &out.ThresholdPercent
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetSpecThresholdRules.
func (in *BillingBudgetSpecThresholdRules) DeepCopy() *BillingBudgetSpecThresholdRules {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetSpecThresholdRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetStatus) DeepCopyInto(out *BillingBudgetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetStatus.
func (in *BillingBudgetStatus) DeepCopy() *BillingBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetStatus)
	in.DeepCopyInto(out)
	return out
}

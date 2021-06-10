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
func (in *DeviceGroup) DeepCopyInto(out *DeviceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGroup.
func (in *DeviceGroup) DeepCopy() *DeviceGroup {
	if in == nil {
		return nil
	}
	out := new(DeviceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceGroupList) DeepCopyInto(out *DeviceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeviceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGroupList.
func (in *DeviceGroupList) DeepCopy() *DeviceGroupList {
	if in == nil {
		return nil
	}
	out := new(DeviceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceGroupSpec) DeepCopyInto(out *DeviceGroupSpec) {
	*out = *in
	in.DeviceGroupSpec2.DeepCopyInto(&out.DeviceGroupSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGroupSpec.
func (in *DeviceGroupSpec) DeepCopy() *DeviceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceGroupSpec2) DeepCopyInto(out *DeviceGroupSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.AllowRule != nil {
		in, out := &in.AllowRule, &out.AllowRule
		*out = new(DeviceGroupSpecAllowRule)
		(*in).DeepCopyInto(*out)
	}
	if in.IothubID != nil {
		in, out := &in.IothubID, &out.IothubID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RangeRule != nil {
		in, out := &in.RangeRule, &out.RangeRule
		*out = make([]DeviceGroupSpecRangeRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGroupSpec2.
func (in *DeviceGroupSpec2) DeepCopy() *DeviceGroupSpec2 {
	if in == nil {
		return nil
	}
	out := new(DeviceGroupSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceGroupSpecAllowRule) DeepCopyInto(out *DeviceGroupSpecAllowRule) {
	*out = *in
	if in.ConnectionToIPNotAllowed != nil {
		in, out := &in.ConnectionToIPNotAllowed, &out.ConnectionToIPNotAllowed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LocalUserNotAllowed != nil {
		in, out := &in.LocalUserNotAllowed, &out.LocalUserNotAllowed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProcessNotAllowed != nil {
		in, out := &in.ProcessNotAllowed, &out.ProcessNotAllowed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGroupSpecAllowRule.
func (in *DeviceGroupSpecAllowRule) DeepCopy() *DeviceGroupSpecAllowRule {
	if in == nil {
		return nil
	}
	out := new(DeviceGroupSpecAllowRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceGroupSpecRangeRule) DeepCopyInto(out *DeviceGroupSpecRangeRule) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGroupSpecRangeRule.
func (in *DeviceGroupSpecRangeRule) DeepCopy() *DeviceGroupSpecRangeRule {
	if in == nil {
		return nil
	}
	out := new(DeviceGroupSpecRangeRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceGroupStatus) DeepCopyInto(out *DeviceGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGroupStatus.
func (in *DeviceGroupStatus) DeepCopy() *DeviceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Solution) DeepCopyInto(out *Solution) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Solution.
func (in *Solution) DeepCopy() *Solution {
	if in == nil {
		return nil
	}
	out := new(Solution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Solution) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionList) DeepCopyInto(out *SolutionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Solution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionList.
func (in *SolutionList) DeepCopy() *SolutionList {
	if in == nil {
		return nil
	}
	out := new(SolutionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolutionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionSpec) DeepCopyInto(out *SolutionSpec) {
	*out = *in
	in.SolutionSpec2.DeepCopyInto(&out.SolutionSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionSpec.
func (in *SolutionSpec) DeepCopy() *SolutionSpec {
	if in == nil {
		return nil
	}
	out := new(SolutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionSpec2) DeepCopyInto(out *SolutionSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventsToExport != nil {
		in, out := &in.EventsToExport, &out.EventsToExport
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IothubIDS != nil {
		in, out := &in.IothubIDS, &out.IothubIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsWorkspaceID != nil {
		in, out := &in.LogAnalyticsWorkspaceID, &out.LogAnalyticsWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.LogUnmaskedIPSEnabled != nil {
		in, out := &in.LogUnmaskedIPSEnabled, &out.LogUnmaskedIPSEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.QueryForResources != nil {
		in, out := &in.QueryForResources, &out.QueryForResources
		*out = new(string)
		**out = **in
	}
	if in.QuerySubscriptionIDS != nil {
		in, out := &in.QuerySubscriptionIDS, &out.QuerySubscriptionIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RecommendationsEnabled != nil {
		in, out := &in.RecommendationsEnabled, &out.RecommendationsEnabled
		*out = new(SolutionSpecRecommendationsEnabled)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionSpec2.
func (in *SolutionSpec2) DeepCopy() *SolutionSpec2 {
	if in == nil {
		return nil
	}
	out := new(SolutionSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionSpecRecommendationsEnabled) DeepCopyInto(out *SolutionSpecRecommendationsEnabled) {
	*out = *in
	if in.AcrAuthentication != nil {
		in, out := &in.AcrAuthentication, &out.AcrAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.AgentSendUnutilizedMsg != nil {
		in, out := &in.AgentSendUnutilizedMsg, &out.AgentSendUnutilizedMsg
		*out = new(bool)
		**out = **in
	}
	if in.Baseline != nil {
		in, out := &in.Baseline, &out.Baseline
		*out = new(bool)
		**out = **in
	}
	if in.EdgeHubMemOptimize != nil {
		in, out := &in.EdgeHubMemOptimize, &out.EdgeHubMemOptimize
		*out = new(bool)
		**out = **in
	}
	if in.EdgeLoggingOption != nil {
		in, out := &in.EdgeLoggingOption, &out.EdgeLoggingOption
		*out = new(bool)
		**out = **in
	}
	if in.InconsistentModuleSettings != nil {
		in, out := &in.InconsistentModuleSettings, &out.InconsistentModuleSettings
		*out = new(bool)
		**out = **in
	}
	if in.InstallAgent != nil {
		in, out := &in.InstallAgent, &out.InstallAgent
		*out = new(bool)
		**out = **in
	}
	if in.IpFilterDenyAll != nil {
		in, out := &in.IpFilterDenyAll, &out.IpFilterDenyAll
		*out = new(bool)
		**out = **in
	}
	if in.IpFilterPermissiveRule != nil {
		in, out := &in.IpFilterPermissiveRule, &out.IpFilterPermissiveRule
		*out = new(bool)
		**out = **in
	}
	if in.OpenPorts != nil {
		in, out := &in.OpenPorts, &out.OpenPorts
		*out = new(bool)
		**out = **in
	}
	if in.PermissiveFirewallPolicy != nil {
		in, out := &in.PermissiveFirewallPolicy, &out.PermissiveFirewallPolicy
		*out = new(bool)
		**out = **in
	}
	if in.PermissiveInputFirewallRules != nil {
		in, out := &in.PermissiveInputFirewallRules, &out.PermissiveInputFirewallRules
		*out = new(bool)
		**out = **in
	}
	if in.PermissiveOutputFirewallRules != nil {
		in, out := &in.PermissiveOutputFirewallRules, &out.PermissiveOutputFirewallRules
		*out = new(bool)
		**out = **in
	}
	if in.PrivilegedDockerOptions != nil {
		in, out := &in.PrivilegedDockerOptions, &out.PrivilegedDockerOptions
		*out = new(bool)
		**out = **in
	}
	if in.SharedCredentials != nil {
		in, out := &in.SharedCredentials, &out.SharedCredentials
		*out = new(bool)
		**out = **in
	}
	if in.VulnerableTLSCipherSuite != nil {
		in, out := &in.VulnerableTLSCipherSuite, &out.VulnerableTLSCipherSuite
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionSpecRecommendationsEnabled.
func (in *SolutionSpecRecommendationsEnabled) DeepCopy() *SolutionSpecRecommendationsEnabled {
	if in == nil {
		return nil
	}
	out := new(SolutionSpecRecommendationsEnabled)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionStatus) DeepCopyInto(out *SolutionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionStatus.
func (in *SolutionStatus) DeepCopy() *SolutionStatus {
	if in == nil {
		return nil
	}
	out := new(SolutionStatus)
	in.DeepCopyInto(out)
	return out
}

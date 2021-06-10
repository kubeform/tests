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
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *App) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppList) DeepCopyInto(out *AppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]App, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppList.
func (in *AppList) DeepCopy() *AppList {
	if in == nil {
		return nil
	}
	out := new(AppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	in.AppSpec2.DeepCopyInto(&out.AppSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec.
func (in *AppSpec) DeepCopy() *AppSpec {
	if in == nil {
		return nil
	}
	out := new(AppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec2) DeepCopyInto(out *AppSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ComputePlatform != nil {
		in, out := &in.ComputePlatform, &out.ComputePlatform
		*out = new(string)
		**out = **in
	}
	if in.GithubAccountName != nil {
		in, out := &in.GithubAccountName, &out.GithubAccountName
		*out = new(string)
		**out = **in
	}
	if in.LinkedToGithub != nil {
		in, out := &in.LinkedToGithub, &out.LinkedToGithub
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec2.
func (in *AppSpec2) DeepCopy() *AppSpec2 {
	if in == nil {
		return nil
	}
	out := new(AppSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfig) DeepCopyInto(out *DeploymentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfig.
func (in *DeploymentConfig) DeepCopy() *DeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigList) DeepCopyInto(out *DeploymentConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigList.
func (in *DeploymentConfigList) DeepCopy() *DeploymentConfigList {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigSpec) DeepCopyInto(out *DeploymentConfigSpec) {
	*out = *in
	in.DeploymentConfigSpec2.DeepCopyInto(&out.DeploymentConfigSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigSpec.
func (in *DeploymentConfigSpec) DeepCopy() *DeploymentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigSpec2) DeepCopyInto(out *DeploymentConfigSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.ComputePlatform != nil {
		in, out := &in.ComputePlatform, &out.ComputePlatform
		*out = new(string)
		**out = **in
	}
	if in.DeploymentConfigID != nil {
		in, out := &in.DeploymentConfigID, &out.DeploymentConfigID
		*out = new(string)
		**out = **in
	}
	if in.DeploymentConfigName != nil {
		in, out := &in.DeploymentConfigName, &out.DeploymentConfigName
		*out = new(string)
		**out = **in
	}
	if in.MinimumHealthyHosts != nil {
		in, out := &in.MinimumHealthyHosts, &out.MinimumHealthyHosts
		*out = new(DeploymentConfigSpecMinimumHealthyHosts)
		(*in).DeepCopyInto(*out)
	}
	if in.TrafficRoutingConfig != nil {
		in, out := &in.TrafficRoutingConfig, &out.TrafficRoutingConfig
		*out = new(DeploymentConfigSpecTrafficRoutingConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigSpec2.
func (in *DeploymentConfigSpec2) DeepCopy() *DeploymentConfigSpec2 {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigSpecMinimumHealthyHosts) DeepCopyInto(out *DeploymentConfigSpecMinimumHealthyHosts) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigSpecMinimumHealthyHosts.
func (in *DeploymentConfigSpecMinimumHealthyHosts) DeepCopy() *DeploymentConfigSpecMinimumHealthyHosts {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigSpecMinimumHealthyHosts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigSpecTrafficRoutingConfig) DeepCopyInto(out *DeploymentConfigSpecTrafficRoutingConfig) {
	*out = *in
	if in.TimeBasedCanary != nil {
		in, out := &in.TimeBasedCanary, &out.TimeBasedCanary
		*out = new(DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeBasedLinear != nil {
		in, out := &in.TimeBasedLinear, &out.TimeBasedLinear
		*out = new(DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigSpecTrafficRoutingConfig.
func (in *DeploymentConfigSpecTrafficRoutingConfig) DeepCopy() *DeploymentConfigSpecTrafficRoutingConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigSpecTrafficRoutingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary) DeepCopyInto(out *DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary.
func (in *DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary) DeepCopy() *DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear) DeepCopyInto(out *DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear.
func (in *DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear) DeepCopy() *DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfigStatus) DeepCopyInto(out *DeploymentConfigStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfigStatus.
func (in *DeploymentConfigStatus) DeepCopy() *DeploymentConfigStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroup) DeepCopyInto(out *DeploymentGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroup.
func (in *DeploymentGroup) DeepCopy() *DeploymentGroup {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupList) DeepCopyInto(out *DeploymentGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupList.
func (in *DeploymentGroupList) DeepCopy() *DeploymentGroupList {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpec) DeepCopyInto(out *DeploymentGroupSpec) {
	*out = *in
	in.DeploymentGroupSpec2.DeepCopyInto(&out.DeploymentGroupSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpec.
func (in *DeploymentGroupSpec) DeepCopy() *DeploymentGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpec2) DeepCopyInto(out *DeploymentGroupSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.AlarmConfiguration != nil {
		in, out := &in.AlarmConfiguration, &out.AlarmConfiguration
		*out = new(DeploymentGroupSpecAlarmConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AppName != nil {
		in, out := &in.AppName, &out.AppName
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoRollbackConfiguration != nil {
		in, out := &in.AutoRollbackConfiguration, &out.AutoRollbackConfiguration
		*out = new(DeploymentGroupSpecAutoRollbackConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoscalingGroups != nil {
		in, out := &in.AutoscalingGroups, &out.AutoscalingGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BlueGreenDeploymentConfig != nil {
		in, out := &in.BlueGreenDeploymentConfig, &out.BlueGreenDeploymentConfig
		*out = new(DeploymentGroupSpecBlueGreenDeploymentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ComputePlatform != nil {
		in, out := &in.ComputePlatform, &out.ComputePlatform
		*out = new(string)
		**out = **in
	}
	if in.DeploymentConfigName != nil {
		in, out := &in.DeploymentConfigName, &out.DeploymentConfigName
		*out = new(string)
		**out = **in
	}
	if in.DeploymentGroupID != nil {
		in, out := &in.DeploymentGroupID, &out.DeploymentGroupID
		*out = new(string)
		**out = **in
	}
	if in.DeploymentGroupName != nil {
		in, out := &in.DeploymentGroupName, &out.DeploymentGroupName
		*out = new(string)
		**out = **in
	}
	if in.DeploymentStyle != nil {
		in, out := &in.DeploymentStyle, &out.DeploymentStyle
		*out = new(DeploymentGroupSpecDeploymentStyle)
		(*in).DeepCopyInto(*out)
	}
	if in.Ec2TagFilter != nil {
		in, out := &in.Ec2TagFilter, &out.Ec2TagFilter
		*out = make([]DeploymentGroupSpecEc2TagFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ec2TagSet != nil {
		in, out := &in.Ec2TagSet, &out.Ec2TagSet
		*out = make([]DeploymentGroupSpecEc2TagSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EcsService != nil {
		in, out := &in.EcsService, &out.EcsService
		*out = new(DeploymentGroupSpecEcsService)
		(*in).DeepCopyInto(*out)
	}
	if in.LoadBalancerInfo != nil {
		in, out := &in.LoadBalancerInfo, &out.LoadBalancerInfo
		*out = new(DeploymentGroupSpecLoadBalancerInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.OnPremisesInstanceTagFilter != nil {
		in, out := &in.OnPremisesInstanceTagFilter, &out.OnPremisesInstanceTagFilter
		*out = make([]DeploymentGroupSpecOnPremisesInstanceTagFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceRoleArn != nil {
		in, out := &in.ServiceRoleArn, &out.ServiceRoleArn
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
	if in.TriggerConfiguration != nil {
		in, out := &in.TriggerConfiguration, &out.TriggerConfiguration
		*out = make([]DeploymentGroupSpecTriggerConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpec2.
func (in *DeploymentGroupSpec2) DeepCopy() *DeploymentGroupSpec2 {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecAlarmConfiguration) DeepCopyInto(out *DeploymentGroupSpecAlarmConfiguration) {
	*out = *in
	if in.Alarms != nil {
		in, out := &in.Alarms, &out.Alarms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IgnorePollAlarmFailure != nil {
		in, out := &in.IgnorePollAlarmFailure, &out.IgnorePollAlarmFailure
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecAlarmConfiguration.
func (in *DeploymentGroupSpecAlarmConfiguration) DeepCopy() *DeploymentGroupSpecAlarmConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecAlarmConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecAutoRollbackConfiguration) DeepCopyInto(out *DeploymentGroupSpecAutoRollbackConfiguration) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecAutoRollbackConfiguration.
func (in *DeploymentGroupSpecAutoRollbackConfiguration) DeepCopy() *DeploymentGroupSpecAutoRollbackConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecAutoRollbackConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfig) DeepCopyInto(out *DeploymentGroupSpecBlueGreenDeploymentConfig) {
	*out = *in
	if in.DeploymentReadyOption != nil {
		in, out := &in.DeploymentReadyOption, &out.DeploymentReadyOption
		*out = new(DeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption)
		(*in).DeepCopyInto(*out)
	}
	if in.GreenFleetProvisioningOption != nil {
		in, out := &in.GreenFleetProvisioningOption, &out.GreenFleetProvisioningOption
		*out = new(DeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption)
		(*in).DeepCopyInto(*out)
	}
	if in.TerminateBlueInstancesOnDeploymentSuccess != nil {
		in, out := &in.TerminateBlueInstancesOnDeploymentSuccess, &out.TerminateBlueInstancesOnDeploymentSuccess
		*out = new(DeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecBlueGreenDeploymentConfig.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfig) DeepCopy() *DeploymentGroupSpecBlueGreenDeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecBlueGreenDeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption) DeepCopyInto(out *DeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption) {
	*out = *in
	if in.ActionOnTimeout != nil {
		in, out := &in.ActionOnTimeout, &out.ActionOnTimeout
		*out = new(string)
		**out = **in
	}
	if in.WaitTimeInMinutes != nil {
		in, out := &in.WaitTimeInMinutes, &out.WaitTimeInMinutes
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption) DeepCopy() *DeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption) DeepCopyInto(out *DeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption) DeepCopy() *DeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess) DeepCopyInto(out *DeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.TerminationWaitTimeInMinutes != nil {
		in, out := &in.TerminationWaitTimeInMinutes, &out.TerminationWaitTimeInMinutes
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.
func (in *DeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess) DeepCopy() *DeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecDeploymentStyle) DeepCopyInto(out *DeploymentGroupSpecDeploymentStyle) {
	*out = *in
	if in.DeploymentOption != nil {
		in, out := &in.DeploymentOption, &out.DeploymentOption
		*out = new(string)
		**out = **in
	}
	if in.DeploymentType != nil {
		in, out := &in.DeploymentType, &out.DeploymentType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecDeploymentStyle.
func (in *DeploymentGroupSpecDeploymentStyle) DeepCopy() *DeploymentGroupSpecDeploymentStyle {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecDeploymentStyle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecEc2TagFilter) DeepCopyInto(out *DeploymentGroupSpecEc2TagFilter) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecEc2TagFilter.
func (in *DeploymentGroupSpecEc2TagFilter) DeepCopy() *DeploymentGroupSpecEc2TagFilter {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecEc2TagFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecEc2TagSet) DeepCopyInto(out *DeploymentGroupSpecEc2TagSet) {
	*out = *in
	if in.Ec2TagFilter != nil {
		in, out := &in.Ec2TagFilter, &out.Ec2TagFilter
		*out = make([]DeploymentGroupSpecEc2TagSetEc2TagFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecEc2TagSet.
func (in *DeploymentGroupSpecEc2TagSet) DeepCopy() *DeploymentGroupSpecEc2TagSet {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecEc2TagSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecEc2TagSetEc2TagFilter) DeepCopyInto(out *DeploymentGroupSpecEc2TagSetEc2TagFilter) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecEc2TagSetEc2TagFilter.
func (in *DeploymentGroupSpecEc2TagSetEc2TagFilter) DeepCopy() *DeploymentGroupSpecEc2TagSetEc2TagFilter {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecEc2TagSetEc2TagFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecEcsService) DeepCopyInto(out *DeploymentGroupSpecEcsService) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecEcsService.
func (in *DeploymentGroupSpecEcsService) DeepCopy() *DeploymentGroupSpecEcsService {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecEcsService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecLoadBalancerInfo) DeepCopyInto(out *DeploymentGroupSpecLoadBalancerInfo) {
	*out = *in
	if in.ElbInfo != nil {
		in, out := &in.ElbInfo, &out.ElbInfo
		*out = make([]DeploymentGroupSpecLoadBalancerInfoElbInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetGroupInfo != nil {
		in, out := &in.TargetGroupInfo, &out.TargetGroupInfo
		*out = make([]DeploymentGroupSpecLoadBalancerInfoTargetGroupInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetGroupPairInfo != nil {
		in, out := &in.TargetGroupPairInfo, &out.TargetGroupPairInfo
		*out = new(DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecLoadBalancerInfo.
func (in *DeploymentGroupSpecLoadBalancerInfo) DeepCopy() *DeploymentGroupSpecLoadBalancerInfo {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecLoadBalancerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecLoadBalancerInfoElbInfo) DeepCopyInto(out *DeploymentGroupSpecLoadBalancerInfoElbInfo) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecLoadBalancerInfoElbInfo.
func (in *DeploymentGroupSpecLoadBalancerInfoElbInfo) DeepCopy() *DeploymentGroupSpecLoadBalancerInfoElbInfo {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecLoadBalancerInfoElbInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupInfo) DeepCopyInto(out *DeploymentGroupSpecLoadBalancerInfoTargetGroupInfo) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecLoadBalancerInfoTargetGroupInfo.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupInfo) DeepCopy() *DeploymentGroupSpecLoadBalancerInfoTargetGroupInfo {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecLoadBalancerInfoTargetGroupInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo) DeepCopyInto(out *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo) {
	*out = *in
	if in.ProdTrafficRoute != nil {
		in, out := &in.ProdTrafficRoute, &out.ProdTrafficRoute
		*out = new(DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetGroup != nil {
		in, out := &in.TargetGroup, &out.TargetGroup
		*out = make([]DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TestTrafficRoute != nil {
		in, out := &in.TestTrafficRoute, &out.TestTrafficRoute
		*out = new(DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo) DeepCopy() *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute) DeepCopyInto(out *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute) {
	*out = *in
	if in.ListenerArns != nil {
		in, out := &in.ListenerArns, &out.ListenerArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute) DeepCopy() *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup) DeepCopyInto(out *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup) DeepCopy() *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute) DeepCopyInto(out *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute) {
	*out = *in
	if in.ListenerArns != nil {
		in, out := &in.ListenerArns, &out.ListenerArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.
func (in *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute) DeepCopy() *DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecOnPremisesInstanceTagFilter) DeepCopyInto(out *DeploymentGroupSpecOnPremisesInstanceTagFilter) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecOnPremisesInstanceTagFilter.
func (in *DeploymentGroupSpecOnPremisesInstanceTagFilter) DeepCopy() *DeploymentGroupSpecOnPremisesInstanceTagFilter {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecOnPremisesInstanceTagFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupSpecTriggerConfiguration) DeepCopyInto(out *DeploymentGroupSpecTriggerConfiguration) {
	*out = *in
	if in.TriggerEvents != nil {
		in, out := &in.TriggerEvents, &out.TriggerEvents
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TriggerName != nil {
		in, out := &in.TriggerName, &out.TriggerName
		*out = new(string)
		**out = **in
	}
	if in.TriggerTargetArn != nil {
		in, out := &in.TriggerTargetArn, &out.TriggerTargetArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupSpecTriggerConfiguration.
func (in *DeploymentGroupSpecTriggerConfiguration) DeepCopy() *DeploymentGroupSpecTriggerConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupSpecTriggerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentGroupStatus) DeepCopyInto(out *DeploymentGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentGroupStatus.
func (in *DeploymentGroupStatus) DeepCopy() *DeploymentGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentGroupStatus)
	in.DeepCopyInto(out)
	return out
}
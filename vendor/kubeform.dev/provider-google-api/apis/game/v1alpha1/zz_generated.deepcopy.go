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
func (in *ServicesGameServerCluster) DeepCopyInto(out *ServicesGameServerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerCluster.
func (in *ServicesGameServerCluster) DeepCopy() *ServicesGameServerCluster {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerClusterList) DeepCopyInto(out *ServicesGameServerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicesGameServerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerClusterList.
func (in *ServicesGameServerClusterList) DeepCopy() *ServicesGameServerClusterList {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerClusterSpec) DeepCopyInto(out *ServicesGameServerClusterSpec) {
	*out = *in
	in.ServicesGameServerClusterSpec2.DeepCopyInto(&out.ServicesGameServerClusterSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerClusterSpec.
func (in *ServicesGameServerClusterSpec) DeepCopy() *ServicesGameServerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerClusterSpec2) DeepCopyInto(out *ServicesGameServerClusterSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ConnectionInfo != nil {
		in, out := &in.ConnectionInfo, &out.ConnectionInfo
		*out = new(ServicesGameServerClusterSpecConnectionInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerClusterSpec2.
func (in *ServicesGameServerClusterSpec2) DeepCopy() *ServicesGameServerClusterSpec2 {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerClusterSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerClusterSpecConnectionInfo) DeepCopyInto(out *ServicesGameServerClusterSpecConnectionInfo) {
	*out = *in
	if in.GkeClusterReference != nil {
		in, out := &in.GkeClusterReference, &out.GkeClusterReference
		*out = new(ServicesGameServerClusterSpecConnectionInfoGkeClusterReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerClusterSpecConnectionInfo.
func (in *ServicesGameServerClusterSpecConnectionInfo) DeepCopy() *ServicesGameServerClusterSpecConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerClusterSpecConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerClusterSpecConnectionInfoGkeClusterReference) DeepCopyInto(out *ServicesGameServerClusterSpecConnectionInfoGkeClusterReference) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerClusterSpecConnectionInfoGkeClusterReference.
func (in *ServicesGameServerClusterSpecConnectionInfoGkeClusterReference) DeepCopy() *ServicesGameServerClusterSpecConnectionInfoGkeClusterReference {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerClusterSpecConnectionInfoGkeClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerClusterStatus) DeepCopyInto(out *ServicesGameServerClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerClusterStatus.
func (in *ServicesGameServerClusterStatus) DeepCopy() *ServicesGameServerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfig) DeepCopyInto(out *ServicesGameServerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfig.
func (in *ServicesGameServerConfig) DeepCopy() *ServicesGameServerConfig {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigList) DeepCopyInto(out *ServicesGameServerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicesGameServerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigList.
func (in *ServicesGameServerConfigList) DeepCopy() *ServicesGameServerConfigList {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigSpec) DeepCopyInto(out *ServicesGameServerConfigSpec) {
	*out = *in
	in.ServicesGameServerConfigSpec2.DeepCopyInto(&out.ServicesGameServerConfigSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigSpec.
func (in *ServicesGameServerConfigSpec) DeepCopy() *ServicesGameServerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigSpec2) DeepCopyInto(out *ServicesGameServerConfigSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.ConfigID != nil {
		in, out := &in.ConfigID, &out.ConfigID
		*out = new(string)
		**out = **in
	}
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FleetConfigs != nil {
		in, out := &in.FleetConfigs, &out.FleetConfigs
		*out = make([]ServicesGameServerConfigSpecFleetConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ScalingConfigs != nil {
		in, out := &in.ScalingConfigs, &out.ScalingConfigs
		*out = make([]ServicesGameServerConfigSpecScalingConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigSpec2.
func (in *ServicesGameServerConfigSpec2) DeepCopy() *ServicesGameServerConfigSpec2 {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigSpecFleetConfigs) DeepCopyInto(out *ServicesGameServerConfigSpecFleetConfigs) {
	*out = *in
	if in.FleetSpec != nil {
		in, out := &in.FleetSpec, &out.FleetSpec
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigSpecFleetConfigs.
func (in *ServicesGameServerConfigSpecFleetConfigs) DeepCopy() *ServicesGameServerConfigSpecFleetConfigs {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigSpecFleetConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigSpecScalingConfigs) DeepCopyInto(out *ServicesGameServerConfigSpecScalingConfigs) {
	*out = *in
	if in.FleetAutoscalerSpec != nil {
		in, out := &in.FleetAutoscalerSpec, &out.FleetAutoscalerSpec
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Schedules != nil {
		in, out := &in.Schedules, &out.Schedules
		*out = make([]ServicesGameServerConfigSpecScalingConfigsSchedules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]ServicesGameServerConfigSpecScalingConfigsSelectors, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigSpecScalingConfigs.
func (in *ServicesGameServerConfigSpecScalingConfigs) DeepCopy() *ServicesGameServerConfigSpecScalingConfigs {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigSpecScalingConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigSpecScalingConfigsSchedules) DeepCopyInto(out *ServicesGameServerConfigSpecScalingConfigsSchedules) {
	*out = *in
	if in.CronJobDuration != nil {
		in, out := &in.CronJobDuration, &out.CronJobDuration
		*out = new(string)
		**out = **in
	}
	if in.CronSpec != nil {
		in, out := &in.CronSpec, &out.CronSpec
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigSpecScalingConfigsSchedules.
func (in *ServicesGameServerConfigSpecScalingConfigsSchedules) DeepCopy() *ServicesGameServerConfigSpecScalingConfigsSchedules {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigSpecScalingConfigsSchedules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigSpecScalingConfigsSelectors) DeepCopyInto(out *ServicesGameServerConfigSpecScalingConfigsSelectors) {
	*out = *in
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigSpecScalingConfigsSelectors.
func (in *ServicesGameServerConfigSpecScalingConfigsSelectors) DeepCopy() *ServicesGameServerConfigSpecScalingConfigsSelectors {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigSpecScalingConfigsSelectors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerConfigStatus) DeepCopyInto(out *ServicesGameServerConfigStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerConfigStatus.
func (in *ServicesGameServerConfigStatus) DeepCopy() *ServicesGameServerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeployment) DeepCopyInto(out *ServicesGameServerDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeployment.
func (in *ServicesGameServerDeployment) DeepCopy() *ServicesGameServerDeployment {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentList) DeepCopyInto(out *ServicesGameServerDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicesGameServerDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentList.
func (in *ServicesGameServerDeploymentList) DeepCopy() *ServicesGameServerDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentRollout) DeepCopyInto(out *ServicesGameServerDeploymentRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentRollout.
func (in *ServicesGameServerDeploymentRollout) DeepCopy() *ServicesGameServerDeploymentRollout {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerDeploymentRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentRolloutList) DeepCopyInto(out *ServicesGameServerDeploymentRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicesGameServerDeploymentRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentRolloutList.
func (in *ServicesGameServerDeploymentRolloutList) DeepCopy() *ServicesGameServerDeploymentRolloutList {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesGameServerDeploymentRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentRolloutSpec) DeepCopyInto(out *ServicesGameServerDeploymentRolloutSpec) {
	*out = *in
	in.ServicesGameServerDeploymentRolloutSpec2.DeepCopyInto(&out.ServicesGameServerDeploymentRolloutSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentRolloutSpec.
func (in *ServicesGameServerDeploymentRolloutSpec) DeepCopy() *ServicesGameServerDeploymentRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentRolloutSpec2) DeepCopyInto(out *ServicesGameServerDeploymentRolloutSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.DefaultGameServerConfig != nil {
		in, out := &in.DefaultGameServerConfig, &out.DefaultGameServerConfig
		*out = new(string)
		**out = **in
	}
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.GameServerConfigOverrides != nil {
		in, out := &in.GameServerConfigOverrides, &out.GameServerConfigOverrides
		*out = make([]ServicesGameServerDeploymentRolloutSpecGameServerConfigOverrides, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentRolloutSpec2.
func (in *ServicesGameServerDeploymentRolloutSpec2) DeepCopy() *ServicesGameServerDeploymentRolloutSpec2 {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentRolloutSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverrides) DeepCopyInto(out *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverrides) {
	*out = *in
	if in.ConfigVersion != nil {
		in, out := &in.ConfigVersion, &out.ConfigVersion
		*out = new(string)
		**out = **in
	}
	if in.RealmsSelector != nil {
		in, out := &in.RealmsSelector, &out.RealmsSelector
		*out = new(ServicesGameServerDeploymentRolloutSpecGameServerConfigOverridesRealmsSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentRolloutSpecGameServerConfigOverrides.
func (in *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverrides) DeepCopy() *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverrides {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentRolloutSpecGameServerConfigOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverridesRealmsSelector) DeepCopyInto(out *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverridesRealmsSelector) {
	*out = *in
	if in.Realms != nil {
		in, out := &in.Realms, &out.Realms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentRolloutSpecGameServerConfigOverridesRealmsSelector.
func (in *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverridesRealmsSelector) DeepCopy() *ServicesGameServerDeploymentRolloutSpecGameServerConfigOverridesRealmsSelector {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentRolloutSpecGameServerConfigOverridesRealmsSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentRolloutStatus) DeepCopyInto(out *ServicesGameServerDeploymentRolloutStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentRolloutStatus.
func (in *ServicesGameServerDeploymentRolloutStatus) DeepCopy() *ServicesGameServerDeploymentRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentSpec) DeepCopyInto(out *ServicesGameServerDeploymentSpec) {
	*out = *in
	in.ServicesGameServerDeploymentSpec2.DeepCopyInto(&out.ServicesGameServerDeploymentSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentSpec.
func (in *ServicesGameServerDeploymentSpec) DeepCopy() *ServicesGameServerDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentSpec2) DeepCopyInto(out *ServicesGameServerDeploymentSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentSpec2.
func (in *ServicesGameServerDeploymentSpec2) DeepCopy() *ServicesGameServerDeploymentSpec2 {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesGameServerDeploymentStatus) DeepCopyInto(out *ServicesGameServerDeploymentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesGameServerDeploymentStatus.
func (in *ServicesGameServerDeploymentStatus) DeepCopy() *ServicesGameServerDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesGameServerDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesRealm) DeepCopyInto(out *ServicesRealm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesRealm.
func (in *ServicesRealm) DeepCopy() *ServicesRealm {
	if in == nil {
		return nil
	}
	out := new(ServicesRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesRealm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesRealmList) DeepCopyInto(out *ServicesRealmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicesRealm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesRealmList.
func (in *ServicesRealmList) DeepCopy() *ServicesRealmList {
	if in == nil {
		return nil
	}
	out := new(ServicesRealmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicesRealmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesRealmSpec) DeepCopyInto(out *ServicesRealmSpec) {
	*out = *in
	in.ServicesRealmSpec2.DeepCopyInto(&out.ServicesRealmSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesRealmSpec.
func (in *ServicesRealmSpec) DeepCopy() *ServicesRealmSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesRealmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesRealmSpec2) DeepCopyInto(out *ServicesRealmSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
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
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesRealmSpec2.
func (in *ServicesRealmSpec2) DeepCopy() *ServicesRealmSpec2 {
	if in == nil {
		return nil
	}
	out := new(ServicesRealmSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesRealmStatus) DeepCopyInto(out *ServicesRealmStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesRealmStatus.
func (in *ServicesRealmStatus) DeepCopy() *ServicesRealmStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesRealmStatus)
	in.DeepCopyInto(out)
	return out
}

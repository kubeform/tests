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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	in.EnvironmentSpec2.DeepCopyInto(&out.EnvironmentSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec2) DeepCopyInto(out *EnvironmentSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.AirflowConfigurationOptions != nil {
		in, out := &in.AirflowConfigurationOptions, &out.AirflowConfigurationOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AirflowVersion != nil {
		in, out := &in.AirflowVersion, &out.AirflowVersion
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.DagS3Path != nil {
		in, out := &in.DagS3Path, &out.DagS3Path
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentClass != nil {
		in, out := &in.EnvironmentClass, &out.EnvironmentClass
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleArn != nil {
		in, out := &in.ExecutionRoleArn, &out.ExecutionRoleArn
		*out = new(string)
		**out = **in
	}
	if in.KmsKey != nil {
		in, out := &in.KmsKey, &out.KmsKey
		*out = new(string)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = make([]EnvironmentSpecLastUpdated, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = new(EnvironmentSpecLoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxWorkers != nil {
		in, out := &in.MaxWorkers, &out.MaxWorkers
		*out = new(int64)
		**out = **in
	}
	if in.MinWorkers != nil {
		in, out := &in.MinWorkers, &out.MinWorkers
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfiguration != nil {
		in, out := &in.NetworkConfiguration, &out.NetworkConfiguration
		*out = new(EnvironmentSpecNetworkConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.PluginsS3ObjectVersion != nil {
		in, out := &in.PluginsS3ObjectVersion, &out.PluginsS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.PluginsS3Path != nil {
		in, out := &in.PluginsS3Path, &out.PluginsS3Path
		*out = new(string)
		**out = **in
	}
	if in.RequirementsS3ObjectVersion != nil {
		in, out := &in.RequirementsS3ObjectVersion, &out.RequirementsS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.RequirementsS3Path != nil {
		in, out := &in.RequirementsS3Path, &out.RequirementsS3Path
		*out = new(string)
		**out = **in
	}
	if in.ServiceRoleArn != nil {
		in, out := &in.ServiceRoleArn, &out.ServiceRoleArn
		*out = new(string)
		**out = **in
	}
	if in.SourceBucketArn != nil {
		in, out := &in.SourceBucketArn, &out.SourceBucketArn
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
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
	if in.WebserverAccessMode != nil {
		in, out := &in.WebserverAccessMode, &out.WebserverAccessMode
		*out = new(string)
		**out = **in
	}
	if in.WebserverURL != nil {
		in, out := &in.WebserverURL, &out.WebserverURL
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindowStart != nil {
		in, out := &in.WeeklyMaintenanceWindowStart, &out.WeeklyMaintenanceWindowStart
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec2.
func (in *EnvironmentSpec2) DeepCopy() *EnvironmentSpec2 {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLastUpdated) DeepCopyInto(out *EnvironmentSpecLastUpdated) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = make([]EnvironmentSpecLastUpdatedError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLastUpdated.
func (in *EnvironmentSpecLastUpdated) DeepCopy() *EnvironmentSpecLastUpdated {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLastUpdated)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLastUpdatedError) DeepCopyInto(out *EnvironmentSpecLastUpdatedError) {
	*out = *in
	if in.ErrorCode != nil {
		in, out := &in.ErrorCode, &out.ErrorCode
		*out = new(string)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLastUpdatedError.
func (in *EnvironmentSpecLastUpdatedError) DeepCopy() *EnvironmentSpecLastUpdatedError {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLastUpdatedError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLoggingConfiguration) DeepCopyInto(out *EnvironmentSpecLoggingConfiguration) {
	*out = *in
	if in.DagProcessingLogs != nil {
		in, out := &in.DagProcessingLogs, &out.DagProcessingLogs
		*out = new(EnvironmentSpecLoggingConfigurationDagProcessingLogs)
		(*in).DeepCopyInto(*out)
	}
	if in.SchedulerLogs != nil {
		in, out := &in.SchedulerLogs, &out.SchedulerLogs
		*out = new(EnvironmentSpecLoggingConfigurationSchedulerLogs)
		(*in).DeepCopyInto(*out)
	}
	if in.TaskLogs != nil {
		in, out := &in.TaskLogs, &out.TaskLogs
		*out = new(EnvironmentSpecLoggingConfigurationTaskLogs)
		(*in).DeepCopyInto(*out)
	}
	if in.WebserverLogs != nil {
		in, out := &in.WebserverLogs, &out.WebserverLogs
		*out = new(EnvironmentSpecLoggingConfigurationWebserverLogs)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerLogs != nil {
		in, out := &in.WorkerLogs, &out.WorkerLogs
		*out = new(EnvironmentSpecLoggingConfigurationWorkerLogs)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLoggingConfiguration.
func (in *EnvironmentSpecLoggingConfiguration) DeepCopy() *EnvironmentSpecLoggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLoggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLoggingConfigurationDagProcessingLogs) DeepCopyInto(out *EnvironmentSpecLoggingConfigurationDagProcessingLogs) {
	*out = *in
	if in.CloudWatchLogGroupArn != nil {
		in, out := &in.CloudWatchLogGroupArn, &out.CloudWatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLoggingConfigurationDagProcessingLogs.
func (in *EnvironmentSpecLoggingConfigurationDagProcessingLogs) DeepCopy() *EnvironmentSpecLoggingConfigurationDagProcessingLogs {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLoggingConfigurationDagProcessingLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLoggingConfigurationSchedulerLogs) DeepCopyInto(out *EnvironmentSpecLoggingConfigurationSchedulerLogs) {
	*out = *in
	if in.CloudWatchLogGroupArn != nil {
		in, out := &in.CloudWatchLogGroupArn, &out.CloudWatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLoggingConfigurationSchedulerLogs.
func (in *EnvironmentSpecLoggingConfigurationSchedulerLogs) DeepCopy() *EnvironmentSpecLoggingConfigurationSchedulerLogs {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLoggingConfigurationSchedulerLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLoggingConfigurationTaskLogs) DeepCopyInto(out *EnvironmentSpecLoggingConfigurationTaskLogs) {
	*out = *in
	if in.CloudWatchLogGroupArn != nil {
		in, out := &in.CloudWatchLogGroupArn, &out.CloudWatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLoggingConfigurationTaskLogs.
func (in *EnvironmentSpecLoggingConfigurationTaskLogs) DeepCopy() *EnvironmentSpecLoggingConfigurationTaskLogs {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLoggingConfigurationTaskLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLoggingConfigurationWebserverLogs) DeepCopyInto(out *EnvironmentSpecLoggingConfigurationWebserverLogs) {
	*out = *in
	if in.CloudWatchLogGroupArn != nil {
		in, out := &in.CloudWatchLogGroupArn, &out.CloudWatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLoggingConfigurationWebserverLogs.
func (in *EnvironmentSpecLoggingConfigurationWebserverLogs) DeepCopy() *EnvironmentSpecLoggingConfigurationWebserverLogs {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLoggingConfigurationWebserverLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecLoggingConfigurationWorkerLogs) DeepCopyInto(out *EnvironmentSpecLoggingConfigurationWorkerLogs) {
	*out = *in
	if in.CloudWatchLogGroupArn != nil {
		in, out := &in.CloudWatchLogGroupArn, &out.CloudWatchLogGroupArn
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecLoggingConfigurationWorkerLogs.
func (in *EnvironmentSpecLoggingConfigurationWorkerLogs) DeepCopy() *EnvironmentSpecLoggingConfigurationWorkerLogs {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecLoggingConfigurationWorkerLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecNetworkConfiguration) DeepCopyInto(out *EnvironmentSpecNetworkConfiguration) {
	*out = *in
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecNetworkConfiguration.
func (in *EnvironmentSpecNetworkConfiguration) DeepCopy() *EnvironmentSpecNetworkConfiguration {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecNetworkConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}
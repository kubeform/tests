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

	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LustreFileSystem) DeepCopyInto(out *LustreFileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LustreFileSystem.
func (in *LustreFileSystem) DeepCopy() *LustreFileSystem {
	if in == nil {
		return nil
	}
	out := new(LustreFileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LustreFileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LustreFileSystemList) DeepCopyInto(out *LustreFileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LustreFileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LustreFileSystemList.
func (in *LustreFileSystemList) DeepCopy() *LustreFileSystemList {
	if in == nil {
		return nil
	}
	out := new(LustreFileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LustreFileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LustreFileSystemSpec) DeepCopyInto(out *LustreFileSystemSpec) {
	*out = *in
	in.LustreFileSystemSpec2.DeepCopyInto(&out.LustreFileSystemSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LustreFileSystemSpec.
func (in *LustreFileSystemSpec) DeepCopy() *LustreFileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(LustreFileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LustreFileSystemSpec2) DeepCopyInto(out *LustreFileSystemSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoImportPolicy != nil {
		in, out := &in.AutoImportPolicy, &out.AutoImportPolicy
		*out = new(string)
		**out = **in
	}
	if in.AutomaticBackupRetentionDays != nil {
		in, out := &in.AutomaticBackupRetentionDays, &out.AutomaticBackupRetentionDays
		*out = new(int64)
		**out = **in
	}
	if in.CopyTagsToBackups != nil {
		in, out := &in.CopyTagsToBackups, &out.CopyTagsToBackups
		*out = new(bool)
		**out = **in
	}
	if in.DailyAutomaticBackupStartTime != nil {
		in, out := &in.DailyAutomaticBackupStartTime, &out.DailyAutomaticBackupStartTime
		*out = new(string)
		**out = **in
	}
	if in.DeploymentType != nil {
		in, out := &in.DeploymentType, &out.DeploymentType
		*out = new(string)
		**out = **in
	}
	if in.DnsName != nil {
		in, out := &in.DnsName, &out.DnsName
		*out = new(string)
		**out = **in
	}
	if in.DriveCacheType != nil {
		in, out := &in.DriveCacheType, &out.DriveCacheType
		*out = new(string)
		**out = **in
	}
	if in.ExportPath != nil {
		in, out := &in.ExportPath, &out.ExportPath
		*out = new(string)
		**out = **in
	}
	if in.ImportPath != nil {
		in, out := &in.ImportPath, &out.ImportPath
		*out = new(string)
		**out = **in
	}
	if in.ImportedFileChunkSize != nil {
		in, out := &in.ImportedFileChunkSize, &out.ImportedFileChunkSize
		*out = new(int64)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.MountName != nil {
		in, out := &in.MountName, &out.MountName
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceIDS != nil {
		in, out := &in.NetworkInterfaceIDS, &out.NetworkInterfaceIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.PerUnitStorageThroughput != nil {
		in, out := &in.PerUnitStorageThroughput, &out.PerUnitStorageThroughput
		*out = new(int64)
		**out = **in
	}
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StorageCapacity != nil {
		in, out := &in.StorageCapacity, &out.StorageCapacity
		*out = new(int64)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceStartTime != nil {
		in, out := &in.WeeklyMaintenanceStartTime, &out.WeeklyMaintenanceStartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LustreFileSystemSpec2.
func (in *LustreFileSystemSpec2) DeepCopy() *LustreFileSystemSpec2 {
	if in == nil {
		return nil
	}
	out := new(LustreFileSystemSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LustreFileSystemStatus) DeepCopyInto(out *LustreFileSystemStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LustreFileSystemStatus.
func (in *LustreFileSystemStatus) DeepCopy() *LustreFileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(LustreFileSystemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsFileSystem) DeepCopyInto(out *WindowsFileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsFileSystem.
func (in *WindowsFileSystem) DeepCopy() *WindowsFileSystem {
	if in == nil {
		return nil
	}
	out := new(WindowsFileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WindowsFileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsFileSystemList) DeepCopyInto(out *WindowsFileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WindowsFileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsFileSystemList.
func (in *WindowsFileSystemList) DeepCopy() *WindowsFileSystemList {
	if in == nil {
		return nil
	}
	out := new(WindowsFileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WindowsFileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsFileSystemSpec) DeepCopyInto(out *WindowsFileSystemSpec) {
	*out = *in
	in.WindowsFileSystemSpec2.DeepCopyInto(&out.WindowsFileSystemSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsFileSystemSpec.
func (in *WindowsFileSystemSpec) DeepCopy() *WindowsFileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(WindowsFileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsFileSystemSpec2) DeepCopyInto(out *WindowsFileSystemSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.ActiveDirectoryID != nil {
		in, out := &in.ActiveDirectoryID, &out.ActiveDirectoryID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutomaticBackupRetentionDays != nil {
		in, out := &in.AutomaticBackupRetentionDays, &out.AutomaticBackupRetentionDays
		*out = new(int64)
		**out = **in
	}
	if in.CopyTagsToBackups != nil {
		in, out := &in.CopyTagsToBackups, &out.CopyTagsToBackups
		*out = new(bool)
		**out = **in
	}
	if in.DailyAutomaticBackupStartTime != nil {
		in, out := &in.DailyAutomaticBackupStartTime, &out.DailyAutomaticBackupStartTime
		*out = new(string)
		**out = **in
	}
	if in.DeploymentType != nil {
		in, out := &in.DeploymentType, &out.DeploymentType
		*out = new(string)
		**out = **in
	}
	if in.DnsName != nil {
		in, out := &in.DnsName, &out.DnsName
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceIDS != nil {
		in, out := &in.NetworkInterfaceIDS, &out.NetworkInterfaceIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.PreferredFileServerIP != nil {
		in, out := &in.PreferredFileServerIP, &out.PreferredFileServerIP
		*out = new(string)
		**out = **in
	}
	if in.PreferredSubnetID != nil {
		in, out := &in.PreferredSubnetID, &out.PreferredSubnetID
		*out = new(string)
		**out = **in
	}
	if in.RemoteAdministrationEndpoint != nil {
		in, out := &in.RemoteAdministrationEndpoint, &out.RemoteAdministrationEndpoint
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SelfManagedActiveDirectory != nil {
		in, out := &in.SelfManagedActiveDirectory, &out.SelfManagedActiveDirectory
		*out = new(WindowsFileSystemSpecSelfManagedActiveDirectory)
		(*in).DeepCopyInto(*out)
	}
	if in.SkipFinalBackup != nil {
		in, out := &in.SkipFinalBackup, &out.SkipFinalBackup
		*out = new(bool)
		**out = **in
	}
	if in.StorageCapacity != nil {
		in, out := &in.StorageCapacity, &out.StorageCapacity
		*out = new(int64)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.ThroughputCapacity != nil {
		in, out := &in.ThroughputCapacity, &out.ThroughputCapacity
		*out = new(int64)
		**out = **in
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceStartTime != nil {
		in, out := &in.WeeklyMaintenanceStartTime, &out.WeeklyMaintenanceStartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsFileSystemSpec2.
func (in *WindowsFileSystemSpec2) DeepCopy() *WindowsFileSystemSpec2 {
	if in == nil {
		return nil
	}
	out := new(WindowsFileSystemSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsFileSystemSpecSelfManagedActiveDirectory) DeepCopyInto(out *WindowsFileSystemSpecSelfManagedActiveDirectory) {
	*out = *in
	if in.DnsIPS != nil {
		in, out := &in.DnsIPS, &out.DnsIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.FileSystemAdministratorsGroup != nil {
		in, out := &in.FileSystemAdministratorsGroup, &out.FileSystemAdministratorsGroup
		*out = new(string)
		**out = **in
	}
	if in.OrganizationalUnitDistinguishedName != nil {
		in, out := &in.OrganizationalUnitDistinguishedName, &out.OrganizationalUnitDistinguishedName
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsFileSystemSpecSelfManagedActiveDirectory.
func (in *WindowsFileSystemSpecSelfManagedActiveDirectory) DeepCopy() *WindowsFileSystemSpecSelfManagedActiveDirectory {
	if in == nil {
		return nil
	}
	out := new(WindowsFileSystemSpecSelfManagedActiveDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowsFileSystemStatus) DeepCopyInto(out *WindowsFileSystemStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowsFileSystemStatus.
func (in *WindowsFileSystemStatus) DeepCopy() *WindowsFileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(WindowsFileSystemStatus)
	in.DeepCopyInto(out)
	return out
}
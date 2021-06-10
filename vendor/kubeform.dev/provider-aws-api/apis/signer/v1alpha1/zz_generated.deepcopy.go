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
func (in *SigningJob) DeepCopyInto(out *SigningJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJob.
func (in *SigningJob) DeepCopy() *SigningJob {
	if in == nil {
		return nil
	}
	out := new(SigningJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobList) DeepCopyInto(out *SigningJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SigningJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobList.
func (in *SigningJobList) DeepCopy() *SigningJobList {
	if in == nil {
		return nil
	}
	out := new(SigningJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpec) DeepCopyInto(out *SigningJobSpec) {
	*out = *in
	in.SigningJobSpec2.DeepCopyInto(&out.SigningJobSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpec.
func (in *SigningJobSpec) DeepCopy() *SigningJobSpec {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpec2) DeepCopyInto(out *SigningJobSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(SigningJobSpecDestination)
		(*in).DeepCopyInto(*out)
	}
	if in.IgnoreSigningJobFailure != nil {
		in, out := &in.IgnoreSigningJobFailure, &out.IgnoreSigningJobFailure
		*out = new(bool)
		**out = **in
	}
	if in.JobID != nil {
		in, out := &in.JobID, &out.JobID
		*out = new(string)
		**out = **in
	}
	if in.JobInvoker != nil {
		in, out := &in.JobInvoker, &out.JobInvoker
		*out = new(string)
		**out = **in
	}
	if in.JobOwner != nil {
		in, out := &in.JobOwner, &out.JobOwner
		*out = new(string)
		**out = **in
	}
	if in.PlatformDisplayName != nil {
		in, out := &in.PlatformDisplayName, &out.PlatformDisplayName
		*out = new(string)
		**out = **in
	}
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.ProfileName != nil {
		in, out := &in.ProfileName, &out.ProfileName
		*out = new(string)
		**out = **in
	}
	if in.ProfileVersion != nil {
		in, out := &in.ProfileVersion, &out.ProfileVersion
		*out = new(string)
		**out = **in
	}
	if in.RequestedBy != nil {
		in, out := &in.RequestedBy, &out.RequestedBy
		*out = new(string)
		**out = **in
	}
	if in.RevocationRecord != nil {
		in, out := &in.RevocationRecord, &out.RevocationRecord
		*out = make([]SigningJobSpecRevocationRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SignatureExpiresAt != nil {
		in, out := &in.SignatureExpiresAt, &out.SignatureExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.SignedObject != nil {
		in, out := &in.SignedObject, &out.SignedObject
		*out = make([]SigningJobSpecSignedObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(SigningJobSpecSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusReason != nil {
		in, out := &in.StatusReason, &out.StatusReason
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpec2.
func (in *SigningJobSpec2) DeepCopy() *SigningJobSpec2 {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpecDestination) DeepCopyInto(out *SigningJobSpecDestination) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(SigningJobSpecDestinationS3)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpecDestination.
func (in *SigningJobSpecDestination) DeepCopy() *SigningJobSpecDestination {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpecDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpecDestinationS3) DeepCopyInto(out *SigningJobSpecDestinationS3) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpecDestinationS3.
func (in *SigningJobSpecDestinationS3) DeepCopy() *SigningJobSpecDestinationS3 {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpecDestinationS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpecRevocationRecord) DeepCopyInto(out *SigningJobSpecRevocationRecord) {
	*out = *in
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = new(string)
		**out = **in
	}
	if in.RevokedBy != nil {
		in, out := &in.RevokedBy, &out.RevokedBy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpecRevocationRecord.
func (in *SigningJobSpecRevocationRecord) DeepCopy() *SigningJobSpecRevocationRecord {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpecRevocationRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpecSignedObject) DeepCopyInto(out *SigningJobSpecSignedObject) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = make([]SigningJobSpecSignedObjectS3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpecSignedObject.
func (in *SigningJobSpecSignedObject) DeepCopy() *SigningJobSpecSignedObject {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpecSignedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpecSignedObjectS3) DeepCopyInto(out *SigningJobSpecSignedObjectS3) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpecSignedObjectS3.
func (in *SigningJobSpecSignedObjectS3) DeepCopy() *SigningJobSpecSignedObjectS3 {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpecSignedObjectS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpecSource) DeepCopyInto(out *SigningJobSpecSource) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(SigningJobSpecSourceS3)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpecSource.
func (in *SigningJobSpecSource) DeepCopy() *SigningJobSpecSource {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpecSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpecSourceS3) DeepCopyInto(out *SigningJobSpecSourceS3) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpecSourceS3.
func (in *SigningJobSpecSourceS3) DeepCopy() *SigningJobSpecSourceS3 {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpecSourceS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobStatus) DeepCopyInto(out *SigningJobStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobStatus.
func (in *SigningJobStatus) DeepCopy() *SigningJobStatus {
	if in == nil {
		return nil
	}
	out := new(SigningJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfile) DeepCopyInto(out *SigningProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfile.
func (in *SigningProfile) DeepCopy() *SigningProfile {
	if in == nil {
		return nil
	}
	out := new(SigningProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileList) DeepCopyInto(out *SigningProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SigningProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileList.
func (in *SigningProfileList) DeepCopy() *SigningProfileList {
	if in == nil {
		return nil
	}
	out := new(SigningProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfilePermission) DeepCopyInto(out *SigningProfilePermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfilePermission.
func (in *SigningProfilePermission) DeepCopy() *SigningProfilePermission {
	if in == nil {
		return nil
	}
	out := new(SigningProfilePermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfilePermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfilePermissionList) DeepCopyInto(out *SigningProfilePermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SigningProfilePermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfilePermissionList.
func (in *SigningProfilePermissionList) DeepCopy() *SigningProfilePermissionList {
	if in == nil {
		return nil
	}
	out := new(SigningProfilePermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfilePermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfilePermissionSpec) DeepCopyInto(out *SigningProfilePermissionSpec) {
	*out = *in
	in.SigningProfilePermissionSpec2.DeepCopyInto(&out.SigningProfilePermissionSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfilePermissionSpec.
func (in *SigningProfilePermissionSpec) DeepCopy() *SigningProfilePermissionSpec {
	if in == nil {
		return nil
	}
	out := new(SigningProfilePermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfilePermissionSpec2) DeepCopyInto(out *SigningProfilePermissionSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(string)
		**out = **in
	}
	if in.ProfileName != nil {
		in, out := &in.ProfileName, &out.ProfileName
		*out = new(string)
		**out = **in
	}
	if in.ProfileVersion != nil {
		in, out := &in.ProfileVersion, &out.ProfileVersion
		*out = new(string)
		**out = **in
	}
	if in.StatementID != nil {
		in, out := &in.StatementID, &out.StatementID
		*out = new(string)
		**out = **in
	}
	if in.StatementIDPrefix != nil {
		in, out := &in.StatementIDPrefix, &out.StatementIDPrefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfilePermissionSpec2.
func (in *SigningProfilePermissionSpec2) DeepCopy() *SigningProfilePermissionSpec2 {
	if in == nil {
		return nil
	}
	out := new(SigningProfilePermissionSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfilePermissionStatus) DeepCopyInto(out *SigningProfilePermissionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfilePermissionStatus.
func (in *SigningProfilePermissionStatus) DeepCopy() *SigningProfilePermissionStatus {
	if in == nil {
		return nil
	}
	out := new(SigningProfilePermissionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileSpec) DeepCopyInto(out *SigningProfileSpec) {
	*out = *in
	in.SigningProfileSpec2.DeepCopyInto(&out.SigningProfileSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileSpec.
func (in *SigningProfileSpec) DeepCopy() *SigningProfileSpec {
	if in == nil {
		return nil
	}
	out := new(SigningProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileSpec2) DeepCopyInto(out *SigningProfileSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.PlatformDisplayName != nil {
		in, out := &in.PlatformDisplayName, &out.PlatformDisplayName
		*out = new(string)
		**out = **in
	}
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.RevocationRecord != nil {
		in, out := &in.RevocationRecord, &out.RevocationRecord
		*out = make([]SigningProfileSpecRevocationRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SignatureValidityPeriod != nil {
		in, out := &in.SignatureValidityPeriod, &out.SignatureValidityPeriod
		*out = new(SigningProfileSpecSignatureValidityPeriod)
		(*in).DeepCopyInto(*out)
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VersionArn != nil {
		in, out := &in.VersionArn, &out.VersionArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileSpec2.
func (in *SigningProfileSpec2) DeepCopy() *SigningProfileSpec2 {
	if in == nil {
		return nil
	}
	out := new(SigningProfileSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileSpecRevocationRecord) DeepCopyInto(out *SigningProfileSpecRevocationRecord) {
	*out = *in
	if in.RevocationEffectiveFrom != nil {
		in, out := &in.RevocationEffectiveFrom, &out.RevocationEffectiveFrom
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = new(string)
		**out = **in
	}
	if in.RevokedBy != nil {
		in, out := &in.RevokedBy, &out.RevokedBy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileSpecRevocationRecord.
func (in *SigningProfileSpecRevocationRecord) DeepCopy() *SigningProfileSpecRevocationRecord {
	if in == nil {
		return nil
	}
	out := new(SigningProfileSpecRevocationRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileSpecSignatureValidityPeriod) DeepCopyInto(out *SigningProfileSpecSignatureValidityPeriod) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileSpecSignatureValidityPeriod.
func (in *SigningProfileSpecSignatureValidityPeriod) DeepCopy() *SigningProfileSpecSignatureValidityPeriod {
	if in == nil {
		return nil
	}
	out := new(SigningProfileSpecSignatureValidityPeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileStatus) DeepCopyInto(out *SigningProfileStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileStatus.
func (in *SigningProfileStatus) DeepCopy() *SigningProfileStatus {
	if in == nil {
		return nil
	}
	out := new(SigningProfileStatus)
	in.DeepCopyInto(out)
	return out
}

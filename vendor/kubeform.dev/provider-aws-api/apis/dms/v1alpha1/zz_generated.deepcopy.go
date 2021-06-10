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

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	in.CertificateSpec2.DeepCopyInto(&out.CertificateSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec2) DeepCopyInto(out *CertificateSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CertificateArn != nil {
		in, out := &in.CertificateArn, &out.CertificateArn
		*out = new(string)
		**out = **in
	}
	if in.CertificateID != nil {
		in, out := &in.CertificateID, &out.CertificateID
		*out = new(string)
		**out = **in
	}
	if in.CertificatePem != nil {
		in, out := &in.CertificatePem, &out.CertificatePem
		*out = new(string)
		**out = **in
	}
	if in.CertificateWallet != nil {
		in, out := &in.CertificateWallet, &out.CertificateWallet
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec2.
func (in *CertificateSpec2) DeepCopy() *CertificateSpec2 {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Endpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointList) DeepCopyInto(out *EndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointList.
func (in *EndpointList) DeepCopy() *EndpointList {
	if in == nil {
		return nil
	}
	out := new(EndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
	in.EndpointSpec2.DeepCopyInto(&out.EndpointSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec2) DeepCopyInto(out *EndpointSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CertificateArn != nil {
		in, out := &in.CertificateArn, &out.CertificateArn
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.ElasticsearchSettings != nil {
		in, out := &in.ElasticsearchSettings, &out.ElasticsearchSettings
		*out = new(EndpointSpecElasticsearchSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.EndpointArn != nil {
		in, out := &in.EndpointArn, &out.EndpointArn
		*out = new(string)
		**out = **in
	}
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.EngineName != nil {
		in, out := &in.EngineName, &out.EngineName
		*out = new(string)
		**out = **in
	}
	if in.ExtraConnectionAttributes != nil {
		in, out := &in.ExtraConnectionAttributes, &out.ExtraConnectionAttributes
		*out = new(string)
		**out = **in
	}
	if in.KafkaSettings != nil {
		in, out := &in.KafkaSettings, &out.KafkaSettings
		*out = new(EndpointSpecKafkaSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.KinesisSettings != nil {
		in, out := &in.KinesisSettings, &out.KinesisSettings
		*out = new(EndpointSpecKinesisSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.KmsKeyArn != nil {
		in, out := &in.KmsKeyArn, &out.KmsKeyArn
		*out = new(string)
		**out = **in
	}
	if in.MongodbSettings != nil {
		in, out := &in.MongodbSettings, &out.MongodbSettings
		*out = new(EndpointSpecMongodbSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.S3Settings != nil {
		in, out := &in.S3Settings, &out.S3Settings
		*out = new(EndpointSpecS3Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerName != nil {
		in, out := &in.ServerName, &out.ServerName
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccessRole != nil {
		in, out := &in.ServiceAccessRole, &out.ServiceAccessRole
		*out = new(string)
		**out = **in
	}
	if in.SslMode != nil {
		in, out := &in.SslMode, &out.SslMode
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
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec2.
func (in *EndpointSpec2) DeepCopy() *EndpointSpec2 {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpecElasticsearchSettings) DeepCopyInto(out *EndpointSpecElasticsearchSettings) {
	*out = *in
	if in.EndpointURI != nil {
		in, out := &in.EndpointURI, &out.EndpointURI
		*out = new(string)
		**out = **in
	}
	if in.ErrorRetryDuration != nil {
		in, out := &in.ErrorRetryDuration, &out.ErrorRetryDuration
		*out = new(int64)
		**out = **in
	}
	if in.FullLoadErrorPercentage != nil {
		in, out := &in.FullLoadErrorPercentage, &out.FullLoadErrorPercentage
		*out = new(int64)
		**out = **in
	}
	if in.ServiceAccessRoleArn != nil {
		in, out := &in.ServiceAccessRoleArn, &out.ServiceAccessRoleArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpecElasticsearchSettings.
func (in *EndpointSpecElasticsearchSettings) DeepCopy() *EndpointSpecElasticsearchSettings {
	if in == nil {
		return nil
	}
	out := new(EndpointSpecElasticsearchSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpecKafkaSettings) DeepCopyInto(out *EndpointSpecKafkaSettings) {
	*out = *in
	if in.Broker != nil {
		in, out := &in.Broker, &out.Broker
		*out = new(string)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpecKafkaSettings.
func (in *EndpointSpecKafkaSettings) DeepCopy() *EndpointSpecKafkaSettings {
	if in == nil {
		return nil
	}
	out := new(EndpointSpecKafkaSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpecKinesisSettings) DeepCopyInto(out *EndpointSpecKinesisSettings) {
	*out = *in
	if in.MessageFormat != nil {
		in, out := &in.MessageFormat, &out.MessageFormat
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccessRoleArn != nil {
		in, out := &in.ServiceAccessRoleArn, &out.ServiceAccessRoleArn
		*out = new(string)
		**out = **in
	}
	if in.StreamArn != nil {
		in, out := &in.StreamArn, &out.StreamArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpecKinesisSettings.
func (in *EndpointSpecKinesisSettings) DeepCopy() *EndpointSpecKinesisSettings {
	if in == nil {
		return nil
	}
	out := new(EndpointSpecKinesisSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpecMongodbSettings) DeepCopyInto(out *EndpointSpecMongodbSettings) {
	*out = *in
	if in.AuthMechanism != nil {
		in, out := &in.AuthMechanism, &out.AuthMechanism
		*out = new(string)
		**out = **in
	}
	if in.AuthSource != nil {
		in, out := &in.AuthSource, &out.AuthSource
		*out = new(string)
		**out = **in
	}
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(string)
		**out = **in
	}
	if in.DocsToInvestigate != nil {
		in, out := &in.DocsToInvestigate, &out.DocsToInvestigate
		*out = new(string)
		**out = **in
	}
	if in.ExtractDocID != nil {
		in, out := &in.ExtractDocID, &out.ExtractDocID
		*out = new(string)
		**out = **in
	}
	if in.NestingLevel != nil {
		in, out := &in.NestingLevel, &out.NestingLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpecMongodbSettings.
func (in *EndpointSpecMongodbSettings) DeepCopy() *EndpointSpecMongodbSettings {
	if in == nil {
		return nil
	}
	out := new(EndpointSpecMongodbSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpecS3Settings) DeepCopyInto(out *EndpointSpecS3Settings) {
	*out = *in
	if in.BucketFolder != nil {
		in, out := &in.BucketFolder, &out.BucketFolder
		*out = new(string)
		**out = **in
	}
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.CompressionType != nil {
		in, out := &in.CompressionType, &out.CompressionType
		*out = new(string)
		**out = **in
	}
	if in.CsvDelimiter != nil {
		in, out := &in.CsvDelimiter, &out.CsvDelimiter
		*out = new(string)
		**out = **in
	}
	if in.CsvRowDelimiter != nil {
		in, out := &in.CsvRowDelimiter, &out.CsvRowDelimiter
		*out = new(string)
		**out = **in
	}
	if in.DatePartitionEnabled != nil {
		in, out := &in.DatePartitionEnabled, &out.DatePartitionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ExternalTableDefinition != nil {
		in, out := &in.ExternalTableDefinition, &out.ExternalTableDefinition
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccessRoleArn != nil {
		in, out := &in.ServiceAccessRoleArn, &out.ServiceAccessRoleArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpecS3Settings.
func (in *EndpointSpecS3Settings) DeepCopy() *EndpointSpecS3Settings {
	if in == nil {
		return nil
	}
	out := new(EndpointSpecS3Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointStatus) DeepCopyInto(out *EndpointStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointStatus.
func (in *EndpointStatus) DeepCopy() *EndpointStatus {
	if in == nil {
		return nil
	}
	out := new(EndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscription) DeepCopyInto(out *EventSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscription.
func (in *EventSubscription) DeepCopy() *EventSubscription {
	if in == nil {
		return nil
	}
	out := new(EventSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionList) DeepCopyInto(out *EventSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionList.
func (in *EventSubscriptionList) DeepCopy() *EventSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionSpec) DeepCopyInto(out *EventSubscriptionSpec) {
	*out = *in
	in.EventSubscriptionSpec2.DeepCopyInto(&out.EventSubscriptionSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionSpec.
func (in *EventSubscriptionSpec) DeepCopy() *EventSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionSpec2) DeepCopyInto(out *EventSubscriptionSpec2) {
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventCategories != nil {
		in, out := &in.EventCategories, &out.EventCategories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SnsTopicArn != nil {
		in, out := &in.SnsTopicArn, &out.SnsTopicArn
		*out = new(string)
		**out = **in
	}
	if in.SourceIDS != nil {
		in, out := &in.SourceIDS, &out.SourceIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionSpec2.
func (in *EventSubscriptionSpec2) DeepCopy() *EventSubscriptionSpec2 {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionStatus) DeepCopyInto(out *EventSubscriptionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionStatus.
func (in *EventSubscriptionStatus) DeepCopy() *EventSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationInstance) DeepCopyInto(out *ReplicationInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationInstance.
func (in *ReplicationInstance) DeepCopy() *ReplicationInstance {
	if in == nil {
		return nil
	}
	out := new(ReplicationInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationInstanceList) DeepCopyInto(out *ReplicationInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationInstanceList.
func (in *ReplicationInstanceList) DeepCopy() *ReplicationInstanceList {
	if in == nil {
		return nil
	}
	out := new(ReplicationInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationInstanceSpec) DeepCopyInto(out *ReplicationInstanceSpec) {
	*out = *in
	in.ReplicationInstanceSpec2.DeepCopyInto(&out.ReplicationInstanceSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationInstanceSpec.
func (in *ReplicationInstanceSpec) DeepCopy() *ReplicationInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationInstanceSpec2) DeepCopyInto(out *ReplicationInstanceSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.AllocatedStorage != nil {
		in, out := &in.AllocatedStorage, &out.AllocatedStorage
		*out = new(int64)
		**out = **in
	}
	if in.AllowMajorVersionUpgrade != nil {
		in, out := &in.AllowMajorVersionUpgrade, &out.AllowMajorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.AutoMinorVersionUpgrade != nil {
		in, out := &in.AutoMinorVersionUpgrade, &out.AutoMinorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyArn != nil {
		in, out := &in.KmsKeyArn, &out.KmsKeyArn
		*out = new(string)
		**out = **in
	}
	if in.MultiAz != nil {
		in, out := &in.MultiAz, &out.MultiAz
		*out = new(bool)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.PubliclyAccessible != nil {
		in, out := &in.PubliclyAccessible, &out.PubliclyAccessible
		*out = new(bool)
		**out = **in
	}
	if in.ReplicationInstanceArn != nil {
		in, out := &in.ReplicationInstanceArn, &out.ReplicationInstanceArn
		*out = new(string)
		**out = **in
	}
	if in.ReplicationInstanceClass != nil {
		in, out := &in.ReplicationInstanceClass, &out.ReplicationInstanceClass
		*out = new(string)
		**out = **in
	}
	if in.ReplicationInstanceID != nil {
		in, out := &in.ReplicationInstanceID, &out.ReplicationInstanceID
		*out = new(string)
		**out = **in
	}
	if in.ReplicationInstancePrivateIPS != nil {
		in, out := &in.ReplicationInstancePrivateIPS, &out.ReplicationInstancePrivateIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplicationInstancePublicIPS != nil {
		in, out := &in.ReplicationInstancePublicIPS, &out.ReplicationInstancePublicIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplicationSubnetGroupID != nil {
		in, out := &in.ReplicationSubnetGroupID, &out.ReplicationSubnetGroupID
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
	if in.VpcSecurityGroupIDS != nil {
		in, out := &in.VpcSecurityGroupIDS, &out.VpcSecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationInstanceSpec2.
func (in *ReplicationInstanceSpec2) DeepCopy() *ReplicationInstanceSpec2 {
	if in == nil {
		return nil
	}
	out := new(ReplicationInstanceSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationInstanceStatus) DeepCopyInto(out *ReplicationInstanceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationInstanceStatus.
func (in *ReplicationInstanceStatus) DeepCopy() *ReplicationInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSubnetGroup) DeepCopyInto(out *ReplicationSubnetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSubnetGroup.
func (in *ReplicationSubnetGroup) DeepCopy() *ReplicationSubnetGroup {
	if in == nil {
		return nil
	}
	out := new(ReplicationSubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationSubnetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSubnetGroupList) DeepCopyInto(out *ReplicationSubnetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationSubnetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSubnetGroupList.
func (in *ReplicationSubnetGroupList) DeepCopy() *ReplicationSubnetGroupList {
	if in == nil {
		return nil
	}
	out := new(ReplicationSubnetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationSubnetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSubnetGroupSpec) DeepCopyInto(out *ReplicationSubnetGroupSpec) {
	*out = *in
	in.ReplicationSubnetGroupSpec2.DeepCopyInto(&out.ReplicationSubnetGroupSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSubnetGroupSpec.
func (in *ReplicationSubnetGroupSpec) DeepCopy() *ReplicationSubnetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSubnetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSubnetGroupSpec2) DeepCopyInto(out *ReplicationSubnetGroupSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.ReplicationSubnetGroupArn != nil {
		in, out := &in.ReplicationSubnetGroupArn, &out.ReplicationSubnetGroupArn
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSubnetGroupDescription != nil {
		in, out := &in.ReplicationSubnetGroupDescription, &out.ReplicationSubnetGroupDescription
		*out = new(string)
		**out = **in
	}
	if in.ReplicationSubnetGroupID != nil {
		in, out := &in.ReplicationSubnetGroupID, &out.ReplicationSubnetGroupID
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSubnetGroupSpec2.
func (in *ReplicationSubnetGroupSpec2) DeepCopy() *ReplicationSubnetGroupSpec2 {
	if in == nil {
		return nil
	}
	out := new(ReplicationSubnetGroupSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSubnetGroupStatus) DeepCopyInto(out *ReplicationSubnetGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSubnetGroupStatus.
func (in *ReplicationSubnetGroupStatus) DeepCopy() *ReplicationSubnetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationSubnetGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationTask) DeepCopyInto(out *ReplicationTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationTask.
func (in *ReplicationTask) DeepCopy() *ReplicationTask {
	if in == nil {
		return nil
	}
	out := new(ReplicationTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationTaskList) DeepCopyInto(out *ReplicationTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationTaskList.
func (in *ReplicationTaskList) DeepCopy() *ReplicationTaskList {
	if in == nil {
		return nil
	}
	out := new(ReplicationTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationTaskSpec) DeepCopyInto(out *ReplicationTaskSpec) {
	*out = *in
	in.ReplicationTaskSpec2.DeepCopyInto(&out.ReplicationTaskSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationTaskSpec.
func (in *ReplicationTaskSpec) DeepCopy() *ReplicationTaskSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationTaskSpec2) DeepCopyInto(out *ReplicationTaskSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.CdcStartTime != nil {
		in, out := &in.CdcStartTime, &out.CdcStartTime
		*out = new(string)
		**out = **in
	}
	if in.MigrationType != nil {
		in, out := &in.MigrationType, &out.MigrationType
		*out = new(string)
		**out = **in
	}
	if in.ReplicationInstanceArn != nil {
		in, out := &in.ReplicationInstanceArn, &out.ReplicationInstanceArn
		*out = new(string)
		**out = **in
	}
	if in.ReplicationTaskArn != nil {
		in, out := &in.ReplicationTaskArn, &out.ReplicationTaskArn
		*out = new(string)
		**out = **in
	}
	if in.ReplicationTaskID != nil {
		in, out := &in.ReplicationTaskID, &out.ReplicationTaskID
		*out = new(string)
		**out = **in
	}
	if in.ReplicationTaskSettings != nil {
		in, out := &in.ReplicationTaskSettings, &out.ReplicationTaskSettings
		*out = new(string)
		**out = **in
	}
	if in.SourceEndpointArn != nil {
		in, out := &in.SourceEndpointArn, &out.SourceEndpointArn
		*out = new(string)
		**out = **in
	}
	if in.TableMappings != nil {
		in, out := &in.TableMappings, &out.TableMappings
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
	if in.TargetEndpointArn != nil {
		in, out := &in.TargetEndpointArn, &out.TargetEndpointArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationTaskSpec2.
func (in *ReplicationTaskSpec2) DeepCopy() *ReplicationTaskSpec2 {
	if in == nil {
		return nil
	}
	out := new(ReplicationTaskSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationTaskStatus) DeepCopyInto(out *ReplicationTaskStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationTaskStatus.
func (in *ReplicationTaskStatus) DeepCopy() *ReplicationTaskStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationTaskStatus)
	in.DeepCopyInto(out)
	return out
}
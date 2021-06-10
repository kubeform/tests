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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type TransferJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferJobSpec   `json:"spec,omitempty"`
	Status            TransferJobStatus `json:"status,omitempty"`
}

type TransferJobSpec struct {
	TransferJobSpec2 `json:",inline"`
	// +optional
	KubeformOutput TransferJobSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type TransferJobSpecScheduleScheduleEndDate struct {
	// Day of month. Must be from 1 to 31 and valid for the year and month.
	Day *int64 `json:"day" tf:"day"`
	// Month of year. Must be from 1 to 12.
	Month *int64 `json:"month" tf:"month"`
	// Year of date. Must be from 1 to 9999.
	Year *int64 `json:"year" tf:"year"`
}

type TransferJobSpecScheduleScheduleStartDate struct {
	// Day of month. Must be from 1 to 31 and valid for the year and month.
	Day *int64 `json:"day" tf:"day"`
	// Month of year. Must be from 1 to 12.
	Month *int64 `json:"month" tf:"month"`
	// Year of date. Must be from 1 to 9999.
	Year *int64 `json:"year" tf:"year"`
}

type TransferJobSpecScheduleStartTimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23.
	Hours *int64 `json:"hours" tf:"hours"`
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes *int64 `json:"minutes" tf:"minutes"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos *int64 `json:"nanos" tf:"nanos"`
	// Seconds of minutes of the time. Must normally be from 0 to 59.
	Seconds *int64 `json:"seconds" tf:"seconds"`
}

type TransferJobSpecSchedule struct {
	// The last day the recurring transfer will be run. If schedule_end_date is the same as schedule_start_date, the transfer will be executed only once.
	// +optional
	ScheduleEndDate *TransferJobSpecScheduleScheduleEndDate `json:"scheduleEndDate,omitempty" tf:"schedule_end_date"`
	// The first day the recurring transfer is scheduled to run. If schedule_start_date is in the past, the transfer will run for the first time on the following day.
	ScheduleStartDate *TransferJobSpecScheduleScheduleStartDate `json:"scheduleStartDate" tf:"schedule_start_date"`
	// The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone.
	// +optional
	StartTimeOfDay *TransferJobSpecScheduleStartTimeOfDay `json:"startTimeOfDay,omitempty" tf:"start_time_of_day"`
}

type TransferJobSpecTransferSpecAwsS3DataSourceAwsAccessKey struct {
	// AWS Key ID.
	AccessKeyID *string `json:"-" sensitive:"true" tf:"access_key_id"`
	// AWS Secret Access Key.
	SecretAccessKey *string `json:"-" sensitive:"true" tf:"secret_access_key"`
}

type TransferJobSpecTransferSpecAwsS3DataSource struct {
	// AWS credentials block.
	AwsAccessKey *TransferJobSpecTransferSpecAwsS3DataSourceAwsAccessKey `json:"awsAccessKey" tf:"aws_access_key"`
	// S3 Bucket name.
	BucketName *string `json:"bucketName" tf:"bucket_name"`
}

type TransferJobSpecTransferSpecGcsDataSink struct {
	// Google Cloud Storage bucket name.
	BucketName *string `json:"bucketName" tf:"bucket_name"`
}

type TransferJobSpecTransferSpecGcsDataSource struct {
	// Google Cloud Storage bucket name.
	BucketName *string `json:"bucketName" tf:"bucket_name"`
}

type TransferJobSpecTransferSpecHttpDataSource struct {
	// The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.
	ListURL *string `json:"listURL" tf:"list_url"`
}

type TransferJobSpecTransferSpecObjectConditions struct {
	// exclude_prefixes must follow the requirements described for include_prefixes.
	// +optional
	ExcludePrefixes []string `json:"excludePrefixes,omitempty" tf:"exclude_prefixes"`
	// If include_refixes is specified, objects that satisfy the object conditions must have names that start with one of the include_prefixes and that do not start with any of the exclude_prefixes. If include_prefixes is not specified, all objects except those that have names starting with one of the exclude_prefixes must satisfy the object conditions.
	// +optional
	IncludePrefixes []string `json:"includePrefixes,omitempty" tf:"include_prefixes"`
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +optional
	MaxTimeElapsedSinceLastModification *string `json:"maxTimeElapsedSinceLastModification,omitempty" tf:"max_time_elapsed_since_last_modification"`
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +optional
	MinTimeElapsedSinceLastModification *string `json:"minTimeElapsedSinceLastModification,omitempty" tf:"min_time_elapsed_since_last_modification"`
}

type TransferJobSpecTransferSpecTransferOptions struct {
	// Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and delete_objects_unique_in_sink are mutually exclusive.
	// +optional
	DeleteObjectsFromSourceAfterTransfer *bool `json:"deleteObjectsFromSourceAfterTransfer,omitempty" tf:"delete_objects_from_source_after_transfer"`
	// Whether objects that exist only in the sink should be deleted. Note that this option and delete_objects_from_source_after_transfer are mutually exclusive.
	// +optional
	DeleteObjectsUniqueInSink *bool `json:"deleteObjectsUniqueInSink,omitempty" tf:"delete_objects_unique_in_sink"`
	// Whether overwriting objects that already exist in the sink is allowed.
	// +optional
	OverwriteObjectsAlreadyExistingInSink *bool `json:"overwriteObjectsAlreadyExistingInSink,omitempty" tf:"overwrite_objects_already_existing_in_sink"`
}

type TransferJobSpecTransferSpec struct {
	// An AWS S3 data source.
	// +optional
	AwsS3DataSource *TransferJobSpecTransferSpecAwsS3DataSource `json:"awsS3DataSource,omitempty" tf:"aws_s3_data_source"`
	// A Google Cloud Storage data sink.
	// +optional
	GcsDataSink *TransferJobSpecTransferSpecGcsDataSink `json:"gcsDataSink,omitempty" tf:"gcs_data_sink"`
	// A Google Cloud Storage data source.
	// +optional
	GcsDataSource *TransferJobSpecTransferSpecGcsDataSource `json:"gcsDataSource,omitempty" tf:"gcs_data_source"`
	// An HTTP URL data source.
	// +optional
	HttpDataSource *TransferJobSpecTransferSpecHttpDataSource `json:"httpDataSource,omitempty" tf:"http_data_source"`
	// Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' last_modification_time do not exclude objects in a data sink.
	// +optional
	ObjectConditions *TransferJobSpecTransferSpecObjectConditions `json:"objectConditions,omitempty" tf:"object_conditions"`
	// Characteristics of how to treat files from datasource and sink during job. If the option delete_objects_unique_in_sink is true, object conditions based on objects' last_modification_time are ignored and do not exclude objects in a data source or a data sink.
	// +optional
	TransferOptions *TransferJobSpecTransferSpecTransferOptions `json:"transferOptions,omitempty" tf:"transfer_options"`
}

type TransferJobSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// When the Transfer Job was created.
	// +optional
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time"`
	// When the Transfer Job was deleted.
	// +optional
	DeletionTime *string `json:"deletionTime,omitempty" tf:"deletion_time"`
	// Unique description to identify the Transfer Job.
	Description *string `json:"description" tf:"description"`
	// When the Transfer Job was last modified.
	// +optional
	LastModificationTime *string `json:"lastModificationTime,omitempty" tf:"last_modification_time"`
	// The name of the Transfer Job.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run.
	Schedule *TransferJobSpecSchedule `json:"schedule" tf:"schedule"`
	// Status of the job. Default: ENABLED. NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Transfer specification.
	TransferSpec *TransferJobSpecTransferSpec `json:"transferSpec" tf:"transfer_spec"`
}

type TransferJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TransferJobList is a list of TransferJobs
type TransferJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TransferJob CRD objects
	Items []TransferJob `json:"items,omitempty"`
}

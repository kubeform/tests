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

type SchedulerJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchedulerJobSpec   `json:"spec,omitempty"`
	Status            SchedulerJobStatus `json:"status,omitempty"`
}

type SchedulerJobSpec struct {
	SchedulerJobSpec2 `json:",inline"`
	// +optional
	KubeformOutput SchedulerJobSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type SchedulerJobSpecAppEngineHTTPTargetAppEngineRouting struct {
	// App instance.
	// By default, the job is sent to an instance which is available when the job is attempted.
	// +optional
	Instance *string `json:"instance,omitempty" tf:"instance"`
	// App service.
	// By default, the job is sent to the service which is the default service when the job is attempted.
	// +optional
	Service *string `json:"service,omitempty" tf:"service"`
	// App version.
	// By default, the job is sent to the version which is the default version when the job is attempted.
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type SchedulerJobSpecAppEngineHTTPTarget struct {
	// App Engine Routing setting for the job.
	// +optional
	AppEngineRouting *SchedulerJobSpecAppEngineHTTPTargetAppEngineRouting `json:"appEngineRouting,omitempty" tf:"app_engine_routing"`
	// HTTP request body.
	// A request body is allowed only if the HTTP method is POST or PUT.
	// It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.
	//
	// A base64-encoded string.
	// +optional
	Body *string `json:"body,omitempty" tf:"body"`
	// HTTP request headers.
	// This map contains the header field names and values.
	// Headers can be set when the job is created.
	// +optional
	Headers *map[string]string `json:"headers,omitempty" tf:"headers"`
	// Which HTTP method to use for the request.
	// +optional
	HttpMethod *string `json:"httpMethod,omitempty" tf:"http_method"`
	// The relative URI.
	// The relative URL must begin with "/" and must be a valid HTTP relative URL.
	// It can contain a path, query string arguments, and \\# fragments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters
	RelativeURI *string `json:"relativeURI" tf:"relative_uri"`
}

type SchedulerJobSpecHttpTargetOauthToken struct {
	// OAuth scope to be used for generating OAuth access token. If not specified,
	// "https://www.googleapis.com/auth/cloud-platform" will be used.
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// Service account email to be used for generating OAuth token.
	// The service account must be within the same project as the job.
	ServiceAccountEmail *string `json:"serviceAccountEmail" tf:"service_account_email"`
}

type SchedulerJobSpecHttpTargetOidcToken struct {
	// Audience to be used when generating OIDC token. If not specified,
	// the URI specified in target will be used.
	// +optional
	Audience *string `json:"audience,omitempty" tf:"audience"`
	// Service account email to be used for generating OAuth token.
	// The service account must be within the same project as the job.
	ServiceAccountEmail *string `json:"serviceAccountEmail" tf:"service_account_email"`
}

type SchedulerJobSpecHttpTarget struct {
	// HTTP request body.
	// A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
	// It is an error to set body on a job with an incompatible HttpMethod.
	//
	// A base64-encoded string.
	// +optional
	Body *string `json:"body,omitempty" tf:"body"`
	// This map contains the header field names and values.
	// Repeated headers are not supported, but a header value can contain commas.
	// +optional
	Headers *map[string]string `json:"headers,omitempty" tf:"headers"`
	// Which HTTP method to use for the request.
	// +optional
	HttpMethod *string `json:"httpMethod,omitempty" tf:"http_method"`
	// Contains information needed for generating an OAuth token.
	// This type of authorization should be used when sending requests to a GCP endpoint.
	// +optional
	OauthToken *SchedulerJobSpecHttpTargetOauthToken `json:"oauthToken,omitempty" tf:"oauth_token"`
	// Contains information needed for generating an OpenID Connect token.
	// This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
	// +optional
	OidcToken *SchedulerJobSpecHttpTargetOidcToken `json:"oidcToken,omitempty" tf:"oidc_token"`
	// The full URI path that the request will be sent to.
	Uri *string `json:"uri" tf:"uri"`
}

type SchedulerJobSpecPubsubTarget struct {
	// Attributes for PubsubMessage.
	// Pubsub message must contain either non-empty data, or at least one attribute.
	// +optional
	Attributes *map[string]string `json:"attributes,omitempty" tf:"attributes"`
	// The message payload for PubsubMessage.
	// Pubsub message must contain either non-empty data, or at least one attribute.
	//
	//  A base64-encoded string.
	// +optional
	Data *string `json:"data,omitempty" tf:"data"`
	// The full resource name for the Cloud Pub/Sub topic to which
	// messages will be published when a job is delivered. ~>**NOTE:**
	// The topic name must be in the same format as required by PubSub's
	// PublishRequest.name, e.g. 'projects/my-project/topics/my-topic'.
	TopicName *string `json:"topicName" tf:"topic_name"`
}

type SchedulerJobSpecRetryConfig struct {
	// The maximum amount of time to wait before retrying a job after it fails.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// +optional
	MaxBackoffDuration *string `json:"maxBackoffDuration,omitempty" tf:"max_backoff_duration"`
	// The time between retries will double maxDoublings times.
	// A job's retry interval starts at minBackoffDuration,
	// then doubles maxDoublings times, then increases linearly,
	// and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
	// +optional
	MaxDoublings *int64 `json:"maxDoublings,omitempty" tf:"max_doublings"`
	// The time limit for retrying a failed job, measured from time when an execution was first attempted.
	// If specified with retryCount, the job will be retried until both limits are reached.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// +optional
	MaxRetryDuration *string `json:"maxRetryDuration,omitempty" tf:"max_retry_duration"`
	// The minimum amount of time to wait before retrying a job after it fails.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// +optional
	MinBackoffDuration *string `json:"minBackoffDuration,omitempty" tf:"min_backoff_duration"`
	// The number of attempts that the system will make to run a
	// job using the exponential backoff procedure described by maxDoublings.
	// Values greater than 5 and negative values are not allowed.
	// +optional
	RetryCount *int64 `json:"retryCount,omitempty" tf:"retry_count"`
}

type SchedulerJobSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// App Engine HTTP target.
	// If the job providers a App Engine HTTP target the cron will
	// send a request to the service instance
	// +optional
	AppEngineHTTPTarget *SchedulerJobSpecAppEngineHTTPTarget `json:"appEngineHTTPTarget,omitempty" tf:"app_engine_http_target"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	// The allowed duration for this deadline is:
	// * For HTTP targets, between 15 seconds and 30 minutes.
	// * For App Engine HTTP targets, between 15 seconds and 24 hours.
	// * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	// +optional
	AttemptDeadline *string `json:"attemptDeadline,omitempty" tf:"attempt_deadline"`
	// A human-readable description for the job.
	// This string must not contain more than 500 characters.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// HTTP target.
	// If the job providers a http_target the cron will
	// send a request to the targeted url
	// +optional
	HttpTarget *SchedulerJobSpecHttpTarget `json:"httpTarget,omitempty" tf:"http_target"`
	// The name of the job.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Pub/Sub target
	// If the job providers a Pub/Sub target the cron will publish
	// a message to the provided topic
	// +optional
	PubsubTarget *SchedulerJobSpecPubsubTarget `json:"pubsubTarget,omitempty" tf:"pubsub_target"`
	// Region where the scheduler job resides. If it is not provided, Terraform will use the provider default.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// By default, if a job does not complete successfully,
	// meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	// +optional
	RetryConfig *SchedulerJobSpecRetryConfig `json:"retryConfig,omitempty" tf:"retry_config"`
	// Describes the schedule on which the job will be executed.
	// +optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule"`
	// Specifies the time zone to be used in interpreting schedule.
	// The value of this field must be a time zone name from the tz database.
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
}

type SchedulerJobStatus struct {
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

// SchedulerJobList is a list of SchedulerJobs
type SchedulerJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SchedulerJob CRD objects
	Items []SchedulerJob `json:"items,omitempty"`
}

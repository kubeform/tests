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

type Topic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec,omitempty"`
	Status            TopicStatus `json:"status,omitempty"`
}

type TopicSpec struct {
	TopicSpec2 `json:",inline"`
	// +optional
	KubeformOutput TopicSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type TopicSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplicationFailureFeedbackRoleArn *string `json:"applicationFailureFeedbackRoleArn,omitempty" tf:"application_failure_feedback_role_arn"`
	// +optional
	ApplicationSuccessFeedbackRoleArn *string `json:"applicationSuccessFeedbackRoleArn,omitempty" tf:"application_success_feedback_role_arn"`
	// +optional
	ApplicationSuccessFeedbackSampleRate *int64 `json:"applicationSuccessFeedbackSampleRate,omitempty" tf:"application_success_feedback_sample_rate"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication"`
	// +optional
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty" tf:"delivery_policy"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FifoTopic *bool `json:"fifoTopic,omitempty" tf:"fifo_topic"`
	// +optional
	HttpFailureFeedbackRoleArn *string `json:"httpFailureFeedbackRoleArn,omitempty" tf:"http_failure_feedback_role_arn"`
	// +optional
	HttpSuccessFeedbackRoleArn *string `json:"httpSuccessFeedbackRoleArn,omitempty" tf:"http_success_feedback_role_arn"`
	// +optional
	HttpSuccessFeedbackSampleRate *int64 `json:"httpSuccessFeedbackSampleRate,omitempty" tf:"http_success_feedback_sample_rate"`
	// +optional
	KmsMasterKeyID *string `json:"kmsMasterKeyID,omitempty" tf:"kms_master_key_id"`
	// +optional
	LambdaFailureFeedbackRoleArn *string `json:"lambdaFailureFeedbackRoleArn,omitempty" tf:"lambda_failure_feedback_role_arn"`
	// +optional
	LambdaSuccessFeedbackRoleArn *string `json:"lambdaSuccessFeedbackRoleArn,omitempty" tf:"lambda_success_feedback_role_arn"`
	// +optional
	LambdaSuccessFeedbackSampleRate *int64 `json:"lambdaSuccessFeedbackSampleRate,omitempty" tf:"lambda_success_feedback_sample_rate"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
	// +optional
	SqsFailureFeedbackRoleArn *string `json:"sqsFailureFeedbackRoleArn,omitempty" tf:"sqs_failure_feedback_role_arn"`
	// +optional
	SqsSuccessFeedbackRoleArn *string `json:"sqsSuccessFeedbackRoleArn,omitempty" tf:"sqs_success_feedback_role_arn"`
	// +optional
	SqsSuccessFeedbackSampleRate *int64 `json:"sqsSuccessFeedbackSampleRate,omitempty" tf:"sqs_success_feedback_sample_rate"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type TopicStatus struct {
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

// TopicList is a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Topic CRD objects
	Items []Topic `json:"items,omitempty"`
}

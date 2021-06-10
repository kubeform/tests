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

type RealtimeLogConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RealtimeLogConfigSpec   `json:"spec,omitempty"`
	Status            RealtimeLogConfigStatus `json:"status,omitempty"`
}

type RealtimeLogConfigSpec struct {
	RealtimeLogConfigSpec2 `json:",inline"`
	// +optional
	KubeformOutput RealtimeLogConfigSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type RealtimeLogConfigSpecEndpointKinesisStreamConfig struct {
	RoleArn   *string `json:"roleArn" tf:"role_arn"`
	StreamArn *string `json:"streamArn" tf:"stream_arn"`
}

type RealtimeLogConfigSpecEndpoint struct {
	KinesisStreamConfig *RealtimeLogConfigSpecEndpointKinesisStreamConfig `json:"kinesisStreamConfig" tf:"kinesis_stream_config"`
	StreamType          *string                                           `json:"streamType" tf:"stream_type"`
}

type RealtimeLogConfigSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn          *string                        `json:"arn,omitempty" tf:"arn"`
	Endpoint     *RealtimeLogConfigSpecEndpoint `json:"endpoint" tf:"endpoint"`
	Fields       []string                       `json:"fields" tf:"fields"`
	Name         *string                        `json:"name" tf:"name"`
	SamplingRate *int64                         `json:"samplingRate" tf:"sampling_rate"`
}

type RealtimeLogConfigStatus struct {
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

// RealtimeLogConfigList is a list of RealtimeLogConfigs
type RealtimeLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RealtimeLogConfig CRD objects
	Items []RealtimeLogConfig `json:"items,omitempty"`
}
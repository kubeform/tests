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

type AnalyticsStreamInputIothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyticsStreamInputIothubSpec   `json:"spec,omitempty"`
	Status            AnalyticsStreamInputIothubStatus `json:"status,omitempty"`
}

type AnalyticsStreamInputIothubSpec struct {
	AnalyticsStreamInputIothubSpec2 `json:",inline"`
	// +optional
	KubeformOutput AnalyticsStreamInputIothubSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type AnalyticsStreamInputIothubSpecSerialization struct {
	// +optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`
	// +optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter"`
	Type           *string `json:"type" tf:"type"`
}

type AnalyticsStreamInputIothubSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	Endpoint                  *string                                      `json:"endpoint" tf:"endpoint"`
	EventhubConsumerGroupName *string                                      `json:"eventhubConsumerGroupName" tf:"eventhub_consumer_group_name"`
	IothubNamespace           *string                                      `json:"iothubNamespace" tf:"iothub_namespace"`
	Name                      *string                                      `json:"name" tf:"name"`
	ResourceGroupName         *string                                      `json:"resourceGroupName" tf:"resource_group_name"`
	Serialization             *AnalyticsStreamInputIothubSpecSerialization `json:"serialization" tf:"serialization"`
	SharedAccessPolicyKey     *string                                      `json:"-" sensitive:"true" tf:"shared_access_policy_key"`
	SharedAccessPolicyName    *string                                      `json:"sharedAccessPolicyName" tf:"shared_access_policy_name"`
	StreamAnalyticsJobName    *string                                      `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
}

type AnalyticsStreamInputIothubStatus struct {
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

// AnalyticsStreamInputIothubList is a list of AnalyticsStreamInputIothubs
type AnalyticsStreamInputIothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AnalyticsStreamInputIothub CRD objects
	Items []AnalyticsStreamInputIothub `json:"items,omitempty"`
}
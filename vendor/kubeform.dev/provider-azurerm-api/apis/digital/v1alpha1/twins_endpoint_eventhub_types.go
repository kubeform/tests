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

type TwinsEndpointEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TwinsEndpointEventhubSpec   `json:"spec,omitempty"`
	Status            TwinsEndpointEventhubStatus `json:"status,omitempty"`
}

type TwinsEndpointEventhubSpec struct {
	TwinsEndpointEventhubSpec2 `json:",inline"`
	// +optional
	KubeformOutput TwinsEndpointEventhubSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type TwinsEndpointEventhubSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	DeadLetterStorageSecret           *string `json:"-" sensitive:"true" tf:"dead_letter_storage_secret"`
	DigitalTwinsID                    *string `json:"digitalTwinsID" tf:"digital_twins_id"`
	EventhubPrimaryConnectionString   *string `json:"-" sensitive:"true" tf:"eventhub_primary_connection_string"`
	EventhubSecondaryConnectionString *string `json:"-" sensitive:"true" tf:"eventhub_secondary_connection_string"`
	Name                              *string `json:"name" tf:"name"`
}

type TwinsEndpointEventhubStatus struct {
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

// TwinsEndpointEventhubList is a list of TwinsEndpointEventhubs
type TwinsEndpointEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TwinsEndpointEventhub CRD objects
	Items []TwinsEndpointEventhub `json:"items,omitempty"`
}
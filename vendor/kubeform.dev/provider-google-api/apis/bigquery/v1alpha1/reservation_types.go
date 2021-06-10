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

type Reservation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReservationSpec   `json:"spec,omitempty"`
	Status            ReservationStatus `json:"status,omitempty"`
}

type ReservationSpec struct {
	ReservationSpec2 `json:",inline"`
	// +optional
	KubeformOutput ReservationSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ReservationSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	// +optional
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty" tf:"ignore_idle_slots"`
	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity *int64 `json:"slotCapacity" tf:"slot_capacity"`
}

type ReservationStatus struct {
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

// ReservationList is a list of Reservations
type ReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Reservation CRD objects
	Items []Reservation `json:"items,omitempty"`
}

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

type NetworkPeeringRoutesConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPeeringRoutesConfigSpec   `json:"spec,omitempty"`
	Status            NetworkPeeringRoutesConfigStatus `json:"status,omitempty"`
}

type NetworkPeeringRoutesConfigSpec struct {
	NetworkPeeringRoutesConfigSpec2 `json:",inline"`
	// +optional
	KubeformOutput NetworkPeeringRoutesConfigSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type NetworkPeeringRoutesConfigSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes *bool `json:"exportCustomRoutes" tf:"export_custom_routes"`
	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes *bool `json:"importCustomRoutes" tf:"import_custom_routes"`
	// The name of the primary network for the peering.
	Network *string `json:"network" tf:"network"`
	// Name of the peering.
	Peering *string `json:"peering" tf:"peering"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type NetworkPeeringRoutesConfigStatus struct {
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

// NetworkPeeringRoutesConfigList is a list of NetworkPeeringRoutesConfigs
type NetworkPeeringRoutesConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkPeeringRoutesConfig CRD objects
	Items []NetworkPeeringRoutesConfig `json:"items,omitempty"`
}
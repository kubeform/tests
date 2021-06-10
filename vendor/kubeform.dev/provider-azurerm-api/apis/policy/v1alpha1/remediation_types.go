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

type Remediation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RemediationSpec   `json:"spec,omitempty"`
	Status            RemediationStatus `json:"status,omitempty"`
}

type RemediationSpec struct {
	RemediationSpec2 `json:",inline"`
	// +optional
	KubeformOutput RemediationSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type RemediationSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	LocationFilters    []string `json:"locationFilters,omitempty" tf:"location_filters"`
	Name               *string  `json:"name" tf:"name"`
	PolicyAssignmentID *string  `json:"policyAssignmentID" tf:"policy_assignment_id"`
	// +optional
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceID,omitempty" tf:"policy_definition_reference_id"`
	// +optional
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode"`
	Scope                 *string `json:"scope" tf:"scope"`
}

type RemediationStatus struct {
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

// RemediationList is a list of Remediations
type RemediationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Remediation CRD objects
	Items []Remediation `json:"items,omitempty"`
}
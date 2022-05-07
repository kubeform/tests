/*
Copyright The Kubeform Authors.

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
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type RoleDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleDefinitionSpec   `json:"spec,omitempty"`
	Status            RoleDefinitionStatus `json:"status,omitempty"`
}

type RoleDefinitionSpecPermissions struct {
	// +optional
	Actions []string `json:"actions,omitempty" tf:"actions,omitempty"`
	// +optional
	DataActions []string `json:"dataActions,omitempty" tf:"data_actions,omitempty"`
	// +optional
	NotActions []string `json:"notActions,omitempty" tf:"not_actions,omitempty"`
	// +optional
	NotDataActions []string `json:"notDataActions,omitempty" tf:"not_data_actions,omitempty"`
}

type RoleDefinitionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AssignableScopes []string `json:"assignableScopes" tf:"assignable_scopes"`
	// +optional
	Description string                          `json:"description,omitempty" tf:"description,omitempty"`
	Name        string                          `json:"name" tf:"name"`
	Permissions []RoleDefinitionSpecPermissions `json:"permissions" tf:"permissions"`
	// +optional
	RoleDefinitionID string `json:"roleDefinitionID,omitempty" tf:"role_definition_id,omitempty"`
	Scope            string `json:"scope" tf:"scope"`
}

type RoleDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RoleDefinitionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RoleDefinitionList is a list of RoleDefinitions
type RoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RoleDefinition CRD objects
	Items []RoleDefinition `json:"items,omitempty"`
}
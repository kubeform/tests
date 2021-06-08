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

type ProjectOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            ProjectOrganizationPolicyStatus `json:"status,omitempty"`
}

type ProjectOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced" tf:"enforced"`
}

type ProjectOrganizationPolicySpecListPolicyAllow struct {
	// +optional
	All bool `json:"all,omitempty" tf:"all,omitempty"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ProjectOrganizationPolicySpecListPolicyDeny struct {
	// +optional
	All bool `json:"all,omitempty" tf:"all,omitempty"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ProjectOrganizationPolicySpecListPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Allow []ProjectOrganizationPolicySpecListPolicyAllow `json:"allow,omitempty" tf:"allow,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Deny []ProjectOrganizationPolicySpecListPolicyDeny `json:"deny,omitempty" tf:"deny,omitempty"`
	// +optional
	InheritFromParent bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`
	// +optional
	SuggestedValue string `json:"suggestedValue,omitempty" tf:"suggested_value,omitempty"`
}

type ProjectOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default" tf:"default"`
}

type ProjectOrganizationPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	BooleanPolicy []ProjectOrganizationPolicySpecBooleanPolicy `json:"booleanPolicy,omitempty" tf:"boolean_policy,omitempty"`
	Constraint    string                                       `json:"constraint" tf:"constraint"`
	// +optional
	Etag string `json:"etag,omitempty" tf:"etag,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ListPolicy []ProjectOrganizationPolicySpecListPolicy `json:"listPolicy,omitempty" tf:"list_policy,omitempty"`
	Project    string                                    `json:"project" tf:"project"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RestorePolicy []ProjectOrganizationPolicySpecRestorePolicy `json:"restorePolicy,omitempty" tf:"restore_policy,omitempty"`
	// +optional
	UpdateTime string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
	// +optional
	Version int64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ProjectOrganizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ProjectOrganizationPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectOrganizationPolicyList is a list of ProjectOrganizationPolicys
type ProjectOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectOrganizationPolicy CRD objects
	Items []ProjectOrganizationPolicy `json:"items,omitempty"`
}
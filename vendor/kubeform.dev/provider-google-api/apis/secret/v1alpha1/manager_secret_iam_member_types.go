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

type ManagerSecretIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerSecretIamMemberSpec   `json:"spec,omitempty"`
	Status            ManagerSecretIamMemberStatus `json:"status,omitempty"`
}

type ManagerSecretIamMemberSpec struct {
	ManagerSecretIamMemberSpec2 `json:",inline"`
	// +optional
	KubeformOutput ManagerSecretIamMemberSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ManagerSecretIamMemberSpecCondition struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Expression  *string `json:"expression" tf:"expression"`
	Title       *string `json:"title" tf:"title"`
}

type ManagerSecretIamMemberSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Condition *ManagerSecretIamMemberSpecCondition `json:"condition,omitempty" tf:"condition"`
	// +optional
	Etag   *string `json:"etag,omitempty" tf:"etag"`
	Member *string `json:"member" tf:"member"`
	// +optional
	Project  *string `json:"project,omitempty" tf:"project"`
	Role     *string `json:"role" tf:"role"`
	SecretID *string `json:"secretID" tf:"secret_id"`
}

type ManagerSecretIamMemberStatus struct {
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

// ManagerSecretIamMemberList is a list of ManagerSecretIamMembers
type ManagerSecretIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagerSecretIamMember CRD objects
	Items []ManagerSecretIamMember `json:"items,omitempty"`
}

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

type RepositoryPermissionsPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositoryPermissionsPolicySpec   `json:"spec,omitempty"`
	Status            RepositoryPermissionsPolicyStatus `json:"status,omitempty"`
}

type RepositoryPermissionsPolicySpec struct {
	RepositoryPermissionsPolicySpec2 `json:",inline"`
	// +optional
	KubeformOutput RepositoryPermissionsPolicySpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type RepositoryPermissionsPolicySpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Domain *string `json:"domain" tf:"domain"`
	// +optional
	DomainOwner    *string `json:"domainOwner,omitempty" tf:"domain_owner"`
	PolicyDocument *string `json:"policyDocument" tf:"policy_document"`
	// +optional
	PolicyRevision *string `json:"policyRevision,omitempty" tf:"policy_revision"`
	Repository     *string `json:"repository" tf:"repository"`
	// +optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn"`
}

type RepositoryPermissionsPolicyStatus struct {
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

// RepositoryPermissionsPolicyList is a list of RepositoryPermissionsPolicys
type RepositoryPermissionsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RepositoryPermissionsPolicy CRD objects
	Items []RepositoryPermissionsPolicy `json:"items,omitempty"`
}
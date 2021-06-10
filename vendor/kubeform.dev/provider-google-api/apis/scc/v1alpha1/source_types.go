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

type Source struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SourceSpec   `json:"spec,omitempty"`
	Status            SourceStatus `json:"status,omitempty"`
}

type SourceSpec struct {
	SourceSpec2 `json:",inline"`
	// +optional
	KubeformOutput SourceSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type SourceSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The description of the source (max of 1024 characters).
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The source’s display name. A source’s display name must be unique
	// amongst its siblings, for example, two sources with the same parent
	// can't share the same display name. The display name must start and end
	// with a letter or digit, may contain letters, digits, spaces, hyphens,
	// and underscores, and can be no longer than 32 characters.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// The resource name of this source, in the format
	// 'organizations/{{organization}}/sources/{{source}}'.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The organization whose Cloud Security Command Center the Source
	// lives in.
	Organization *string `json:"organization" tf:"organization"`
}

type SourceStatus struct {
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

// SourceList is a list of Sources
type SourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Source CRD objects
	Items []Source `json:"items,omitempty"`
}
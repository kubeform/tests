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

type FolderExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderExclusionSpec   `json:"spec,omitempty"`
	Status            FolderExclusionStatus `json:"status,omitempty"`
}

type FolderExclusionSpec struct {
	FolderExclusionSpec2 `json:",inline"`
	// +optional
	KubeformOutput FolderExclusionSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type FolderExclusionSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A human-readable description.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter *string `json:"filter" tf:"filter"`
	Folder *string `json:"folder" tf:"folder"`
	// The name of the logging exclusion.
	Name *string `json:"name" tf:"name"`
}

type FolderExclusionStatus struct {
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

// FolderExclusionList is a list of FolderExclusions
type FolderExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FolderExclusion CRD objects
	Items []FolderExclusion `json:"items,omitempty"`
}
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

type FirestoreIndex struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirestoreIndexSpec   `json:"spec,omitempty"`
	Status            FirestoreIndexStatus `json:"status,omitempty"`
}

type FirestoreIndexSpecFields struct {
	// +optional
	ArrayConfig string `json:"arrayConfig,omitempty" tf:"array_config,omitempty"`
	// +optional
	FieldPath string `json:"fieldPath,omitempty" tf:"field_path,omitempty"`
	// +optional
	Order string `json:"order,omitempty" tf:"order,omitempty"`
}

type FirestoreIndexSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Collection string `json:"collection" tf:"collection"`
	// +optional
	Database string `json:"database,omitempty" tf:"database,omitempty"`
	// +kubebuilder:validation:MinItems=2
	Fields []FirestoreIndexSpecFields `json:"fields" tf:"fields"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	QueryScope string `json:"queryScope,omitempty" tf:"query_scope,omitempty"`
}

type FirestoreIndexStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FirestoreIndexSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirestoreIndexList is a list of FirestoreIndexs
type FirestoreIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FirestoreIndex CRD objects
	Items []FirestoreIndex `json:"items,omitempty"`
}

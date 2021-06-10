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

type Volume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec,omitempty"`
	Status            VolumeStatus `json:"status,omitempty"`
}

type VolumeSpec struct {
	VolumeSpec2 `json:",inline"`
	// +optional
	KubeformOutput *VolumeSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type VolumeSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DropletIDS []int64 `json:"dropletIDS,omitempty" tf:"droplet_ids"`
	// +optional
	FilesystemLabel *string `json:"filesystemLabel,omitempty" tf:"filesystem_label"`
	// +optional
	// Deprecated
	FilesystemType *string `json:"filesystemType,omitempty" tf:"filesystem_type"`
	// +optional
	InitialFilesystemLabel *string `json:"initialFilesystemLabel,omitempty" tf:"initial_filesystem_label"`
	// +optional
	InitialFilesystemType *string `json:"initialFilesystemType,omitempty" tf:"initial_filesystem_type"`
	Name                  *string `json:"name" tf:"name"`
	Region                *string `json:"region" tf:"region"`
	Size                  *int64  `json:"size" tf:"size"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// the uniform resource name for the volume.
	// +optional
	Urn *string `json:"urn,omitempty" tf:"urn"`
}

type VolumeStatus struct {
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

// VolumeList is a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Volume CRD objects
	Items []Volume `json:"items,omitempty"`
}

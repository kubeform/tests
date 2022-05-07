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

type StorageBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBlobSpec   `json:"spec,omitempty"`
	Status            StorageBlobStatus `json:"status,omitempty"`
}

type StorageBlobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessTier string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`
	// +optional
	// Deprecated
	Attempts int64 `json:"attempts,omitempty" tf:"attempts,omitempty"`
	// +optional
	ContentType string `json:"contentType,omitempty" tf:"content_type,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	Name     string            `json:"name" tf:"name"`
	// +optional
	Parallelism int64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`
	// +optional
	// Deprecated
	ResourceGroupName string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	// +optional
	Size int64 `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
	// +optional
	SourceContent string `json:"sourceContent,omitempty" tf:"source_content,omitempty"`
	// +optional
	SourceURI            string `json:"sourceURI,omitempty" tf:"source_uri,omitempty"`
	StorageAccountName   string `json:"storageAccountName" tf:"storage_account_name"`
	StorageContainerName string `json:"storageContainerName" tf:"storage_container_name"`
	Type                 string `json:"type" tf:"type"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
}

type StorageBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageBlobSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBlobList is a list of StorageBlobs
type StorageBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBlob CRD objects
	Items []StorageBlob `json:"items,omitempty"`
}
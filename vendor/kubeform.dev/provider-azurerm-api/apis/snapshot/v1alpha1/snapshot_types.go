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

type Snapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec,omitempty"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

type SnapshotSpec struct {
	SnapshotSpec2 `json:",inline"`
	// +optional
	KubeformOutput SnapshotSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type SnapshotSpecEncryptionSettingsDiskEncryptionKey struct {
	SecretURL     *string `json:"secretURL" tf:"secret_url"`
	SourceVaultID *string `json:"sourceVaultID" tf:"source_vault_id"`
}

type SnapshotSpecEncryptionSettingsKeyEncryptionKey struct {
	KeyURL        *string `json:"keyURL" tf:"key_url"`
	SourceVaultID *string `json:"sourceVaultID" tf:"source_vault_id"`
}

type SnapshotSpecEncryptionSettings struct {
	// +optional
	DiskEncryptionKey *SnapshotSpecEncryptionSettingsDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key"`
	Enabled           *bool                                            `json:"enabled" tf:"enabled"`
	// +optional
	KeyEncryptionKey *SnapshotSpecEncryptionSettingsKeyEncryptionKey `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key"`
}

type SnapshotSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CreateOption *string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	// +optional
	EncryptionSettings *SnapshotSpecEncryptionSettings `json:"encryptionSettings,omitempty" tf:"encryption_settings"`
	Location           *string                         `json:"location" tf:"location"`
	Name               *string                         `json:"name" tf:"name"`
	ResourceGroupName  *string                         `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SourceResourceID *string `json:"sourceResourceID,omitempty" tf:"source_resource_id"`
	// +optional
	SourceURI *string `json:"sourceURI,omitempty" tf:"source_uri"`
	// +optional
	StorageAccountID *string `json:"storageAccountID,omitempty" tf:"storage_account_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type SnapshotStatus struct {
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

// SnapshotList is a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Snapshot CRD objects
	Items []Snapshot `json:"items,omitempty"`
}
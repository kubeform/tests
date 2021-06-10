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

type CachedIscsiVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CachedIscsiVolumeSpec   `json:"spec,omitempty"`
	Status            CachedIscsiVolumeStatus `json:"status,omitempty"`
}

type CachedIscsiVolumeSpec struct {
	CachedIscsiVolumeSpec2 `json:",inline"`
	// +optional
	KubeformOutput CachedIscsiVolumeSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type CachedIscsiVolumeSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ChapEnabled *bool   `json:"chapEnabled,omitempty" tf:"chap_enabled"`
	GatewayArn  *string `json:"gatewayArn" tf:"gateway_arn"`
	// +optional
	KmsEncrypted *bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted"`
	// +optional
	KmsKey *string `json:"kmsKey,omitempty" tf:"kms_key"`
	// +optional
	LunNumber          *int64  `json:"lunNumber,omitempty" tf:"lun_number"`
	NetworkInterfaceID *string `json:"networkInterfaceID" tf:"network_interface_id"`
	// +optional
	NetworkInterfacePort *int64 `json:"networkInterfacePort,omitempty" tf:"network_interface_port"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
	// +optional
	SourceVolumeArn *string `json:"sourceVolumeArn,omitempty" tf:"source_volume_arn"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TargetArn  *string `json:"targetArn,omitempty" tf:"target_arn"`
	TargetName *string `json:"targetName" tf:"target_name"`
	// +optional
	VolumeArn *string `json:"volumeArn,omitempty" tf:"volume_arn"`
	// +optional
	VolumeID          *string `json:"volumeID,omitempty" tf:"volume_id"`
	VolumeSizeInBytes *int64  `json:"volumeSizeInBytes" tf:"volume_size_in_bytes"`
}

type CachedIscsiVolumeStatus struct {
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

// CachedIscsiVolumeList is a list of CachedIscsiVolumes
type CachedIscsiVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CachedIscsiVolume CRD objects
	Items []CachedIscsiVolume `json:"items,omitempty"`
}
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

type StorageKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageKeySpec   `json:"spec,omitempty"`
	Status            StorageKeyStatus `json:"status,omitempty"`
}

type StorageKeySpec struct {
	StorageKeySpec2 `json:",inline"`
	// +optional
	KubeformOutput *StorageKeySpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type StorageKeySpecBucketAccess struct {
	// The unique label of the bucket to which the key will grant limited access.
	BucketName *string `json:"bucketName" tf:"bucket_name"`
	// The Object Storage cluster where a bucket to which the key is granting access is hosted.
	Cluster *string `json:"cluster" tf:"cluster"`
	// This Limited Access Key’s permissions for the selected bucket.
	Permissions *string `json:"permissions" tf:"permissions"`
}

type StorageKeySpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// This keypair's access key. This is not secret.
	// +optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key"`
	// +optional
	BucketAccess []StorageKeySpecBucketAccess `json:"bucketAccess,omitempty" tf:"bucket_access"`
	// The label given to this key. For display purposes only.
	Label *string `json:"label" tf:"label"`
	// Whether or not this key is a limited access key.
	// +optional
	Limited *bool `json:"limited,omitempty" tf:"limited"`
	// This keypair's secret key.
	// +optional
	SecretKey *string `json:"-" sensitive:"true" tf:"secret_key"`
}

type StorageKeyStatus struct {
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

// StorageKeyList is a list of StorageKeys
type StorageKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageKey CRD objects
	Items []StorageKey `json:"items,omitempty"`
}

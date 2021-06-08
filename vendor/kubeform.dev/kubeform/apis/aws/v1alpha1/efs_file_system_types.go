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

type EfsFileSystem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EfsFileSystemSpec   `json:"spec,omitempty"`
	Status            EfsFileSystemStatus `json:"status,omitempty"`
}

type EfsFileSystemSpecLifecyclePolicy struct {
	// +optional
	TransitionToIa string `json:"transitionToIa,omitempty" tf:"transition_to_ia,omitempty"`
}

type EfsFileSystemSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CreationToken string `json:"creationToken,omitempty" tf:"creation_token,omitempty"`
	// +optional
	DnsName string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LifecyclePolicy []EfsFileSystemSpecLifecyclePolicy `json:"lifecyclePolicy,omitempty" tf:"lifecycle_policy,omitempty"`
	// +optional
	PerformanceMode string `json:"performanceMode,omitempty" tf:"performance_mode,omitempty"`
	// +optional
	ProvisionedThroughputInMibps float64 `json:"provisionedThroughputInMibps,omitempty" tf:"provisioned_throughput_in_mibps,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ThroughputMode string `json:"throughputMode,omitempty" tf:"throughput_mode,omitempty"`
}

type EfsFileSystemStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EfsFileSystemSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EfsFileSystemList is a list of EfsFileSystems
type EfsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EfsFileSystem CRD objects
	Items []EfsFileSystem `json:"items,omitempty"`
}
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

type ServicebusQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusQueueSpec   `json:"spec,omitempty"`
	Status            ServicebusQueueStatus `json:"status,omitempty"`
}

type ServicebusQueueSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoDeleteOnIdle string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`
	// +optional
	DeadLetteringOnMessageExpiration bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`
	// +optional
	DefaultMessageTtl string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`
	// +optional
	DuplicateDetectionHistoryTimeWindow string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`
	// +optional
	// Deprecated
	EnableBatchedOperations bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`
	// +optional
	EnableExpress bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`
	// +optional
	EnablePartitioning bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	LockDuration string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`
	// +optional
	MaxDeliveryCount int64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
	// +optional
	MaxSizeInMegabytes int64  `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`
	Name               string `json:"name" tf:"name"`
	NamespaceName      string `json:"namespaceName" tf:"namespace_name"`
	// +optional
	RequiresDuplicateDetection bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`
	// +optional
	RequiresSession   bool   `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// Deprecated
	SupportOrdering bool `json:"supportOrdering,omitempty" tf:"support_ordering,omitempty"`
}

type ServicebusQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServicebusQueueSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusQueueList is a list of ServicebusQueues
type ServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusQueue CRD objects
	Items []ServicebusQueue `json:"items,omitempty"`
}
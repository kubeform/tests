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

type RedisInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisInstanceSpec   `json:"spec,omitempty"`
	Status            RedisInstanceStatus `json:"status,omitempty"`
}

type RedisInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AlternativeLocationID string `json:"alternativeLocationID,omitempty" tf:"alternative_location_id,omitempty"`
	// +optional
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`
	// +optional
	CreateTime string `json:"createTime,omitempty" tf:"create_time,omitempty"`
	// +optional
	CurrentLocationID string `json:"currentLocationID,omitempty" tf:"current_location_id,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LocationID   string `json:"locationID,omitempty" tf:"location_id,omitempty"`
	MemorySizeGb int64  `json:"memorySizeGb" tf:"memory_size_gb"`
	Name         string `json:"name" tf:"name"`
	// +optional
	Port int64 `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	RedisConfigs map[string]string `json:"redisConfigs,omitempty" tf:"redis_configs,omitempty"`
	// +optional
	RedisVersion string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	ReservedIPRange string `json:"reservedIPRange,omitempty" tf:"reserved_ip_range,omitempty"`
	// +optional
	Tier string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type RedisInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RedisInstanceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedisInstanceList is a list of RedisInstances
type RedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedisInstance CRD objects
	Items []RedisInstance `json:"items,omitempty"`
}

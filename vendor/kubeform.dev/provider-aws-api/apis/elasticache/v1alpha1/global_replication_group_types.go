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

type GlobalReplicationGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalReplicationGroupSpec   `json:"spec,omitempty"`
	Status            GlobalReplicationGroupStatus `json:"status,omitempty"`
}

type GlobalReplicationGroupSpec struct {
	GlobalReplicationGroupSpec2 `json:",inline"`
	// +optional
	KubeformOutput GlobalReplicationGroupSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type GlobalReplicationGroupSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActualEngineVersion *string `json:"actualEngineVersion,omitempty" tf:"actual_engine_version"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled"`
	// +optional
	AuthTokenEnabled *bool `json:"authTokenEnabled,omitempty" tf:"auth_token_enabled"`
	// +optional
	CacheNodeType *string `json:"cacheNodeType,omitempty" tf:"cache_node_type"`
	// +optional
	ClusterEnabled *bool `json:"clusterEnabled,omitempty" tf:"cluster_enabled"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty" tf:"global_replication_group_description"`
	// +optional
	GlobalReplicationGroupID       *string `json:"globalReplicationGroupID,omitempty" tf:"global_replication_group_id"`
	GlobalReplicationGroupIDSuffix *string `json:"globalReplicationGroupIDSuffix" tf:"global_replication_group_id_suffix"`
	PrimaryReplicationGroupID      *string `json:"primaryReplicationGroupID" tf:"primary_replication_group_id"`
	// +optional
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled"`
}

type GlobalReplicationGroupStatus struct {
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

// GlobalReplicationGroupList is a list of GlobalReplicationGroups
type GlobalReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlobalReplicationGroup CRD objects
	Items []GlobalReplicationGroup `json:"items,omitempty"`
}

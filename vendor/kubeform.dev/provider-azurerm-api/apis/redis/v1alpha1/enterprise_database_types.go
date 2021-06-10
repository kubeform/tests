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

type EnterpriseDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnterpriseDatabaseSpec   `json:"spec,omitempty"`
	Status            EnterpriseDatabaseStatus `json:"status,omitempty"`
}

type EnterpriseDatabaseSpec struct {
	EnterpriseDatabaseSpec2 `json:",inline"`
	// +optional
	KubeformOutput EnterpriseDatabaseSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type EnterpriseDatabaseSpecModule struct {
	// +optional
	Args *string `json:"args,omitempty" tf:"args"`
	Name *string `json:"name" tf:"name"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type EnterpriseDatabaseSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	ClientProtocol *string `json:"clientProtocol,omitempty" tf:"client_protocol"`
	ClusterID      *string `json:"clusterID" tf:"cluster_id"`
	// +optional
	ClusteringPolicy *string `json:"clusteringPolicy,omitempty" tf:"clustering_policy"`
	// +optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	Module []EnterpriseDatabaseSpecModule `json:"module,omitempty" tf:"module"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PrimaryAccessKey  *string `json:"-" sensitive:"true" tf:"primary_access_key"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey *string `json:"-" sensitive:"true" tf:"secondary_access_key"`
}

type EnterpriseDatabaseStatus struct {
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

// EnterpriseDatabaseList is a list of EnterpriseDatabases
type EnterpriseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EnterpriseDatabase CRD objects
	Items []EnterpriseDatabase `json:"items,omitempty"`
}

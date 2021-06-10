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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpec struct {
	ClusterSpec2 `json:",inline"`
	// +optional
	KubeformOutput ClusterSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ClusterSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type ClusterSpecOptimizedAutoScale struct {
	MaximumInstances *int64 `json:"maximumInstances" tf:"maximum_instances"`
	MinimumInstances *int64 `json:"minimumInstances" tf:"minimum_instances"`
}

type ClusterSpecSku struct {
	// +optional
	Capacity *int64  `json:"capacity,omitempty" tf:"capacity"`
	Name     *string `json:"name" tf:"name"`
}

type ClusterSpecVirtualNetworkConfiguration struct {
	DataManagementPublicIPID *string `json:"dataManagementPublicIPID" tf:"data_management_public_ip_id"`
	EnginePublicIPID         *string `json:"enginePublicIPID" tf:"engine_public_ip_id"`
	SubnetID                 *string `json:"subnetID" tf:"subnet_id"`
}

type ClusterSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DataIngestionURI *string `json:"dataIngestionURI,omitempty" tf:"data_ingestion_uri"`
	// +optional
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty" tf:"double_encryption_enabled"`
	// +optional
	EnableDiskEncryption *bool `json:"enableDiskEncryption,omitempty" tf:"enable_disk_encryption"`
	// +optional
	EnablePurge *bool `json:"enablePurge,omitempty" tf:"enable_purge"`
	// +optional
	EnableStreamingIngest *bool `json:"enableStreamingIngest,omitempty" tf:"enable_streaming_ingest"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	Identity *ClusterSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	LanguageExtensions []string `json:"languageExtensions,omitempty" tf:"language_extensions"`
	Location           *string  `json:"location" tf:"location"`
	Name               *string  `json:"name" tf:"name"`
	// +optional
	OptimizedAutoScale *ClusterSpecOptimizedAutoScale `json:"optimizedAutoScale,omitempty" tf:"optimized_auto_scale"`
	ResourceGroupName  *string                        `json:"resourceGroupName" tf:"resource_group_name"`
	Sku                *ClusterSpecSku                `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TrustedExternalTenants []string `json:"trustedExternalTenants,omitempty" tf:"trusted_external_tenants"`
	// +optional
	Uri *string `json:"uri,omitempty" tf:"uri"`
	// +optional
	VirtualNetworkConfiguration *ClusterSpecVirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}

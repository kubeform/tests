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

type ServiceEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceEnvironmentSpec   `json:"spec,omitempty"`
	Status            ServiceEnvironmentStatus `json:"status,omitempty"`
}

type ServiceEnvironmentSpec struct {
	ServiceEnvironmentSpec2 `json:",inline"`
	// +optional
	KubeformOutput ServiceEnvironmentSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ServiceEnvironmentSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccessEndpointType *string `json:"accessEndpointType" tf:"access_endpoint_type"`
	// +optional
	ConnectorEndpointIPAddresses []string `json:"connectorEndpointIPAddresses,omitempty" tf:"connector_endpoint_ip_addresses"`
	// +optional
	ConnectorOutboundIPAddresses []string `json:"connectorOutboundIPAddresses,omitempty" tf:"connector_outbound_ip_addresses"`
	Location                     *string  `json:"location" tf:"location"`
	Name                         *string  `json:"name" tf:"name"`
	ResourceGroupName            *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:MinItems=4
	VirtualNetworkSubnetIDS []string `json:"virtualNetworkSubnetIDS" tf:"virtual_network_subnet_ids"`
	// +optional
	WorkflowEndpointIPAddresses []string `json:"workflowEndpointIPAddresses,omitempty" tf:"workflow_endpoint_ip_addresses"`
	// +optional
	WorkflowOutboundIPAddresses []string `json:"workflowOutboundIPAddresses,omitempty" tf:"workflow_outbound_ip_addresses"`
}

type ServiceEnvironmentStatus struct {
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

// ServiceEnvironmentList is a list of ServiceEnvironments
type ServiceEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceEnvironment CRD objects
	Items []ServiceEnvironment `json:"items,omitempty"`
}

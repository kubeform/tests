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

type NetworkRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkRuleCollectionSpec   `json:"spec,omitempty"`
	Status            NetworkRuleCollectionStatus `json:"status,omitempty"`
}

type NetworkRuleCollectionSpec struct {
	NetworkRuleCollectionSpec2 `json:",inline"`
	// +optional
	KubeformOutput NetworkRuleCollectionSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type NetworkRuleCollectionSpecRule struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DestinationAddresses []string `json:"destinationAddresses,omitempty" tf:"destination_addresses"`
	// +optional
	DestinationFqdns []string `json:"destinationFqdns,omitempty" tf:"destination_fqdns"`
	// +optional
	DestinationIPGroups []string `json:"destinationIPGroups,omitempty" tf:"destination_ip_groups"`
	DestinationPorts    []string `json:"destinationPorts" tf:"destination_ports"`
	Name                *string  `json:"name" tf:"name"`
	Protocols           []string `json:"protocols" tf:"protocols"`
	// +optional
	SourceAddresses []string `json:"sourceAddresses,omitempty" tf:"source_addresses"`
	// +optional
	SourceIPGroups []string `json:"sourceIPGroups,omitempty" tf:"source_ip_groups"`
}

type NetworkRuleCollectionSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action            *string `json:"action" tf:"action"`
	AzureFirewallName *string `json:"azureFirewallName" tf:"azure_firewall_name"`
	Name              *string `json:"name" tf:"name"`
	Priority          *int64  `json:"priority" tf:"priority"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	Rule []NetworkRuleCollectionSpecRule `json:"rule" tf:"rule"`
}

type NetworkRuleCollectionStatus struct {
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

// NetworkRuleCollectionList is a list of NetworkRuleCollections
type NetworkRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkRuleCollection CRD objects
	Items []NetworkRuleCollection `json:"items,omitempty"`
}

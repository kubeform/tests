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

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpec struct {
	FirewallSpec2 `json:",inline"`
	// +optional
	KubeformOutput FirewallSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type FirewallSpecIpConfiguration struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress  *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	PublicIPAddressID *string `json:"publicIPAddressID" tf:"public_ip_address_id"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type FirewallSpecManagementIPConfiguration struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress  *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	PublicIPAddressID *string `json:"publicIPAddressID" tf:"public_ip_address_id"`
	SubnetID          *string `json:"subnetID" tf:"subnet_id"`
}

type FirewallSpecVirtualHub struct {
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses"`
	// +optional
	PublicIPCount *int64  `json:"publicIPCount,omitempty" tf:"public_ip_count"`
	VirtualHubID  *string `json:"virtualHubID" tf:"virtual_hub_id"`
}

type FirewallSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	DnsServers []string `json:"dnsServers,omitempty" tf:"dns_servers"`
	// +optional
	FirewallPolicyID *string `json:"firewallPolicyID,omitempty" tf:"firewall_policy_id"`
	// +optional
	IpConfiguration []FirewallSpecIpConfiguration `json:"ipConfiguration,omitempty" tf:"ip_configuration"`
	Location        *string                       `json:"location" tf:"location"`
	// +optional
	ManagementIPConfiguration *FirewallSpecManagementIPConfiguration `json:"managementIPConfiguration,omitempty" tf:"management_ip_configuration"`
	Name                      *string                                `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	PrivateIPRanges   []string `json:"privateIPRanges,omitempty" tf:"private_ip_ranges"`
	ResourceGroupName *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name"`
	// +optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	ThreatIntelMode *string `json:"threatIntelMode,omitempty" tf:"threat_intel_mode"`
	// +optional
	VirtualHub *FirewallSpecVirtualHub `json:"virtualHub,omitempty" tf:"virtual_hub"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type FirewallStatus struct {
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

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}
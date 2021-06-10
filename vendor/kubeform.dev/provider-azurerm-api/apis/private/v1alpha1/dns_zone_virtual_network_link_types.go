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

type DnsZoneVirtualNetworkLink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsZoneVirtualNetworkLinkSpec   `json:"spec,omitempty"`
	Status            DnsZoneVirtualNetworkLinkStatus `json:"status,omitempty"`
}

type DnsZoneVirtualNetworkLinkSpec struct {
	DnsZoneVirtualNetworkLinkSpec2 `json:",inline"`
	// +optional
	KubeformOutput DnsZoneVirtualNetworkLinkSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type DnsZoneVirtualNetworkLinkSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name               *string `json:"name" tf:"name"`
	PrivateDNSZoneName *string `json:"privateDNSZoneName" tf:"private_dns_zone_name"`
	// +optional
	RegistrationEnabled *bool   `json:"registrationEnabled,omitempty" tf:"registration_enabled"`
	ResourceGroupName   *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags             *map[string]string `json:"tags,omitempty" tf:"tags"`
	VirtualNetworkID *string            `json:"virtualNetworkID" tf:"virtual_network_id"`
}

type DnsZoneVirtualNetworkLinkStatus struct {
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

// DnsZoneVirtualNetworkLinkList is a list of DnsZoneVirtualNetworkLinks
type DnsZoneVirtualNetworkLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsZoneVirtualNetworkLink CRD objects
	Items []DnsZoneVirtualNetworkLink `json:"items,omitempty"`
}

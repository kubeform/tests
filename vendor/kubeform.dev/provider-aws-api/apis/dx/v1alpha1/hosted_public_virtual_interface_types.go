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

type HostedPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedPublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            HostedPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type HostedPublicVirtualInterfaceSpec struct {
	HostedPublicVirtualInterfaceSpec2 `json:",inline"`
	// +optional
	KubeformOutput HostedPublicVirtualInterfaceSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type HostedPublicVirtualInterfaceSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressFamily *string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`
	// +optional
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device"`
	BgpAsn    *int64  `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey   *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`
	ConnectionID *string `json:"connectionID" tf:"connection_id"`
	// +optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`
	Name            *string `json:"name" tf:"name"`
	OwnerAccountID  *string `json:"ownerAccountID" tf:"owner_account_id"`
	// +kubebuilder:validation:MinItems=1
	RouteFilterPrefixes []string `json:"routeFilterPrefixes" tf:"route_filter_prefixes"`
	Vlan                *int64   `json:"vlan" tf:"vlan"`
}

type HostedPublicVirtualInterfaceStatus struct {
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

// HostedPublicVirtualInterfaceList is a list of HostedPublicVirtualInterfaces
type HostedPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HostedPublicVirtualInterface CRD objects
	Items []HostedPublicVirtualInterface `json:"items,omitempty"`
}

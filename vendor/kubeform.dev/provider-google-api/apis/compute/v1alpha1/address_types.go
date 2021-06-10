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

type Address struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddressSpec   `json:"spec,omitempty"`
	Status            AddressStatus `json:"status,omitempty"`
}

type AddressSpec struct {
	AddressSpec2 `json:",inline"`
	// +optional
	KubeformOutput AddressSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type AddressSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// The type of address to reserve. Default value: "EXTERNAL" Possible values: ["INTERNAL", "EXTERNAL"]
	// +optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type"`
	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp"`
	// An optional description of this resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM. Possible values: ["PREMIUM", "STANDARD"]
	// +optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The purpose of this resource. Possible values include:
	//
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	//
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
	//
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	// +optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose"`
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	// +optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork"`
	// The URLs of the resources that are using this address.
	// +optional
	Users []string `json:"users,omitempty" tf:"users"`
}

type AddressStatus struct {
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

// AddressList is a list of Addresss
type AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Address CRD objects
	Items []Address `json:"items,omitempty"`
}

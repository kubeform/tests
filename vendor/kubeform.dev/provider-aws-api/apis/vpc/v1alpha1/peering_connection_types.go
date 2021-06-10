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

type PeeringConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PeeringConnectionSpec   `json:"spec,omitempty"`
	Status            PeeringConnectionStatus `json:"status,omitempty"`
}

type PeeringConnectionSpec struct {
	PeeringConnectionSpec2 `json:",inline"`
	// +optional
	KubeformOutput PeeringConnectionSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type PeeringConnectionSpecAccepter struct {
	// +optional
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc"`
	// +optional
	AllowRemoteVpcDNSResolution *bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution"`
	// +optional
	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link"`
}

type PeeringConnectionSpecRequester struct {
	// +optional
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc"`
	// +optional
	AllowRemoteVpcDNSResolution *bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution"`
	// +optional
	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link"`
}

type PeeringConnectionSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AcceptStatus *string `json:"acceptStatus,omitempty" tf:"accept_status"`
	// +optional
	Accepter *PeeringConnectionSpecAccepter `json:"accepter,omitempty" tf:"accepter"`
	// +optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept"`
	// +optional
	PeerOwnerID *string `json:"peerOwnerID,omitempty" tf:"peer_owner_id"`
	// +optional
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region"`
	PeerVpcID  *string `json:"peerVpcID" tf:"peer_vpc_id"`
	// +optional
	Requester *PeeringConnectionSpecRequester `json:"requester,omitempty" tf:"requester"`
	// +optional
	Tags  *map[string]string `json:"tags,omitempty" tf:"tags"`
	VpcID *string            `json:"vpcID" tf:"vpc_id"`
}

type PeeringConnectionStatus struct {
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

// PeeringConnectionList is a list of PeeringConnections
type PeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PeeringConnection CRD objects
	Items []PeeringConnection `json:"items,omitempty"`
}

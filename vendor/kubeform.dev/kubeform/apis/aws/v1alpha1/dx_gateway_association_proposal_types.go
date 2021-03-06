/*
Copyright The Kubeform Authors.

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
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewayAssociationProposalSpec   `json:"spec,omitempty"`
	Status            DxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

type DxGatewayAssociationProposalSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowedPrefixes []string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`
	// +optional
	AssociatedGatewayID string `json:"associatedGatewayID,omitempty" tf:"associated_gateway_id,omitempty"`
	// +optional
	AssociatedGatewayOwnerAccountID string `json:"associatedGatewayOwnerAccountID,omitempty" tf:"associated_gateway_owner_account_id,omitempty"`
	// +optional
	AssociatedGatewayType   string `json:"associatedGatewayType,omitempty" tf:"associated_gateway_type,omitempty"`
	DxGatewayID             string `json:"dxGatewayID" tf:"dx_gateway_id"`
	DxGatewayOwnerAccountID string `json:"dxGatewayOwnerAccountID" tf:"dx_gateway_owner_account_id"`
	// +optional
	// Deprecated
	VpnGatewayID string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type DxGatewayAssociationProposalStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxGatewayAssociationProposalSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxGatewayAssociationProposalList is a list of DxGatewayAssociationProposals
type DxGatewayAssociationProposalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxGatewayAssociationProposal CRD objects
	Items []DxGatewayAssociationProposal `json:"items,omitempty"`
}

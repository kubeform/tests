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

type Group struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec,omitempty"`
	Status            GroupStatus `json:"status,omitempty"`
}

type GroupSpec struct {
	GroupSpec2 `json:",inline"`
	// +optional
	KubeformOutput GroupSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type GroupSpecEgress struct {
	// +optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	FromPort    *int64  `json:"fromPort" tf:"from_port"`
	// +optional
	Ipv6CIDRBlocks []string `json:"ipv6CIDRBlocks,omitempty" tf:"ipv6_cidr_blocks"`
	// +optional
	PrefixListIDS []string `json:"prefixListIDS,omitempty" tf:"prefix_list_ids"`
	Protocol      *string  `json:"protocol" tf:"protocol"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	Self   *bool  `json:"self,omitempty" tf:"self"`
	ToPort *int64 `json:"toPort" tf:"to_port"`
}

type GroupSpecIngress struct {
	// +optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	FromPort    *int64  `json:"fromPort" tf:"from_port"`
	// +optional
	Ipv6CIDRBlocks []string `json:"ipv6CIDRBlocks,omitempty" tf:"ipv6_cidr_blocks"`
	// +optional
	PrefixListIDS []string `json:"prefixListIDS,omitempty" tf:"prefix_list_ids"`
	Protocol      *string  `json:"protocol" tf:"protocol"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	// +optional
	Self   *bool  `json:"self,omitempty" tf:"self"`
	ToPort *int64 `json:"toPort" tf:"to_port"`
}

type GroupSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Egress []GroupSpecEgress `json:"egress,omitempty" tf:"egress"`
	// +optional
	Ingress []GroupSpecIngress `json:"ingress,omitempty" tf:"ingress"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	RevokeRulesOnDelete *bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type GroupStatus struct {
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

// GroupList is a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Group CRD objects
	Items []Group `json:"items,omitempty"`
}
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

type AccessPoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessPointSpec   `json:"spec,omitempty"`
	Status            AccessPointStatus `json:"status,omitempty"`
}

type AccessPointSpec struct {
	AccessPointSpec2 `json:",inline"`
	// +optional
	KubeformOutput AccessPointSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type AccessPointSpecPublicAccessBlockConfiguration struct {
	// +optional
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls"`
	// +optional
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy"`
	// +optional
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls"`
	// +optional
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets"`
}

type AccessPointSpecVpcConfiguration struct {
	VpcID *string `json:"vpcID" tf:"vpc_id"`
}

type AccessPointSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// +optional
	Arn    *string `json:"arn,omitempty" tf:"arn"`
	Bucket *string `json:"bucket" tf:"bucket"`
	// +optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name"`
	// +optional
	HasPublicAccessPolicy *bool   `json:"hasPublicAccessPolicy,omitempty" tf:"has_public_access_policy"`
	Name                  *string `json:"name" tf:"name"`
	// +optional
	NetworkOrigin *string `json:"networkOrigin,omitempty" tf:"network_origin"`
	// +optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
	// +optional
	PublicAccessBlockConfiguration *AccessPointSpecPublicAccessBlockConfiguration `json:"publicAccessBlockConfiguration,omitempty" tf:"public_access_block_configuration"`
	// +optional
	VpcConfiguration *AccessPointSpecVpcConfiguration `json:"vpcConfiguration,omitempty" tf:"vpc_configuration"`
}

type AccessPointStatus struct {
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

// AccessPointList is a list of AccessPoints
type AccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AccessPoint CRD objects
	Items []AccessPoint `json:"items,omitempty"`
}
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

type Organization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationSpec   `json:"spec,omitempty"`
	Status            OrganizationStatus `json:"status,omitempty"`
}

type OrganizationSpec struct {
	OrganizationSpec2 `json:",inline"`
	// +optional
	KubeformOutput OrganizationSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type OrganizationSpecAccounts struct {
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type OrganizationSpecNonMasterAccounts struct {
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type OrganizationSpecRootsPolicyTypes struct {
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type OrganizationSpecRoots struct {
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PolicyTypes []OrganizationSpecRootsPolicyTypes `json:"policyTypes,omitempty" tf:"policy_types"`
}

type OrganizationSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Accounts []OrganizationSpecAccounts `json:"accounts,omitempty" tf:"accounts"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AwsServiceAccessPrincipals []string `json:"awsServiceAccessPrincipals,omitempty" tf:"aws_service_access_principals"`
	// +optional
	EnabledPolicyTypes []string `json:"enabledPolicyTypes,omitempty" tf:"enabled_policy_types"`
	// +optional
	FeatureSet *string `json:"featureSet,omitempty" tf:"feature_set"`
	// +optional
	MasterAccountArn *string `json:"masterAccountArn,omitempty" tf:"master_account_arn"`
	// +optional
	MasterAccountEmail *string `json:"masterAccountEmail,omitempty" tf:"master_account_email"`
	// +optional
	MasterAccountID *string `json:"masterAccountID,omitempty" tf:"master_account_id"`
	// +optional
	NonMasterAccounts []OrganizationSpecNonMasterAccounts `json:"nonMasterAccounts,omitempty" tf:"non_master_accounts"`
	// +optional
	Roots []OrganizationSpecRoots `json:"roots,omitempty" tf:"roots"`
}

type OrganizationStatus struct {
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

// OrganizationList is a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Organization CRD objects
	Items []Organization `json:"items,omitempty"`
}

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

type OrganizationsAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsAccountSpec   `json:"spec,omitempty"`
	Status            OrganizationsAccountStatus `json:"status,omitempty"`
}

type OrganizationsAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn   string `json:"arn,omitempty" tf:"arn,omitempty"`
	Email string `json:"email" tf:"email"`
	// +optional
	IamUserAccessToBilling string `json:"iamUserAccessToBilling,omitempty" tf:"iam_user_access_to_billing,omitempty"`
	// +optional
	JoinedMethod string `json:"joinedMethod,omitempty" tf:"joined_method,omitempty"`
	// +optional
	JoinedTimestamp string `json:"joinedTimestamp,omitempty" tf:"joined_timestamp,omitempty"`
	Name            string `json:"name" tf:"name"`
	// +optional
	ParentID string `json:"parentID,omitempty" tf:"parent_id,omitempty"`
	// +optional
	RoleName string `json:"roleName,omitempty" tf:"role_name,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OrganizationsAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OrganizationsAccountSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationsAccountList is a list of OrganizationsAccounts
type OrganizationsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationsAccount CRD objects
	Items []OrganizationsAccount `json:"items,omitempty"`
}
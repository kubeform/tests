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

type IamCustomRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamCustomRoleSpec   `json:"spec,omitempty"`
	Status            IamCustomRoleStatus `json:"status,omitempty"`
}

type IamCustomRoleSpec struct {
	IamCustomRoleSpec2 `json:",inline"`
	// +optional
	KubeformOutput IamCustomRoleSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type IamCustomRoleSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The current deleted state of the role.
	// +optional
	Deleted *bool `json:"deleted,omitempty" tf:"deleted"`
	// A human-readable description for the role.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The name of the role in the format organizations/{{org_id}}/roles/{{role_id}}. Like id, this field can be used as a reference in other resources such as IAM role bindings.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgID *string `json:"orgID" tf:"org_id"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +kubebuilder:validation:MinItems=1
	Permissions []string `json:"permissions" tf:"permissions"`
	// The role id to use for this role.
	RoleID *string `json:"roleID" tf:"role_id"`
	// The current launch stage of the role. Defaults to GA.
	// +optional
	Stage *string `json:"stage,omitempty" tf:"stage"`
	// A human-readable title for the role.
	Title *string `json:"title" tf:"title"`
}

type IamCustomRoleStatus struct {
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

// IamCustomRoleList is a list of IamCustomRoles
type IamCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamCustomRole CRD objects
	Items []IamCustomRole `json:"items,omitempty"`
}
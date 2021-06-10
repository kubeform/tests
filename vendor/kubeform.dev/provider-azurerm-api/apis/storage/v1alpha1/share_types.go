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

type Share struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShareSpec   `json:"spec,omitempty"`
	Status            ShareStatus `json:"status,omitempty"`
}

type ShareSpec struct {
	ShareSpec2 `json:",inline"`
	// +optional
	KubeformOutput ShareSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ShareSpecAclAccessPolicy struct {
	// +optional
	Expiry      *string `json:"expiry,omitempty" tf:"expiry"`
	Permissions *string `json:"permissions" tf:"permissions"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type ShareSpecAcl struct {
	// +optional
	AccessPolicy []ShareSpecAclAccessPolicy `json:"accessPolicy,omitempty" tf:"access_policy"`
	ID           *string                    `json:"ID" tf:"id"`
}

type ShareSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Acl []ShareSpecAcl `json:"acl,omitempty" tf:"acl"`
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
	Name     *string            `json:"name" tf:"name"`
	// +optional
	Quota *int64 `json:"quota,omitempty" tf:"quota"`
	// +optional
	ResourceManagerID  *string `json:"resourceManagerID,omitempty" tf:"resource_manager_id"`
	StorageAccountName *string `json:"storageAccountName" tf:"storage_account_name"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type ShareStatus struct {
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

// ShareList is a list of Shares
type ShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Share CRD objects
	Items []Share `json:"items,omitempty"`
}

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

type CacheAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CacheAccessPolicySpec   `json:"spec,omitempty"`
	Status            CacheAccessPolicyStatus `json:"status,omitempty"`
}

type CacheAccessPolicySpec struct {
	CacheAccessPolicySpec2 `json:",inline"`
	// +optional
	KubeformOutput CacheAccessPolicySpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type CacheAccessPolicySpecAccessRule struct {
	Access *string `json:"access" tf:"access"`
	// +optional
	AnonymousGid *int64 `json:"anonymousGid,omitempty" tf:"anonymous_gid"`
	// +optional
	AnonymousUid *int64 `json:"anonymousUid,omitempty" tf:"anonymous_uid"`
	// +optional
	Filter *string `json:"filter,omitempty" tf:"filter"`
	// +optional
	RootSquashEnabled *bool   `json:"rootSquashEnabled,omitempty" tf:"root_squash_enabled"`
	Scope             *string `json:"scope" tf:"scope"`
	// +optional
	SubmountAccessEnabled *bool `json:"submountAccessEnabled,omitempty" tf:"submount_access_enabled"`
	// +optional
	SuidEnabled *bool `json:"suidEnabled,omitempty" tf:"suid_enabled"`
}

type CacheAccessPolicySpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=3
	// +kubebuilder:validation:MinItems=1
	AccessRule []CacheAccessPolicySpecAccessRule `json:"accessRule" tf:"access_rule"`
	HpcCacheID *string                           `json:"hpcCacheID" tf:"hpc_cache_id"`
	Name       *string                           `json:"name" tf:"name"`
}

type CacheAccessPolicyStatus struct {
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

// CacheAccessPolicyList is a list of CacheAccessPolicys
type CacheAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CacheAccessPolicy CRD objects
	Items []CacheAccessPolicy `json:"items,omitempty"`
}

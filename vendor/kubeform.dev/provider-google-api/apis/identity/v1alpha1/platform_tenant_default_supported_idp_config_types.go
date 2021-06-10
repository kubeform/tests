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

type PlatformTenantDefaultSupportedIdpConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlatformTenantDefaultSupportedIdpConfigSpec   `json:"spec,omitempty"`
	Status            PlatformTenantDefaultSupportedIdpConfigStatus `json:"status,omitempty"`
}

type PlatformTenantDefaultSupportedIdpConfigSpec struct {
	PlatformTenantDefaultSupportedIdpConfigSpec2 `json:",inline"`
	// +optional
	KubeformOutput PlatformTenantDefaultSupportedIdpConfigSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type PlatformTenantDefaultSupportedIdpConfigSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// OAuth client ID
	ClientID *string `json:"clientID" tf:"client_id"`
	// OAuth client secret
	ClientSecret *string `json:"clientSecret" tf:"client_secret"`
	// If this IDP allows the user to sign in
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// ID of the IDP. Possible values include:
	//
	// * 'apple.com'
	//
	// * 'facebook.com'
	//
	// * 'gc.apple.com'
	//
	// * 'github.com'
	//
	// * 'google.com'
	//
	// * 'linkedin.com'
	//
	// * 'microsoft.com'
	//
	// * 'playgames.google.com'
	//
	// * 'twitter.com'
	//
	// * 'yahoo.com'
	IdpID *string `json:"idpID" tf:"idp_id"`
	// The name of the default supported IDP config resource
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The name of the tenant where this DefaultSupportedIdpConfig resource exists
	Tenant *string `json:"tenant" tf:"tenant"`
}

type PlatformTenantDefaultSupportedIdpConfigStatus struct {
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

// PlatformTenantDefaultSupportedIdpConfigList is a list of PlatformTenantDefaultSupportedIdpConfigs
type PlatformTenantDefaultSupportedIdpConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PlatformTenantDefaultSupportedIdpConfig CRD objects
	Items []PlatformTenantDefaultSupportedIdpConfig `json:"items,omitempty"`
}

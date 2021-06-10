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

type OriginRequestPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OriginRequestPolicySpec   `json:"spec,omitempty"`
	Status            OriginRequestPolicyStatus `json:"status,omitempty"`
}

type OriginRequestPolicySpec struct {
	OriginRequestPolicySpec2 `json:",inline"`
	// +optional
	KubeformOutput OriginRequestPolicySpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type OriginRequestPolicySpecCookiesConfigCookies struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type OriginRequestPolicySpecCookiesConfig struct {
	CookieBehavior *string `json:"cookieBehavior" tf:"cookie_behavior"`
	// +optional
	Cookies *OriginRequestPolicySpecCookiesConfigCookies `json:"cookies,omitempty" tf:"cookies"`
}

type OriginRequestPolicySpecHeadersConfigHeaders struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type OriginRequestPolicySpecHeadersConfig struct {
	// +optional
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior"`
	// +optional
	Headers *OriginRequestPolicySpecHeadersConfigHeaders `json:"headers,omitempty" tf:"headers"`
}

type OriginRequestPolicySpecQueryStringsConfigQueryStrings struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type OriginRequestPolicySpecQueryStringsConfig struct {
	QueryStringBehavior *string `json:"queryStringBehavior" tf:"query_string_behavior"`
	// +optional
	QueryStrings *OriginRequestPolicySpecQueryStringsConfigQueryStrings `json:"queryStrings,omitempty" tf:"query_strings"`
}

type OriginRequestPolicySpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Comment       *string                               `json:"comment,omitempty" tf:"comment"`
	CookiesConfig *OriginRequestPolicySpecCookiesConfig `json:"cookiesConfig" tf:"cookies_config"`
	// +optional
	Etag               *string                                    `json:"etag,omitempty" tf:"etag"`
	HeadersConfig      *OriginRequestPolicySpecHeadersConfig      `json:"headersConfig" tf:"headers_config"`
	Name               *string                                    `json:"name" tf:"name"`
	QueryStringsConfig *OriginRequestPolicySpecQueryStringsConfig `json:"queryStringsConfig" tf:"query_strings_config"`
}

type OriginRequestPolicyStatus struct {
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

// OriginRequestPolicyList is a list of OriginRequestPolicys
type OriginRequestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OriginRequestPolicy CRD objects
	Items []OriginRequestPolicy `json:"items,omitempty"`
}
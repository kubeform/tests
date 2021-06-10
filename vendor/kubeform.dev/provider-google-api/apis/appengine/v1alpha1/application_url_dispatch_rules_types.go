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

type ApplicationURLDispatchRules struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationURLDispatchRulesSpec   `json:"spec,omitempty"`
	Status            ApplicationURLDispatchRulesStatus `json:"status,omitempty"`
}

type ApplicationURLDispatchRulesSpec struct {
	ApplicationURLDispatchRulesSpec2 `json:",inline"`
	// +optional
	KubeformOutput ApplicationURLDispatchRulesSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ApplicationURLDispatchRulesSpecDispatchRules struct {
	// Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
	// Defaults to matching all domains: "*".
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	Path *string `json:"path" tf:"path"`
	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	Service *string `json:"service" tf:"service"`
}

type ApplicationURLDispatchRulesSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Rules to match an HTTP request and dispatch that request to a service.
	DispatchRules []ApplicationURLDispatchRulesSpecDispatchRules `json:"dispatchRules" tf:"dispatch_rules"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type ApplicationURLDispatchRulesStatus struct {
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

// ApplicationURLDispatchRulesList is a list of ApplicationURLDispatchRuless
type ApplicationURLDispatchRulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationURLDispatchRules CRD objects
	Items []ApplicationURLDispatchRules `json:"items,omitempty"`
}
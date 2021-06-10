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

type UsagePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UsagePlanSpec   `json:"spec,omitempty"`
	Status            UsagePlanStatus `json:"status,omitempty"`
}

type UsagePlanSpec struct {
	UsagePlanSpec2 `json:",inline"`
	// +optional
	KubeformOutput UsagePlanSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type UsagePlanSpecApiStages struct {
	ApiID *string `json:"apiID" tf:"api_id"`
	Stage *string `json:"stage" tf:"stage"`
}

type UsagePlanSpecQuotaSettings struct {
	Limit *int64 `json:"limit" tf:"limit"`
	// +optional
	Offset *int64  `json:"offset,omitempty" tf:"offset"`
	Period *string `json:"period" tf:"period"`
}

type UsagePlanSpecThrottleSettings struct {
	// +optional
	BurstLimit *int64 `json:"burstLimit,omitempty" tf:"burst_limit"`
	// +optional
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit"`
}

type UsagePlanSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApiStages []UsagePlanSpecApiStages `json:"apiStages,omitempty" tf:"api_stages"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	// +optional
	ProductCode *string `json:"productCode,omitempty" tf:"product_code"`
	// +optional
	QuotaSettings *UsagePlanSpecQuotaSettings `json:"quotaSettings,omitempty" tf:"quota_settings"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	ThrottleSettings *UsagePlanSpecThrottleSettings `json:"throttleSettings,omitempty" tf:"throttle_settings"`
}

type UsagePlanStatus struct {
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

// UsagePlanList is a list of UsagePlans
type UsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of UsagePlan CRD objects
	Items []UsagePlan `json:"items,omitempty"`
}

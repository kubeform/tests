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

type Profile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProfileSpec   `json:"spec,omitempty"`
	Status            ProfileStatus `json:"status,omitempty"`
}

type ProfileSpec struct {
	ProfileSpec2 `json:",inline"`
	// +optional
	KubeformOutput ProfileSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ProfileSpecDnsConfig struct {
	RelativeName *string `json:"relativeName" tf:"relative_name"`
	Ttl          *int64  `json:"ttl" tf:"ttl"`
}

type ProfileSpecMonitorConfigCustomHeader struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type ProfileSpecMonitorConfig struct {
	// +optional
	CustomHeader []ProfileSpecMonitorConfigCustomHeader `json:"customHeader,omitempty" tf:"custom_header"`
	// +optional
	ExpectedStatusCodeRanges []string `json:"expectedStatusCodeRanges,omitempty" tf:"expected_status_code_ranges"`
	// +optional
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds"`
	// +optional
	Path     *string `json:"path,omitempty" tf:"path"`
	Port     *int64  `json:"port" tf:"port"`
	Protocol *string `json:"protocol" tf:"protocol"`
	// +optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`
	// +optional
	ToleratedNumberOfFailures *int64 `json:"toleratedNumberOfFailures,omitempty" tf:"tolerated_number_of_failures"`
}

type ProfileSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DnsConfig *ProfileSpecDnsConfig `json:"dnsConfig" tf:"dns_config"`
	// +optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`
	// +optional
	MaxReturn     *int64                    `json:"maxReturn,omitempty" tf:"max_return"`
	MonitorConfig *ProfileSpecMonitorConfig `json:"monitorConfig" tf:"monitor_config"`
	Name          *string                   `json:"name" tf:"name"`
	// +optional
	ProfileStatus     *string `json:"profileStatus,omitempty" tf:"profile_status"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags                 *map[string]string `json:"tags,omitempty" tf:"tags"`
	TrafficRoutingMethod *string            `json:"trafficRoutingMethod" tf:"traffic_routing_method"`
	// +optional
	TrafficViewEnabled *bool `json:"trafficViewEnabled,omitempty" tf:"traffic_view_enabled"`
}

type ProfileStatus struct {
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

// ProfileList is a list of Profiles
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Profile CRD objects
	Items []Profile `json:"items,omitempty"`
}

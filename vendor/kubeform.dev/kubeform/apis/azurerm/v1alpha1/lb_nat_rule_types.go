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

type LbNATRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbNATRuleSpec   `json:"spec,omitempty"`
	Status            LbNATRuleStatus `json:"status,omitempty"`
}

type LbNATRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BackendIPConfigurationID string `json:"backendIPConfigurationID,omitempty" tf:"backend_ip_configuration_id,omitempty"`
	BackendPort              int64  `json:"backendPort" tf:"backend_port"`
	// +optional
	EnableFloatingIP bool `json:"enableFloatingIP,omitempty" tf:"enable_floating_ip,omitempty"`
	// +optional
	EnableTcpReset bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`
	// +optional
	FrontendIPConfigurationID   string `json:"frontendIPConfigurationID,omitempty" tf:"frontend_ip_configuration_id,omitempty"`
	FrontendIPConfigurationName string `json:"frontendIPConfigurationName" tf:"frontend_ip_configuration_name"`
	FrontendPort                int64  `json:"frontendPort" tf:"frontend_port"`
	// +optional
	IdleTimeoutInMinutes int64  `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`
	LoadbalancerID       string `json:"loadbalancerID" tf:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty" tf:"location,omitempty"`
	Name              string `json:"name" tf:"name"`
	Protocol          string `json:"protocol" tf:"protocol"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type LbNATRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbNATRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbNATRuleList is a list of LbNATRules
type LbNATRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbNATRule CRD objects
	Items []LbNATRule `json:"items,omitempty"`
}
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

type EndpointGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointGroupSpec   `json:"spec,omitempty"`
	Status            EndpointGroupStatus `json:"status,omitempty"`
}

type EndpointGroupSpec struct {
	EndpointGroupSpec2 `json:",inline"`
	// +optional
	KubeformOutput EndpointGroupSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type EndpointGroupSpecEndpointConfiguration struct {
	// +optional
	ClientIPPreservationEnabled *bool `json:"clientIPPreservationEnabled,omitempty" tf:"client_ip_preservation_enabled"`
	// +optional
	EndpointID *string `json:"endpointID,omitempty" tf:"endpoint_id"`
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type EndpointGroupSpecPortOverride struct {
	EndpointPort *int64 `json:"endpointPort" tf:"endpoint_port"`
	ListenerPort *int64 `json:"listenerPort" tf:"listener_port"`
}

type EndpointGroupSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	EndpointConfiguration []EndpointGroupSpecEndpointConfiguration `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration"`
	// +optional
	EndpointGroupRegion *string `json:"endpointGroupRegion,omitempty" tf:"endpoint_group_region"`
	// +optional
	HealthCheckIntervalSeconds *int64 `json:"healthCheckIntervalSeconds,omitempty" tf:"health_check_interval_seconds"`
	// +optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path"`
	// +optional
	HealthCheckPort *int64 `json:"healthCheckPort,omitempty" tf:"health_check_port"`
	// +optional
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty" tf:"health_check_protocol"`
	ListenerArn         *string `json:"listenerArn" tf:"listener_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	PortOverride []EndpointGroupSpecPortOverride `json:"portOverride,omitempty" tf:"port_override"`
	// +optional
	ThresholdCount *int64 `json:"thresholdCount,omitempty" tf:"threshold_count"`
	// +optional
	TrafficDialPercentage *float64 `json:"trafficDialPercentage,omitempty" tf:"traffic_dial_percentage"`
}

type EndpointGroupStatus struct {
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

// EndpointGroupList is a list of EndpointGroups
type EndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EndpointGroup CRD objects
	Items []EndpointGroup `json:"items,omitempty"`
}
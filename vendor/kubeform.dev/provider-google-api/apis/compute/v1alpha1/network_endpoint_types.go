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

type NetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkEndpointSpec   `json:"spec,omitempty"`
	Status            NetworkEndpointStatus `json:"status,omitempty"`
}

type NetworkEndpointSpec struct {
	NetworkEndpointSpec2 `json:",inline"`
	// +optional
	KubeformOutput NetworkEndpointSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type NetworkEndpointSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	Instance *string `json:"instance" tf:"instance"`
	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IpAddress *string `json:"ipAddress" tf:"ip_address"`
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup *string `json:"networkEndpointGroup" tf:"network_endpoint_group"`
	// Port number of network endpoint.
	Port *int64 `json:"port" tf:"port"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Zone where the containing network endpoint group is located.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type NetworkEndpointStatus struct {
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

// NetworkEndpointList is a list of NetworkEndpoints
type NetworkEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkEndpoint CRD objects
	Items []NetworkEndpoint `json:"items,omitempty"`
}
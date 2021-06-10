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

type Workspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec,omitempty"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

type WorkspaceSpec struct {
	WorkspaceSpec2 `json:",inline"`
	// +optional
	KubeformOutput WorkspaceSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type WorkspaceSpecCustomParameters struct {
	// +optional
	NoPublicIP *bool `json:"noPublicIP,omitempty" tf:"no_public_ip"`
	// +optional
	PrivateSubnetName *string `json:"privateSubnetName,omitempty" tf:"private_subnet_name"`
	// +optional
	PublicSubnetName *string `json:"publicSubnetName,omitempty" tf:"public_subnet_name"`
	// +optional
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
}

type WorkspaceSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomParameters *WorkspaceSpecCustomParameters `json:"customParameters,omitempty" tf:"custom_parameters"`
	Location         *string                        `json:"location" tf:"location"`
	// +optional
	ManagedResourceGroupID *string `json:"managedResourceGroupID,omitempty" tf:"managed_resource_group_id"`
	// +optional
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name"`
	Name                     *string `json:"name" tf:"name"`
	ResourceGroupName        *string `json:"resourceGroupName" tf:"resource_group_name"`
	Sku                      *string `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	WorkspaceID *string `json:"workspaceID,omitempty" tf:"workspace_id"`
	// +optional
	WorkspaceURL *string `json:"workspaceURL,omitempty" tf:"workspace_url"`
}

type WorkspaceStatus struct {
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

// WorkspaceList is a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Workspace CRD objects
	Items []Workspace `json:"items,omitempty"`
}

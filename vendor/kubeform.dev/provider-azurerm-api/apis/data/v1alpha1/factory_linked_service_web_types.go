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

type FactoryLinkedServiceWeb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceWebSpec   `json:"spec,omitempty"`
	Status            FactoryLinkedServiceWebStatus `json:"status,omitempty"`
}

type FactoryLinkedServiceWebSpec struct {
	FactoryLinkedServiceWebSpec2 `json:",inline"`
	// +optional
	KubeformOutput FactoryLinkedServiceWebSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type FactoryLinkedServiceWebSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`
	// +optional
	Annotations        []string `json:"annotations,omitempty" tf:"annotations"`
	AuthenticationType *string  `json:"authenticationType" tf:"authentication_type"`
	DataFactoryName    *string  `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name"`
	Name                   *string `json:"name" tf:"name"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	Password          *string `json:"-" sensitive:"true" tf:"password"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	Url               *string `json:"url" tf:"url"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type FactoryLinkedServiceWebStatus struct {
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

// FactoryLinkedServiceWebList is a list of FactoryLinkedServiceWebs
type FactoryLinkedServiceWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FactoryLinkedServiceWeb CRD objects
	Items []FactoryLinkedServiceWeb `json:"items,omitempty"`
}

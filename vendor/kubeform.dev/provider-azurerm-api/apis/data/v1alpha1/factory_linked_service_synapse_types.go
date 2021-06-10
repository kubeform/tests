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

type FactoryLinkedServiceSynapse struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceSynapseSpec   `json:"spec,omitempty"`
	Status            FactoryLinkedServiceSynapseStatus `json:"status,omitempty"`
}

type FactoryLinkedServiceSynapseSpec struct {
	FactoryLinkedServiceSynapseSpec2 `json:",inline"`
	// +optional
	KubeformOutput FactoryLinkedServiceSynapseSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type FactoryLinkedServiceSynapseSpecKeyVaultPassword struct {
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name"`
	SecretName        *string `json:"secretName" tf:"secret_name"`
}

type FactoryLinkedServiceSynapseSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`
	// +optional
	Annotations      []string `json:"annotations,omitempty" tf:"annotations"`
	ConnectionString *string  `json:"connectionString" tf:"connection_string"`
	DataFactoryName  *string  `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name"`
	// +optional
	KeyVaultPassword *FactoryLinkedServiceSynapseSpecKeyVaultPassword `json:"keyVaultPassword,omitempty" tf:"key_vault_password"`
	Name             *string                                          `json:"name" tf:"name"`
	// +optional
	Parameters        *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	ResourceGroupName *string            `json:"resourceGroupName" tf:"resource_group_name"`
}

type FactoryLinkedServiceSynapseStatus struct {
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

// FactoryLinkedServiceSynapseList is a list of FactoryLinkedServiceSynapses
type FactoryLinkedServiceSynapseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FactoryLinkedServiceSynapse CRD objects
	Items []FactoryLinkedServiceSynapse `json:"items,omitempty"`
}
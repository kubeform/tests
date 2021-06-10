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

type Connection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec,omitempty"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

type ConnectionSpec struct {
	ConnectionSpec2 `json:",inline"`
	// +optional
	KubeformOutput ConnectionSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ConnectionSpecPhysicalConnectionRequirements struct {
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	SecurityGroupIDList []string `json:"securityGroupIDList,omitempty" tf:"security_group_id_list"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type ConnectionSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CatalogID            *string           `json:"catalogID,omitempty" tf:"catalog_id"`
	ConnectionProperties map[string]string `json:"-" sensitive:"true" tf:"connection_properties"`
	// +optional
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	MatchCriteria []string `json:"matchCriteria,omitempty" tf:"match_criteria"`
	Name          *string  `json:"name" tf:"name"`
	// +optional
	PhysicalConnectionRequirements *ConnectionSpecPhysicalConnectionRequirements `json:"physicalConnectionRequirements,omitempty" tf:"physical_connection_requirements"`
}

type ConnectionStatus struct {
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

// ConnectionList is a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Connection CRD objects
	Items []Connection `json:"items,omitempty"`
}

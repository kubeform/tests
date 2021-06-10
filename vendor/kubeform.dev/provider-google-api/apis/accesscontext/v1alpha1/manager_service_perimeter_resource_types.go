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

type ManagerServicePerimeterResource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerServicePerimeterResourceSpec   `json:"spec,omitempty"`
	Status            ManagerServicePerimeterResourceStatus `json:"status,omitempty"`
}

type ManagerServicePerimeterResourceSpec struct {
	ManagerServicePerimeterResourceSpec2 `json:",inline"`
	// +optional
	KubeformOutput ManagerServicePerimeterResourceSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ManagerServicePerimeterResourceSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Service Perimeter to add this resource to.
	PerimeterName *string `json:"perimeterName" tf:"perimeter_name"`
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource *string `json:"resource" tf:"resource"`
}

type ManagerServicePerimeterResourceStatus struct {
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

// ManagerServicePerimeterResourceList is a list of ManagerServicePerimeterResources
type ManagerServicePerimeterResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagerServicePerimeterResource CRD objects
	Items []ManagerServicePerimeterResource `json:"items,omitempty"`
}
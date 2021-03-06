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

type AccessContextManagerServicePerimeter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessContextManagerServicePerimeterSpec   `json:"spec,omitempty"`
	Status            AccessContextManagerServicePerimeterStatus `json:"status,omitempty"`
}

type AccessContextManagerServicePerimeterSpecStatus struct {
	// +optional
	AccessLevels []string `json:"accessLevels,omitempty" tf:"access_levels,omitempty"`
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources,omitempty"`
	// +optional
	RestrictedServices []string `json:"restrictedServices,omitempty" tf:"restricted_services,omitempty"`
}

type AccessContextManagerServicePerimeterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreateTime string `json:"createTime,omitempty" tf:"create_time,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Parent      string `json:"parent" tf:"parent"`
	// +optional
	PerimeterType string `json:"perimeterType,omitempty" tf:"perimeter_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Status []AccessContextManagerServicePerimeterSpecStatus `json:"status,omitempty" tf:"status,omitempty"`
	Title  string                                           `json:"title" tf:"title"`
	// +optional
	UpdateTime string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type AccessContextManagerServicePerimeterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AccessContextManagerServicePerimeterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AccessContextManagerServicePerimeterList is a list of AccessContextManagerServicePerimeters
type AccessContextManagerServicePerimeterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AccessContextManagerServicePerimeter CRD objects
	Items []AccessContextManagerServicePerimeter `json:"items,omitempty"`
}

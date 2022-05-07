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

type AccessContextManagerAccessLevel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessContextManagerAccessLevelSpec   `json:"spec,omitempty"`
	Status            AccessContextManagerAccessLevelStatus `json:"status,omitempty"`
}

type AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicyOsConstraints struct {
	// +optional
	MinimumVersion string `json:"minimumVersion,omitempty" tf:"minimum_version,omitempty"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
}

type AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicy struct {
	// +optional
	AllowedDeviceManagementLevels []string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels,omitempty"`
	// +optional
	AllowedEncryptionStatuses []string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses,omitempty"`
	// +optional
	OsConstraints []AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicyOsConstraints `json:"osConstraints,omitempty" tf:"os_constraints,omitempty"`
	// +optional
	RequireScreenLock bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock,omitempty"`
}

type AccessContextManagerAccessLevelSpecBasicConditions struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DevicePolicy []AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicy `json:"devicePolicy,omitempty" tf:"device_policy,omitempty"`
	// +optional
	IpSubnetworks []string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks,omitempty"`
	// +optional
	Members []string `json:"members,omitempty" tf:"members,omitempty"`
	// +optional
	Negate bool `json:"negate,omitempty" tf:"negate,omitempty"`
	// +optional
	RequiredAccessLevels []string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels,omitempty"`
}

type AccessContextManagerAccessLevelSpecBasic struct {
	// +optional
	CombiningFunction string `json:"combiningFunction,omitempty" tf:"combining_function,omitempty"`
	// +kubebuilder:validation:MinItems=1
	Conditions []AccessContextManagerAccessLevelSpecBasicConditions `json:"conditions" tf:"conditions"`
}

type AccessContextManagerAccessLevelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Basic []AccessContextManagerAccessLevelSpecBasic `json:"basic,omitempty" tf:"basic,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Parent      string `json:"parent" tf:"parent"`
	Title       string `json:"title" tf:"title"`
}

type AccessContextManagerAccessLevelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AccessContextManagerAccessLevelSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AccessContextManagerAccessLevelList is a list of AccessContextManagerAccessLevels
type AccessContextManagerAccessLevelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AccessContextManagerAccessLevel CRD objects
	Items []AccessContextManagerAccessLevel `json:"items,omitempty"`
}
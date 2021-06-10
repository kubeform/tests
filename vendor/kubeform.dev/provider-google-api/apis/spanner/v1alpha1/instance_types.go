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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpec struct {
	InstanceSpec2 `json:",inline"`
	// +optional
	KubeformOutput InstanceSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type InstanceSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the instance's configuration (similar but not
	// quite the same as a region) which defines defines the geographic placement and
	// replication of your databases in this instance. It determines where your data
	// is stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	Config *string `json:"config" tf:"config"`
	// The descriptive name for this instance as it appears in UIs. Must be
	// unique per project and between 4 and 30 characters in length.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// A unique identifier for the instance, which cannot be changed after
	// the instance is created. The name must be between 6 and 30 characters
	// in length.
	//
	//
	// If not provided, a random string starting with 'tf-' will be selected.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The number of nodes allocated to this instance.
	// +optional
	NumNodes *int64 `json:"numNodes,omitempty" tf:"num_nodes"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Instance status: 'CREATING' or 'READY'.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
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

type ImageRecipe struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageRecipeSpec   `json:"spec,omitempty"`
	Status            ImageRecipeStatus `json:"status,omitempty"`
}

type ImageRecipeSpec struct {
	ImageRecipeSpec2 `json:",inline"`
	// +optional
	KubeformOutput ImageRecipeSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ImageRecipeSpecBlockDeviceMappingEbs struct {
	// +optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	// +optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
	// +optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`
	// +optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type ImageRecipeSpecBlockDeviceMapping struct {
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// +optional
	Ebs *ImageRecipeSpecBlockDeviceMappingEbs `json:"ebs,omitempty" tf:"ebs"`
	// +optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device"`
	// +optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name"`
}

type ImageRecipeSpecComponent struct {
	ComponentArn *string `json:"componentArn" tf:"component_arn"`
}

type ImageRecipeSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	BlockDeviceMapping []ImageRecipeSpecBlockDeviceMapping `json:"blockDeviceMapping,omitempty" tf:"block_device_mapping"`
	// +kubebuilder:validation:MinItems=1
	Component []ImageRecipeSpecComponent `json:"component" tf:"component"`
	// +optional
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	// +optional
	Owner       *string `json:"owner,omitempty" tf:"owner"`
	ParentImage *string `json:"parentImage" tf:"parent_image"`
	// +optional
	Platform *string `json:"platform,omitempty" tf:"platform"`
	// +optional
	Tags    *map[string]string `json:"tags,omitempty" tf:"tags"`
	Version *string            `json:"version" tf:"version"`
	// +optional
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory"`
}

type ImageRecipeStatus struct {
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

// ImageRecipeList is a list of ImageRecipes
type ImageRecipeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ImageRecipe CRD objects
	Items []ImageRecipe `json:"items,omitempty"`
}
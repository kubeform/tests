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

type CatalogTagTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogTagTemplateSpec   `json:"spec,omitempty"`
	Status            CatalogTagTemplateStatus `json:"status,omitempty"`
}

type CatalogTagTemplateSpec struct {
	CatalogTagTemplateSpec2 `json:",inline"`
	// +optional
	KubeformOutput CatalogTagTemplateSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type CatalogTagTemplateSpecFieldsTypeEnumTypeAllowedValues struct {
	// The display name of the enum value.
	DisplayName *string `json:"displayName" tf:"display_name"`
}

type CatalogTagTemplateSpecFieldsTypeEnumType struct {
	// The set of allowed values for this enum. The display names of the
	// values must be case-insensitively unique within this set. Currently,
	// enum values can only be added to the list of allowed values. Deletion
	// and renaming of enum values are not supported.
	// Can have up to 500 allowed values.
	AllowedValues []CatalogTagTemplateSpecFieldsTypeEnumTypeAllowedValues `json:"allowedValues" tf:"allowed_values"`
}

type CatalogTagTemplateSpecFieldsType struct {
	// Represents an enum type.
	//  Exactly one of 'primitive_type' or 'enum_type' must be set
	// +optional
	EnumType *CatalogTagTemplateSpecFieldsTypeEnumType `json:"enumType,omitempty" tf:"enum_type"`
	// Represents primitive types - string, bool etc.
	//  Exactly one of 'primitive_type' or 'enum_type' must be set Possible values: ["DOUBLE", "STRING", "BOOL", "TIMESTAMP"]
	// +optional
	PrimitiveType *string `json:"primitiveType,omitempty" tf:"primitive_type"`
}

type CatalogTagTemplateSpecFields struct {
	// A description for this field.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The display name for this field.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	FieldID     *string `json:"fieldID" tf:"field_id"`
	// Whether this is a required field. Defaults to false.
	// +optional
	IsRequired *bool `json:"isRequired,omitempty" tf:"is_required"`
	// The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The order of this field with respect to other fields in this tag template.
	// A higher value indicates a more important field. The value can be negative.
	// Multiple fields can have the same order, and field orders within a tag do not have to be sequential.
	// +optional
	Order *int64 `json:"order,omitempty" tf:"order"`
	// The type of value this tag field can contain.
	Type *CatalogTagTemplateSpecFieldsType `json:"type" tf:"type"`
}

type CatalogTagTemplateSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name for this template.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
	Fields []CatalogTagTemplateSpecFields `json:"fields" tf:"fields"`
	// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
	// +optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete"`
	// The resource name of the tag template in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Template location region.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The id of the tag template to create.
	TagTemplateID *string `json:"tagTemplateID" tf:"tag_template_id"`
}

type CatalogTagTemplateStatus struct {
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

// CatalogTagTemplateList is a list of CatalogTagTemplates
type CatalogTagTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CatalogTagTemplate CRD objects
	Items []CatalogTagTemplate `json:"items,omitempty"`
}

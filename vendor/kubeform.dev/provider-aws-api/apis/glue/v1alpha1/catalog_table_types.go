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

type CatalogTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogTableSpec   `json:"spec,omitempty"`
	Status            CatalogTableStatus `json:"status,omitempty"`
}

type CatalogTableSpec struct {
	CatalogTableSpec2 `json:",inline"`
	// +optional
	KubeformOutput CatalogTableSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type CatalogTableSpecPartitionIndex struct {
	IndexName *string `json:"indexName" tf:"index_name"`
	// +optional
	IndexStatus *string `json:"indexStatus,omitempty" tf:"index_status"`
	// +kubebuilder:validation:MinItems=1
	Keys []string `json:"keys" tf:"keys"`
}

type CatalogTableSpecPartitionKeys struct {
	// +optional
	Comment *string `json:"comment,omitempty" tf:"comment"`
	Name    *string `json:"name" tf:"name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type CatalogTableSpecStorageDescriptorColumns struct {
	// +optional
	Comment *string `json:"comment,omitempty" tf:"comment"`
	Name    *string `json:"name" tf:"name"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type CatalogTableSpecStorageDescriptorSchemaReferenceSchemaID struct {
	// +optional
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name"`
	// +optional
	SchemaArn *string `json:"schemaArn,omitempty" tf:"schema_arn"`
	// +optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name"`
}

type CatalogTableSpecStorageDescriptorSchemaReference struct {
	// +optional
	SchemaID *CatalogTableSpecStorageDescriptorSchemaReferenceSchemaID `json:"schemaID,omitempty" tf:"schema_id"`
	// +optional
	SchemaVersionID     *string `json:"schemaVersionID,omitempty" tf:"schema_version_id"`
	SchemaVersionNumber *int64  `json:"schemaVersionNumber" tf:"schema_version_number"`
}

type CatalogTableSpecStorageDescriptorSerDeInfo struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	SerializationLibrary *string `json:"serializationLibrary,omitempty" tf:"serialization_library"`
}

type CatalogTableSpecStorageDescriptorSkewedInfo struct {
	// +optional
	SkewedColumnNames []string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names"`
	// +optional
	SkewedColumnValueLocationMaps *map[string]string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps"`
	// +optional
	SkewedColumnValues []string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values"`
}

type CatalogTableSpecStorageDescriptorSortColumns struct {
	Column    *string `json:"column" tf:"column"`
	SortOrder *int64  `json:"sortOrder" tf:"sort_order"`
}

type CatalogTableSpecStorageDescriptor struct {
	// +optional
	BucketColumns []string `json:"bucketColumns,omitempty" tf:"bucket_columns"`
	// +optional
	Columns []CatalogTableSpecStorageDescriptorColumns `json:"columns,omitempty" tf:"columns"`
	// +optional
	Compressed *bool `json:"compressed,omitempty" tf:"compressed"`
	// +optional
	InputFormat *string `json:"inputFormat,omitempty" tf:"input_format"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	NumberOfBuckets *int64 `json:"numberOfBuckets,omitempty" tf:"number_of_buckets"`
	// +optional
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	SchemaReference *CatalogTableSpecStorageDescriptorSchemaReference `json:"schemaReference,omitempty" tf:"schema_reference"`
	// +optional
	SerDeInfo *CatalogTableSpecStorageDescriptorSerDeInfo `json:"serDeInfo,omitempty" tf:"ser_de_info"`
	// +optional
	SkewedInfo *CatalogTableSpecStorageDescriptorSkewedInfo `json:"skewedInfo,omitempty" tf:"skewed_info"`
	// +optional
	SortColumns []CatalogTableSpecStorageDescriptorSortColumns `json:"sortColumns,omitempty" tf:"sort_columns"`
	// +optional
	StoredAsSubDirectories *bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories"`
}

type CatalogTableSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CatalogID    *string `json:"catalogID,omitempty" tf:"catalog_id"`
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Name        *string `json:"name" tf:"name"`
	// +optional
	Owner *string `json:"owner,omitempty" tf:"owner"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	PartitionIndex []CatalogTableSpecPartitionIndex `json:"partitionIndex,omitempty" tf:"partition_index"`
	// +optional
	PartitionKeys []CatalogTableSpecPartitionKeys `json:"partitionKeys,omitempty" tf:"partition_keys"`
	// +optional
	Retention *int64 `json:"retention,omitempty" tf:"retention"`
	// +optional
	StorageDescriptor *CatalogTableSpecStorageDescriptor `json:"storageDescriptor,omitempty" tf:"storage_descriptor"`
	// +optional
	TableType *string `json:"tableType,omitempty" tf:"table_type"`
	// +optional
	ViewExpandedText *string `json:"viewExpandedText,omitempty" tf:"view_expanded_text"`
	// +optional
	ViewOriginalText *string `json:"viewOriginalText,omitempty" tf:"view_original_text"`
}

type CatalogTableStatus struct {
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

// CatalogTableList is a list of CatalogTables
type CatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CatalogTable CRD objects
	Items []CatalogTable `json:"items,omitempty"`
}

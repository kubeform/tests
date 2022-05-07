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

type StorageBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketSpec   `json:"spec,omitempty"`
	Status            StorageBucketStatus `json:"status,omitempty"`
}

type StorageBucketSpecCors struct {
	// +optional
	MaxAgeSeconds int64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
	// +optional
	Method []string `json:"method,omitempty" tf:"method,omitempty"`
	// +optional
	Origin []string `json:"origin,omitempty" tf:"origin,omitempty"`
	// +optional
	ResponseHeader []string `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type StorageBucketSpecEncryption struct {
	DefaultKmsKeyName string `json:"defaultKmsKeyName" tf:"default_kms_key_name"`
}

type StorageBucketSpecLifecycleRuleAction struct {
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
	Type         string `json:"type" tf:"type"`
}

type StorageBucketSpecLifecycleRuleCondition struct {
	// +optional
	Age int64 `json:"age,omitempty" tf:"age,omitempty"`
	// +optional
	CreatedBefore string `json:"createdBefore,omitempty" tf:"created_before,omitempty"`
	// +optional
	// Deprecated
	IsLive bool `json:"isLive,omitempty" tf:"is_live,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty" tf:"matches_storage_class,omitempty"`
	// +optional
	NumNewerVersions int64 `json:"numNewerVersions,omitempty" tf:"num_newer_versions,omitempty"`
	// +optional
	WithState string `json:"withState,omitempty" tf:"with_state,omitempty"`
}

type StorageBucketSpecLifecycleRule struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Action []StorageBucketSpecLifecycleRuleAction `json:"action" tf:"action"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Condition []StorageBucketSpecLifecycleRuleCondition `json:"condition" tf:"condition"`
}

type StorageBucketSpecLogging struct {
	LogBucket string `json:"logBucket" tf:"log_bucket"`
	// +optional
	LogObjectPrefix string `json:"logObjectPrefix,omitempty" tf:"log_object_prefix,omitempty"`
}

type StorageBucketSpecRetentionPolicy struct {
	// +optional
	IsLocked        bool  `json:"isLocked,omitempty" tf:"is_locked,omitempty"`
	RetentionPeriod int64 `json:"retentionPeriod" tf:"retention_period"`
}

type StorageBucketSpecVersioning struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type StorageBucketSpecWebsite struct {
	// +optional
	MainPageSuffix string `json:"mainPageSuffix,omitempty" tf:"main_page_suffix,omitempty"`
	// +optional
	NotFoundPage string `json:"notFoundPage,omitempty" tf:"not_found_page,omitempty"`
}

type StorageBucketSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BucketPolicyOnly bool `json:"bucketPolicyOnly,omitempty" tf:"bucket_policy_only,omitempty"`
	// +optional
	Cors []StorageBucketSpecCors `json:"cors,omitempty" tf:"cors,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Encryption []StorageBucketSpecEncryption `json:"encryption,omitempty" tf:"encryption,omitempty"`
	// +optional
	ForceDestroy bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	LifecycleRule []StorageBucketSpecLifecycleRule `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging []StorageBucketSpecLogging `json:"logging,omitempty" tf:"logging,omitempty"`
	Name    string                     `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	RequesterPays bool `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []StorageBucketSpecRetentionPolicy `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Versioning []StorageBucketSpecVersioning `json:"versioning,omitempty" tf:"versioning,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Website []StorageBucketSpecWebsite `json:"website,omitempty" tf:"website,omitempty"`
}

type StorageBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageBucketSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketList is a list of StorageBuckets
type StorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucket CRD objects
	Items []StorageBucket `json:"items,omitempty"`
}
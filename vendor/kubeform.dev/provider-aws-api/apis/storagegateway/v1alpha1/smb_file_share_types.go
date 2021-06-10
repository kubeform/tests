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

type SmbFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SmbFileShareSpec   `json:"spec,omitempty"`
	Status            SmbFileShareStatus `json:"status,omitempty"`
}

type SmbFileShareSpec struct {
	SmbFileShareSpec2 `json:",inline"`
	// +optional
	KubeformOutput SmbFileShareSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type SmbFileShareSpecCacheAttributes struct {
	// +optional
	CacheStaleTimeoutInSeconds *int64 `json:"cacheStaleTimeoutInSeconds,omitempty" tf:"cache_stale_timeout_in_seconds"`
}

type SmbFileShareSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessBasedEnumeration *bool `json:"accessBasedEnumeration,omitempty" tf:"access_based_enumeration"`
	// +optional
	AdminUserList []string `json:"adminUserList,omitempty" tf:"admin_user_list"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AuditDestinationArn *string `json:"auditDestinationArn,omitempty" tf:"audit_destination_arn"`
	// +optional
	Authentication *string `json:"authentication,omitempty" tf:"authentication"`
	// +optional
	CacheAttributes *SmbFileShareSpecCacheAttributes `json:"cacheAttributes,omitempty" tf:"cache_attributes"`
	// +optional
	CaseSensitivity *string `json:"caseSensitivity,omitempty" tf:"case_sensitivity"`
	// +optional
	DefaultStorageClass *string `json:"defaultStorageClass,omitempty" tf:"default_storage_class"`
	// +optional
	FileShareName *string `json:"fileShareName,omitempty" tf:"file_share_name"`
	// +optional
	FileshareID *string `json:"fileshareID,omitempty" tf:"fileshare_id"`
	GatewayArn  *string `json:"gatewayArn" tf:"gateway_arn"`
	// +optional
	GuessMimeTypeEnabled *bool `json:"guessMimeTypeEnabled,omitempty" tf:"guess_mime_type_enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	InvalidUserList []string `json:"invalidUserList,omitempty" tf:"invalid_user_list"`
	// +optional
	KmsEncrypted *bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted"`
	// +optional
	KmsKeyArn   *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	LocationArn *string `json:"locationArn" tf:"location_arn"`
	// +optional
	NotificationPolicy *string `json:"notificationPolicy,omitempty" tf:"notification_policy"`
	// +optional
	ObjectACL *string `json:"objectACL,omitempty" tf:"object_acl"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`
	// +optional
	RequesterPays *bool   `json:"requesterPays,omitempty" tf:"requester_pays"`
	RoleArn       *string `json:"roleArn" tf:"role_arn"`
	// +optional
	SmbACLEnabled *bool `json:"smbACLEnabled,omitempty" tf:"smb_acl_enabled"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	ValidUserList []string `json:"validUserList,omitempty" tf:"valid_user_list"`
}

type SmbFileShareStatus struct {
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

// SmbFileShareList is a list of SmbFileShares
type SmbFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SmbFileShare CRD objects
	Items []SmbFileShare `json:"items,omitempty"`
}

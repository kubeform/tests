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

type AccountKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountKeySpec   `json:"spec,omitempty"`
	Status            AccountKeyStatus `json:"status,omitempty"`
}

type AccountKeySpec struct {
	AccountKeySpec2 `json:",inline"`
	// +optional
	KubeformOutput AccountKeySpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type AccountKeySpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// Arbitrary map of values that, when changed, will trigger recreation of resource.
	// +optional
	Keepers map[string]string `json:"keepers,omitempty" tf:"keepers"`
	// The algorithm used to generate the key, used only on create. KEY_ALG_RSA_2048 is the default algorithm. Valid values are: "KEY_ALG_RSA_1024", "KEY_ALG_RSA_2048".
	// +optional
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm"`
	// The name used for this key pair
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating service account keys through the CLI or web console. This is only populated when creating a new key.
	// +optional
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	PrivateKeyType *string `json:"privateKeyType,omitempty" tf:"private_key_type"`
	// The public key, base64 encoded
	// +optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key"`
	// A field that allows clients to upload their own public key. If set, use this public key data to create a service account key for given service account. Please note, the expected format for this field is a base64 encoded X509_PEM.
	// +optional
	PublicKeyData *string `json:"publicKeyData,omitempty" tf:"public_key_data"`
	// +optional
	PublicKeyType *string `json:"publicKeyType,omitempty" tf:"public_key_type"`
	// The ID of the parent service account of the key. This can be a string in the format {ACCOUNT} or projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}, where {ACCOUNT} is the email address or unique id of the service account. If the {ACCOUNT} syntax is used, the project will be inferred from the account.
	ServiceAccountID *string `json:"serviceAccountID" tf:"service_account_id"`
	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	ValidAfter *string `json:"validAfter,omitempty" tf:"valid_after"`
	// The key can be used before this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	ValidBefore *string `json:"validBefore,omitempty" tf:"valid_before"`
}

type AccountKeyStatus struct {
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

// AccountKeyList is a list of AccountKeys
type AccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AccountKey CRD objects
	Items []AccountKey `json:"items,omitempty"`
}
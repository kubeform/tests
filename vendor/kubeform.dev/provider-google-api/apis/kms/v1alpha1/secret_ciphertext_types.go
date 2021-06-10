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

type SecretCiphertext struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretCiphertextSpec   `json:"spec,omitempty"`
	Status            SecretCiphertextStatus `json:"status,omitempty"`
}

type SecretCiphertextSpec struct {
	SecretCiphertextSpec2 `json:",inline"`
	// +optional
	KubeformOutput SecretCiphertextSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type SecretCiphertextSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// The additional authenticated data used for integrity checks during encryption and decryption.
	// +optional
	AdditionalAuthenticatedData *string `json:"-" sensitive:"true" tf:"additional_authenticated_data"`
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	// +optional
	Ciphertext *string `json:"ciphertext,omitempty" tf:"ciphertext"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
	CryptoKey *string `json:"cryptoKey" tf:"crypto_key"`
	// The plaintext to be encrypted.
	Plaintext *string `json:"-" sensitive:"true" tf:"plaintext"`
}

type SecretCiphertextStatus struct {
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

// SecretCiphertextList is a list of SecretCiphertexts
type SecretCiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecretCiphertext CRD objects
	Items []SecretCiphertext `json:"items,omitempty"`
}

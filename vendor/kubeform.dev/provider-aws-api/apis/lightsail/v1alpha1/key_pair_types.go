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

type KeyPair struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyPairSpec   `json:"spec,omitempty"`
	Status            KeyPairStatus `json:"status,omitempty"`
}

type KeyPairSpec struct {
	KeyPairSpec2 `json:",inline"`
	// +optional
	KubeformOutput KeyPairSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type KeyPairSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	EncryptedFingerprint *string `json:"encryptedFingerprint,omitempty" tf:"encrypted_fingerprint"`
	// +optional
	EncryptedPrivateKey *string `json:"encryptedPrivateKey,omitempty" tf:"encrypted_private_key"`
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key"`
	// +optional
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key"`
	// +optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key"`
}

type KeyPairStatus struct {
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

// KeyPairList is a list of KeyPairs
type KeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyPair CRD objects
	Items []KeyPair `json:"items,omitempty"`
}

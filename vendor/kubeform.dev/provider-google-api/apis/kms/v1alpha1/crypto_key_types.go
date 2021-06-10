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

type CryptoKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CryptoKeySpec   `json:"spec,omitempty"`
	Status            CryptoKeyStatus `json:"status,omitempty"`
}

type CryptoKeySpec struct {
	CryptoKeySpec2 `json:",inline"`
	// +optional
	KubeformOutput CryptoKeySpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type CryptoKeySpecVersionTemplate struct {
	// The algorithm to use when creating a version based on this template.
	// See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
	Algorithm *string `json:"algorithm" tf:"algorithm"`
	// The protection level to use when creating a version based on this template. Default value: "SOFTWARE" Possible values: ["SOFTWARE", "HSM"]
	// +optional
	ProtectionLevel *string `json:"protectionLevel,omitempty" tf:"protection_level"`
}

type CryptoKeySpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The KeyRing that this key belongs to.
	// Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	KeyRing *string `json:"keyRing" tf:"key_ring"`
	// Labels with user-defined metadata to apply to this resource.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The resource name for the CryptoKey.
	Name *string `json:"name" tf:"name"`
	// The immutable purpose of this CryptoKey. See the
	// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	// for possible inputs. Default value: "ENCRYPT_DECRYPT" Possible values: ["ENCRYPT_DECRYPT", "ASYMMETRIC_SIGN", "ASYMMETRIC_DECRYPT"]
	// +optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose"`
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter 's' (seconds). It must be greater than a day (ie, 86400).
	// +optional
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period"`
	// The self link of the created KeyRing in the format projects/{project}/locations/{location}/keyRings/{name}.
	// +optional
	// Deprecated
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the 'google_kms_key_ring_import_job' resource to import the CryptoKeyVersion.
	// +optional
	SkipInitialVersionCreation *bool `json:"skipInitialVersionCreation,omitempty" tf:"skip_initial_version_creation"`
	// A template describing settings for new crypto key versions.
	// +optional
	VersionTemplate *CryptoKeySpecVersionTemplate `json:"versionTemplate,omitempty" tf:"version_template"`
}

type CryptoKeyStatus struct {
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

// CryptoKeyList is a list of CryptoKeys
type CryptoKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CryptoKey CRD objects
	Items []CryptoKey `json:"items,omitempty"`
}
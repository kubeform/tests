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

type ServerCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerCertificateSpec   `json:"spec,omitempty"`
	Status            ServerCertificateStatus `json:"status,omitempty"`
}

type ServerCertificateSpec struct {
	ServerCertificateSpec2 `json:",inline"`
	// +optional
	KubeformOutput ServerCertificateSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ServerCertificateSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Arn             *string `json:"arn,omitempty" tf:"arn"`
	CertificateBody *string `json:"certificateBody" tf:"certificate_body"`
	// +optional
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain"`
	// +optional
	Expiration *string `json:"expiration,omitempty" tf:"expiration"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Path       *string `json:"path,omitempty" tf:"path"`
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UploadDate *string `json:"uploadDate,omitempty" tf:"upload_date"`
}

type ServerCertificateStatus struct {
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

// ServerCertificateList is a list of ServerCertificates
type ServerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServerCertificate CRD objects
	Items []ServerCertificate `json:"items,omitempty"`
}
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

type CustomHTTPSConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomHTTPSConfigurationSpec   `json:"spec,omitempty"`
	Status            CustomHTTPSConfigurationStatus `json:"status,omitempty"`
}

type CustomHTTPSConfigurationSpec struct {
	CustomHTTPSConfigurationSpec2 `json:",inline"`
	// +optional
	KubeformOutput CustomHTTPSConfigurationSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type CustomHTTPSConfigurationSpecCustomHTTPSConfiguration struct {
	// +optional
	AzureKeyVaultCertificateSecretName *string `json:"azureKeyVaultCertificateSecretName,omitempty" tf:"azure_key_vault_certificate_secret_name"`
	// +optional
	AzureKeyVaultCertificateSecretVersion *string `json:"azureKeyVaultCertificateSecretVersion,omitempty" tf:"azure_key_vault_certificate_secret_version"`
	// +optional
	AzureKeyVaultCertificateVaultID *string `json:"azureKeyVaultCertificateVaultID,omitempty" tf:"azure_key_vault_certificate_vault_id"`
	// +optional
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source"`
	// +optional
	MinimumTLSVersion *string `json:"minimumTLSVersion,omitempty" tf:"minimum_tls_version"`
	// +optional
	ProvisioningState *string `json:"provisioningState,omitempty" tf:"provisioning_state"`
	// +optional
	ProvisioningSubstate *string `json:"provisioningSubstate,omitempty" tf:"provisioning_substate"`
}

type CustomHTTPSConfigurationSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomHTTPSConfiguration       *CustomHTTPSConfigurationSpecCustomHTTPSConfiguration `json:"customHTTPSConfiguration,omitempty" tf:"custom_https_configuration"`
	CustomHTTPSProvisioningEnabled *bool                                                 `json:"customHTTPSProvisioningEnabled" tf:"custom_https_provisioning_enabled"`
	FrontendEndpointID             *string                                               `json:"frontendEndpointID" tf:"frontend_endpoint_id"`
}

type CustomHTTPSConfigurationStatus struct {
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

// CustomHTTPSConfigurationList is a list of CustomHTTPSConfigurations
type CustomHTTPSConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomHTTPSConfiguration CRD objects
	Items []CustomHTTPSConfiguration `json:"items,omitempty"`
}

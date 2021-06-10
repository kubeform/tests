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

type GraphqlAPI struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GraphqlAPISpec   `json:"spec,omitempty"`
	Status            GraphqlAPIStatus `json:"status,omitempty"`
}

type GraphqlAPISpec struct {
	GraphqlAPISpec2 `json:",inline"`
	// +optional
	KubeformOutput GraphqlAPISpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig struct {
	// +optional
	AuthTtl *int64 `json:"authTtl,omitempty" tf:"auth_ttl"`
	// +optional
	ClientID *string `json:"clientID,omitempty" tf:"client_id"`
	// +optional
	IatTtl *int64  `json:"iatTtl,omitempty" tf:"iat_ttl"`
	Issuer *string `json:"issuer" tf:"issuer"`
}

type GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig struct {
	// +optional
	AppIDClientRegex *string `json:"appIDClientRegex,omitempty" tf:"app_id_client_regex"`
	// +optional
	AwsRegion  *string `json:"awsRegion,omitempty" tf:"aws_region"`
	UserPoolID *string `json:"userPoolID" tf:"user_pool_id"`
}

type GraphqlAPISpecAdditionalAuthenticationProvider struct {
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type"`
	// +optional
	OpenidConnectConfig *GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig `json:"openidConnectConfig,omitempty" tf:"openid_connect_config"`
	// +optional
	UserPoolConfig *GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig `json:"userPoolConfig,omitempty" tf:"user_pool_config"`
}

type GraphqlAPISpecLogConfig struct {
	CloudwatchLogsRoleArn *string `json:"cloudwatchLogsRoleArn" tf:"cloudwatch_logs_role_arn"`
	// +optional
	ExcludeVerboseContent *bool   `json:"excludeVerboseContent,omitempty" tf:"exclude_verbose_content"`
	FieldLogLevel         *string `json:"fieldLogLevel" tf:"field_log_level"`
}

type GraphqlAPISpecOpenidConnectConfig struct {
	// +optional
	AuthTtl *int64 `json:"authTtl,omitempty" tf:"auth_ttl"`
	// +optional
	ClientID *string `json:"clientID,omitempty" tf:"client_id"`
	// +optional
	IatTtl *int64  `json:"iatTtl,omitempty" tf:"iat_ttl"`
	Issuer *string `json:"issuer" tf:"issuer"`
}

type GraphqlAPISpecUserPoolConfig struct {
	// +optional
	AppIDClientRegex *string `json:"appIDClientRegex,omitempty" tf:"app_id_client_regex"`
	// +optional
	AwsRegion     *string `json:"awsRegion,omitempty" tf:"aws_region"`
	DefaultAction *string `json:"defaultAction" tf:"default_action"`
	UserPoolID    *string `json:"userPoolID" tf:"user_pool_id"`
}

type GraphqlAPISpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalAuthenticationProvider []GraphqlAPISpecAdditionalAuthenticationProvider `json:"additionalAuthenticationProvider,omitempty" tf:"additional_authentication_provider"`
	// +optional
	Arn                *string `json:"arn,omitempty" tf:"arn"`
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type"`
	// +optional
	LogConfig *GraphqlAPISpecLogConfig `json:"logConfig,omitempty" tf:"log_config"`
	Name      *string                  `json:"name" tf:"name"`
	// +optional
	OpenidConnectConfig *GraphqlAPISpecOpenidConnectConfig `json:"openidConnectConfig,omitempty" tf:"openid_connect_config"`
	// +optional
	Schema *string `json:"schema,omitempty" tf:"schema"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Uris *map[string]string `json:"uris,omitempty" tf:"uris"`
	// +optional
	UserPoolConfig *GraphqlAPISpecUserPoolConfig `json:"userPoolConfig,omitempty" tf:"user_pool_config"`
	// +optional
	XrayEnabled *bool `json:"xrayEnabled,omitempty" tf:"xray_enabled"`
}

type GraphqlAPIStatus struct {
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

// GraphqlAPIList is a list of GraphqlAPIs
type GraphqlAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GraphqlAPI CRD objects
	Items []GraphqlAPI `json:"items,omitempty"`
}
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

type AppEngineStandardAppVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppEngineStandardAppVersionSpec   `json:"spec,omitempty"`
	Status            AppEngineStandardAppVersionStatus `json:"status,omitempty"`
}

type AppEngineStandardAppVersionSpecDeploymentFiles struct {
	Name string `json:"name" tf:"name"`
	// +optional
	Sha1Sum string `json:"sha1Sum,omitempty" tf:"sha1_sum,omitempty"`
	// +optional
	SourceURL string `json:"sourceURL,omitempty" tf:"source_url,omitempty"`
}

type AppEngineStandardAppVersionSpecDeploymentZip struct {
	// +optional
	FilesCount int64 `json:"filesCount,omitempty" tf:"files_count,omitempty"`
	// +optional
	SourceURL string `json:"sourceURL,omitempty" tf:"source_url,omitempty"`
}

type AppEngineStandardAppVersionSpecDeployment struct {
	// +optional
	Files []AppEngineStandardAppVersionSpecDeploymentFiles `json:"files,omitempty" tf:"files,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zip []AppEngineStandardAppVersionSpecDeploymentZip `json:"zip,omitempty" tf:"zip,omitempty"`
}

type AppEngineStandardAppVersionSpecEntrypoint struct {
	// +optional
	Shell string `json:"shell,omitempty" tf:"shell,omitempty"`
}

type AppEngineStandardAppVersionSpecHandlersScript struct {
	// +optional
	ScriptPath string `json:"scriptPath,omitempty" tf:"script_path,omitempty"`
}

type AppEngineStandardAppVersionSpecHandlersStaticFiles struct {
	// +optional
	ApplicationReadable bool `json:"applicationReadable,omitempty" tf:"application_readable,omitempty"`
	// +optional
	Expiration string `json:"expiration,omitempty" tf:"expiration,omitempty"`
	// +optional
	HttpHeaders map[string]string `json:"httpHeaders,omitempty" tf:"http_headers,omitempty"`
	// +optional
	MimeType string `json:"mimeType,omitempty" tf:"mime_type,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	RequireMatchingFile bool `json:"requireMatchingFile,omitempty" tf:"require_matching_file,omitempty"`
	// +optional
	UploadPathRegex string `json:"uploadPathRegex,omitempty" tf:"upload_path_regex,omitempty"`
}

type AppEngineStandardAppVersionSpecHandlers struct {
	// +optional
	AuthFailAction string `json:"authFailAction,omitempty" tf:"auth_fail_action,omitempty"`
	// +optional
	Login string `json:"login,omitempty" tf:"login,omitempty"`
	// +optional
	RedirectHTTPResponseCode string `json:"redirectHTTPResponseCode,omitempty" tf:"redirect_http_response_code,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Script []AppEngineStandardAppVersionSpecHandlersScript `json:"script,omitempty" tf:"script,omitempty"`
	// +optional
	SecurityLevel string `json:"securityLevel,omitempty" tf:"security_level,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StaticFiles []AppEngineStandardAppVersionSpecHandlersStaticFiles `json:"staticFiles,omitempty" tf:"static_files,omitempty"`
	// +optional
	UrlRegex string `json:"urlRegex,omitempty" tf:"url_regex,omitempty"`
}

type AppEngineStandardAppVersionSpecLibraries struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type AppEngineStandardAppVersionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Deployment []AppEngineStandardAppVersionSpecDeployment `json:"deployment,omitempty" tf:"deployment,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Entrypoint []AppEngineStandardAppVersionSpecEntrypoint `json:"entrypoint,omitempty" tf:"entrypoint,omitempty"`
	// +optional
	EnvVariables map[string]string `json:"envVariables,omitempty" tf:"env_variables,omitempty"`
	// +optional
	Handlers []AppEngineStandardAppVersionSpecHandlers `json:"handlers,omitempty" tf:"handlers,omitempty"`
	// +optional
	Libraries []AppEngineStandardAppVersionSpecLibraries `json:"libraries,omitempty" tf:"libraries,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NoopOnDestroy bool `json:"noopOnDestroy,omitempty" tf:"noop_on_destroy,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	Runtime string `json:"runtime" tf:"runtime"`
	// +optional
	RuntimeAPIVersion string `json:"runtimeAPIVersion,omitempty" tf:"runtime_api_version,omitempty"`
	// +optional
	Service string `json:"service,omitempty" tf:"service,omitempty"`
	// +optional
	Threadsafe bool `json:"threadsafe,omitempty" tf:"threadsafe,omitempty"`
	// +optional
	VersionID string `json:"versionID,omitempty" tf:"version_id,omitempty"`
}

type AppEngineStandardAppVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppEngineStandardAppVersionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppEngineStandardAppVersionList is a list of AppEngineStandardAppVersions
type AppEngineStandardAppVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppEngineStandardAppVersion CRD objects
	Items []AppEngineStandardAppVersion `json:"items,omitempty"`
}

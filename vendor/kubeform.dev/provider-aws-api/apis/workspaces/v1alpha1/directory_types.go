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

type Directory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectorySpec   `json:"spec,omitempty"`
	Status            DirectoryStatus `json:"status,omitempty"`
}

type DirectorySpec struct {
	DirectorySpec2 `json:",inline"`
	// +optional
	KubeformOutput DirectorySpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type DirectorySpecSelfServicePermissions struct {
	// +optional
	ChangeComputeType *bool `json:"changeComputeType,omitempty" tf:"change_compute_type"`
	// +optional
	IncreaseVolumeSize *bool `json:"increaseVolumeSize,omitempty" tf:"increase_volume_size"`
	// +optional
	RebuildWorkspace *bool `json:"rebuildWorkspace,omitempty" tf:"rebuild_workspace"`
	// +optional
	RestartWorkspace *bool `json:"restartWorkspace,omitempty" tf:"restart_workspace"`
	// +optional
	SwitchRunningMode *bool `json:"switchRunningMode,omitempty" tf:"switch_running_mode"`
}

type DirectorySpecWorkspaceAccessProperties struct {
	// +optional
	DeviceTypeAndroid *string `json:"deviceTypeAndroid,omitempty" tf:"device_type_android"`
	// +optional
	DeviceTypeChromeos *string `json:"deviceTypeChromeos,omitempty" tf:"device_type_chromeos"`
	// +optional
	DeviceTypeIos *string `json:"deviceTypeIos,omitempty" tf:"device_type_ios"`
	// +optional
	DeviceTypeOsx *string `json:"deviceTypeOsx,omitempty" tf:"device_type_osx"`
	// +optional
	DeviceTypeWeb *string `json:"deviceTypeWeb,omitempty" tf:"device_type_web"`
	// +optional
	DeviceTypeWindows *string `json:"deviceTypeWindows,omitempty" tf:"device_type_windows"`
	// +optional
	DeviceTypeZeroclient *string `json:"deviceTypeZeroclient,omitempty" tf:"device_type_zeroclient"`
}

type DirectorySpecWorkspaceCreationProperties struct {
	// +optional
	CustomSecurityGroupID *string `json:"customSecurityGroupID,omitempty" tf:"custom_security_group_id"`
	// +optional
	DefaultOu *string `json:"defaultOu,omitempty" tf:"default_ou"`
	// +optional
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty" tf:"enable_internet_access"`
	// +optional
	EnableMaintenanceMode *bool `json:"enableMaintenanceMode,omitempty" tf:"enable_maintenance_mode"`
	// +optional
	UserEnabledAsLocalAdministrator *bool `json:"userEnabledAsLocalAdministrator,omitempty" tf:"user_enabled_as_local_administrator"`
}

type DirectorySpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Alias *string `json:"alias,omitempty" tf:"alias"`
	// +optional
	CustomerUserName *string `json:"customerUserName,omitempty" tf:"customer_user_name"`
	DirectoryID      *string `json:"directoryID" tf:"directory_id"`
	// +optional
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name"`
	// +optional
	DirectoryType *string `json:"directoryType,omitempty" tf:"directory_type"`
	// +optional
	DnsIPAddresses []string `json:"dnsIPAddresses,omitempty" tf:"dns_ip_addresses"`
	// +optional
	IamRoleID *string `json:"iamRoleID,omitempty" tf:"iam_role_id"`
	// +optional
	IpGroupIDS []string `json:"ipGroupIDS,omitempty" tf:"ip_group_ids"`
	// +optional
	RegistrationCode *string `json:"registrationCode,omitempty" tf:"registration_code"`
	// +optional
	SelfServicePermissions *DirectorySpecSelfServicePermissions `json:"selfServicePermissions,omitempty" tf:"self_service_permissions"`
	// +optional
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	WorkspaceAccessProperties *DirectorySpecWorkspaceAccessProperties `json:"workspaceAccessProperties,omitempty" tf:"workspace_access_properties"`
	// +optional
	WorkspaceCreationProperties *DirectorySpecWorkspaceCreationProperties `json:"workspaceCreationProperties,omitempty" tf:"workspace_creation_properties"`
	// +optional
	WorkspaceSecurityGroupID *string `json:"workspaceSecurityGroupID,omitempty" tf:"workspace_security_group_id"`
}

type DirectoryStatus struct {
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

// DirectoryList is a list of Directorys
type DirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Directory CRD objects
	Items []Directory `json:"items,omitempty"`
}

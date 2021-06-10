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

type VirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineScaleSetSpec   `json:"spec,omitempty"`
	Status            VirtualMachineScaleSetStatus `json:"status,omitempty"`
}

type VirtualMachineScaleSetSpec struct {
	VirtualMachineScaleSetSpec2 `json:",inline"`
	// +optional
	KubeformOutput VirtualMachineScaleSetSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type VirtualMachineScaleSetSpecAdditionalCapabilities struct {
	// +optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled"`
}

type VirtualMachineScaleSetSpecAdditionalUnattendContent struct {
	Content *string `json:"-" sensitive:"true" tf:"content"`
	Setting *string `json:"setting" tf:"setting"`
}

type VirtualMachineScaleSetSpecAutomaticInstanceRepair struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	GracePeriod *string `json:"gracePeriod,omitempty" tf:"grace_period"`
}

type VirtualMachineScaleSetSpecAutomaticOsUpgradePolicy struct {
	DisableAutomaticRollback *bool `json:"disableAutomaticRollback" tf:"disable_automatic_rollback"`
	EnableAutomaticOsUpgrade *bool `json:"enableAutomaticOsUpgrade" tf:"enable_automatic_os_upgrade"`
}

type VirtualMachineScaleSetSpecBootDiagnostics struct {
	// +optional
	StorageAccountURI *string `json:"storageAccountURI,omitempty" tf:"storage_account_uri"`
}

type VirtualMachineScaleSetSpecDataDisk struct {
	Caching *string `json:"caching" tf:"caching"`
	// +optional
	CreateOption *string `json:"createOption,omitempty" tf:"create_option"`
	// +optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id"`
	// +optional
	DiskIopsReadWrite *int64 `json:"diskIopsReadWrite,omitempty" tf:"disk_iops_read_write"`
	// +optional
	DiskMbpsReadWrite  *int64  `json:"diskMbpsReadWrite,omitempty" tf:"disk_mbps_read_write"`
	DiskSizeGb         *int64  `json:"diskSizeGb" tf:"disk_size_gb"`
	Lun                *int64  `json:"lun" tf:"lun"`
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type"`
	// +optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled"`
}

type VirtualMachineScaleSetSpecExtension struct {
	// +optional
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version"`
	// +optional
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag"`
	Name           *string `json:"name" tf:"name"`
	// +optional
	ProtectedSettings *string `json:"-" sensitive:"true" tf:"protected_settings"`
	// +optional
	ProvisionAfterExtensions []string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions"`
	Publisher                *string  `json:"publisher" tf:"publisher"`
	// +optional
	Settings           *string `json:"settings,omitempty" tf:"settings"`
	Type               *string `json:"type" tf:"type"`
	TypeHandlerVersion *string `json:"typeHandlerVersion" tf:"type_handler_version"`
}

type VirtualMachineScaleSetSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	Type        *string `json:"type" tf:"type"`
}

type VirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddressIpTag struct {
	Tag  *string `json:"tag" tf:"tag"`
	Type *string `json:"type" tf:"type"`
}

type VirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddress struct {
	// +optional
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label"`
	// +optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes"`
	// +optional
	IpTag []VirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddressIpTag `json:"ipTag,omitempty" tf:"ip_tag"`
	Name  *string                                                                         `json:"name" tf:"name"`
	// +optional
	PublicIPPrefixID *string `json:"publicIPPrefixID,omitempty" tf:"public_ip_prefix_id"`
}

type VirtualMachineScaleSetSpecNetworkInterfaceIpConfiguration struct {
	// +optional
	ApplicationGatewayBackendAddressPoolIDS []string `json:"applicationGatewayBackendAddressPoolIDS,omitempty" tf:"application_gateway_backend_address_pool_ids"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ApplicationSecurityGroupIDS []string `json:"applicationSecurityGroupIDS,omitempty" tf:"application_security_group_ids"`
	// +optional
	LoadBalancerBackendAddressPoolIDS []string `json:"loadBalancerBackendAddressPoolIDS,omitempty" tf:"load_balancer_backend_address_pool_ids"`
	// +optional
	LoadBalancerInboundNATRulesIDS []string `json:"loadBalancerInboundNATRulesIDS,omitempty" tf:"load_balancer_inbound_nat_rules_ids"`
	Name                           *string  `json:"name" tf:"name"`
	// +optional
	Primary *bool `json:"primary,omitempty" tf:"primary"`
	// +optional
	PublicIPAddress []VirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddress `json:"publicIPAddress,omitempty" tf:"public_ip_address"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type VirtualMachineScaleSetSpecNetworkInterface struct {
	// +optional
	DnsServers []string `json:"dnsServers,omitempty" tf:"dns_servers"`
	// +optional
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking"`
	// +optional
	EnableIPForwarding *bool                                                       `json:"enableIPForwarding,omitempty" tf:"enable_ip_forwarding"`
	IpConfiguration    []VirtualMachineScaleSetSpecNetworkInterfaceIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Name               *string                                                     `json:"name" tf:"name"`
	// +optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupID,omitempty" tf:"network_security_group_id"`
	// +optional
	Primary *bool `json:"primary,omitempty" tf:"primary"`
}

type VirtualMachineScaleSetSpecOsDiskDiffDiskSettings struct {
	Option *string `json:"option" tf:"option"`
}

type VirtualMachineScaleSetSpecOsDisk struct {
	Caching *string `json:"caching" tf:"caching"`
	// +optional
	DiffDiskSettings *VirtualMachineScaleSetSpecOsDiskDiffDiskSettings `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings"`
	// +optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id"`
	// +optional
	DiskSizeGb         *int64  `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type"`
	// +optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled"`
}

type VirtualMachineScaleSetSpecPlan struct {
	Name      *string `json:"name" tf:"name"`
	Product   *string `json:"product" tf:"product"`
	Publisher *string `json:"publisher" tf:"publisher"`
}

type VirtualMachineScaleSetSpecRollingUpgradePolicy struct {
	MaxBatchInstancePercent             *int64  `json:"maxBatchInstancePercent" tf:"max_batch_instance_percent"`
	MaxUnhealthyInstancePercent         *int64  `json:"maxUnhealthyInstancePercent" tf:"max_unhealthy_instance_percent"`
	MaxUnhealthyUpgradedInstancePercent *int64  `json:"maxUnhealthyUpgradedInstancePercent" tf:"max_unhealthy_upgraded_instance_percent"`
	PauseTimeBetweenBatches             *string `json:"pauseTimeBetweenBatches" tf:"pause_time_between_batches"`
}

type VirtualMachineScaleSetSpecSecretCertificate struct {
	Store *string `json:"store" tf:"store"`
	Url   *string `json:"url" tf:"url"`
}

type VirtualMachineScaleSetSpecSecret struct {
	// +kubebuilder:validation:MinItems=1
	Certificate []VirtualMachineScaleSetSpecSecretCertificate `json:"certificate" tf:"certificate"`
	KeyVaultID  *string                                       `json:"keyVaultID" tf:"key_vault_id"`
}

type VirtualMachineScaleSetSpecSourceImageReference struct {
	Offer     *string `json:"offer" tf:"offer"`
	Publisher *string `json:"publisher" tf:"publisher"`
	Sku       *string `json:"sku" tf:"sku"`
	Version   *string `json:"version" tf:"version"`
}

type VirtualMachineScaleSetSpecTerminateNotification struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout"`
}

type VirtualMachineScaleSetSpecWinrmListener struct {
	// +optional
	CertificateURL *string `json:"certificateURL,omitempty" tf:"certificate_url"`
	Protocol       *string `json:"protocol" tf:"protocol"`
}

type VirtualMachineScaleSetSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AdditionalCapabilities *VirtualMachineScaleSetSpecAdditionalCapabilities `json:"additionalCapabilities,omitempty" tf:"additional_capabilities"`
	// +optional
	AdditionalUnattendContent []VirtualMachineScaleSetSpecAdditionalUnattendContent `json:"additionalUnattendContent,omitempty" tf:"additional_unattend_content"`
	AdminPassword             *string                                               `json:"-" sensitive:"true" tf:"admin_password"`
	AdminUsername             *string                                               `json:"adminUsername" tf:"admin_username"`
	// +optional
	AutomaticInstanceRepair *VirtualMachineScaleSetSpecAutomaticInstanceRepair `json:"automaticInstanceRepair,omitempty" tf:"automatic_instance_repair"`
	// +optional
	AutomaticOsUpgradePolicy *VirtualMachineScaleSetSpecAutomaticOsUpgradePolicy `json:"automaticOsUpgradePolicy,omitempty" tf:"automatic_os_upgrade_policy"`
	// +optional
	BootDiagnostics *VirtualMachineScaleSetSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics"`
	// +optional
	ComputerNamePrefix *string `json:"computerNamePrefix,omitempty" tf:"computer_name_prefix"`
	// +optional
	CustomData *string `json:"-" sensitive:"true" tf:"custom_data"`
	// +optional
	DataDisk []VirtualMachineScaleSetSpecDataDisk `json:"dataDisk,omitempty" tf:"data_disk"`
	// +optional
	DoNotRunExtensionsOnOverprovisionedMachines *bool `json:"doNotRunExtensionsOnOverprovisionedMachines,omitempty" tf:"do_not_run_extensions_on_overprovisioned_machines"`
	// +optional
	EnableAutomaticUpdates *bool `json:"enableAutomaticUpdates,omitempty" tf:"enable_automatic_updates"`
	// +optional
	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled"`
	// +optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`
	// +optional
	Extension []VirtualMachineScaleSetSpecExtension `json:"extension,omitempty" tf:"extension"`
	// +optional
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget"`
	// +optional
	HealthProbeID *string `json:"healthProbeID,omitempty" tf:"health_probe_id"`
	// +optional
	Identity  *VirtualMachineScaleSetSpecIdentity `json:"identity,omitempty" tf:"identity"`
	Instances *int64                              `json:"instances" tf:"instances"`
	// +optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`
	Location    *string `json:"location" tf:"location"`
	// +optional
	MaxBidPrice      *float64                                     `json:"maxBidPrice,omitempty" tf:"max_bid_price"`
	Name             *string                                      `json:"name" tf:"name"`
	NetworkInterface []VirtualMachineScaleSetSpecNetworkInterface `json:"networkInterface" tf:"network_interface"`
	OsDisk           *VirtualMachineScaleSetSpecOsDisk            `json:"osDisk" tf:"os_disk"`
	// +optional
	Overprovision *bool `json:"overprovision,omitempty" tf:"overprovision"`
	// +optional
	Plan *VirtualMachineScaleSetSpecPlan `json:"plan,omitempty" tf:"plan"`
	// +optional
	PlatformFaultDomainCount *int64 `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count"`
	// +optional
	Priority *string `json:"priority,omitempty" tf:"priority"`
	// +optional
	ProvisionVmAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent"`
	// +optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty" tf:"proximity_placement_group_id"`
	ResourceGroupName         *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RollingUpgradePolicy *VirtualMachineScaleSetSpecRollingUpgradePolicy `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy"`
	// +optional
	ScaleInPolicy *string `json:"scaleInPolicy,omitempty" tf:"scale_in_policy"`
	// +optional
	Secret []VirtualMachineScaleSetSpecSecret `json:"secret,omitempty" tf:"secret"`
	// +optional
	SinglePlacementGroup *bool   `json:"singlePlacementGroup,omitempty" tf:"single_placement_group"`
	Sku                  *string `json:"sku" tf:"sku"`
	// +optional
	SourceImageID *string `json:"sourceImageID,omitempty" tf:"source_image_id"`
	// +optional
	SourceImageReference *VirtualMachineScaleSetSpecSourceImageReference `json:"sourceImageReference,omitempty" tf:"source_image_reference"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TerminateNotification *VirtualMachineScaleSetSpecTerminateNotification `json:"terminateNotification,omitempty" tf:"terminate_notification"`
	// +optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone"`
	// +optional
	UniqueID *string `json:"uniqueID,omitempty" tf:"unique_id"`
	// +optional
	UpgradeMode *string `json:"upgradeMode,omitempty" tf:"upgrade_mode"`
	// +optional
	WinrmListener []VirtualMachineScaleSetSpecWinrmListener `json:"winrmListener,omitempty" tf:"winrm_listener"`
	// +optional
	ZoneBalance *bool `json:"zoneBalance,omitempty" tf:"zone_balance"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type VirtualMachineScaleSetStatus struct {
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

// VirtualMachineScaleSetList is a list of VirtualMachineScaleSets
type VirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachineScaleSet CRD objects
	Items []VirtualMachineScaleSet `json:"items,omitempty"`
}

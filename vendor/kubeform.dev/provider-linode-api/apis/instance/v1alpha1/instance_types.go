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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpec struct {
	InstanceSpec2 `json:",inline"`
	// +optional
	KubeformOutput *InstanceSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type InstanceSpecAlerts struct {
	// The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. If this is set to 0, the alert is disabled.
	// +optional
	Cpu *int64 `json:"cpu,omitempty" tf:"cpu"`
	// The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0, this alert is disabled.
	// +optional
	Io *int64 `json:"io,omitempty" tf:"io"`
	// The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.
	// +optional
	NetworkIn *int64 `json:"networkIn,omitempty" tf:"network_in"`
	// The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.
	// +optional
	NetworkOut *int64 `json:"networkOut,omitempty" tf:"network_out"`
	// The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.
	// +optional
	TransferQuota *int64 `json:"transferQuota,omitempty" tf:"transfer_quota"`
}

type InstanceSpecBackupsSchedule struct {
	// The day ('Sunday'-'Saturday') of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period.  If not set manually, then when backups are initially enabled, this may come back as 'Scheduling' until the day is automatically selected.
	// +optional
	Day *string `json:"day,omitempty" tf:"day"`
	// The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically.  If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected.
	// +optional
	Window *string `json:"window,omitempty" tf:"window"`
}

type InstanceSpecBackups struct {
	// If this Linode has the Backup service enabled.
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	Schedule []InstanceSpecBackupsSchedule `json:"schedule,omitempty" tf:"schedule"`
}

type InstanceSpecConfigDevicesSda struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevicesSdb struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevicesSdc struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevicesSdd struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevicesSde struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevicesSdf struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevicesSdg struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevicesSdh struct {
	// The Disk ID to map to this disk slot
	// +optional
	DiskID *int64 `json:"diskID,omitempty" tf:"disk_id"`
	// The `label` of the `disk` to map to this `device` slot.
	// +optional
	DiskLabel *string `json:"diskLabel,omitempty" tf:"disk_label"`
	// The Block Storage volume ID to map to this disk slot
	// +optional
	VolumeID *int64 `json:"volumeID,omitempty" tf:"volume_id"`
}

type InstanceSpecConfigDevices struct {
	// +optional
	Sda *InstanceSpecConfigDevicesSda `json:"sda,omitempty" tf:"sda"`
	// +optional
	Sdb *InstanceSpecConfigDevicesSdb `json:"sdb,omitempty" tf:"sdb"`
	// +optional
	Sdc *InstanceSpecConfigDevicesSdc `json:"sdc,omitempty" tf:"sdc"`
	// +optional
	Sdd *InstanceSpecConfigDevicesSdd `json:"sdd,omitempty" tf:"sdd"`
	// +optional
	Sde *InstanceSpecConfigDevicesSde `json:"sde,omitempty" tf:"sde"`
	// +optional
	Sdf *InstanceSpecConfigDevicesSdf `json:"sdf,omitempty" tf:"sdf"`
	// +optional
	Sdg *InstanceSpecConfigDevicesSdg `json:"sdg,omitempty" tf:"sdg"`
	// +optional
	Sdh *InstanceSpecConfigDevicesSdh `json:"sdh,omitempty" tf:"sdh"`
}

type InstanceSpecConfigHelpers struct {
	// Populates the /dev directory early during boot without udev. Defaults to false.
	// +optional
	DevtmpfsAutomount *bool `json:"devtmpfsAutomount,omitempty" tf:"devtmpfs_automount"`
	// Controls the behavior of the Linode Config's Distribution Helper setting.
	// +optional
	Distro *bool `json:"distro,omitempty" tf:"distro"`
	// Creates a modules dependency file for the Kernel you run.
	// +optional
	ModulesDep *bool `json:"modulesDep,omitempty" tf:"modules_dep"`
	// Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.
	// +optional
	Network *bool `json:"network,omitempty" tf:"network"`
	// Disables updatedb cron job to avoid disk thrashing.
	// +optional
	UpdatedbDisabled *bool `json:"updatedbDisabled,omitempty" tf:"updatedb_disabled"`
}

type InstanceSpecConfig struct {
	// Optional field for arbitrary User comments on this Config.
	// +optional
	Comments *string `json:"comments,omitempty" tf:"comments"`
	// Device sda-sdh can be either a Disk or Volume identified by disk_label or volume_id. Only one type per slot allowed.
	// +optional
	Devices *InstanceSpecConfigDevices `json:"devices,omitempty" tf:"devices"`
	// Helpers enabled when booting to this Linode Config.
	// +optional
	Helpers *InstanceSpecConfigHelpers `json:"helpers,omitempty" tf:"helpers"`
	// A Kernel ID to boot a Linode with. Default is based on image choice. (examples: linode/latest-64bit, linode/grub2, linode/direct-disk)
	// +optional
	Kernel *string `json:"kernel,omitempty" tf:"kernel"`
	// The Config's label for display purposes.  Also used by `boot_config_label`.
	Label *string `json:"label" tf:"label"`
	// Defaults to the total RAM of the Linode
	// +optional
	MemoryLimit *int64 `json:"memoryLimit,omitempty" tf:"memory_limit"`
	// The root device to boot. The corresponding disk must be attached.
	// +optional
	RootDevice *string `json:"rootDevice,omitempty" tf:"root_device"`
	// Defines the state of your Linode after booting. Defaults to default.
	// +optional
	RunLevel *string `json:"runLevel,omitempty" tf:"run_level"`
	// Controls the virtualization mode. Defaults to paravirt.
	// +optional
	VirtMode *string `json:"virtMode,omitempty" tf:"virt_mode"`
}

type InstanceSpecDisk struct {
	// A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys"`
	// A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users"`
	// The Disk filesystem can be one of: raw, swap, ext3, ext4, initrd (max 32mb)
	// +optional
	Filesystem *string `json:"filesystem,omitempty" tf:"filesystem"`
	// The ID of the Disk (for use in Linode Image resources and Linode Instance Config Devices)
	// +optional
	ID *int64 `json:"ID,omitempty" tf:"id"`
	// An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/.
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// The disks label, which acts as an identifier in Terraform.
	Label *string `json:"label" tf:"label"`
	// If true, this Disk is read-only.
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only"`
	// The password that will be initialially assigned to the 'root' user account.
	// +optional
	RootPass *string `json:"-" sensitive:"true" tf:"root_pass"`
	// The size of the Disk in MB.
	Size *int64 `json:"size" tf:"size"`
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
	// +optional
	StackscriptData map[string]string `json:"-" sensitive:"true" tf:"stackscript_data"`
	// The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.
	// +optional
	StackscriptID *int64 `json:"stackscriptID,omitempty" tf:"stackscript_id"`
}

type InstanceSpecSpecs struct {
	// The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image without specifying disks.
	// +optional
	Disk *int64 `json:"disk,omitempty" tf:"disk"`
	// The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.
	// +optional
	Memory *int64 `json:"memory,omitempty" tf:"memory"`
	// The amount of network transfer this Linode is allotted each month.
	// +optional
	Transfer *int64 `json:"transfer,omitempty" tf:"transfer"`
	// The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.
	// +optional
	Vcpus *int64 `json:"vcpus,omitempty" tf:"vcpus"`
}

type InstanceSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Alerts *InstanceSpecAlerts `json:"alerts,omitempty" tf:"alerts"`
	// A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is provided.
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys"`
	// A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users"`
	// A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive.
	// +optional
	BackupID *int64 `json:"backupID,omitempty" tf:"backup_id"`
	// Information about this Linode's backups status.
	// +optional
	Backups []InstanceSpecBackups `json:"backups,omitempty" tf:"backups"`
	// If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
	// +optional
	BackupsEnabled *bool `json:"backupsEnabled,omitempty" tf:"backups_enabled"`
	// The Label of the Instance Config that should be used to boot the Linode instance.
	// +optional
	BootConfigLabel *string `json:"bootConfigLabel,omitempty" tf:"boot_config_label"`
	// Configuration profiles define the VM settings and boot behavior of the Linode Instance.
	// +optional
	Config []InstanceSpecConfig `json:"config,omitempty" tf:"config"`
	// +optional
	Disk []InstanceSpecDisk `json:"disk,omitempty" tf:"disk"`
	// The display group of the Linode instance.
	// +optional
	Group *string `json:"group,omitempty" tf:"group"`
	// An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use.
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary address will be used for this field.
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.
	// +optional
	Ipv4 []string `json:"ipv4,omitempty" tf:"ipv4"`
	// This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.
	// +optional
	Ipv6 *string `json:"ipv6,omitempty" tf:"ipv6"`
	// The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region.
	// +optional
	PrivateIP *bool `json:"privateIP,omitempty" tf:"private_ip"`
	// This Linode's Private IPv4 Address.  The regional private IP address range is 192.168.128/17 address shared by all Linode Instances in a region.
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// This is the location where the Linode was deployed. This cannot be changed without opening a support ticket.
	Region *string `json:"region" tf:"region"`
	// The password that will be initialially assigned to the 'root' user account.
	// +optional
	RootPass *string `json:"-" sensitive:"true" tf:"root_pass"`
	// +optional
	Specs []InstanceSpecSpecs `json:"specs,omitempty" tf:"specs"`
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
	// +optional
	StackscriptData map[string]string `json:"-" sensitive:"true" tf:"stackscript_data"`
	// The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.
	// +optional
	StackscriptID *int64 `json:"stackscriptID,omitempty" tf:"stackscript_id"`
	// The status of the instance, indicating the current readiness state.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.
	// +optional
	SwapSize *int64 `json:"swapSize,omitempty" tf:"swap_size"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The type of instance to be deployed, determining the price and size.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
	// +optional
	WatchdogEnabled *bool `json:"watchdogEnabled,omitempty" tf:"watchdog_enabled"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}

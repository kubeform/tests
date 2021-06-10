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

type FromInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FromInstanceSpec   `json:"spec,omitempty"`
	Status            FromInstanceStatus `json:"status,omitempty"`
}

type FromInstanceSpec struct {
	FromInstanceSpec2 `json:",inline"`
	// +optional
	KubeformOutput FromInstanceSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type FromInstanceSpecEbsBlockDevice struct {
	// +optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// +optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty" tf:"snapshot_id"`
	// +optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
	// +optional
	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`
	// +optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type FromInstanceSpecEphemeralBlockDevice struct {
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// +optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name"`
}

type FromInstanceSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	EbsBlockDevice []FromInstanceSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device"`
	// +optional
	EnaSupport *bool `json:"enaSupport,omitempty" tf:"ena_support"`
	// +optional
	EphemeralBlockDevice []FromInstanceSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device"`
	// +optional
	Hypervisor *string `json:"hypervisor,omitempty" tf:"hypervisor"`
	// +optional
	ImageLocation *string `json:"imageLocation,omitempty" tf:"image_location"`
	// +optional
	ImageOwnerAlias *string `json:"imageOwnerAlias,omitempty" tf:"image_owner_alias"`
	// +optional
	ImageType *string `json:"imageType,omitempty" tf:"image_type"`
	// +optional
	KernelID *string `json:"kernelID,omitempty" tf:"kernel_id"`
	// +optional
	ManageEbsSnapshots *bool   `json:"manageEbsSnapshots,omitempty" tf:"manage_ebs_snapshots"`
	Name               *string `json:"name" tf:"name"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	Platform *string `json:"platform,omitempty" tf:"platform"`
	// +optional
	PlatformDetails *string `json:"platformDetails,omitempty" tf:"platform_details"`
	// +optional
	Public *bool `json:"public,omitempty" tf:"public"`
	// +optional
	RamdiskID *string `json:"ramdiskID,omitempty" tf:"ramdisk_id"`
	// +optional
	RootDeviceName *string `json:"rootDeviceName,omitempty" tf:"root_device_name"`
	// +optional
	RootSnapshotID *string `json:"rootSnapshotID,omitempty" tf:"root_snapshot_id"`
	// +optional
	SnapshotWithoutReboot *bool   `json:"snapshotWithoutReboot,omitempty" tf:"snapshot_without_reboot"`
	SourceInstanceID      *string `json:"sourceInstanceID" tf:"source_instance_id"`
	// +optional
	SriovNetSupport *string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UsageOperation *string `json:"usageOperation,omitempty" tf:"usage_operation"`
	// +optional
	VirtualizationType *string `json:"virtualizationType,omitempty" tf:"virtualization_type"`
}

type FromInstanceStatus struct {
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

// FromInstanceList is a list of FromInstances
type FromInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FromInstance CRD objects
	Items []FromInstance `json:"items,omitempty"`
}

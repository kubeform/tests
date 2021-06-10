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

type Image struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec,omitempty"`
	Status            ImageStatus `json:"status,omitempty"`
}

type ImageSpec struct {
	ImageSpec2 `json:",inline"`
	// +optional
	KubeformOutput ImageSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ImageSpecGuestOsFeatures struct {
	// The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: ["MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC"]
	Type *string `json:"type" tf:"type"`
}

type ImageSpecRawDisk struct {
	// The format used to encode and transmit the block device, which
	// should be TAR. This is just a container and transmission format
	// and not a runtime format. Provided by the client when the disk
	// image is created. Default value: "TAR" Possible values: ["TAR"]
	// +optional
	ContainerType *string `json:"containerType,omitempty" tf:"container_type"`
	// An optional SHA1 checksum of the disk image before unpackaging.
	// This is provided by the client when the disk image is created.
	// +optional
	Sha1 *string `json:"sha1,omitempty" tf:"sha1"`
	// The full Google Cloud Storage URL where disk storage is stored
	// You must provide either this property or the sourceDisk property
	// but not both.
	Source *string `json:"source" tf:"source"`
}

type ImageSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Size of the image tar.gz archive stored in Google Cloud Storage (in
	// bytes).
	// +optional
	ArchiveSizeBytes *int64 `json:"archiveSizeBytes,omitempty" tf:"archive_size_bytes"`
	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	// +optional
	Family *string `json:"family,omitempty" tf:"family"`
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// +optional
	GuestOsFeatures []ImageSpecGuestOsFeatures `json:"guestOsFeatures,omitempty" tf:"guest_os_features"`
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	// +optional
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint"`
	// Labels to apply to this Image.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Any applicable license URI.
	// +optional
	Licenses []string `json:"licenses,omitempty" tf:"licenses"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The parameters of the raw disk image.
	// +optional
	RawDisk *ImageSpecRawDisk `json:"rawDisk,omitempty" tf:"raw_disk"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	// +optional
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	//
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	// +optional
	SourceImage *string `json:"sourceImage,omitempty" tf:"source_image"`
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	//
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	// +optional
	SourceSnapshot *string `json:"sourceSnapshot,omitempty" tf:"source_snapshot"`
}

type ImageStatus struct {
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

// ImageList is a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Image CRD objects
	Items []Image `json:"items,omitempty"`
}

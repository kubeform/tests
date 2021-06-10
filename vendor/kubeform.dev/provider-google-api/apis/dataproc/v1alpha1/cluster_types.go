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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpec struct {
	ClusterSpec2 `json:",inline"`
	// +optional
	KubeformOutput ClusterSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ClusterSpecClusterConfigAutoscalingConfig struct {
	// The autoscaling policy used by the cluster.
	PolicyURI *string `json:"policyURI" tf:"policy_uri"`
}

type ClusterSpecClusterConfigEncryptionConfig struct {
	// The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.
	KmsKeyName *string `json:"kmsKeyName" tf:"kms_key_name"`
}

type ClusterSpecClusterConfigGceClusterConfig struct {
	// By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. If set to true, all instances in the cluster will only have internal IP addresses. Note: Private Google Access (also known as privateIpGoogleAccess) must be enabled on the subnetwork that the cluster will be launched in.
	// +optional
	InternalIPOnly *bool `json:"internalIPOnly,omitempty" tf:"internal_ip_only"`
	// A map of the Compute Engine metadata entries to add to all instances
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
	// The name or self_link of the Google Compute Engine network to the cluster will be part of. Conflicts with subnetwork. If neither is specified, this defaults to the "default" network.
	// +optional
	Network *string `json:"network,omitempty" tf:"network"`
	// The service account to be used by the Node VMs. If not specified, the "default" service account is used.
	// +optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account"`
	// The set of Google API scopes to be made available on all of the node VMs under the service_account specified. These can be either FQDNs, or scope aliases.
	// +optional
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty" tf:"service_account_scopes"`
	// The name or self_link of the Google Compute Engine subnetwork the cluster will be part of. Conflicts with network.
	// +optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork"`
	// The list of instance tags applied to instances in the cluster. Tags are used to identify valid sources or targets for network firewalls.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The GCP zone where your data is stored and used (i.e. where the master and the worker nodes will be created in). If region is set to 'global' (default) then zone is mandatory, otherwise GCP is able to make use of Auto Zone Placement to determine this automatically for you. Note: This setting additionally determines and restricts which computing resources are available for use with other configs such as cluster_config.master_config.machine_type and cluster_config.worker_config.machine_type.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type ClusterSpecClusterConfigInitializationAction struct {
	// The script to be executed during initialization of the cluster. The script must be a GCS file with a gs:// prefix.
	Script *string `json:"script" tf:"script"`
	// The maximum duration (in seconds) which script is allowed to take to execute its action. GCP will default to a predetermined computed value if not set (currently 300).
	// +optional
	TimeoutSec *int64 `json:"timeoutSec,omitempty" tf:"timeout_sec"`
}

type ClusterSpecClusterConfigMasterConfigAccelerators struct {
	// The number of the accelerator cards of this type exposed to this instance. Often restricted to one of 1, 2, 4, or 8.
	AcceleratorCount *int64 `json:"acceleratorCount" tf:"accelerator_count"`
	// The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.
	AcceleratorType *string `json:"acceleratorType" tf:"accelerator_type"`
}

type ClusterSpecClusterConfigMasterConfigDiskConfig struct {
	// Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.
	// +optional
	BootDiskSizeGb *int64 `json:"bootDiskSizeGb,omitempty" tf:"boot_disk_size_gb"`
	// The disk type of the primary disk attached to each node. One of "pd-ssd" or "pd-standard". Defaults to "pd-standard".
	// +optional
	BootDiskType *string `json:"bootDiskType,omitempty" tf:"boot_disk_type"`
	// The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.
	// +optional
	NumLocalSsds *int64 `json:"numLocalSsds,omitempty" tf:"num_local_ssds"`
}

type ClusterSpecClusterConfigMasterConfig struct {
	// The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.
	// +optional
	Accelerators []ClusterSpecClusterConfigMasterConfigAccelerators `json:"accelerators,omitempty" tf:"accelerators"`
	// Disk Config
	// +optional
	DiskConfig *ClusterSpecClusterConfigMasterConfigDiskConfig `json:"diskConfig,omitempty" tf:"disk_config"`
	// The URI for the image to use for this master/worker
	// +optional
	ImageURI *string `json:"imageURI,omitempty" tf:"image_uri"`
	// List of master/worker instance names which have been assigned to the cluster.
	// +optional
	InstanceNames []string `json:"instanceNames,omitempty" tf:"instance_names"`
	// The name of a Google Compute Engine machine type to create for the master/worker
	// +optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type"`
	// The name of a minimum generation of CPU family for the master/worker. If not specified, GCP will default to a predetermined computed value for each zone.
	// +optional
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform"`
	// Specifies the number of master/worker nodes to create. If not specified, GCP will default to a predetermined computed value.
	// +optional
	NumInstances *int64 `json:"numInstances,omitempty" tf:"num_instances"`
}

type ClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig struct {
	// Size of the primary disk attached to each preemptible worker node, specified in GB. The smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.
	// +optional
	BootDiskSizeGb *int64 `json:"bootDiskSizeGb,omitempty" tf:"boot_disk_size_gb"`
	// The disk type of the primary disk attached to each preemptible worker node. One of "pd-ssd" or "pd-standard". Defaults to "pd-standard".
	// +optional
	BootDiskType *string `json:"bootDiskType,omitempty" tf:"boot_disk_type"`
	// The amount of local SSD disks that will be attached to each preemptible worker node. Defaults to 0.
	// +optional
	NumLocalSsds *int64 `json:"numLocalSsds,omitempty" tf:"num_local_ssds"`
}

type ClusterSpecClusterConfigPreemptibleWorkerConfig struct {
	// Disk Config
	// +optional
	DiskConfig *ClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig `json:"diskConfig,omitempty" tf:"disk_config"`
	// List of preemptible instance names which have been assigned to the cluster.
	// +optional
	InstanceNames []string `json:"instanceNames,omitempty" tf:"instance_names"`
	// Specifies the number of preemptible nodes to create. Defaults to 0.
	// +optional
	NumInstances *int64 `json:"numInstances,omitempty" tf:"num_instances"`
}

type ClusterSpecClusterConfigSecurityConfigKerberosConfig struct {
	// The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship.
	// +optional
	CrossRealmTrustAdminServer *string `json:"crossRealmTrustAdminServer,omitempty" tf:"cross_realm_trust_admin_server"`
	// The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship.
	// +optional
	CrossRealmTrustKdc *string `json:"crossRealmTrustKdc,omitempty" tf:"cross_realm_trust_kdc"`
	// The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust.
	// +optional
	CrossRealmTrustRealm *string `json:"crossRealmTrustRealm,omitempty" tf:"cross_realm_trust_realm"`
	// The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster
	// Kerberos realm and the remote trusted realm, in a cross realm trust relationship.
	// +optional
	CrossRealmTrustSharedPasswordURI *string `json:"crossRealmTrustSharedPasswordURI,omitempty" tf:"cross_realm_trust_shared_password_uri"`
	// Flag to indicate whether to Kerberize the cluster.
	// +optional
	EnableKerberos *bool `json:"enableKerberos,omitempty" tf:"enable_kerberos"`
	// The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database.
	// +optional
	KdcDbKeyURI *string `json:"kdcDbKeyURI,omitempty" tf:"kdc_db_key_uri"`
	// The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc.
	// +optional
	KeyPasswordURI *string `json:"keyPasswordURI,omitempty" tf:"key_password_uri"`
	// The Cloud Storage URI of a KMS encrypted file containing
	// the password to the user provided keystore. For the self-signed certificate, this password is generated
	// by Dataproc
	// +optional
	KeystorePasswordURI *string `json:"keystorePasswordURI,omitempty" tf:"keystore_password_uri"`
	// The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.
	// +optional
	KeystoreURI *string `json:"keystoreURI,omitempty" tf:"keystore_uri"`
	// The uri of the KMS key used to encrypt various sensitive files.
	KmsKeyURI *string `json:"kmsKeyURI" tf:"kms_key_uri"`
	// The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm.
	// +optional
	Realm *string `json:"realm,omitempty" tf:"realm"`
	// The cloud Storage URI of a KMS encrypted file containing the root principal password.
	RootPrincipalPasswordURI *string `json:"rootPrincipalPasswordURI" tf:"root_principal_password_uri"`
	// The lifetime of the ticket granting ticket, in hours.
	// +optional
	TgtLifetimeHours *int64 `json:"tgtLifetimeHours,omitempty" tf:"tgt_lifetime_hours"`
	// The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc.
	// +optional
	TruststorePasswordURI *string `json:"truststorePasswordURI,omitempty" tf:"truststore_password_uri"`
	// The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.
	// +optional
	TruststoreURI *string `json:"truststoreURI,omitempty" tf:"truststore_uri"`
}

type ClusterSpecClusterConfigSecurityConfig struct {
	// Kerberos related configuration
	KerberosConfig *ClusterSpecClusterConfigSecurityConfigKerberosConfig `json:"kerberosConfig" tf:"kerberos_config"`
}

type ClusterSpecClusterConfigSoftwareConfig struct {
	// The Cloud Dataproc image version to use for the cluster - this controls the sets of software versions installed onto the nodes when you create clusters. If not specified, defaults to the latest version.
	// +optional
	ImageVersion *string `json:"imageVersion,omitempty" tf:"image_version"`
	// The set of optional components to activate on the cluster.
	// +optional
	OptionalComponents []string `json:"optionalComponents,omitempty" tf:"optional_components"`
	// A list of override and additional properties (key/value pairs) used to modify various aspects of the common configuration files used when creating a cluster.
	// +optional
	OverrideProperties *map[string]string `json:"overrideProperties,omitempty" tf:"override_properties"`
	// A list of the properties used to set the daemon config files. This will include any values supplied by the user via cluster_config.software_config.override_properties
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
}

type ClusterSpecClusterConfigWorkerConfigAccelerators struct {
	// The number of the accelerator cards of this type exposed to this instance. Often restricted to one of 1, 2, 4, or 8.
	AcceleratorCount *int64 `json:"acceleratorCount" tf:"accelerator_count"`
	// The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.
	AcceleratorType *string `json:"acceleratorType" tf:"accelerator_type"`
}

type ClusterSpecClusterConfigWorkerConfigDiskConfig struct {
	// Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.
	// +optional
	BootDiskSizeGb *int64 `json:"bootDiskSizeGb,omitempty" tf:"boot_disk_size_gb"`
	// The disk type of the primary disk attached to each node. One of "pd-ssd" or "pd-standard". Defaults to "pd-standard".
	// +optional
	BootDiskType *string `json:"bootDiskType,omitempty" tf:"boot_disk_type"`
	// The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.
	// +optional
	NumLocalSsds *int64 `json:"numLocalSsds,omitempty" tf:"num_local_ssds"`
}

type ClusterSpecClusterConfigWorkerConfig struct {
	// The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.
	// +optional
	Accelerators []ClusterSpecClusterConfigWorkerConfigAccelerators `json:"accelerators,omitempty" tf:"accelerators"`
	// Disk Config
	// +optional
	DiskConfig *ClusterSpecClusterConfigWorkerConfigDiskConfig `json:"diskConfig,omitempty" tf:"disk_config"`
	// The URI for the image to use for this master/worker
	// +optional
	ImageURI *string `json:"imageURI,omitempty" tf:"image_uri"`
	// List of master/worker instance names which have been assigned to the cluster.
	// +optional
	InstanceNames []string `json:"instanceNames,omitempty" tf:"instance_names"`
	// The name of a Google Compute Engine machine type to create for the master/worker
	// +optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type"`
	// The name of a minimum generation of CPU family for the master/worker. If not specified, GCP will default to a predetermined computed value for each zone.
	// +optional
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform"`
	// Specifies the number of master/worker nodes to create. If not specified, GCP will default to a predetermined computed value.
	// +optional
	NumInstances *int64 `json:"numInstances,omitempty" tf:"num_instances"`
}

type ClusterSpecClusterConfig struct {
	// The autoscaling policy config associated with the cluster.
	// +optional
	AutoscalingConfig *ClusterSpecClusterConfigAutoscalingConfig `json:"autoscalingConfig,omitempty" tf:"autoscaling_config"`
	//  The name of the cloud storage bucket ultimately used to house the staging data for the cluster. If staging_bucket is specified, it will contain this value, otherwise it will be the auto generated name.
	// +optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`
	// The Customer managed encryption keys settings for the cluster.
	// +optional
	EncryptionConfig *ClusterSpecClusterConfigEncryptionConfig `json:"encryptionConfig,omitempty" tf:"encryption_config"`
	// Common config settings for resources of Google Compute Engine cluster instances, applicable to all instances in the cluster.
	// +optional
	GceClusterConfig *ClusterSpecClusterConfigGceClusterConfig `json:"gceClusterConfig,omitempty" tf:"gce_cluster_config"`
	// Commands to execute on each node after config is completed. You can specify multiple versions of these.
	// +optional
	InitializationAction []ClusterSpecClusterConfigInitializationAction `json:"initializationAction,omitempty" tf:"initialization_action"`
	// The Google Compute Engine config settings for the master/worker instances in a cluster.
	// +optional
	MasterConfig *ClusterSpecClusterConfigMasterConfig `json:"masterConfig,omitempty" tf:"master_config"`
	// The Google Compute Engine config settings for the additional (aka preemptible) instances in a cluster.
	// +optional
	PreemptibleWorkerConfig *ClusterSpecClusterConfigPreemptibleWorkerConfig `json:"preemptibleWorkerConfig,omitempty" tf:"preemptible_worker_config"`
	// Security related configuration.
	// +optional
	SecurityConfig *ClusterSpecClusterConfigSecurityConfig `json:"securityConfig,omitempty" tf:"security_config"`
	// The config settings for software inside the cluster.
	// +optional
	SoftwareConfig *ClusterSpecClusterConfigSoftwareConfig `json:"softwareConfig,omitempty" tf:"software_config"`
	// The Cloud Storage staging bucket used to stage files, such as Hadoop jars, between client machines and the cluster. Note: If you don't explicitly specify a staging_bucket then GCP will auto create / assign one for you. However, you are not guaranteed an auto generated bucket which is solely dedicated to your cluster; it may be shared with other clusters in the same region/zone also choosing to use the auto generation option.
	// +optional
	StagingBucket *string `json:"stagingBucket,omitempty" tf:"staging_bucket"`
	// The Cloud Storage temp bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. Note: If you don't explicitly specify a temp_bucket then GCP will auto create / assign one for you.
	// +optional
	TempBucket *string `json:"tempBucket,omitempty" tf:"temp_bucket"`
	// The Google Compute Engine config settings for the master/worker instances in a cluster.
	// +optional
	WorkerConfig *ClusterSpecClusterConfigWorkerConfig `json:"workerConfig,omitempty" tf:"worker_config"`
}

type ClusterSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Allows you to configure various aspects of the cluster.
	// +optional
	ClusterConfig *ClusterSpecClusterConfig `json:"clusterConfig,omitempty" tf:"cluster_config"`
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a terraform apply
	// +optional
	GracefulDecommissionTimeout *string `json:"gracefulDecommissionTimeout,omitempty" tf:"graceful_decommission_timeout"`
	// The list of labels (key/value pairs) to be applied to instances in the cluster. GCP generates some itself including goog-dataproc-cluster-name which is the name of the cluster.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The name of the cluster, unique within the project and zone.
	Name *string `json:"name" tf:"name"`
	// The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The region in which the cluster and associated nodes will be created in. Defaults to global.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
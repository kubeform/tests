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

type RailsAppLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RailsAppLayerSpec   `json:"spec,omitempty"`
	Status            RailsAppLayerStatus `json:"status,omitempty"`
}

type RailsAppLayerSpec struct {
	RailsAppLayerSpec2 `json:",inline"`
	// +optional
	KubeformOutput RailsAppLayerSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type RailsAppLayerSpecEbsVolume struct {
	// +optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`
	// +optional
	Iops          *int64  `json:"iops,omitempty" tf:"iops"`
	MountPoint    *string `json:"mountPoint" tf:"mount_point"`
	NumberOfDisks *int64  `json:"numberOfDisks" tf:"number_of_disks"`
	// +optional
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level"`
	Size      *int64  `json:"size" tf:"size"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type RailsAppLayerSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppServer *string `json:"appServer,omitempty" tf:"app_server"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutoAssignElasticIPS *bool `json:"autoAssignElasticIPS,omitempty" tf:"auto_assign_elastic_ips"`
	// +optional
	AutoAssignPublicIPS *bool `json:"autoAssignPublicIPS,omitempty" tf:"auto_assign_public_ips"`
	// +optional
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing"`
	// +optional
	BundlerVersion *string `json:"bundlerVersion,omitempty" tf:"bundler_version"`
	// +optional
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes"`
	// +optional
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes"`
	// +optional
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn"`
	// +optional
	CustomJSON *string `json:"customJSON,omitempty" tf:"custom_json"`
	// +optional
	CustomSecurityGroupIDS []string `json:"customSecurityGroupIDS,omitempty" tf:"custom_security_group_ids"`
	// +optional
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes"`
	// +optional
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes"`
	// +optional
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes"`
	// +optional
	DrainElbOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown"`
	// +optional
	EbsVolume []RailsAppLayerSpecEbsVolume `json:"ebsVolume,omitempty" tf:"ebs_volume"`
	// +optional
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer"`
	// +optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot"`
	// +optional
	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout"`
	// +optional
	ManageBundler *bool `json:"manageBundler,omitempty" tf:"manage_bundler"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PassengerVersion *string `json:"passengerVersion,omitempty" tf:"passenger_version"`
	// +optional
	RubyVersion *string `json:"rubyVersion,omitempty" tf:"ruby_version"`
	// +optional
	RubygemsVersion *string `json:"rubygemsVersion,omitempty" tf:"rubygems_version"`
	StackID         *string `json:"stackID" tf:"stack_id"`
	// +optional
	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UseEbsOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances"`
}

type RailsAppLayerStatus struct {
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

// RailsAppLayerList is a list of RailsAppLayers
type RailsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RailsAppLayer CRD objects
	Items []RailsAppLayer `json:"items,omitempty"`
}

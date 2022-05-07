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

type Elb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElbSpec   `json:"spec,omitempty"`
	Status            ElbStatus `json:"status,omitempty"`
}

type ElbSpecAccessLogs struct {
	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	BucketPrefix string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Interval int64 `json:"interval,omitempty" tf:"interval,omitempty"`
}

type ElbSpecHealthCheck struct {
	HealthyThreshold   int64  `json:"healthyThreshold" tf:"healthy_threshold"`
	Interval           int64  `json:"interval" tf:"interval"`
	Target             string `json:"target" tf:"target"`
	Timeout            int64  `json:"timeout" tf:"timeout"`
	UnhealthyThreshold int64  `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type ElbSpecListener struct {
	InstancePort     int64  `json:"instancePort" tf:"instance_port"`
	InstanceProtocol string `json:"instanceProtocol" tf:"instance_protocol"`
	LbPort           int64  `json:"lbPort" tf:"lb_port"`
	LbProtocol       string `json:"lbProtocol" tf:"lb_protocol"`
	// +optional
	SslCertificateID string `json:"sslCertificateID,omitempty" tf:"ssl_certificate_id,omitempty"`
}

type ElbSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogs []ElbSpecAccessLogs `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	ConnectionDraining bool `json:"connectionDraining,omitempty" tf:"connection_draining,omitempty"`
	// +optional
	ConnectionDrainingTimeout int64 `json:"connectionDrainingTimeout,omitempty" tf:"connection_draining_timeout,omitempty"`
	// +optional
	CrossZoneLoadBalancing bool `json:"crossZoneLoadBalancing,omitempty" tf:"cross_zone_load_balancing,omitempty"`
	// +optional
	DnsName string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheck []ElbSpecHealthCheck `json:"healthCheck,omitempty" tf:"health_check,omitempty"`
	// +optional
	IdleTimeout int64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`
	// +optional
	Instances []string `json:"instances,omitempty" tf:"instances,omitempty"`
	// +optional
	Internal bool              `json:"internal,omitempty" tf:"internal,omitempty"`
	Listener []ElbSpecListener `json:"listener" tf:"listener"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SourceSecurityGroup string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`
	// +optional
	SourceSecurityGroupID string `json:"sourceSecurityGroupID,omitempty" tf:"source_security_group_id,omitempty"`
	// +optional
	Subnets []string `json:"subnets,omitempty" tf:"subnets,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ZoneID string `json:"zoneID,omitempty" tf:"zone_id,omitempty"`
}

type ElbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElbSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElbList is a list of Elbs
type ElbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Elb CRD objects
	Items []Elb `json:"items,omitempty"`
}
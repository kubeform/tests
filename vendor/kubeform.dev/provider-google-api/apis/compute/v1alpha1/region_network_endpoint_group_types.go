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

type RegionNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status            RegionNetworkEndpointGroupStatus `json:"status,omitempty"`
}

type RegionNetworkEndpointGroupSpec struct {
	RegionNetworkEndpointGroupSpec2 `json:",inline"`
	// +optional
	KubeformOutput RegionNetworkEndpointGroupSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type RegionNetworkEndpointGroupSpecAppEngine struct {
	// Optional serving service.
	// The service name must be 1-63 characters long, and comply with RFC1035.
	// Example value: "default", "my-service".
	// +optional
	Service *string `json:"service,omitempty" tf:"service"`
	// A template to parse service and version fields from a request URL.
	// URL mask allows for routing to multiple App Engine services without
	// having to create multiple Network Endpoint Groups and backend services.
	//
	// For example, the request URLs "foo1-dot-appname.appspot.com/v1" and
	// "foo1-dot-appname.appspot.com/v2" can be backed by the same Serverless NEG with
	// URL mask "-dot-appname.appspot.com/". The URL mask will parse
	// them to { service = "foo1", version = "v1" } and { service = "foo1", version = "v2" } respectively.
	// +optional
	UrlMask *string `json:"urlMask,omitempty" tf:"url_mask"`
	// Optional serving version.
	// The version must be 1-63 characters long, and comply with RFC1035.
	// Example value: "v1", "v2".
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type RegionNetworkEndpointGroupSpecCloudFunction struct {
	// A user-defined name of the Cloud Function.
	// The function name is case-sensitive and must be 1-63 characters long.
	// Example value: "func1".
	// +optional
	Function *string `json:"function,omitempty" tf:"function"`
	// A template to parse function field from a request URL. URL mask allows
	// for routing to multiple Cloud Functions without having to create
	// multiple Network Endpoint Groups and backend services.
	//
	// For example, request URLs "mydomain.com/function1" and "mydomain.com/function2"
	// can be backed by the same Serverless NEG with URL mask "/". The URL mask
	// will parse them to { function = "function1" } and { function = "function2" } respectively.
	// +optional
	UrlMask *string `json:"urlMask,omitempty" tf:"url_mask"`
}

type RegionNetworkEndpointGroupSpecCloudRun struct {
	// Cloud Run service is the main resource of Cloud Run.
	// The service must be 1-63 characters long, and comply with RFC1035.
	// Example value: "run-service".
	// +optional
	Service *string `json:"service,omitempty" tf:"service"`
	// Cloud Run tag represents the "named-revision" to provide
	// additional fine-grained traffic routing information.
	// The tag must be 1-63 characters long, and comply with RFC1035.
	// Example value: "revision-0010".
	// +optional
	Tag *string `json:"tag,omitempty" tf:"tag"`
	// A template to parse service and tag fields from a request URL.
	// URL mask allows for routing to multiple Run services without having
	// to create multiple network endpoint groups and backend services.
	//
	// For example, request URLs "foo1.domain.com/bar1" and "foo1.domain.com/bar2"
	// an be backed by the same Serverless Network Endpoint Group (NEG) with
	// URL mask ".domain.com/". The URL mask will parse them to { service="bar1", tag="foo1" }
	// and { service="bar2", tag="foo2" } respectively.
	// +optional
	UrlMask *string `json:"urlMask,omitempty" tf:"url_mask"`
}

type RegionNetworkEndpointGroupSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine or cloud_function may be set.
	// +optional
	AppEngine *RegionNetworkEndpointGroupSpecAppEngine `json:"appEngine,omitempty" tf:"app_engine"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine or cloud_function may be set.
	// +optional
	CloudFunction *RegionNetworkEndpointGroupSpecCloudFunction `json:"cloudFunction,omitempty" tf:"cloud_function"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine or cloud_function may be set.
	// +optional
	CloudRun *RegionNetworkEndpointGroupSpecCloudRun `json:"cloudRun,omitempty" tf:"cloud_run"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS Default value: "SERVERLESS" Possible values: ["SERVERLESS"]
	// +optional
	NetworkEndpointType *string `json:"networkEndpointType,omitempty" tf:"network_endpoint_type"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// A reference to the region where the Serverless NEGs Reside.
	Region *string `json:"region" tf:"region"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
}

type RegionNetworkEndpointGroupStatus struct {
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

// RegionNetworkEndpointGroupList is a list of RegionNetworkEndpointGroups
type RegionNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RegionNetworkEndpointGroup CRD objects
	Items []RegionNetworkEndpointGroup `json:"items,omitempty"`
}

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

type RunDomainMapping struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RunDomainMappingSpec   `json:"spec,omitempty"`
	Status            RunDomainMappingStatus `json:"status,omitempty"`
}

type RunDomainMappingSpec struct {
	RunDomainMappingSpec2 `json:",inline"`
	// +optional
	KubeformOutput RunDomainMappingSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type RunDomainMappingSpecMetadata struct {
	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: http://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	// +optional
	Annotations *map[string]string `json:"annotations,omitempty" tf:"annotations"`
	// A sequence number representing a specific generation of the desired state.
	// +optional
	Generation *int64 `json:"generation,omitempty" tf:"generation"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// In Cloud Run the namespace must be equal to either the
	// project ID or project number.
	Namespace *string `json:"namespace" tf:"namespace"`
	// An opaque value that represents the internal version of this object that
	// can be used by clients to determine when objects have changed. May be used
	// for optimistic concurrency, change detection, and the watch operation on a
	// resource or set of resources. They may only be valid for a
	// particular resource or set of resources.
	//
	// More info:
	// https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version"`
	// SelfLink is a URL representing this object.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// UID is a unique id generated by the server on successful creation of a resource and is not
	// allowed to change on PUT operations.
	//
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
}

type RunDomainMappingSpecSpec struct {
	// The mode of the certificate. Default value: "AUTOMATIC" Possible values: ["NONE", "AUTOMATIC"]
	// +optional
	CertificateMode *string `json:"certificateMode,omitempty" tf:"certificate_mode"`
	// If set, the mapping will override any mapping set before this spec was set.
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	// +optional
	ForceOverride *bool `json:"forceOverride,omitempty" tf:"force_override"`
	// The name of the Cloud Run Service that this DomainMapping applies to.
	// The route must exist.
	RouteName *string `json:"routeName" tf:"route_name"`
}

type RunDomainMappingSpecStatusConditions struct {
	// Human readable message indicating details about the current status.
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// One-word CamelCase reason for the condition's current status.
	// +optional
	Reason *string `json:"reason,omitempty" tf:"reason"`
	// Status of the condition, one of True, False, Unknown.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Type of domain mapping condition.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type RunDomainMappingSpecStatusResourceRecords struct {
	// Relative name of the object affected by this record. Only applicable for
	// 'CNAME' records. Example: 'www'.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Data for this record. Values vary by record type, as defined in RFC 1035
	// (section 5) and RFC 1034 (section 3.6.1).
	// +optional
	Rrdata *string `json:"rrdata,omitempty" tf:"rrdata"`
	// Resource record type. Example: 'AAAA'. Possible values: ["A", "AAAA", "CNAME"]
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type RunDomainMappingSpecStatus struct {
	// Array of observed DomainMappingConditions, indicating the current state
	// of the DomainMapping.
	// +optional
	Conditions []RunDomainMappingSpecStatusConditions `json:"conditions,omitempty" tf:"conditions"`
	// The name of the route that the mapping currently points to.
	// +optional
	MappedRouteName *string `json:"mappedRouteName,omitempty" tf:"mapped_route_name"`
	// ObservedGeneration is the 'Generation' of the DomainMapping that
	// was last processed by the controller.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty" tf:"observed_generation"`
	// The resource records required to configure this domain mapping. These
	// records must be added to the domain's DNS configuration in order to
	// serve the application via this domain mapping.
	// +optional
	ResourceRecords []RunDomainMappingSpecStatusResourceRecords `json:"resourceRecords,omitempty" tf:"resource_records"`
}

type RunDomainMappingSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the cloud run instance. eg us-central1
	Location *string `json:"location" tf:"location"`
	// Metadata associated with this DomainMapping.
	Metadata *RunDomainMappingSpecMetadata `json:"metadata" tf:"metadata"`
	// Name should be a verified domain
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The spec for this DomainMapping.
	Spec *RunDomainMappingSpecSpec `json:"spec" tf:"spec"`
	// The current status of the DomainMapping.
	// +optional
	Status []RunDomainMappingSpecStatus `json:"status,omitempty" tf:"status"`
}

type RunDomainMappingStatus struct {
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

// RunDomainMappingList is a list of RunDomainMappings
type RunDomainMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RunDomainMapping CRD objects
	Items []RunDomainMapping `json:"items,omitempty"`
}

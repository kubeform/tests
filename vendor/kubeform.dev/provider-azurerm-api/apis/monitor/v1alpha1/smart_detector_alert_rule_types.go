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

type SmartDetectorAlertRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SmartDetectorAlertRuleSpec   `json:"spec,omitempty"`
	Status            SmartDetectorAlertRuleStatus `json:"status,omitempty"`
}

type SmartDetectorAlertRuleSpec struct {
	SmartDetectorAlertRuleSpec2 `json:",inline"`
	// +optional
	KubeformOutput SmartDetectorAlertRuleSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type SmartDetectorAlertRuleSpecActionGroup struct {
	// +optional
	EmailSubject *string  `json:"emailSubject,omitempty" tf:"email_subject"`
	Ids          []string `json:"ids" tf:"ids"`
	// +optional
	WebhookPayload *string `json:"webhookPayload,omitempty" tf:"webhook_payload"`
}

type SmartDetectorAlertRuleSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ActionGroup *SmartDetectorAlertRuleSpecActionGroup `json:"actionGroup" tf:"action_group"`
	// +optional
	Description  *string `json:"description,omitempty" tf:"description"`
	DetectorType *string `json:"detectorType" tf:"detector_type"`
	// +optional
	Enabled           *bool    `json:"enabled,omitempty" tf:"enabled"`
	Frequency         *string  `json:"frequency" tf:"frequency"`
	Name              *string  `json:"name" tf:"name"`
	ResourceGroupName *string  `json:"resourceGroupName" tf:"resource_group_name"`
	ScopeResourceIDS  []string `json:"scopeResourceIDS" tf:"scope_resource_ids"`
	Severity          *string  `json:"severity" tf:"severity"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	ThrottlingDuration *string `json:"throttlingDuration,omitempty" tf:"throttling_duration"`
}

type SmartDetectorAlertRuleStatus struct {
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

// SmartDetectorAlertRuleList is a list of SmartDetectorAlertRules
type SmartDetectorAlertRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SmartDetectorAlertRule CRD objects
	Items []SmartDetectorAlertRule `json:"items,omitempty"`
}

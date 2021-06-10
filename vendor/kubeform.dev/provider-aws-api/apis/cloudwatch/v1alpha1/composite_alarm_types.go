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

type CompositeAlarm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CompositeAlarmSpec   `json:"spec,omitempty"`
	Status            CompositeAlarmStatus `json:"status,omitempty"`
}

type CompositeAlarmSpec struct {
	CompositeAlarmSpec2 `json:",inline"`
	// +optional
	KubeformOutput CompositeAlarmSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type CompositeAlarmSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	AlarmActions []string `json:"alarmActions,omitempty" tf:"alarm_actions"`
	// +optional
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description"`
	AlarmName        *string `json:"alarmName" tf:"alarm_name"`
	AlarmRule        *string `json:"alarmRule" tf:"alarm_rule"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	InsufficientDataActions []string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	OkActions []string `json:"okActions,omitempty" tf:"ok_actions"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type CompositeAlarmStatus struct {
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

// CompositeAlarmList is a list of CompositeAlarms
type CompositeAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CompositeAlarm CRD objects
	Items []CompositeAlarm `json:"items,omitempty"`
}

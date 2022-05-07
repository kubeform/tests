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

type AppautoscalingScheduledAction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppautoscalingScheduledActionSpec   `json:"spec,omitempty"`
	Status            AppautoscalingScheduledActionStatus `json:"status,omitempty"`
}

type AppautoscalingScheduledActionSpecScalableTargetAction struct {
	// +optional
	MaxCapacity int64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`
	// +optional
	MinCapacity int64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

type AppautoscalingScheduledActionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	EndTime    string `json:"endTime,omitempty" tf:"end_time,omitempty"`
	Name       string `json:"name" tf:"name"`
	ResourceID string `json:"resourceID" tf:"resource_id"`
	// +optional
	ScalableDimension string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ScalableTargetAction []AppautoscalingScheduledActionSpecScalableTargetAction `json:"scalableTargetAction,omitempty" tf:"scalable_target_action,omitempty"`
	// +optional
	Schedule         string `json:"schedule,omitempty" tf:"schedule,omitempty"`
	ServiceNamespace string `json:"serviceNamespace" tf:"service_namespace"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type AppautoscalingScheduledActionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppautoscalingScheduledActionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppautoscalingScheduledActionList is a list of AppautoscalingScheduledActions
type AppautoscalingScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppautoscalingScheduledAction CRD objects
	Items []AppautoscalingScheduledAction `json:"items,omitempty"`
}
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

type Intent struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntentSpec   `json:"spec,omitempty"`
	Status            IntentStatus `json:"status,omitempty"`
}

type IntentSpec struct {
	IntentSpec2 `json:",inline"`
	// +optional
	KubeformOutput IntentSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type IntentSpecFollowupIntentInfo struct {
	// The unique identifier of the followup intent.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	// +optional
	FollowupIntentName *string `json:"followupIntentName,omitempty" tf:"followup_intent_name"`
	// The unique identifier of the followup intent's parent.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	// +optional
	ParentFollowupIntentName *string `json:"parentFollowupIntentName,omitempty" tf:"parent_followup_intent_name"`
}

type IntentSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform). Possible values: ["FACEBOOK", "SLACK", "TELEGRAM", "KIK", "SKYPE", "LINE", "VIBER", "ACTIONS_ON_GOOGLE", "GOOGLE_HANGOUTS"]
	// +optional
	DefaultResponsePlatforms []string `json:"defaultResponsePlatforms,omitempty" tf:"default_response_platforms"`
	// The name of this intent to be displayed on the console.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	// +optional
	Events []string `json:"events,omitempty" tf:"events"`
	// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field
	// only in the output.
	// +optional
	FollowupIntentInfo []IntentSpecFollowupIntentInfo `json:"followupIntentInfo,omitempty" tf:"followup_intent_info"`
	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	// +optional
	InputContextNames []string `json:"inputContextNames,omitempty" tf:"input_context_names"`
	// Indicates whether this is a fallback intent.
	// +optional
	IsFallback *bool `json:"isFallback,omitempty" tf:"is_fallback"`
	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	// +optional
	MlDisabled *bool `json:"mlDisabled,omitempty" tf:"ml_disabled"`
	// The unique identifier of this intent.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	// +optional
	ParentFollowupIntentName *string `json:"parentFollowupIntentName,omitempty" tf:"parent_followup_intent_name"`
	// The priority of this intent. Higher numbers represent higher priorities.
	//   - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	//   to the Normal priority in the console.
	//   - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	// +optional
	ResetContexts *bool `json:"resetContexts,omitempty" tf:"reset_contexts"`
	// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup
	// intents chain for this intent.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	// +optional
	RootFollowupIntentName *string `json:"rootFollowupIntentName,omitempty" tf:"root_followup_intent_name"`
	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	// filling prompt is forwarded to the webhook. Possible values: ["WEBHOOK_STATE_ENABLED", "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING"]
	// +optional
	WebhookState *string `json:"webhookState,omitempty" tf:"webhook_state"`
}

type IntentStatus struct {
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

// IntentList is a list of Intents
type IntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Intent CRD objects
	Items []Intent `json:"items,omitempty"`
}
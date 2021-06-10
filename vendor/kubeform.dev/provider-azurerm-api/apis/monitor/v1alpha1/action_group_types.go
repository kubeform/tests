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

type ActionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionGroupSpec   `json:"spec,omitempty"`
	Status            ActionGroupStatus `json:"status,omitempty"`
}

type ActionGroupSpec struct {
	ActionGroupSpec2 `json:",inline"`
	// +optional
	KubeformOutput ActionGroupSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ActionGroupSpecArmRoleReceiver struct {
	Name   *string `json:"name" tf:"name"`
	RoleID *string `json:"roleID" tf:"role_id"`
	// +optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema"`
}

type ActionGroupSpecAutomationRunbookReceiver struct {
	AutomationAccountID *string `json:"automationAccountID" tf:"automation_account_id"`
	IsGlobalRunbook     *bool   `json:"isGlobalRunbook" tf:"is_global_runbook"`
	Name                *string `json:"name" tf:"name"`
	RunbookName         *string `json:"runbookName" tf:"runbook_name"`
	ServiceURI          *string `json:"serviceURI" tf:"service_uri"`
	// +optional
	UseCommonAlertSchema *bool   `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema"`
	WebhookResourceID    *string `json:"webhookResourceID" tf:"webhook_resource_id"`
}

type ActionGroupSpecAzureAppPushReceiver struct {
	EmailAddress *string `json:"emailAddress" tf:"email_address"`
	Name         *string `json:"name" tf:"name"`
}

type ActionGroupSpecAzureFunctionReceiver struct {
	FunctionAppResourceID *string `json:"functionAppResourceID" tf:"function_app_resource_id"`
	FunctionName          *string `json:"functionName" tf:"function_name"`
	HttpTriggerURL        *string `json:"httpTriggerURL" tf:"http_trigger_url"`
	Name                  *string `json:"name" tf:"name"`
	// +optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema"`
}

type ActionGroupSpecEmailReceiver struct {
	EmailAddress *string `json:"emailAddress" tf:"email_address"`
	Name         *string `json:"name" tf:"name"`
	// +optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema"`
}

type ActionGroupSpecItsmReceiver struct {
	ConnectionID        *string `json:"connectionID" tf:"connection_id"`
	Name                *string `json:"name" tf:"name"`
	Region              *string `json:"region" tf:"region"`
	TicketConfiguration *string `json:"ticketConfiguration" tf:"ticket_configuration"`
	WorkspaceID         *string `json:"workspaceID" tf:"workspace_id"`
}

type ActionGroupSpecLogicAppReceiver struct {
	CallbackURL *string `json:"callbackURL" tf:"callback_url"`
	Name        *string `json:"name" tf:"name"`
	ResourceID  *string `json:"resourceID" tf:"resource_id"`
	// +optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema"`
}

type ActionGroupSpecSmsReceiver struct {
	CountryCode *string `json:"countryCode" tf:"country_code"`
	Name        *string `json:"name" tf:"name"`
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number"`
}

type ActionGroupSpecVoiceReceiver struct {
	CountryCode *string `json:"countryCode" tf:"country_code"`
	Name        *string `json:"name" tf:"name"`
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number"`
}

type ActionGroupSpecWebhookReceiverAadAuth struct {
	// +optional
	IdentifierURI *string `json:"identifierURI,omitempty" tf:"identifier_uri"`
	ObjectID      *string `json:"objectID" tf:"object_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
}

type ActionGroupSpecWebhookReceiver struct {
	// +optional
	AadAuth    *ActionGroupSpecWebhookReceiverAadAuth `json:"aadAuth,omitempty" tf:"aad_auth"`
	Name       *string                                `json:"name" tf:"name"`
	ServiceURI *string                                `json:"serviceURI" tf:"service_uri"`
	// +optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema"`
}

type ActionGroupSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ArmRoleReceiver []ActionGroupSpecArmRoleReceiver `json:"armRoleReceiver,omitempty" tf:"arm_role_receiver"`
	// +optional
	AutomationRunbookReceiver []ActionGroupSpecAutomationRunbookReceiver `json:"automationRunbookReceiver,omitempty" tf:"automation_runbook_receiver"`
	// +optional
	AzureAppPushReceiver []ActionGroupSpecAzureAppPushReceiver `json:"azureAppPushReceiver,omitempty" tf:"azure_app_push_receiver"`
	// +optional
	AzureFunctionReceiver []ActionGroupSpecAzureFunctionReceiver `json:"azureFunctionReceiver,omitempty" tf:"azure_function_receiver"`
	// +optional
	EmailReceiver []ActionGroupSpecEmailReceiver `json:"emailReceiver,omitempty" tf:"email_receiver"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	ItsmReceiver []ActionGroupSpecItsmReceiver `json:"itsmReceiver,omitempty" tf:"itsm_receiver"`
	// +optional
	LogicAppReceiver  []ActionGroupSpecLogicAppReceiver `json:"logicAppReceiver,omitempty" tf:"logic_app_receiver"`
	Name              *string                           `json:"name" tf:"name"`
	ResourceGroupName *string                           `json:"resourceGroupName" tf:"resource_group_name"`
	ShortName         *string                           `json:"shortName" tf:"short_name"`
	// +optional
	SmsReceiver []ActionGroupSpecSmsReceiver `json:"smsReceiver,omitempty" tf:"sms_receiver"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VoiceReceiver []ActionGroupSpecVoiceReceiver `json:"voiceReceiver,omitempty" tf:"voice_receiver"`
	// +optional
	WebhookReceiver []ActionGroupSpecWebhookReceiver `json:"webhookReceiver,omitempty" tf:"webhook_receiver"`
}

type ActionGroupStatus struct {
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

// ActionGroupList is a list of ActionGroups
type ActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ActionGroup CRD objects
	Items []ActionGroup `json:"items,omitempty"`
}
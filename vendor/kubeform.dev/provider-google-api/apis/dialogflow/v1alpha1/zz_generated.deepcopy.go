// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Agent.
func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Agent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentList) DeepCopyInto(out *AgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Agent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentList.
func (in *AgentList) DeepCopy() *AgentList {
	if in == nil {
		return nil
	}
	out := new(AgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	in.AgentSpec2.DeepCopyInto(&out.AgentSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec.
func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec2) DeepCopyInto(out *AgentSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.ApiVersion != nil {
		in, out := &in.ApiVersion, &out.ApiVersion
		*out = new(string)
		**out = **in
	}
	if in.AvatarURI != nil {
		in, out := &in.AvatarURI, &out.AvatarURI
		*out = new(string)
		**out = **in
	}
	if in.AvatarURIBackend != nil {
		in, out := &in.AvatarURIBackend, &out.AvatarURIBackend
		*out = new(string)
		**out = **in
	}
	if in.ClassificationThreshold != nil {
		in, out := &in.ClassificationThreshold, &out.ClassificationThreshold
		*out = new(float64)
		**out = **in
	}
	if in.DefaultLanguageCode != nil {
		in, out := &in.DefaultLanguageCode, &out.DefaultLanguageCode
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EnableLogging != nil {
		in, out := &in.EnableLogging, &out.EnableLogging
		*out = new(bool)
		**out = **in
	}
	if in.MatchMode != nil {
		in, out := &in.MatchMode, &out.MatchMode
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.SupportedLanguageCodes != nil {
		in, out := &in.SupportedLanguageCodes, &out.SupportedLanguageCodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec2.
func (in *AgentSpec2) DeepCopy() *AgentSpec2 {
	if in == nil {
		return nil
	}
	out := new(AgentSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentStatus) DeepCopyInto(out *AgentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentStatus.
func (in *AgentStatus) DeepCopy() *AgentStatus {
	if in == nil {
		return nil
	}
	out := new(AgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityType) DeepCopyInto(out *EntityType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityType.
func (in *EntityType) DeepCopy() *EntityType {
	if in == nil {
		return nil
	}
	out := new(EntityType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntityType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityTypeList) DeepCopyInto(out *EntityTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EntityType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityTypeList.
func (in *EntityTypeList) DeepCopy() *EntityTypeList {
	if in == nil {
		return nil
	}
	out := new(EntityTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntityTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityTypeSpec) DeepCopyInto(out *EntityTypeSpec) {
	*out = *in
	in.EntityTypeSpec2.DeepCopyInto(&out.EntityTypeSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityTypeSpec.
func (in *EntityTypeSpec) DeepCopy() *EntityTypeSpec {
	if in == nil {
		return nil
	}
	out := new(EntityTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityTypeSpec2) DeepCopyInto(out *EntityTypeSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EnableFuzzyExtraction != nil {
		in, out := &in.EnableFuzzyExtraction, &out.EnableFuzzyExtraction
		*out = new(bool)
		**out = **in
	}
	if in.Entities != nil {
		in, out := &in.Entities, &out.Entities
		*out = make([]EntityTypeSpecEntities, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityTypeSpec2.
func (in *EntityTypeSpec2) DeepCopy() *EntityTypeSpec2 {
	if in == nil {
		return nil
	}
	out := new(EntityTypeSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityTypeSpecEntities) DeepCopyInto(out *EntityTypeSpecEntities) {
	*out = *in
	if in.Synonyms != nil {
		in, out := &in.Synonyms, &out.Synonyms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityTypeSpecEntities.
func (in *EntityTypeSpecEntities) DeepCopy() *EntityTypeSpecEntities {
	if in == nil {
		return nil
	}
	out := new(EntityTypeSpecEntities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityTypeStatus) DeepCopyInto(out *EntityTypeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityTypeStatus.
func (in *EntityTypeStatus) DeepCopy() *EntityTypeStatus {
	if in == nil {
		return nil
	}
	out := new(EntityTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Intent) DeepCopyInto(out *Intent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Intent.
func (in *Intent) DeepCopy() *Intent {
	if in == nil {
		return nil
	}
	out := new(Intent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Intent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntentList) DeepCopyInto(out *IntentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Intent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntentList.
func (in *IntentList) DeepCopy() *IntentList {
	if in == nil {
		return nil
	}
	out := new(IntentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IntentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntentSpec) DeepCopyInto(out *IntentSpec) {
	*out = *in
	in.IntentSpec2.DeepCopyInto(&out.IntentSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntentSpec.
func (in *IntentSpec) DeepCopy() *IntentSpec {
	if in == nil {
		return nil
	}
	out := new(IntentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntentSpec2) DeepCopyInto(out *IntentSpec2) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	out.ProviderRef = in.ProviderRef
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.DefaultResponsePlatforms != nil {
		in, out := &in.DefaultResponsePlatforms, &out.DefaultResponsePlatforms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FollowupIntentInfo != nil {
		in, out := &in.FollowupIntentInfo, &out.FollowupIntentInfo
		*out = make([]IntentSpecFollowupIntentInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InputContextNames != nil {
		in, out := &in.InputContextNames, &out.InputContextNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IsFallback != nil {
		in, out := &in.IsFallback, &out.IsFallback
		*out = new(bool)
		**out = **in
	}
	if in.MlDisabled != nil {
		in, out := &in.MlDisabled, &out.MlDisabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ParentFollowupIntentName != nil {
		in, out := &in.ParentFollowupIntentName, &out.ParentFollowupIntentName
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ResetContexts != nil {
		in, out := &in.ResetContexts, &out.ResetContexts
		*out = new(bool)
		**out = **in
	}
	if in.RootFollowupIntentName != nil {
		in, out := &in.RootFollowupIntentName, &out.RootFollowupIntentName
		*out = new(string)
		**out = **in
	}
	if in.WebhookState != nil {
		in, out := &in.WebhookState, &out.WebhookState
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntentSpec2.
func (in *IntentSpec2) DeepCopy() *IntentSpec2 {
	if in == nil {
		return nil
	}
	out := new(IntentSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntentSpecFollowupIntentInfo) DeepCopyInto(out *IntentSpecFollowupIntentInfo) {
	*out = *in
	if in.FollowupIntentName != nil {
		in, out := &in.FollowupIntentName, &out.FollowupIntentName
		*out = new(string)
		**out = **in
	}
	if in.ParentFollowupIntentName != nil {
		in, out := &in.ParentFollowupIntentName, &out.ParentFollowupIntentName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntentSpecFollowupIntentInfo.
func (in *IntentSpecFollowupIntentInfo) DeepCopy() *IntentSpecFollowupIntentInfo {
	if in == nil {
		return nil
	}
	out := new(IntentSpecFollowupIntentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntentStatus) DeepCopyInto(out *IntentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntentStatus.
func (in *IntentStatus) DeepCopy() *IntentStatus {
	if in == nil {
		return nil
	}
	out := new(IntentStatus)
	in.DeepCopyInto(out)
	return out
}
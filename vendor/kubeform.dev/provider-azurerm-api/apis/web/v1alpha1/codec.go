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
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecManagedRules{}).Type1()):   ApplicationFirewallPolicySpecManagedRulesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecPolicySettings{}).Type1()): ApplicationFirewallPolicySpecPolicySettingsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecManagedRules{}).Type1()):   ApplicationFirewallPolicySpecManagedRulesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecPolicySettings{}).Type1()): ApplicationFirewallPolicySpecPolicySettingsCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type ApplicationFirewallPolicySpecManagedRulesCodec struct {
}

func (ApplicationFirewallPolicySpecManagedRulesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ApplicationFirewallPolicySpecManagedRules)(ptr) == nil
}

func (ApplicationFirewallPolicySpecManagedRulesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ApplicationFirewallPolicySpecManagedRules)(ptr)
	var objs []ApplicationFirewallPolicySpecManagedRules
	if obj != nil {
		objs = []ApplicationFirewallPolicySpecManagedRules{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecManagedRules{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ApplicationFirewallPolicySpecManagedRulesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ApplicationFirewallPolicySpecManagedRules)(ptr) = ApplicationFirewallPolicySpecManagedRules{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ApplicationFirewallPolicySpecManagedRules

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecManagedRules{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ApplicationFirewallPolicySpecManagedRules)(ptr) = objs[0]
			} else {
				*(*ApplicationFirewallPolicySpecManagedRules)(ptr) = ApplicationFirewallPolicySpecManagedRules{}
			}
		} else {
			*(*ApplicationFirewallPolicySpecManagedRules)(ptr) = ApplicationFirewallPolicySpecManagedRules{}
		}
	default:
		iter.ReportError("decode ApplicationFirewallPolicySpecManagedRules", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ApplicationFirewallPolicySpecPolicySettingsCodec struct {
}

func (ApplicationFirewallPolicySpecPolicySettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ApplicationFirewallPolicySpecPolicySettings)(ptr) == nil
}

func (ApplicationFirewallPolicySpecPolicySettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ApplicationFirewallPolicySpecPolicySettings)(ptr)
	var objs []ApplicationFirewallPolicySpecPolicySettings
	if obj != nil {
		objs = []ApplicationFirewallPolicySpecPolicySettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecPolicySettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ApplicationFirewallPolicySpecPolicySettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ApplicationFirewallPolicySpecPolicySettings)(ptr) = ApplicationFirewallPolicySpecPolicySettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ApplicationFirewallPolicySpecPolicySettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ApplicationFirewallPolicySpecPolicySettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ApplicationFirewallPolicySpecPolicySettings)(ptr) = objs[0]
			} else {
				*(*ApplicationFirewallPolicySpecPolicySettings)(ptr) = ApplicationFirewallPolicySpecPolicySettings{}
			}
		} else {
			*(*ApplicationFirewallPolicySpecPolicySettings)(ptr) = ApplicationFirewallPolicySpecPolicySettings{}
		}
	default:
		iter.ReportError("decode ApplicationFirewallPolicySpecPolicySettings", "unexpected JSON type")
	}
}
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
		jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamBindingSpecCondition{}).Type1()): ConfigIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamMemberSpecCondition{}).Type1()):  ConfigIamMemberSpecConditionCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamBindingSpecCondition{}).Type1()): ConfigIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamMemberSpecCondition{}).Type1()):  ConfigIamMemberSpecConditionCodec{},
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
type ConfigIamBindingSpecConditionCodec struct {
}

func (ConfigIamBindingSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConfigIamBindingSpecCondition)(ptr) == nil
}

func (ConfigIamBindingSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConfigIamBindingSpecCondition)(ptr)
	var objs []ConfigIamBindingSpecCondition
	if obj != nil {
		objs = []ConfigIamBindingSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamBindingSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConfigIamBindingSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConfigIamBindingSpecCondition)(ptr) = ConfigIamBindingSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConfigIamBindingSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamBindingSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConfigIamBindingSpecCondition)(ptr) = objs[0]
			} else {
				*(*ConfigIamBindingSpecCondition)(ptr) = ConfigIamBindingSpecCondition{}
			}
		} else {
			*(*ConfigIamBindingSpecCondition)(ptr) = ConfigIamBindingSpecCondition{}
		}
	default:
		iter.ReportError("decode ConfigIamBindingSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ConfigIamMemberSpecConditionCodec struct {
}

func (ConfigIamMemberSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConfigIamMemberSpecCondition)(ptr) == nil
}

func (ConfigIamMemberSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConfigIamMemberSpecCondition)(ptr)
	var objs []ConfigIamMemberSpecCondition
	if obj != nil {
		objs = []ConfigIamMemberSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamMemberSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConfigIamMemberSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConfigIamMemberSpecCondition)(ptr) = ConfigIamMemberSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConfigIamMemberSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConfigIamMemberSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConfigIamMemberSpecCondition)(ptr) = objs[0]
			} else {
				*(*ConfigIamMemberSpecCondition)(ptr) = ConfigIamMemberSpecCondition{}
			}
		} else {
			*(*ConfigIamMemberSpecCondition)(ptr) = ConfigIamMemberSpecCondition{}
		}
	default:
		iter.ReportError("decode ConfigIamMemberSpecCondition", "unexpected JSON type")
	}
}
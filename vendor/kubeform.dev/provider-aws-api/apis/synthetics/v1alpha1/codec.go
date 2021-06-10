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
		jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecRunConfig{}).Type1()): CanarySpecRunConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecSchedule{}).Type1()):  CanarySpecScheduleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecVpcConfig{}).Type1()): CanarySpecVpcConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecRunConfig{}).Type1()): CanarySpecRunConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecSchedule{}).Type1()):  CanarySpecScheduleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecVpcConfig{}).Type1()): CanarySpecVpcConfigCodec{},
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
type CanarySpecRunConfigCodec struct {
}

func (CanarySpecRunConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CanarySpecRunConfig)(ptr) == nil
}

func (CanarySpecRunConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CanarySpecRunConfig)(ptr)
	var objs []CanarySpecRunConfig
	if obj != nil {
		objs = []CanarySpecRunConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecRunConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CanarySpecRunConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CanarySpecRunConfig)(ptr) = CanarySpecRunConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CanarySpecRunConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecRunConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CanarySpecRunConfig)(ptr) = objs[0]
			} else {
				*(*CanarySpecRunConfig)(ptr) = CanarySpecRunConfig{}
			}
		} else {
			*(*CanarySpecRunConfig)(ptr) = CanarySpecRunConfig{}
		}
	default:
		iter.ReportError("decode CanarySpecRunConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CanarySpecScheduleCodec struct {
}

func (CanarySpecScheduleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CanarySpecSchedule)(ptr) == nil
}

func (CanarySpecScheduleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CanarySpecSchedule)(ptr)
	var objs []CanarySpecSchedule
	if obj != nil {
		objs = []CanarySpecSchedule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecSchedule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CanarySpecScheduleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CanarySpecSchedule)(ptr) = CanarySpecSchedule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CanarySpecSchedule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecSchedule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CanarySpecSchedule)(ptr) = objs[0]
			} else {
				*(*CanarySpecSchedule)(ptr) = CanarySpecSchedule{}
			}
		} else {
			*(*CanarySpecSchedule)(ptr) = CanarySpecSchedule{}
		}
	default:
		iter.ReportError("decode CanarySpecSchedule", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CanarySpecVpcConfigCodec struct {
}

func (CanarySpecVpcConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CanarySpecVpcConfig)(ptr) == nil
}

func (CanarySpecVpcConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CanarySpecVpcConfig)(ptr)
	var objs []CanarySpecVpcConfig
	if obj != nil {
		objs = []CanarySpecVpcConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecVpcConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CanarySpecVpcConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CanarySpecVpcConfig)(ptr) = CanarySpecVpcConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CanarySpecVpcConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CanarySpecVpcConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CanarySpecVpcConfig)(ptr) = objs[0]
			} else {
				*(*CanarySpecVpcConfig)(ptr) = CanarySpecVpcConfig{}
			}
		} else {
			*(*CanarySpecVpcConfig)(ptr) = CanarySpecVpcConfig{}
		}
	default:
		iter.ReportError("decode CanarySpecVpcConfig", "unexpected JSON type")
	}
}

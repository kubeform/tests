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
		jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeySpecVersionTemplate{}).Type1()):     CryptoKeySpecVersionTemplateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamBindingSpecCondition{}).Type1()): CryptoKeyIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamMemberSpecCondition{}).Type1()):  CryptoKeyIamMemberSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamBindingSpecCondition{}).Type1()):   KeyRingIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamMemberSpecCondition{}).Type1()):    KeyRingIamMemberSpecConditionCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeySpecVersionTemplate{}).Type1()):     CryptoKeySpecVersionTemplateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamBindingSpecCondition{}).Type1()): CryptoKeyIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamMemberSpecCondition{}).Type1()):  CryptoKeyIamMemberSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamBindingSpecCondition{}).Type1()):   KeyRingIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamMemberSpecCondition{}).Type1()):    KeyRingIamMemberSpecConditionCodec{},
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
type CryptoKeySpecVersionTemplateCodec struct {
}

func (CryptoKeySpecVersionTemplateCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CryptoKeySpecVersionTemplate)(ptr) == nil
}

func (CryptoKeySpecVersionTemplateCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CryptoKeySpecVersionTemplate)(ptr)
	var objs []CryptoKeySpecVersionTemplate
	if obj != nil {
		objs = []CryptoKeySpecVersionTemplate{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeySpecVersionTemplate{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CryptoKeySpecVersionTemplateCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CryptoKeySpecVersionTemplate)(ptr) = CryptoKeySpecVersionTemplate{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CryptoKeySpecVersionTemplate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeySpecVersionTemplate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CryptoKeySpecVersionTemplate)(ptr) = objs[0]
			} else {
				*(*CryptoKeySpecVersionTemplate)(ptr) = CryptoKeySpecVersionTemplate{}
			}
		} else {
			*(*CryptoKeySpecVersionTemplate)(ptr) = CryptoKeySpecVersionTemplate{}
		}
	default:
		iter.ReportError("decode CryptoKeySpecVersionTemplate", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CryptoKeyIamBindingSpecConditionCodec struct {
}

func (CryptoKeyIamBindingSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CryptoKeyIamBindingSpecCondition)(ptr) == nil
}

func (CryptoKeyIamBindingSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CryptoKeyIamBindingSpecCondition)(ptr)
	var objs []CryptoKeyIamBindingSpecCondition
	if obj != nil {
		objs = []CryptoKeyIamBindingSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamBindingSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CryptoKeyIamBindingSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CryptoKeyIamBindingSpecCondition)(ptr) = CryptoKeyIamBindingSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CryptoKeyIamBindingSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamBindingSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CryptoKeyIamBindingSpecCondition)(ptr) = objs[0]
			} else {
				*(*CryptoKeyIamBindingSpecCondition)(ptr) = CryptoKeyIamBindingSpecCondition{}
			}
		} else {
			*(*CryptoKeyIamBindingSpecCondition)(ptr) = CryptoKeyIamBindingSpecCondition{}
		}
	default:
		iter.ReportError("decode CryptoKeyIamBindingSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CryptoKeyIamMemberSpecConditionCodec struct {
}

func (CryptoKeyIamMemberSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CryptoKeyIamMemberSpecCondition)(ptr) == nil
}

func (CryptoKeyIamMemberSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CryptoKeyIamMemberSpecCondition)(ptr)
	var objs []CryptoKeyIamMemberSpecCondition
	if obj != nil {
		objs = []CryptoKeyIamMemberSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamMemberSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CryptoKeyIamMemberSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CryptoKeyIamMemberSpecCondition)(ptr) = CryptoKeyIamMemberSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CryptoKeyIamMemberSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CryptoKeyIamMemberSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CryptoKeyIamMemberSpecCondition)(ptr) = objs[0]
			} else {
				*(*CryptoKeyIamMemberSpecCondition)(ptr) = CryptoKeyIamMemberSpecCondition{}
			}
		} else {
			*(*CryptoKeyIamMemberSpecCondition)(ptr) = CryptoKeyIamMemberSpecCondition{}
		}
	default:
		iter.ReportError("decode CryptoKeyIamMemberSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type KeyRingIamBindingSpecConditionCodec struct {
}

func (KeyRingIamBindingSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*KeyRingIamBindingSpecCondition)(ptr) == nil
}

func (KeyRingIamBindingSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*KeyRingIamBindingSpecCondition)(ptr)
	var objs []KeyRingIamBindingSpecCondition
	if obj != nil {
		objs = []KeyRingIamBindingSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamBindingSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (KeyRingIamBindingSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*KeyRingIamBindingSpecCondition)(ptr) = KeyRingIamBindingSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []KeyRingIamBindingSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamBindingSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*KeyRingIamBindingSpecCondition)(ptr) = objs[0]
			} else {
				*(*KeyRingIamBindingSpecCondition)(ptr) = KeyRingIamBindingSpecCondition{}
			}
		} else {
			*(*KeyRingIamBindingSpecCondition)(ptr) = KeyRingIamBindingSpecCondition{}
		}
	default:
		iter.ReportError("decode KeyRingIamBindingSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type KeyRingIamMemberSpecConditionCodec struct {
}

func (KeyRingIamMemberSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*KeyRingIamMemberSpecCondition)(ptr) == nil
}

func (KeyRingIamMemberSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*KeyRingIamMemberSpecCondition)(ptr)
	var objs []KeyRingIamMemberSpecCondition
	if obj != nil {
		objs = []KeyRingIamMemberSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamMemberSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (KeyRingIamMemberSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*KeyRingIamMemberSpecCondition)(ptr) = KeyRingIamMemberSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []KeyRingIamMemberSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KeyRingIamMemberSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*KeyRingIamMemberSpecCondition)(ptr) = objs[0]
			} else {
				*(*KeyRingIamMemberSpecCondition)(ptr) = KeyRingIamMemberSpecCondition{}
			}
		} else {
			*(*KeyRingIamMemberSpecCondition)(ptr) = KeyRingIamMemberSpecCondition{}
		}
	default:
		iter.ReportError("decode KeyRingIamMemberSpecCondition", "unexpected JSON type")
	}
}

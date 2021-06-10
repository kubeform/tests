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
		jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFallbackRoute{}).Type1()): IothubSpecFallbackRouteCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFileUpload{}).Type1()):    IothubSpecFileUploadCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecSku{}).Type1()):           IothubSpecSkuCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DpsSpecSku{}).Type1()):              DpsSpecSkuCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFallbackRoute{}).Type1()): IothubSpecFallbackRouteCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFileUpload{}).Type1()):    IothubSpecFileUploadCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecSku{}).Type1()):           IothubSpecSkuCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DpsSpecSku{}).Type1()):              DpsSpecSkuCodec{},
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
type IothubSpecFallbackRouteCodec struct {
}

func (IothubSpecFallbackRouteCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*IothubSpecFallbackRoute)(ptr) == nil
}

func (IothubSpecFallbackRouteCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*IothubSpecFallbackRoute)(ptr)
	var objs []IothubSpecFallbackRoute
	if obj != nil {
		objs = []IothubSpecFallbackRoute{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFallbackRoute{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (IothubSpecFallbackRouteCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*IothubSpecFallbackRoute)(ptr) = IothubSpecFallbackRoute{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []IothubSpecFallbackRoute

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFallbackRoute{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*IothubSpecFallbackRoute)(ptr) = objs[0]
			} else {
				*(*IothubSpecFallbackRoute)(ptr) = IothubSpecFallbackRoute{}
			}
		} else {
			*(*IothubSpecFallbackRoute)(ptr) = IothubSpecFallbackRoute{}
		}
	default:
		iter.ReportError("decode IothubSpecFallbackRoute", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type IothubSpecFileUploadCodec struct {
}

func (IothubSpecFileUploadCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*IothubSpecFileUpload)(ptr) == nil
}

func (IothubSpecFileUploadCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*IothubSpecFileUpload)(ptr)
	var objs []IothubSpecFileUpload
	if obj != nil {
		objs = []IothubSpecFileUpload{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFileUpload{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (IothubSpecFileUploadCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*IothubSpecFileUpload)(ptr) = IothubSpecFileUpload{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []IothubSpecFileUpload

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecFileUpload{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*IothubSpecFileUpload)(ptr) = objs[0]
			} else {
				*(*IothubSpecFileUpload)(ptr) = IothubSpecFileUpload{}
			}
		} else {
			*(*IothubSpecFileUpload)(ptr) = IothubSpecFileUpload{}
		}
	default:
		iter.ReportError("decode IothubSpecFileUpload", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type IothubSpecSkuCodec struct {
}

func (IothubSpecSkuCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*IothubSpecSku)(ptr) == nil
}

func (IothubSpecSkuCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*IothubSpecSku)(ptr)
	var objs []IothubSpecSku
	if obj != nil {
		objs = []IothubSpecSku{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecSku{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (IothubSpecSkuCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*IothubSpecSku)(ptr) = IothubSpecSku{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []IothubSpecSku

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IothubSpecSku{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*IothubSpecSku)(ptr) = objs[0]
			} else {
				*(*IothubSpecSku)(ptr) = IothubSpecSku{}
			}
		} else {
			*(*IothubSpecSku)(ptr) = IothubSpecSku{}
		}
	default:
		iter.ReportError("decode IothubSpecSku", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type DpsSpecSkuCodec struct {
}

func (DpsSpecSkuCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DpsSpecSku)(ptr) == nil
}

func (DpsSpecSkuCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DpsSpecSku)(ptr)
	var objs []DpsSpecSku
	if obj != nil {
		objs = []DpsSpecSku{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DpsSpecSku{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DpsSpecSkuCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DpsSpecSku)(ptr) = DpsSpecSku{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DpsSpecSku

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DpsSpecSku{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DpsSpecSku)(ptr) = objs[0]
			} else {
				*(*DpsSpecSku)(ptr) = DpsSpecSku{}
			}
		} else {
			*(*DpsSpecSku)(ptr) = DpsSpecSku{}
		}
	default:
		iter.ReportError("decode DpsSpecSku", "unexpected JSON type")
	}
}

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
		jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecIdentity{}).Type1()):              ServerSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecStorageProfile{}).Type1()):        ServerSpecStorageProfileCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecThreatDetectionPolicy{}).Type1()): ServerSpecThreatDetectionPolicyCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecIdentity{}).Type1()):              ServerSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecStorageProfile{}).Type1()):        ServerSpecStorageProfileCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecThreatDetectionPolicy{}).Type1()): ServerSpecThreatDetectionPolicyCodec{},
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
type ServerSpecIdentityCodec struct {
}

func (ServerSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServerSpecIdentity)(ptr) == nil
}

func (ServerSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServerSpecIdentity)(ptr)
	var objs []ServerSpecIdentity
	if obj != nil {
		objs = []ServerSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServerSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServerSpecIdentity)(ptr) = ServerSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServerSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServerSpecIdentity)(ptr) = objs[0]
			} else {
				*(*ServerSpecIdentity)(ptr) = ServerSpecIdentity{}
			}
		} else {
			*(*ServerSpecIdentity)(ptr) = ServerSpecIdentity{}
		}
	default:
		iter.ReportError("decode ServerSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServerSpecStorageProfileCodec struct {
}

func (ServerSpecStorageProfileCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServerSpecStorageProfile)(ptr) == nil
}

func (ServerSpecStorageProfileCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServerSpecStorageProfile)(ptr)
	var objs []ServerSpecStorageProfile
	if obj != nil {
		objs = []ServerSpecStorageProfile{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecStorageProfile{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServerSpecStorageProfileCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServerSpecStorageProfile)(ptr) = ServerSpecStorageProfile{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServerSpecStorageProfile

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecStorageProfile{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServerSpecStorageProfile)(ptr) = objs[0]
			} else {
				*(*ServerSpecStorageProfile)(ptr) = ServerSpecStorageProfile{}
			}
		} else {
			*(*ServerSpecStorageProfile)(ptr) = ServerSpecStorageProfile{}
		}
	default:
		iter.ReportError("decode ServerSpecStorageProfile", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServerSpecThreatDetectionPolicyCodec struct {
}

func (ServerSpecThreatDetectionPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServerSpecThreatDetectionPolicy)(ptr) == nil
}

func (ServerSpecThreatDetectionPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServerSpecThreatDetectionPolicy)(ptr)
	var objs []ServerSpecThreatDetectionPolicy
	if obj != nil {
		objs = []ServerSpecThreatDetectionPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecThreatDetectionPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServerSpecThreatDetectionPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServerSpecThreatDetectionPolicy)(ptr) = ServerSpecThreatDetectionPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServerSpecThreatDetectionPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerSpecThreatDetectionPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServerSpecThreatDetectionPolicy)(ptr) = objs[0]
			} else {
				*(*ServerSpecThreatDetectionPolicy)(ptr) = ServerSpecThreatDetectionPolicy{}
			}
		} else {
			*(*ServerSpecThreatDetectionPolicy)(ptr) = ServerSpecThreatDetectionPolicy{}
		}
	default:
		iter.ReportError("decode ServerSpecThreatDetectionPolicy", "unexpected JSON type")
	}
}

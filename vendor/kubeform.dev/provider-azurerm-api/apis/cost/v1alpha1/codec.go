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
		jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecDeliveryInfo{}).Type1()): ManagementExportResourceGroupSpecDeliveryInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecQuery{}).Type1()):        ManagementExportResourceGroupSpecQueryCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecDeliveryInfo{}).Type1()): ManagementExportResourceGroupSpecDeliveryInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecQuery{}).Type1()):        ManagementExportResourceGroupSpecQueryCodec{},
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
type ManagementExportResourceGroupSpecDeliveryInfoCodec struct {
}

func (ManagementExportResourceGroupSpecDeliveryInfoCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagementExportResourceGroupSpecDeliveryInfo)(ptr) == nil
}

func (ManagementExportResourceGroupSpecDeliveryInfoCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagementExportResourceGroupSpecDeliveryInfo)(ptr)
	var objs []ManagementExportResourceGroupSpecDeliveryInfo
	if obj != nil {
		objs = []ManagementExportResourceGroupSpecDeliveryInfo{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecDeliveryInfo{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagementExportResourceGroupSpecDeliveryInfoCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagementExportResourceGroupSpecDeliveryInfo)(ptr) = ManagementExportResourceGroupSpecDeliveryInfo{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagementExportResourceGroupSpecDeliveryInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecDeliveryInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagementExportResourceGroupSpecDeliveryInfo)(ptr) = objs[0]
			} else {
				*(*ManagementExportResourceGroupSpecDeliveryInfo)(ptr) = ManagementExportResourceGroupSpecDeliveryInfo{}
			}
		} else {
			*(*ManagementExportResourceGroupSpecDeliveryInfo)(ptr) = ManagementExportResourceGroupSpecDeliveryInfo{}
		}
	default:
		iter.ReportError("decode ManagementExportResourceGroupSpecDeliveryInfo", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ManagementExportResourceGroupSpecQueryCodec struct {
}

func (ManagementExportResourceGroupSpecQueryCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagementExportResourceGroupSpecQuery)(ptr) == nil
}

func (ManagementExportResourceGroupSpecQueryCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagementExportResourceGroupSpecQuery)(ptr)
	var objs []ManagementExportResourceGroupSpecQuery
	if obj != nil {
		objs = []ManagementExportResourceGroupSpecQuery{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecQuery{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagementExportResourceGroupSpecQueryCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagementExportResourceGroupSpecQuery)(ptr) = ManagementExportResourceGroupSpecQuery{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagementExportResourceGroupSpecQuery

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagementExportResourceGroupSpecQuery{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagementExportResourceGroupSpecQuery)(ptr) = objs[0]
			} else {
				*(*ManagementExportResourceGroupSpecQuery)(ptr) = ManagementExportResourceGroupSpecQuery{}
			}
		} else {
			*(*ManagementExportResourceGroupSpecQuery)(ptr) = ManagementExportResourceGroupSpecQuery{}
		}
	default:
		iter.ReportError("decode ManagementExportResourceGroupSpecQuery", "unexpected JSON type")
	}
}
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
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFilters{}).Type1()):                                            InsightSpecFiltersCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersCreatedAtDateRange{}).Type1()):                          InsightSpecFiltersCreatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersFirstObservedAtDateRange{}).Type1()):                    InsightSpecFiltersFirstObservedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersLastObservedAtDateRange{}).Type1()):                     InsightSpecFiltersLastObservedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersNoteUpdatedAtDateRange{}).Type1()):                      InsightSpecFiltersNoteUpdatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessLaunchedAtDateRange{}).Type1()):                  InsightSpecFiltersProcessLaunchedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessTerminatedAtDateRange{}).Type1()):                InsightSpecFiltersProcessTerminatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{}).Type1()):   InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{}).Type1()):   InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceContainerLaunchedAtDateRange{}).Type1()):        InsightSpecFiltersResourceContainerLaunchedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{}).Type1()): InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersUpdatedAtDateRange{}).Type1()):                          InsightSpecFiltersUpdatedAtDateRangeCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFilters{}).Type1()):                                            InsightSpecFiltersCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersCreatedAtDateRange{}).Type1()):                          InsightSpecFiltersCreatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersFirstObservedAtDateRange{}).Type1()):                    InsightSpecFiltersFirstObservedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersLastObservedAtDateRange{}).Type1()):                     InsightSpecFiltersLastObservedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersNoteUpdatedAtDateRange{}).Type1()):                      InsightSpecFiltersNoteUpdatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessLaunchedAtDateRange{}).Type1()):                  InsightSpecFiltersProcessLaunchedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessTerminatedAtDateRange{}).Type1()):                InsightSpecFiltersProcessTerminatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{}).Type1()):   InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{}).Type1()):   InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceContainerLaunchedAtDateRange{}).Type1()):        InsightSpecFiltersResourceContainerLaunchedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{}).Type1()): InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersUpdatedAtDateRange{}).Type1()):                          InsightSpecFiltersUpdatedAtDateRangeCodec{},
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
type InsightSpecFiltersCodec struct {
}

func (InsightSpecFiltersCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFilters)(ptr) == nil
}

func (InsightSpecFiltersCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFilters)(ptr)
	var objs []InsightSpecFilters
	if obj != nil {
		objs = []InsightSpecFilters{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFilters{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFilters)(ptr) = InsightSpecFilters{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFilters

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFilters{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFilters)(ptr) = objs[0]
			} else {
				*(*InsightSpecFilters)(ptr) = InsightSpecFilters{}
			}
		} else {
			*(*InsightSpecFilters)(ptr) = InsightSpecFilters{}
		}
	default:
		iter.ReportError("decode InsightSpecFilters", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersCreatedAtDateRangeCodec struct {
}

func (InsightSpecFiltersCreatedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersCreatedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersCreatedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersCreatedAtDateRange)(ptr)
	var objs []InsightSpecFiltersCreatedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersCreatedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersCreatedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersCreatedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersCreatedAtDateRange)(ptr) = InsightSpecFiltersCreatedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersCreatedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersCreatedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersCreatedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersCreatedAtDateRange)(ptr) = InsightSpecFiltersCreatedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersCreatedAtDateRange)(ptr) = InsightSpecFiltersCreatedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersCreatedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersFirstObservedAtDateRangeCodec struct {
}

func (InsightSpecFiltersFirstObservedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersFirstObservedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersFirstObservedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersFirstObservedAtDateRange)(ptr)
	var objs []InsightSpecFiltersFirstObservedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersFirstObservedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersFirstObservedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersFirstObservedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersFirstObservedAtDateRange)(ptr) = InsightSpecFiltersFirstObservedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersFirstObservedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersFirstObservedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersFirstObservedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersFirstObservedAtDateRange)(ptr) = InsightSpecFiltersFirstObservedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersFirstObservedAtDateRange)(ptr) = InsightSpecFiltersFirstObservedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersFirstObservedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersLastObservedAtDateRangeCodec struct {
}

func (InsightSpecFiltersLastObservedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersLastObservedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersLastObservedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersLastObservedAtDateRange)(ptr)
	var objs []InsightSpecFiltersLastObservedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersLastObservedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersLastObservedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersLastObservedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersLastObservedAtDateRange)(ptr) = InsightSpecFiltersLastObservedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersLastObservedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersLastObservedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersLastObservedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersLastObservedAtDateRange)(ptr) = InsightSpecFiltersLastObservedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersLastObservedAtDateRange)(ptr) = InsightSpecFiltersLastObservedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersLastObservedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersNoteUpdatedAtDateRangeCodec struct {
}

func (InsightSpecFiltersNoteUpdatedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersNoteUpdatedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersNoteUpdatedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersNoteUpdatedAtDateRange)(ptr)
	var objs []InsightSpecFiltersNoteUpdatedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersNoteUpdatedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersNoteUpdatedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersNoteUpdatedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersNoteUpdatedAtDateRange)(ptr) = InsightSpecFiltersNoteUpdatedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersNoteUpdatedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersNoteUpdatedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersNoteUpdatedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersNoteUpdatedAtDateRange)(ptr) = InsightSpecFiltersNoteUpdatedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersNoteUpdatedAtDateRange)(ptr) = InsightSpecFiltersNoteUpdatedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersNoteUpdatedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersProcessLaunchedAtDateRangeCodec struct {
}

func (InsightSpecFiltersProcessLaunchedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersProcessLaunchedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersProcessLaunchedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersProcessLaunchedAtDateRange)(ptr)
	var objs []InsightSpecFiltersProcessLaunchedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersProcessLaunchedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessLaunchedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersProcessLaunchedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersProcessLaunchedAtDateRange)(ptr) = InsightSpecFiltersProcessLaunchedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersProcessLaunchedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessLaunchedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersProcessLaunchedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersProcessLaunchedAtDateRange)(ptr) = InsightSpecFiltersProcessLaunchedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersProcessLaunchedAtDateRange)(ptr) = InsightSpecFiltersProcessLaunchedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersProcessLaunchedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersProcessTerminatedAtDateRangeCodec struct {
}

func (InsightSpecFiltersProcessTerminatedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersProcessTerminatedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersProcessTerminatedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersProcessTerminatedAtDateRange)(ptr)
	var objs []InsightSpecFiltersProcessTerminatedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersProcessTerminatedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessTerminatedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersProcessTerminatedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersProcessTerminatedAtDateRange)(ptr) = InsightSpecFiltersProcessTerminatedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersProcessTerminatedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersProcessTerminatedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersProcessTerminatedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersProcessTerminatedAtDateRange)(ptr) = InsightSpecFiltersProcessTerminatedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersProcessTerminatedAtDateRange)(ptr) = InsightSpecFiltersProcessTerminatedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersProcessTerminatedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRangeCodec struct {
}

func (InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange)(ptr)
	var objs []InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange)(ptr) = InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange)(ptr) = InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange)(ptr) = InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRangeCodec struct {
}

func (InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange)(ptr)
	var objs []InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange)(ptr) = InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange)(ptr) = InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange)(ptr) = InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersResourceContainerLaunchedAtDateRangeCodec struct {
}

func (InsightSpecFiltersResourceContainerLaunchedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersResourceContainerLaunchedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersResourceContainerLaunchedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersResourceContainerLaunchedAtDateRange)(ptr)
	var objs []InsightSpecFiltersResourceContainerLaunchedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersResourceContainerLaunchedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceContainerLaunchedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersResourceContainerLaunchedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersResourceContainerLaunchedAtDateRange)(ptr) = InsightSpecFiltersResourceContainerLaunchedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersResourceContainerLaunchedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersResourceContainerLaunchedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersResourceContainerLaunchedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersResourceContainerLaunchedAtDateRange)(ptr) = InsightSpecFiltersResourceContainerLaunchedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersResourceContainerLaunchedAtDateRange)(ptr) = InsightSpecFiltersResourceContainerLaunchedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersResourceContainerLaunchedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRangeCodec struct {
}

func (InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange)(ptr)
	var objs []InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange)(ptr) = InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange)(ptr) = InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange)(ptr) = InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InsightSpecFiltersUpdatedAtDateRangeCodec struct {
}

func (InsightSpecFiltersUpdatedAtDateRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InsightSpecFiltersUpdatedAtDateRange)(ptr) == nil
}

func (InsightSpecFiltersUpdatedAtDateRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InsightSpecFiltersUpdatedAtDateRange)(ptr)
	var objs []InsightSpecFiltersUpdatedAtDateRange
	if obj != nil {
		objs = []InsightSpecFiltersUpdatedAtDateRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersUpdatedAtDateRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InsightSpecFiltersUpdatedAtDateRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InsightSpecFiltersUpdatedAtDateRange)(ptr) = InsightSpecFiltersUpdatedAtDateRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InsightSpecFiltersUpdatedAtDateRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InsightSpecFiltersUpdatedAtDateRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InsightSpecFiltersUpdatedAtDateRange)(ptr) = objs[0]
			} else {
				*(*InsightSpecFiltersUpdatedAtDateRange)(ptr) = InsightSpecFiltersUpdatedAtDateRange{}
			}
		} else {
			*(*InsightSpecFiltersUpdatedAtDateRange)(ptr) = InsightSpecFiltersUpdatedAtDateRange{}
		}
	default:
		iter.ReportError("decode InsightSpecFiltersUpdatedAtDateRange", "unexpected JSON type")
	}
}

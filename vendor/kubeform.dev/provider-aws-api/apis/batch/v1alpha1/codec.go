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
		jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResources{}).Type1()):               ComputeEnvironmentSpecComputeResourcesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResourcesLaunchTemplate{}).Type1()): ComputeEnvironmentSpecComputeResourcesLaunchTemplateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecRetryStrategy{}).Type1()):                       JobDefinitionSpecRetryStrategyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecTimeout{}).Type1()):                             JobDefinitionSpecTimeoutCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResources{}).Type1()):               ComputeEnvironmentSpecComputeResourcesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResourcesLaunchTemplate{}).Type1()): ComputeEnvironmentSpecComputeResourcesLaunchTemplateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecRetryStrategy{}).Type1()):                       JobDefinitionSpecRetryStrategyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecTimeout{}).Type1()):                             JobDefinitionSpecTimeoutCodec{},
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
type ComputeEnvironmentSpecComputeResourcesCodec struct {
}

func (ComputeEnvironmentSpecComputeResourcesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ComputeEnvironmentSpecComputeResources)(ptr) == nil
}

func (ComputeEnvironmentSpecComputeResourcesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ComputeEnvironmentSpecComputeResources)(ptr)
	var objs []ComputeEnvironmentSpecComputeResources
	if obj != nil {
		objs = []ComputeEnvironmentSpecComputeResources{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResources{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ComputeEnvironmentSpecComputeResourcesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ComputeEnvironmentSpecComputeResources)(ptr) = ComputeEnvironmentSpecComputeResources{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ComputeEnvironmentSpecComputeResources

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResources{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ComputeEnvironmentSpecComputeResources)(ptr) = objs[0]
			} else {
				*(*ComputeEnvironmentSpecComputeResources)(ptr) = ComputeEnvironmentSpecComputeResources{}
			}
		} else {
			*(*ComputeEnvironmentSpecComputeResources)(ptr) = ComputeEnvironmentSpecComputeResources{}
		}
	default:
		iter.ReportError("decode ComputeEnvironmentSpecComputeResources", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ComputeEnvironmentSpecComputeResourcesLaunchTemplateCodec struct {
}

func (ComputeEnvironmentSpecComputeResourcesLaunchTemplateCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ComputeEnvironmentSpecComputeResourcesLaunchTemplate)(ptr) == nil
}

func (ComputeEnvironmentSpecComputeResourcesLaunchTemplateCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ComputeEnvironmentSpecComputeResourcesLaunchTemplate)(ptr)
	var objs []ComputeEnvironmentSpecComputeResourcesLaunchTemplate
	if obj != nil {
		objs = []ComputeEnvironmentSpecComputeResourcesLaunchTemplate{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResourcesLaunchTemplate{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ComputeEnvironmentSpecComputeResourcesLaunchTemplateCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ComputeEnvironmentSpecComputeResourcesLaunchTemplate)(ptr) = ComputeEnvironmentSpecComputeResourcesLaunchTemplate{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ComputeEnvironmentSpecComputeResourcesLaunchTemplate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ComputeEnvironmentSpecComputeResourcesLaunchTemplate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ComputeEnvironmentSpecComputeResourcesLaunchTemplate)(ptr) = objs[0]
			} else {
				*(*ComputeEnvironmentSpecComputeResourcesLaunchTemplate)(ptr) = ComputeEnvironmentSpecComputeResourcesLaunchTemplate{}
			}
		} else {
			*(*ComputeEnvironmentSpecComputeResourcesLaunchTemplate)(ptr) = ComputeEnvironmentSpecComputeResourcesLaunchTemplate{}
		}
	default:
		iter.ReportError("decode ComputeEnvironmentSpecComputeResourcesLaunchTemplate", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type JobDefinitionSpecRetryStrategyCodec struct {
}

func (JobDefinitionSpecRetryStrategyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*JobDefinitionSpecRetryStrategy)(ptr) == nil
}

func (JobDefinitionSpecRetryStrategyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*JobDefinitionSpecRetryStrategy)(ptr)
	var objs []JobDefinitionSpecRetryStrategy
	if obj != nil {
		objs = []JobDefinitionSpecRetryStrategy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecRetryStrategy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (JobDefinitionSpecRetryStrategyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*JobDefinitionSpecRetryStrategy)(ptr) = JobDefinitionSpecRetryStrategy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []JobDefinitionSpecRetryStrategy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecRetryStrategy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*JobDefinitionSpecRetryStrategy)(ptr) = objs[0]
			} else {
				*(*JobDefinitionSpecRetryStrategy)(ptr) = JobDefinitionSpecRetryStrategy{}
			}
		} else {
			*(*JobDefinitionSpecRetryStrategy)(ptr) = JobDefinitionSpecRetryStrategy{}
		}
	default:
		iter.ReportError("decode JobDefinitionSpecRetryStrategy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type JobDefinitionSpecTimeoutCodec struct {
}

func (JobDefinitionSpecTimeoutCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*JobDefinitionSpecTimeout)(ptr) == nil
}

func (JobDefinitionSpecTimeoutCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*JobDefinitionSpecTimeout)(ptr)
	var objs []JobDefinitionSpecTimeout
	if obj != nil {
		objs = []JobDefinitionSpecTimeout{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecTimeout{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (JobDefinitionSpecTimeoutCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*JobDefinitionSpecTimeout)(ptr) = JobDefinitionSpecTimeout{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []JobDefinitionSpecTimeout

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(JobDefinitionSpecTimeout{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*JobDefinitionSpecTimeout)(ptr) = objs[0]
			} else {
				*(*JobDefinitionSpecTimeout)(ptr) = JobDefinitionSpecTimeout{}
			}
		} else {
			*(*JobDefinitionSpecTimeout)(ptr) = JobDefinitionSpecTimeout{}
		}
	default:
		iter.ReportError("decode JobDefinitionSpecTimeout", "unexpected JSON type")
	}
}
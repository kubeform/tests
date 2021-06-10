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
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestination{}).Type1()):                 SigningJobSpecDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestinationS3{}).Type1()):               SigningJobSpecDestinationS3Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSource{}).Type1()):                      SigningJobSpecSourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSourceS3{}).Type1()):                    SigningJobSpecSourceS3Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningProfileSpecSignatureValidityPeriod{}).Type1()): SigningProfileSpecSignatureValidityPeriodCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestination{}).Type1()):                 SigningJobSpecDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestinationS3{}).Type1()):               SigningJobSpecDestinationS3Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSource{}).Type1()):                      SigningJobSpecSourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSourceS3{}).Type1()):                    SigningJobSpecSourceS3Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SigningProfileSpecSignatureValidityPeriod{}).Type1()): SigningProfileSpecSignatureValidityPeriodCodec{},
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
type SigningJobSpecDestinationCodec struct {
}

func (SigningJobSpecDestinationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SigningJobSpecDestination)(ptr) == nil
}

func (SigningJobSpecDestinationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SigningJobSpecDestination)(ptr)
	var objs []SigningJobSpecDestination
	if obj != nil {
		objs = []SigningJobSpecDestination{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestination{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SigningJobSpecDestinationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SigningJobSpecDestination)(ptr) = SigningJobSpecDestination{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SigningJobSpecDestination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SigningJobSpecDestination)(ptr) = objs[0]
			} else {
				*(*SigningJobSpecDestination)(ptr) = SigningJobSpecDestination{}
			}
		} else {
			*(*SigningJobSpecDestination)(ptr) = SigningJobSpecDestination{}
		}
	default:
		iter.ReportError("decode SigningJobSpecDestination", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SigningJobSpecDestinationS3Codec struct {
}

func (SigningJobSpecDestinationS3Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SigningJobSpecDestinationS3)(ptr) == nil
}

func (SigningJobSpecDestinationS3Codec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SigningJobSpecDestinationS3)(ptr)
	var objs []SigningJobSpecDestinationS3
	if obj != nil {
		objs = []SigningJobSpecDestinationS3{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestinationS3{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SigningJobSpecDestinationS3Codec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SigningJobSpecDestinationS3)(ptr) = SigningJobSpecDestinationS3{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SigningJobSpecDestinationS3

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecDestinationS3{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SigningJobSpecDestinationS3)(ptr) = objs[0]
			} else {
				*(*SigningJobSpecDestinationS3)(ptr) = SigningJobSpecDestinationS3{}
			}
		} else {
			*(*SigningJobSpecDestinationS3)(ptr) = SigningJobSpecDestinationS3{}
		}
	default:
		iter.ReportError("decode SigningJobSpecDestinationS3", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SigningJobSpecSourceCodec struct {
}

func (SigningJobSpecSourceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SigningJobSpecSource)(ptr) == nil
}

func (SigningJobSpecSourceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SigningJobSpecSource)(ptr)
	var objs []SigningJobSpecSource
	if obj != nil {
		objs = []SigningJobSpecSource{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSource{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SigningJobSpecSourceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SigningJobSpecSource)(ptr) = SigningJobSpecSource{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SigningJobSpecSource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SigningJobSpecSource)(ptr) = objs[0]
			} else {
				*(*SigningJobSpecSource)(ptr) = SigningJobSpecSource{}
			}
		} else {
			*(*SigningJobSpecSource)(ptr) = SigningJobSpecSource{}
		}
	default:
		iter.ReportError("decode SigningJobSpecSource", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SigningJobSpecSourceS3Codec struct {
}

func (SigningJobSpecSourceS3Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SigningJobSpecSourceS3)(ptr) == nil
}

func (SigningJobSpecSourceS3Codec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SigningJobSpecSourceS3)(ptr)
	var objs []SigningJobSpecSourceS3
	if obj != nil {
		objs = []SigningJobSpecSourceS3{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSourceS3{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SigningJobSpecSourceS3Codec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SigningJobSpecSourceS3)(ptr) = SigningJobSpecSourceS3{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SigningJobSpecSourceS3

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningJobSpecSourceS3{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SigningJobSpecSourceS3)(ptr) = objs[0]
			} else {
				*(*SigningJobSpecSourceS3)(ptr) = SigningJobSpecSourceS3{}
			}
		} else {
			*(*SigningJobSpecSourceS3)(ptr) = SigningJobSpecSourceS3{}
		}
	default:
		iter.ReportError("decode SigningJobSpecSourceS3", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SigningProfileSpecSignatureValidityPeriodCodec struct {
}

func (SigningProfileSpecSignatureValidityPeriodCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SigningProfileSpecSignatureValidityPeriod)(ptr) == nil
}

func (SigningProfileSpecSignatureValidityPeriodCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SigningProfileSpecSignatureValidityPeriod)(ptr)
	var objs []SigningProfileSpecSignatureValidityPeriod
	if obj != nil {
		objs = []SigningProfileSpecSignatureValidityPeriod{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningProfileSpecSignatureValidityPeriod{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SigningProfileSpecSignatureValidityPeriodCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SigningProfileSpecSignatureValidityPeriod)(ptr) = SigningProfileSpecSignatureValidityPeriod{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SigningProfileSpecSignatureValidityPeriod

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SigningProfileSpecSignatureValidityPeriod{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SigningProfileSpecSignatureValidityPeriod)(ptr) = objs[0]
			} else {
				*(*SigningProfileSpecSignatureValidityPeriod)(ptr) = SigningProfileSpecSignatureValidityPeriod{}
			}
		} else {
			*(*SigningProfileSpecSignatureValidityPeriod)(ptr) = SigningProfileSpecSignatureValidityPeriod{}
		}
	default:
		iter.ReportError("decode SigningProfileSpecSignatureValidityPeriod", "unexpected JSON type")
	}
}

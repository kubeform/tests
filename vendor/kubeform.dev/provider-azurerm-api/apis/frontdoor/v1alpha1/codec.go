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
		jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleForwardingConfiguration{}).Type1()):      FrontdoorSpecRoutingRuleForwardingConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleRedirectConfiguration{}).Type1()):        FrontdoorSpecRoutingRuleRedirectConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{}).Type1()): CustomHTTPSConfigurationSpecCustomHTTPSConfigurationCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleForwardingConfiguration{}).Type1()):      FrontdoorSpecRoutingRuleForwardingConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleRedirectConfiguration{}).Type1()):        FrontdoorSpecRoutingRuleRedirectConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{}).Type1()): CustomHTTPSConfigurationSpecCustomHTTPSConfigurationCodec{},
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
type FrontdoorSpecRoutingRuleForwardingConfigurationCodec struct {
}

func (FrontdoorSpecRoutingRuleForwardingConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FrontdoorSpecRoutingRuleForwardingConfiguration)(ptr) == nil
}

func (FrontdoorSpecRoutingRuleForwardingConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FrontdoorSpecRoutingRuleForwardingConfiguration)(ptr)
	var objs []FrontdoorSpecRoutingRuleForwardingConfiguration
	if obj != nil {
		objs = []FrontdoorSpecRoutingRuleForwardingConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleForwardingConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FrontdoorSpecRoutingRuleForwardingConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FrontdoorSpecRoutingRuleForwardingConfiguration)(ptr) = FrontdoorSpecRoutingRuleForwardingConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FrontdoorSpecRoutingRuleForwardingConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleForwardingConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FrontdoorSpecRoutingRuleForwardingConfiguration)(ptr) = objs[0]
			} else {
				*(*FrontdoorSpecRoutingRuleForwardingConfiguration)(ptr) = FrontdoorSpecRoutingRuleForwardingConfiguration{}
			}
		} else {
			*(*FrontdoorSpecRoutingRuleForwardingConfiguration)(ptr) = FrontdoorSpecRoutingRuleForwardingConfiguration{}
		}
	default:
		iter.ReportError("decode FrontdoorSpecRoutingRuleForwardingConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FrontdoorSpecRoutingRuleRedirectConfigurationCodec struct {
}

func (FrontdoorSpecRoutingRuleRedirectConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FrontdoorSpecRoutingRuleRedirectConfiguration)(ptr) == nil
}

func (FrontdoorSpecRoutingRuleRedirectConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FrontdoorSpecRoutingRuleRedirectConfiguration)(ptr)
	var objs []FrontdoorSpecRoutingRuleRedirectConfiguration
	if obj != nil {
		objs = []FrontdoorSpecRoutingRuleRedirectConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleRedirectConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FrontdoorSpecRoutingRuleRedirectConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FrontdoorSpecRoutingRuleRedirectConfiguration)(ptr) = FrontdoorSpecRoutingRuleRedirectConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FrontdoorSpecRoutingRuleRedirectConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FrontdoorSpecRoutingRuleRedirectConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FrontdoorSpecRoutingRuleRedirectConfiguration)(ptr) = objs[0]
			} else {
				*(*FrontdoorSpecRoutingRuleRedirectConfiguration)(ptr) = FrontdoorSpecRoutingRuleRedirectConfiguration{}
			}
		} else {
			*(*FrontdoorSpecRoutingRuleRedirectConfiguration)(ptr) = FrontdoorSpecRoutingRuleRedirectConfiguration{}
		}
	default:
		iter.ReportError("decode FrontdoorSpecRoutingRuleRedirectConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CustomHTTPSConfigurationSpecCustomHTTPSConfigurationCodec struct {
}

func (CustomHTTPSConfigurationSpecCustomHTTPSConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CustomHTTPSConfigurationSpecCustomHTTPSConfiguration)(ptr) == nil
}

func (CustomHTTPSConfigurationSpecCustomHTTPSConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CustomHTTPSConfigurationSpecCustomHTTPSConfiguration)(ptr)
	var objs []CustomHTTPSConfigurationSpecCustomHTTPSConfiguration
	if obj != nil {
		objs = []CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CustomHTTPSConfigurationSpecCustomHTTPSConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CustomHTTPSConfigurationSpecCustomHTTPSConfiguration)(ptr) = CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CustomHTTPSConfigurationSpecCustomHTTPSConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CustomHTTPSConfigurationSpecCustomHTTPSConfiguration)(ptr) = objs[0]
			} else {
				*(*CustomHTTPSConfigurationSpecCustomHTTPSConfiguration)(ptr) = CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{}
			}
		} else {
			*(*CustomHTTPSConfigurationSpecCustomHTTPSConfiguration)(ptr) = CustomHTTPSConfigurationSpecCustomHTTPSConfiguration{}
		}
	default:
		iter.ReportError("decode CustomHTTPSConfigurationSpecCustomHTTPSConfiguration", "unexpected JSON type")
	}
}
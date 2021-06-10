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
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecDnssecConfig{}).Type1()):               ManagedZoneSpecDnssecConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecForwardingConfig{}).Type1()):           ManagedZoneSpecForwardingConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfig{}).Type1()):              ManagedZoneSpecPeeringConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfigTargetNetwork{}).Type1()): ManagedZoneSpecPeeringConfigTargetNetworkCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPrivateVisibilityConfig{}).Type1()):    ManagedZoneSpecPrivateVisibilityConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecAlternativeNameServerConfig{}).Type1()):     PolicySpecAlternativeNameServerConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecDnssecConfig{}).Type1()):               ManagedZoneSpecDnssecConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecForwardingConfig{}).Type1()):           ManagedZoneSpecForwardingConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfig{}).Type1()):              ManagedZoneSpecPeeringConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfigTargetNetwork{}).Type1()): ManagedZoneSpecPeeringConfigTargetNetworkCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPrivateVisibilityConfig{}).Type1()):    ManagedZoneSpecPrivateVisibilityConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecAlternativeNameServerConfig{}).Type1()):     PolicySpecAlternativeNameServerConfigCodec{},
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
type ManagedZoneSpecDnssecConfigCodec struct {
}

func (ManagedZoneSpecDnssecConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagedZoneSpecDnssecConfig)(ptr) == nil
}

func (ManagedZoneSpecDnssecConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagedZoneSpecDnssecConfig)(ptr)
	var objs []ManagedZoneSpecDnssecConfig
	if obj != nil {
		objs = []ManagedZoneSpecDnssecConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecDnssecConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagedZoneSpecDnssecConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagedZoneSpecDnssecConfig)(ptr) = ManagedZoneSpecDnssecConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagedZoneSpecDnssecConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecDnssecConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagedZoneSpecDnssecConfig)(ptr) = objs[0]
			} else {
				*(*ManagedZoneSpecDnssecConfig)(ptr) = ManagedZoneSpecDnssecConfig{}
			}
		} else {
			*(*ManagedZoneSpecDnssecConfig)(ptr) = ManagedZoneSpecDnssecConfig{}
		}
	default:
		iter.ReportError("decode ManagedZoneSpecDnssecConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ManagedZoneSpecForwardingConfigCodec struct {
}

func (ManagedZoneSpecForwardingConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagedZoneSpecForwardingConfig)(ptr) == nil
}

func (ManagedZoneSpecForwardingConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagedZoneSpecForwardingConfig)(ptr)
	var objs []ManagedZoneSpecForwardingConfig
	if obj != nil {
		objs = []ManagedZoneSpecForwardingConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecForwardingConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagedZoneSpecForwardingConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagedZoneSpecForwardingConfig)(ptr) = ManagedZoneSpecForwardingConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagedZoneSpecForwardingConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecForwardingConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagedZoneSpecForwardingConfig)(ptr) = objs[0]
			} else {
				*(*ManagedZoneSpecForwardingConfig)(ptr) = ManagedZoneSpecForwardingConfig{}
			}
		} else {
			*(*ManagedZoneSpecForwardingConfig)(ptr) = ManagedZoneSpecForwardingConfig{}
		}
	default:
		iter.ReportError("decode ManagedZoneSpecForwardingConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ManagedZoneSpecPeeringConfigCodec struct {
}

func (ManagedZoneSpecPeeringConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagedZoneSpecPeeringConfig)(ptr) == nil
}

func (ManagedZoneSpecPeeringConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagedZoneSpecPeeringConfig)(ptr)
	var objs []ManagedZoneSpecPeeringConfig
	if obj != nil {
		objs = []ManagedZoneSpecPeeringConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagedZoneSpecPeeringConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagedZoneSpecPeeringConfig)(ptr) = ManagedZoneSpecPeeringConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagedZoneSpecPeeringConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagedZoneSpecPeeringConfig)(ptr) = objs[0]
			} else {
				*(*ManagedZoneSpecPeeringConfig)(ptr) = ManagedZoneSpecPeeringConfig{}
			}
		} else {
			*(*ManagedZoneSpecPeeringConfig)(ptr) = ManagedZoneSpecPeeringConfig{}
		}
	default:
		iter.ReportError("decode ManagedZoneSpecPeeringConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ManagedZoneSpecPeeringConfigTargetNetworkCodec struct {
}

func (ManagedZoneSpecPeeringConfigTargetNetworkCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagedZoneSpecPeeringConfigTargetNetwork)(ptr) == nil
}

func (ManagedZoneSpecPeeringConfigTargetNetworkCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagedZoneSpecPeeringConfigTargetNetwork)(ptr)
	var objs []ManagedZoneSpecPeeringConfigTargetNetwork
	if obj != nil {
		objs = []ManagedZoneSpecPeeringConfigTargetNetwork{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfigTargetNetwork{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagedZoneSpecPeeringConfigTargetNetworkCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagedZoneSpecPeeringConfigTargetNetwork)(ptr) = ManagedZoneSpecPeeringConfigTargetNetwork{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagedZoneSpecPeeringConfigTargetNetwork

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPeeringConfigTargetNetwork{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagedZoneSpecPeeringConfigTargetNetwork)(ptr) = objs[0]
			} else {
				*(*ManagedZoneSpecPeeringConfigTargetNetwork)(ptr) = ManagedZoneSpecPeeringConfigTargetNetwork{}
			}
		} else {
			*(*ManagedZoneSpecPeeringConfigTargetNetwork)(ptr) = ManagedZoneSpecPeeringConfigTargetNetwork{}
		}
	default:
		iter.ReportError("decode ManagedZoneSpecPeeringConfigTargetNetwork", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ManagedZoneSpecPrivateVisibilityConfigCodec struct {
}

func (ManagedZoneSpecPrivateVisibilityConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagedZoneSpecPrivateVisibilityConfig)(ptr) == nil
}

func (ManagedZoneSpecPrivateVisibilityConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagedZoneSpecPrivateVisibilityConfig)(ptr)
	var objs []ManagedZoneSpecPrivateVisibilityConfig
	if obj != nil {
		objs = []ManagedZoneSpecPrivateVisibilityConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPrivateVisibilityConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagedZoneSpecPrivateVisibilityConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagedZoneSpecPrivateVisibilityConfig)(ptr) = ManagedZoneSpecPrivateVisibilityConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagedZoneSpecPrivateVisibilityConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedZoneSpecPrivateVisibilityConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagedZoneSpecPrivateVisibilityConfig)(ptr) = objs[0]
			} else {
				*(*ManagedZoneSpecPrivateVisibilityConfig)(ptr) = ManagedZoneSpecPrivateVisibilityConfig{}
			}
		} else {
			*(*ManagedZoneSpecPrivateVisibilityConfig)(ptr) = ManagedZoneSpecPrivateVisibilityConfig{}
		}
	default:
		iter.ReportError("decode ManagedZoneSpecPrivateVisibilityConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecAlternativeNameServerConfigCodec struct {
}

func (PolicySpecAlternativeNameServerConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecAlternativeNameServerConfig)(ptr) == nil
}

func (PolicySpecAlternativeNameServerConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecAlternativeNameServerConfig)(ptr)
	var objs []PolicySpecAlternativeNameServerConfig
	if obj != nil {
		objs = []PolicySpecAlternativeNameServerConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecAlternativeNameServerConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecAlternativeNameServerConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecAlternativeNameServerConfig)(ptr) = PolicySpecAlternativeNameServerConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecAlternativeNameServerConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecAlternativeNameServerConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecAlternativeNameServerConfig)(ptr) = objs[0]
			} else {
				*(*PolicySpecAlternativeNameServerConfig)(ptr) = PolicySpecAlternativeNameServerConfig{}
			}
		} else {
			*(*PolicySpecAlternativeNameServerConfig)(ptr) = PolicySpecAlternativeNameServerConfig{}
		}
	default:
		iter.ReportError("decode PolicySpecAlternativeNameServerConfig", "unexpected JSON type")
	}
}

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
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptions{}).Type1()):                  ElasticsearchDomainSpecAdvancedSecurityOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{}).Type1()): ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfig{}).Type1()):                            ElasticsearchDomainSpecClusterConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{}).Type1()):         ElasticsearchDomainSpecClusterConfigZoneAwarenessConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecCognitoOptions{}).Type1()):                           ElasticsearchDomainSpecCognitoOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecDomainEndpointOptions{}).Type1()):                    ElasticsearchDomainSpecDomainEndpointOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEbsOptions{}).Type1()):                               ElasticsearchDomainSpecEbsOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEncryptAtRest{}).Type1()):                            ElasticsearchDomainSpecEncryptAtRestCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecNodeToNodeEncryption{}).Type1()):                     ElasticsearchDomainSpecNodeToNodeEncryptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecSnapshotOptions{}).Type1()):                          ElasticsearchDomainSpecSnapshotOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecVpcOptions{}).Type1()):                               ElasticsearchDomainSpecVpcOptionsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptions{}).Type1()):                  ElasticsearchDomainSpecAdvancedSecurityOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{}).Type1()): ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfig{}).Type1()):                            ElasticsearchDomainSpecClusterConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{}).Type1()):         ElasticsearchDomainSpecClusterConfigZoneAwarenessConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecCognitoOptions{}).Type1()):                           ElasticsearchDomainSpecCognitoOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecDomainEndpointOptions{}).Type1()):                    ElasticsearchDomainSpecDomainEndpointOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEbsOptions{}).Type1()):                               ElasticsearchDomainSpecEbsOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEncryptAtRest{}).Type1()):                            ElasticsearchDomainSpecEncryptAtRestCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecNodeToNodeEncryption{}).Type1()):                     ElasticsearchDomainSpecNodeToNodeEncryptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecSnapshotOptions{}).Type1()):                          ElasticsearchDomainSpecSnapshotOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecVpcOptions{}).Type1()):                               ElasticsearchDomainSpecVpcOptionsCodec{},
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
type ElasticsearchDomainSpecAdvancedSecurityOptionsCodec struct {
}

func (ElasticsearchDomainSpecAdvancedSecurityOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecAdvancedSecurityOptions)(ptr) == nil
}

func (ElasticsearchDomainSpecAdvancedSecurityOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecAdvancedSecurityOptions)(ptr)
	var objs []ElasticsearchDomainSpecAdvancedSecurityOptions
	if obj != nil {
		objs = []ElasticsearchDomainSpecAdvancedSecurityOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecAdvancedSecurityOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecAdvancedSecurityOptions)(ptr) = ElasticsearchDomainSpecAdvancedSecurityOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecAdvancedSecurityOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecAdvancedSecurityOptions)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecAdvancedSecurityOptions)(ptr) = ElasticsearchDomainSpecAdvancedSecurityOptions{}
			}
		} else {
			*(*ElasticsearchDomainSpecAdvancedSecurityOptions)(ptr) = ElasticsearchDomainSpecAdvancedSecurityOptions{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecAdvancedSecurityOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptionsCodec struct {
}

func (ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)(ptr) == nil
}

func (ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)(ptr)
	var objs []ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions
	if obj != nil {
		objs = []ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)(ptr) = ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)(ptr) = ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{}
			}
		} else {
			*(*ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions)(ptr) = ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecAdvancedSecurityOptionsMasterUserOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecClusterConfigCodec struct {
}

func (ElasticsearchDomainSpecClusterConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecClusterConfig)(ptr) == nil
}

func (ElasticsearchDomainSpecClusterConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecClusterConfig)(ptr)
	var objs []ElasticsearchDomainSpecClusterConfig
	if obj != nil {
		objs = []ElasticsearchDomainSpecClusterConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecClusterConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecClusterConfig)(ptr) = ElasticsearchDomainSpecClusterConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecClusterConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecClusterConfig)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecClusterConfig)(ptr) = ElasticsearchDomainSpecClusterConfig{}
			}
		} else {
			*(*ElasticsearchDomainSpecClusterConfig)(ptr) = ElasticsearchDomainSpecClusterConfig{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecClusterConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecClusterConfigZoneAwarenessConfigCodec struct {
}

func (ElasticsearchDomainSpecClusterConfigZoneAwarenessConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)(ptr) == nil
}

func (ElasticsearchDomainSpecClusterConfigZoneAwarenessConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)(ptr)
	var objs []ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig
	if obj != nil {
		objs = []ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecClusterConfigZoneAwarenessConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)(ptr) = ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)(ptr) = ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{}
			}
		} else {
			*(*ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig)(ptr) = ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecClusterConfigZoneAwarenessConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecCognitoOptionsCodec struct {
}

func (ElasticsearchDomainSpecCognitoOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecCognitoOptions)(ptr) == nil
}

func (ElasticsearchDomainSpecCognitoOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecCognitoOptions)(ptr)
	var objs []ElasticsearchDomainSpecCognitoOptions
	if obj != nil {
		objs = []ElasticsearchDomainSpecCognitoOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecCognitoOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecCognitoOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecCognitoOptions)(ptr) = ElasticsearchDomainSpecCognitoOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecCognitoOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecCognitoOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecCognitoOptions)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecCognitoOptions)(ptr) = ElasticsearchDomainSpecCognitoOptions{}
			}
		} else {
			*(*ElasticsearchDomainSpecCognitoOptions)(ptr) = ElasticsearchDomainSpecCognitoOptions{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecCognitoOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecDomainEndpointOptionsCodec struct {
}

func (ElasticsearchDomainSpecDomainEndpointOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecDomainEndpointOptions)(ptr) == nil
}

func (ElasticsearchDomainSpecDomainEndpointOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecDomainEndpointOptions)(ptr)
	var objs []ElasticsearchDomainSpecDomainEndpointOptions
	if obj != nil {
		objs = []ElasticsearchDomainSpecDomainEndpointOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecDomainEndpointOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecDomainEndpointOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecDomainEndpointOptions)(ptr) = ElasticsearchDomainSpecDomainEndpointOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecDomainEndpointOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecDomainEndpointOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecDomainEndpointOptions)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecDomainEndpointOptions)(ptr) = ElasticsearchDomainSpecDomainEndpointOptions{}
			}
		} else {
			*(*ElasticsearchDomainSpecDomainEndpointOptions)(ptr) = ElasticsearchDomainSpecDomainEndpointOptions{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecDomainEndpointOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecEbsOptionsCodec struct {
}

func (ElasticsearchDomainSpecEbsOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecEbsOptions)(ptr) == nil
}

func (ElasticsearchDomainSpecEbsOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecEbsOptions)(ptr)
	var objs []ElasticsearchDomainSpecEbsOptions
	if obj != nil {
		objs = []ElasticsearchDomainSpecEbsOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEbsOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecEbsOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecEbsOptions)(ptr) = ElasticsearchDomainSpecEbsOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecEbsOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEbsOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecEbsOptions)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecEbsOptions)(ptr) = ElasticsearchDomainSpecEbsOptions{}
			}
		} else {
			*(*ElasticsearchDomainSpecEbsOptions)(ptr) = ElasticsearchDomainSpecEbsOptions{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecEbsOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecEncryptAtRestCodec struct {
}

func (ElasticsearchDomainSpecEncryptAtRestCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecEncryptAtRest)(ptr) == nil
}

func (ElasticsearchDomainSpecEncryptAtRestCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecEncryptAtRest)(ptr)
	var objs []ElasticsearchDomainSpecEncryptAtRest
	if obj != nil {
		objs = []ElasticsearchDomainSpecEncryptAtRest{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEncryptAtRest{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecEncryptAtRestCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecEncryptAtRest)(ptr) = ElasticsearchDomainSpecEncryptAtRest{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecEncryptAtRest

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecEncryptAtRest{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecEncryptAtRest)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecEncryptAtRest)(ptr) = ElasticsearchDomainSpecEncryptAtRest{}
			}
		} else {
			*(*ElasticsearchDomainSpecEncryptAtRest)(ptr) = ElasticsearchDomainSpecEncryptAtRest{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecEncryptAtRest", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecNodeToNodeEncryptionCodec struct {
}

func (ElasticsearchDomainSpecNodeToNodeEncryptionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecNodeToNodeEncryption)(ptr) == nil
}

func (ElasticsearchDomainSpecNodeToNodeEncryptionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecNodeToNodeEncryption)(ptr)
	var objs []ElasticsearchDomainSpecNodeToNodeEncryption
	if obj != nil {
		objs = []ElasticsearchDomainSpecNodeToNodeEncryption{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecNodeToNodeEncryption{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecNodeToNodeEncryptionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecNodeToNodeEncryption)(ptr) = ElasticsearchDomainSpecNodeToNodeEncryption{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecNodeToNodeEncryption

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecNodeToNodeEncryption{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecNodeToNodeEncryption)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecNodeToNodeEncryption)(ptr) = ElasticsearchDomainSpecNodeToNodeEncryption{}
			}
		} else {
			*(*ElasticsearchDomainSpecNodeToNodeEncryption)(ptr) = ElasticsearchDomainSpecNodeToNodeEncryption{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecNodeToNodeEncryption", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecSnapshotOptionsCodec struct {
}

func (ElasticsearchDomainSpecSnapshotOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecSnapshotOptions)(ptr) == nil
}

func (ElasticsearchDomainSpecSnapshotOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecSnapshotOptions)(ptr)
	var objs []ElasticsearchDomainSpecSnapshotOptions
	if obj != nil {
		objs = []ElasticsearchDomainSpecSnapshotOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecSnapshotOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecSnapshotOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecSnapshotOptions)(ptr) = ElasticsearchDomainSpecSnapshotOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecSnapshotOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecSnapshotOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecSnapshotOptions)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecSnapshotOptions)(ptr) = ElasticsearchDomainSpecSnapshotOptions{}
			}
		} else {
			*(*ElasticsearchDomainSpecSnapshotOptions)(ptr) = ElasticsearchDomainSpecSnapshotOptions{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecSnapshotOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ElasticsearchDomainSpecVpcOptionsCodec struct {
}

func (ElasticsearchDomainSpecVpcOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ElasticsearchDomainSpecVpcOptions)(ptr) == nil
}

func (ElasticsearchDomainSpecVpcOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ElasticsearchDomainSpecVpcOptions)(ptr)
	var objs []ElasticsearchDomainSpecVpcOptions
	if obj != nil {
		objs = []ElasticsearchDomainSpecVpcOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecVpcOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ElasticsearchDomainSpecVpcOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ElasticsearchDomainSpecVpcOptions)(ptr) = ElasticsearchDomainSpecVpcOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ElasticsearchDomainSpecVpcOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ElasticsearchDomainSpecVpcOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ElasticsearchDomainSpecVpcOptions)(ptr) = objs[0]
			} else {
				*(*ElasticsearchDomainSpecVpcOptions)(ptr) = ElasticsearchDomainSpecVpcOptions{}
			}
		} else {
			*(*ElasticsearchDomainSpecVpcOptions)(ptr) = ElasticsearchDomainSpecVpcOptions{}
		}
	default:
		iter.ReportError("decode ElasticsearchDomainSpecVpcOptions", "unexpected JSON type")
	}
}

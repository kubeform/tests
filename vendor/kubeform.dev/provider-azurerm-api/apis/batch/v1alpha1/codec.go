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
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecKeyVaultReference{}).Type1()):          AccountSpecKeyVaultReferenceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecAutoScale{}).Type1()):                     PoolSpecAutoScaleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecContainerConfiguration{}).Type1()):        PoolSpecContainerConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecFixedScale{}).Type1()):                    PoolSpecFixedScaleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecNetworkConfiguration{}).Type1()):          PoolSpecNetworkConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTask{}).Type1()):                     PoolSpecStartTaskCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentity{}).Type1()):         PoolSpecStartTaskUserIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentityAutoUser{}).Type1()): PoolSpecStartTaskUserIdentityAutoUserCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStorageImageReference{}).Type1()):         PoolSpecStorageImageReferenceCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecKeyVaultReference{}).Type1()):          AccountSpecKeyVaultReferenceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecAutoScale{}).Type1()):                     PoolSpecAutoScaleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecContainerConfiguration{}).Type1()):        PoolSpecContainerConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecFixedScale{}).Type1()):                    PoolSpecFixedScaleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecNetworkConfiguration{}).Type1()):          PoolSpecNetworkConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTask{}).Type1()):                     PoolSpecStartTaskCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentity{}).Type1()):         PoolSpecStartTaskUserIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentityAutoUser{}).Type1()): PoolSpecStartTaskUserIdentityAutoUserCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStorageImageReference{}).Type1()):         PoolSpecStorageImageReferenceCodec{},
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
type AccountSpecKeyVaultReferenceCodec struct {
}

func (AccountSpecKeyVaultReferenceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccountSpecKeyVaultReference)(ptr) == nil
}

func (AccountSpecKeyVaultReferenceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccountSpecKeyVaultReference)(ptr)
	var objs []AccountSpecKeyVaultReference
	if obj != nil {
		objs = []AccountSpecKeyVaultReference{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecKeyVaultReference{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccountSpecKeyVaultReferenceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccountSpecKeyVaultReference)(ptr) = AccountSpecKeyVaultReference{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccountSpecKeyVaultReference

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecKeyVaultReference{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccountSpecKeyVaultReference)(ptr) = objs[0]
			} else {
				*(*AccountSpecKeyVaultReference)(ptr) = AccountSpecKeyVaultReference{}
			}
		} else {
			*(*AccountSpecKeyVaultReference)(ptr) = AccountSpecKeyVaultReference{}
		}
	default:
		iter.ReportError("decode AccountSpecKeyVaultReference", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecAutoScaleCodec struct {
}

func (PoolSpecAutoScaleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecAutoScale)(ptr) == nil
}

func (PoolSpecAutoScaleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecAutoScale)(ptr)
	var objs []PoolSpecAutoScale
	if obj != nil {
		objs = []PoolSpecAutoScale{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecAutoScale{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecAutoScaleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecAutoScale)(ptr) = PoolSpecAutoScale{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecAutoScale

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecAutoScale{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecAutoScale)(ptr) = objs[0]
			} else {
				*(*PoolSpecAutoScale)(ptr) = PoolSpecAutoScale{}
			}
		} else {
			*(*PoolSpecAutoScale)(ptr) = PoolSpecAutoScale{}
		}
	default:
		iter.ReportError("decode PoolSpecAutoScale", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecContainerConfigurationCodec struct {
}

func (PoolSpecContainerConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecContainerConfiguration)(ptr) == nil
}

func (PoolSpecContainerConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecContainerConfiguration)(ptr)
	var objs []PoolSpecContainerConfiguration
	if obj != nil {
		objs = []PoolSpecContainerConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecContainerConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecContainerConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecContainerConfiguration)(ptr) = PoolSpecContainerConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecContainerConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecContainerConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecContainerConfiguration)(ptr) = objs[0]
			} else {
				*(*PoolSpecContainerConfiguration)(ptr) = PoolSpecContainerConfiguration{}
			}
		} else {
			*(*PoolSpecContainerConfiguration)(ptr) = PoolSpecContainerConfiguration{}
		}
	default:
		iter.ReportError("decode PoolSpecContainerConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecFixedScaleCodec struct {
}

func (PoolSpecFixedScaleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecFixedScale)(ptr) == nil
}

func (PoolSpecFixedScaleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecFixedScale)(ptr)
	var objs []PoolSpecFixedScale
	if obj != nil {
		objs = []PoolSpecFixedScale{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecFixedScale{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecFixedScaleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecFixedScale)(ptr) = PoolSpecFixedScale{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecFixedScale

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecFixedScale{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecFixedScale)(ptr) = objs[0]
			} else {
				*(*PoolSpecFixedScale)(ptr) = PoolSpecFixedScale{}
			}
		} else {
			*(*PoolSpecFixedScale)(ptr) = PoolSpecFixedScale{}
		}
	default:
		iter.ReportError("decode PoolSpecFixedScale", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecNetworkConfigurationCodec struct {
}

func (PoolSpecNetworkConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecNetworkConfiguration)(ptr) == nil
}

func (PoolSpecNetworkConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecNetworkConfiguration)(ptr)
	var objs []PoolSpecNetworkConfiguration
	if obj != nil {
		objs = []PoolSpecNetworkConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecNetworkConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecNetworkConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecNetworkConfiguration)(ptr) = PoolSpecNetworkConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecNetworkConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecNetworkConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecNetworkConfiguration)(ptr) = objs[0]
			} else {
				*(*PoolSpecNetworkConfiguration)(ptr) = PoolSpecNetworkConfiguration{}
			}
		} else {
			*(*PoolSpecNetworkConfiguration)(ptr) = PoolSpecNetworkConfiguration{}
		}
	default:
		iter.ReportError("decode PoolSpecNetworkConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecStartTaskCodec struct {
}

func (PoolSpecStartTaskCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecStartTask)(ptr) == nil
}

func (PoolSpecStartTaskCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecStartTask)(ptr)
	var objs []PoolSpecStartTask
	if obj != nil {
		objs = []PoolSpecStartTask{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTask{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecStartTaskCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecStartTask)(ptr) = PoolSpecStartTask{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecStartTask

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTask{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecStartTask)(ptr) = objs[0]
			} else {
				*(*PoolSpecStartTask)(ptr) = PoolSpecStartTask{}
			}
		} else {
			*(*PoolSpecStartTask)(ptr) = PoolSpecStartTask{}
		}
	default:
		iter.ReportError("decode PoolSpecStartTask", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecStartTaskUserIdentityCodec struct {
}

func (PoolSpecStartTaskUserIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecStartTaskUserIdentity)(ptr) == nil
}

func (PoolSpecStartTaskUserIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecStartTaskUserIdentity)(ptr)
	var objs []PoolSpecStartTaskUserIdentity
	if obj != nil {
		objs = []PoolSpecStartTaskUserIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecStartTaskUserIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecStartTaskUserIdentity)(ptr) = PoolSpecStartTaskUserIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecStartTaskUserIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecStartTaskUserIdentity)(ptr) = objs[0]
			} else {
				*(*PoolSpecStartTaskUserIdentity)(ptr) = PoolSpecStartTaskUserIdentity{}
			}
		} else {
			*(*PoolSpecStartTaskUserIdentity)(ptr) = PoolSpecStartTaskUserIdentity{}
		}
	default:
		iter.ReportError("decode PoolSpecStartTaskUserIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecStartTaskUserIdentityAutoUserCodec struct {
}

func (PoolSpecStartTaskUserIdentityAutoUserCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecStartTaskUserIdentityAutoUser)(ptr) == nil
}

func (PoolSpecStartTaskUserIdentityAutoUserCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecStartTaskUserIdentityAutoUser)(ptr)
	var objs []PoolSpecStartTaskUserIdentityAutoUser
	if obj != nil {
		objs = []PoolSpecStartTaskUserIdentityAutoUser{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentityAutoUser{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecStartTaskUserIdentityAutoUserCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecStartTaskUserIdentityAutoUser)(ptr) = PoolSpecStartTaskUserIdentityAutoUser{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecStartTaskUserIdentityAutoUser

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStartTaskUserIdentityAutoUser{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecStartTaskUserIdentityAutoUser)(ptr) = objs[0]
			} else {
				*(*PoolSpecStartTaskUserIdentityAutoUser)(ptr) = PoolSpecStartTaskUserIdentityAutoUser{}
			}
		} else {
			*(*PoolSpecStartTaskUserIdentityAutoUser)(ptr) = PoolSpecStartTaskUserIdentityAutoUser{}
		}
	default:
		iter.ReportError("decode PoolSpecStartTaskUserIdentityAutoUser", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PoolSpecStorageImageReferenceCodec struct {
}

func (PoolSpecStorageImageReferenceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PoolSpecStorageImageReference)(ptr) == nil
}

func (PoolSpecStorageImageReferenceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PoolSpecStorageImageReference)(ptr)
	var objs []PoolSpecStorageImageReference
	if obj != nil {
		objs = []PoolSpecStorageImageReference{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStorageImageReference{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PoolSpecStorageImageReferenceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PoolSpecStorageImageReference)(ptr) = PoolSpecStorageImageReference{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PoolSpecStorageImageReference

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PoolSpecStorageImageReference{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PoolSpecStorageImageReference)(ptr) = objs[0]
			} else {
				*(*PoolSpecStorageImageReference)(ptr) = PoolSpecStorageImageReference{}
			}
		} else {
			*(*PoolSpecStorageImageReference)(ptr) = PoolSpecStorageImageReference{}
		}
	default:
		iter.ReportError("decode PoolSpecStorageImageReference", "unexpected JSON type")
	}
}

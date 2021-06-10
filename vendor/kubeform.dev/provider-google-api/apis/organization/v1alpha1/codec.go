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
		jsoniter.MustGetKind(reflect2.TypeOf(IamBindingSpecCondition{}).Type1()):   IamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IamMemberSpecCondition{}).Type1()):    IamMemberSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecBooleanPolicy{}).Type1()):   PolicySpecBooleanPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicy{}).Type1()):      PolicySpecListPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyAllow{}).Type1()): PolicySpecListPolicyAllowCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyDeny{}).Type1()):  PolicySpecListPolicyDenyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecRestorePolicy{}).Type1()):   PolicySpecRestorePolicyCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(IamBindingSpecCondition{}).Type1()):   IamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IamMemberSpecCondition{}).Type1()):    IamMemberSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecBooleanPolicy{}).Type1()):   PolicySpecBooleanPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicy{}).Type1()):      PolicySpecListPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyAllow{}).Type1()): PolicySpecListPolicyAllowCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyDeny{}).Type1()):  PolicySpecListPolicyDenyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecRestorePolicy{}).Type1()):   PolicySpecRestorePolicyCodec{},
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
type IamBindingSpecConditionCodec struct {
}

func (IamBindingSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*IamBindingSpecCondition)(ptr) == nil
}

func (IamBindingSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*IamBindingSpecCondition)(ptr)
	var objs []IamBindingSpecCondition
	if obj != nil {
		objs = []IamBindingSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IamBindingSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (IamBindingSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*IamBindingSpecCondition)(ptr) = IamBindingSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []IamBindingSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IamBindingSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*IamBindingSpecCondition)(ptr) = objs[0]
			} else {
				*(*IamBindingSpecCondition)(ptr) = IamBindingSpecCondition{}
			}
		} else {
			*(*IamBindingSpecCondition)(ptr) = IamBindingSpecCondition{}
		}
	default:
		iter.ReportError("decode IamBindingSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type IamMemberSpecConditionCodec struct {
}

func (IamMemberSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*IamMemberSpecCondition)(ptr) == nil
}

func (IamMemberSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*IamMemberSpecCondition)(ptr)
	var objs []IamMemberSpecCondition
	if obj != nil {
		objs = []IamMemberSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IamMemberSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (IamMemberSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*IamMemberSpecCondition)(ptr) = IamMemberSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []IamMemberSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IamMemberSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*IamMemberSpecCondition)(ptr) = objs[0]
			} else {
				*(*IamMemberSpecCondition)(ptr) = IamMemberSpecCondition{}
			}
		} else {
			*(*IamMemberSpecCondition)(ptr) = IamMemberSpecCondition{}
		}
	default:
		iter.ReportError("decode IamMemberSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecBooleanPolicyCodec struct {
}

func (PolicySpecBooleanPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecBooleanPolicy)(ptr) == nil
}

func (PolicySpecBooleanPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecBooleanPolicy)(ptr)
	var objs []PolicySpecBooleanPolicy
	if obj != nil {
		objs = []PolicySpecBooleanPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecBooleanPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecBooleanPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecBooleanPolicy)(ptr) = PolicySpecBooleanPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecBooleanPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecBooleanPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecBooleanPolicy)(ptr) = objs[0]
			} else {
				*(*PolicySpecBooleanPolicy)(ptr) = PolicySpecBooleanPolicy{}
			}
		} else {
			*(*PolicySpecBooleanPolicy)(ptr) = PolicySpecBooleanPolicy{}
		}
	default:
		iter.ReportError("decode PolicySpecBooleanPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecListPolicyCodec struct {
}

func (PolicySpecListPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecListPolicy)(ptr) == nil
}

func (PolicySpecListPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecListPolicy)(ptr)
	var objs []PolicySpecListPolicy
	if obj != nil {
		objs = []PolicySpecListPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecListPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecListPolicy)(ptr) = PolicySpecListPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecListPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecListPolicy)(ptr) = objs[0]
			} else {
				*(*PolicySpecListPolicy)(ptr) = PolicySpecListPolicy{}
			}
		} else {
			*(*PolicySpecListPolicy)(ptr) = PolicySpecListPolicy{}
		}
	default:
		iter.ReportError("decode PolicySpecListPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecListPolicyAllowCodec struct {
}

func (PolicySpecListPolicyAllowCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecListPolicyAllow)(ptr) == nil
}

func (PolicySpecListPolicyAllowCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecListPolicyAllow)(ptr)
	var objs []PolicySpecListPolicyAllow
	if obj != nil {
		objs = []PolicySpecListPolicyAllow{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyAllow{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecListPolicyAllowCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecListPolicyAllow)(ptr) = PolicySpecListPolicyAllow{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecListPolicyAllow

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyAllow{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecListPolicyAllow)(ptr) = objs[0]
			} else {
				*(*PolicySpecListPolicyAllow)(ptr) = PolicySpecListPolicyAllow{}
			}
		} else {
			*(*PolicySpecListPolicyAllow)(ptr) = PolicySpecListPolicyAllow{}
		}
	default:
		iter.ReportError("decode PolicySpecListPolicyAllow", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecListPolicyDenyCodec struct {
}

func (PolicySpecListPolicyDenyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecListPolicyDeny)(ptr) == nil
}

func (PolicySpecListPolicyDenyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecListPolicyDeny)(ptr)
	var objs []PolicySpecListPolicyDeny
	if obj != nil {
		objs = []PolicySpecListPolicyDeny{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyDeny{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecListPolicyDenyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecListPolicyDeny)(ptr) = PolicySpecListPolicyDeny{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecListPolicyDeny

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecListPolicyDeny{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecListPolicyDeny)(ptr) = objs[0]
			} else {
				*(*PolicySpecListPolicyDeny)(ptr) = PolicySpecListPolicyDeny{}
			}
		} else {
			*(*PolicySpecListPolicyDeny)(ptr) = PolicySpecListPolicyDeny{}
		}
	default:
		iter.ReportError("decode PolicySpecListPolicyDeny", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecRestorePolicyCodec struct {
}

func (PolicySpecRestorePolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecRestorePolicy)(ptr) == nil
}

func (PolicySpecRestorePolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecRestorePolicy)(ptr)
	var objs []PolicySpecRestorePolicy
	if obj != nil {
		objs = []PolicySpecRestorePolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecRestorePolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecRestorePolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecRestorePolicy)(ptr) = PolicySpecRestorePolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecRestorePolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecRestorePolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecRestorePolicy)(ptr) = objs[0]
			} else {
				*(*PolicySpecRestorePolicy)(ptr) = PolicySpecRestorePolicy{}
			}
		} else {
			*(*PolicySpecRestorePolicy)(ptr) = PolicySpecRestorePolicy{}
		}
	default:
		iter.ReportError("decode PolicySpecRestorePolicy", "unexpected JSON type")
	}
}
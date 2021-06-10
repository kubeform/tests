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
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecBackup{}).Type1()):                        AccountSpecBackupCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecConsistencyPolicy{}).Type1()):             AccountSpecConsistencyPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecCorsRule{}).Type1()):                      AccountSpecCorsRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecIdentity{}).Type1()):                      AccountSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CassandraKeyspaceSpecAutoscaleSettings{}).Type1()):   CassandraKeyspaceSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecAutoscaleSettings{}).Type1()):      CassandraTableSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecSchema{}).Type1()):                 CassandraTableSpecSchemaCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GremlinDatabaseSpecAutoscaleSettings{}).Type1()):     GremlinDatabaseSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecAutoscaleSettings{}).Type1()):        GremlinGraphSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecConflictResolutionPolicy{}).Type1()): GremlinGraphSpecConflictResolutionPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MongoCollectionSpecAutoscaleSettings{}).Type1()):     MongoCollectionSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MongoDatabaseSpecAutoscaleSettings{}).Type1()):       MongoDatabaseSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecAutoscaleSettings{}).Type1()):        SqlContainerSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecConflictResolutionPolicy{}).Type1()): SqlContainerSpecConflictResolutionPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecIndexingPolicy{}).Type1()):           SqlContainerSpecIndexingPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlDatabaseSpecAutoscaleSettings{}).Type1()):         SqlDatabaseSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecAutoscaleSettings{}).Type1()):               TableSpecAutoscaleSettingsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecBackup{}).Type1()):                        AccountSpecBackupCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecConsistencyPolicy{}).Type1()):             AccountSpecConsistencyPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecCorsRule{}).Type1()):                      AccountSpecCorsRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecIdentity{}).Type1()):                      AccountSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CassandraKeyspaceSpecAutoscaleSettings{}).Type1()):   CassandraKeyspaceSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecAutoscaleSettings{}).Type1()):      CassandraTableSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecSchema{}).Type1()):                 CassandraTableSpecSchemaCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GremlinDatabaseSpecAutoscaleSettings{}).Type1()):     GremlinDatabaseSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecAutoscaleSettings{}).Type1()):        GremlinGraphSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecConflictResolutionPolicy{}).Type1()): GremlinGraphSpecConflictResolutionPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MongoCollectionSpecAutoscaleSettings{}).Type1()):     MongoCollectionSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MongoDatabaseSpecAutoscaleSettings{}).Type1()):       MongoDatabaseSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecAutoscaleSettings{}).Type1()):        SqlContainerSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecConflictResolutionPolicy{}).Type1()): SqlContainerSpecConflictResolutionPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecIndexingPolicy{}).Type1()):           SqlContainerSpecIndexingPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlDatabaseSpecAutoscaleSettings{}).Type1()):         SqlDatabaseSpecAutoscaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecAutoscaleSettings{}).Type1()):               TableSpecAutoscaleSettingsCodec{},
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
type AccountSpecBackupCodec struct {
}

func (AccountSpecBackupCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccountSpecBackup)(ptr) == nil
}

func (AccountSpecBackupCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccountSpecBackup)(ptr)
	var objs []AccountSpecBackup
	if obj != nil {
		objs = []AccountSpecBackup{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecBackup{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccountSpecBackupCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccountSpecBackup)(ptr) = AccountSpecBackup{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccountSpecBackup

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecBackup{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccountSpecBackup)(ptr) = objs[0]
			} else {
				*(*AccountSpecBackup)(ptr) = AccountSpecBackup{}
			}
		} else {
			*(*AccountSpecBackup)(ptr) = AccountSpecBackup{}
		}
	default:
		iter.ReportError("decode AccountSpecBackup", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AccountSpecConsistencyPolicyCodec struct {
}

func (AccountSpecConsistencyPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccountSpecConsistencyPolicy)(ptr) == nil
}

func (AccountSpecConsistencyPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccountSpecConsistencyPolicy)(ptr)
	var objs []AccountSpecConsistencyPolicy
	if obj != nil {
		objs = []AccountSpecConsistencyPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecConsistencyPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccountSpecConsistencyPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccountSpecConsistencyPolicy)(ptr) = AccountSpecConsistencyPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccountSpecConsistencyPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecConsistencyPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccountSpecConsistencyPolicy)(ptr) = objs[0]
			} else {
				*(*AccountSpecConsistencyPolicy)(ptr) = AccountSpecConsistencyPolicy{}
			}
		} else {
			*(*AccountSpecConsistencyPolicy)(ptr) = AccountSpecConsistencyPolicy{}
		}
	default:
		iter.ReportError("decode AccountSpecConsistencyPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AccountSpecCorsRuleCodec struct {
}

func (AccountSpecCorsRuleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccountSpecCorsRule)(ptr) == nil
}

func (AccountSpecCorsRuleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccountSpecCorsRule)(ptr)
	var objs []AccountSpecCorsRule
	if obj != nil {
		objs = []AccountSpecCorsRule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecCorsRule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccountSpecCorsRuleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccountSpecCorsRule)(ptr) = AccountSpecCorsRule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccountSpecCorsRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecCorsRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccountSpecCorsRule)(ptr) = objs[0]
			} else {
				*(*AccountSpecCorsRule)(ptr) = AccountSpecCorsRule{}
			}
		} else {
			*(*AccountSpecCorsRule)(ptr) = AccountSpecCorsRule{}
		}
	default:
		iter.ReportError("decode AccountSpecCorsRule", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AccountSpecIdentityCodec struct {
}

func (AccountSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccountSpecIdentity)(ptr) == nil
}

func (AccountSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccountSpecIdentity)(ptr)
	var objs []AccountSpecIdentity
	if obj != nil {
		objs = []AccountSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccountSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccountSpecIdentity)(ptr) = AccountSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccountSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccountSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccountSpecIdentity)(ptr) = objs[0]
			} else {
				*(*AccountSpecIdentity)(ptr) = AccountSpecIdentity{}
			}
		} else {
			*(*AccountSpecIdentity)(ptr) = AccountSpecIdentity{}
		}
	default:
		iter.ReportError("decode AccountSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CassandraKeyspaceSpecAutoscaleSettingsCodec struct {
}

func (CassandraKeyspaceSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CassandraKeyspaceSpecAutoscaleSettings)(ptr) == nil
}

func (CassandraKeyspaceSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CassandraKeyspaceSpecAutoscaleSettings)(ptr)
	var objs []CassandraKeyspaceSpecAutoscaleSettings
	if obj != nil {
		objs = []CassandraKeyspaceSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CassandraKeyspaceSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CassandraKeyspaceSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CassandraKeyspaceSpecAutoscaleSettings)(ptr) = CassandraKeyspaceSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CassandraKeyspaceSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CassandraKeyspaceSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CassandraKeyspaceSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*CassandraKeyspaceSpecAutoscaleSettings)(ptr) = CassandraKeyspaceSpecAutoscaleSettings{}
			}
		} else {
			*(*CassandraKeyspaceSpecAutoscaleSettings)(ptr) = CassandraKeyspaceSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode CassandraKeyspaceSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CassandraTableSpecAutoscaleSettingsCodec struct {
}

func (CassandraTableSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CassandraTableSpecAutoscaleSettings)(ptr) == nil
}

func (CassandraTableSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CassandraTableSpecAutoscaleSettings)(ptr)
	var objs []CassandraTableSpecAutoscaleSettings
	if obj != nil {
		objs = []CassandraTableSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CassandraTableSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CassandraTableSpecAutoscaleSettings)(ptr) = CassandraTableSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CassandraTableSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CassandraTableSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*CassandraTableSpecAutoscaleSettings)(ptr) = CassandraTableSpecAutoscaleSettings{}
			}
		} else {
			*(*CassandraTableSpecAutoscaleSettings)(ptr) = CassandraTableSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode CassandraTableSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CassandraTableSpecSchemaCodec struct {
}

func (CassandraTableSpecSchemaCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CassandraTableSpecSchema)(ptr) == nil
}

func (CassandraTableSpecSchemaCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CassandraTableSpecSchema)(ptr)
	var objs []CassandraTableSpecSchema
	if obj != nil {
		objs = []CassandraTableSpecSchema{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecSchema{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CassandraTableSpecSchemaCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CassandraTableSpecSchema)(ptr) = CassandraTableSpecSchema{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CassandraTableSpecSchema

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CassandraTableSpecSchema{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CassandraTableSpecSchema)(ptr) = objs[0]
			} else {
				*(*CassandraTableSpecSchema)(ptr) = CassandraTableSpecSchema{}
			}
		} else {
			*(*CassandraTableSpecSchema)(ptr) = CassandraTableSpecSchema{}
		}
	default:
		iter.ReportError("decode CassandraTableSpecSchema", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GremlinDatabaseSpecAutoscaleSettingsCodec struct {
}

func (GremlinDatabaseSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GremlinDatabaseSpecAutoscaleSettings)(ptr) == nil
}

func (GremlinDatabaseSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GremlinDatabaseSpecAutoscaleSettings)(ptr)
	var objs []GremlinDatabaseSpecAutoscaleSettings
	if obj != nil {
		objs = []GremlinDatabaseSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GremlinDatabaseSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GremlinDatabaseSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GremlinDatabaseSpecAutoscaleSettings)(ptr) = GremlinDatabaseSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GremlinDatabaseSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GremlinDatabaseSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GremlinDatabaseSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*GremlinDatabaseSpecAutoscaleSettings)(ptr) = GremlinDatabaseSpecAutoscaleSettings{}
			}
		} else {
			*(*GremlinDatabaseSpecAutoscaleSettings)(ptr) = GremlinDatabaseSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode GremlinDatabaseSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GremlinGraphSpecAutoscaleSettingsCodec struct {
}

func (GremlinGraphSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GremlinGraphSpecAutoscaleSettings)(ptr) == nil
}

func (GremlinGraphSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GremlinGraphSpecAutoscaleSettings)(ptr)
	var objs []GremlinGraphSpecAutoscaleSettings
	if obj != nil {
		objs = []GremlinGraphSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GremlinGraphSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GremlinGraphSpecAutoscaleSettings)(ptr) = GremlinGraphSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GremlinGraphSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GremlinGraphSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*GremlinGraphSpecAutoscaleSettings)(ptr) = GremlinGraphSpecAutoscaleSettings{}
			}
		} else {
			*(*GremlinGraphSpecAutoscaleSettings)(ptr) = GremlinGraphSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode GremlinGraphSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GremlinGraphSpecConflictResolutionPolicyCodec struct {
}

func (GremlinGraphSpecConflictResolutionPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GremlinGraphSpecConflictResolutionPolicy)(ptr) == nil
}

func (GremlinGraphSpecConflictResolutionPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GremlinGraphSpecConflictResolutionPolicy)(ptr)
	var objs []GremlinGraphSpecConflictResolutionPolicy
	if obj != nil {
		objs = []GremlinGraphSpecConflictResolutionPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecConflictResolutionPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GremlinGraphSpecConflictResolutionPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GremlinGraphSpecConflictResolutionPolicy)(ptr) = GremlinGraphSpecConflictResolutionPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GremlinGraphSpecConflictResolutionPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GremlinGraphSpecConflictResolutionPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GremlinGraphSpecConflictResolutionPolicy)(ptr) = objs[0]
			} else {
				*(*GremlinGraphSpecConflictResolutionPolicy)(ptr) = GremlinGraphSpecConflictResolutionPolicy{}
			}
		} else {
			*(*GremlinGraphSpecConflictResolutionPolicy)(ptr) = GremlinGraphSpecConflictResolutionPolicy{}
		}
	default:
		iter.ReportError("decode GremlinGraphSpecConflictResolutionPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MongoCollectionSpecAutoscaleSettingsCodec struct {
}

func (MongoCollectionSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MongoCollectionSpecAutoscaleSettings)(ptr) == nil
}

func (MongoCollectionSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MongoCollectionSpecAutoscaleSettings)(ptr)
	var objs []MongoCollectionSpecAutoscaleSettings
	if obj != nil {
		objs = []MongoCollectionSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MongoCollectionSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MongoCollectionSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MongoCollectionSpecAutoscaleSettings)(ptr) = MongoCollectionSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MongoCollectionSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MongoCollectionSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MongoCollectionSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*MongoCollectionSpecAutoscaleSettings)(ptr) = MongoCollectionSpecAutoscaleSettings{}
			}
		} else {
			*(*MongoCollectionSpecAutoscaleSettings)(ptr) = MongoCollectionSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode MongoCollectionSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MongoDatabaseSpecAutoscaleSettingsCodec struct {
}

func (MongoDatabaseSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MongoDatabaseSpecAutoscaleSettings)(ptr) == nil
}

func (MongoDatabaseSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MongoDatabaseSpecAutoscaleSettings)(ptr)
	var objs []MongoDatabaseSpecAutoscaleSettings
	if obj != nil {
		objs = []MongoDatabaseSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MongoDatabaseSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MongoDatabaseSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MongoDatabaseSpecAutoscaleSettings)(ptr) = MongoDatabaseSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MongoDatabaseSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MongoDatabaseSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MongoDatabaseSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*MongoDatabaseSpecAutoscaleSettings)(ptr) = MongoDatabaseSpecAutoscaleSettings{}
			}
		} else {
			*(*MongoDatabaseSpecAutoscaleSettings)(ptr) = MongoDatabaseSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode MongoDatabaseSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SqlContainerSpecAutoscaleSettingsCodec struct {
}

func (SqlContainerSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SqlContainerSpecAutoscaleSettings)(ptr) == nil
}

func (SqlContainerSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SqlContainerSpecAutoscaleSettings)(ptr)
	var objs []SqlContainerSpecAutoscaleSettings
	if obj != nil {
		objs = []SqlContainerSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SqlContainerSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SqlContainerSpecAutoscaleSettings)(ptr) = SqlContainerSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SqlContainerSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SqlContainerSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*SqlContainerSpecAutoscaleSettings)(ptr) = SqlContainerSpecAutoscaleSettings{}
			}
		} else {
			*(*SqlContainerSpecAutoscaleSettings)(ptr) = SqlContainerSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode SqlContainerSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SqlContainerSpecConflictResolutionPolicyCodec struct {
}

func (SqlContainerSpecConflictResolutionPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SqlContainerSpecConflictResolutionPolicy)(ptr) == nil
}

func (SqlContainerSpecConflictResolutionPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SqlContainerSpecConflictResolutionPolicy)(ptr)
	var objs []SqlContainerSpecConflictResolutionPolicy
	if obj != nil {
		objs = []SqlContainerSpecConflictResolutionPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecConflictResolutionPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SqlContainerSpecConflictResolutionPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SqlContainerSpecConflictResolutionPolicy)(ptr) = SqlContainerSpecConflictResolutionPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SqlContainerSpecConflictResolutionPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecConflictResolutionPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SqlContainerSpecConflictResolutionPolicy)(ptr) = objs[0]
			} else {
				*(*SqlContainerSpecConflictResolutionPolicy)(ptr) = SqlContainerSpecConflictResolutionPolicy{}
			}
		} else {
			*(*SqlContainerSpecConflictResolutionPolicy)(ptr) = SqlContainerSpecConflictResolutionPolicy{}
		}
	default:
		iter.ReportError("decode SqlContainerSpecConflictResolutionPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SqlContainerSpecIndexingPolicyCodec struct {
}

func (SqlContainerSpecIndexingPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SqlContainerSpecIndexingPolicy)(ptr) == nil
}

func (SqlContainerSpecIndexingPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SqlContainerSpecIndexingPolicy)(ptr)
	var objs []SqlContainerSpecIndexingPolicy
	if obj != nil {
		objs = []SqlContainerSpecIndexingPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecIndexingPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SqlContainerSpecIndexingPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SqlContainerSpecIndexingPolicy)(ptr) = SqlContainerSpecIndexingPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SqlContainerSpecIndexingPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlContainerSpecIndexingPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SqlContainerSpecIndexingPolicy)(ptr) = objs[0]
			} else {
				*(*SqlContainerSpecIndexingPolicy)(ptr) = SqlContainerSpecIndexingPolicy{}
			}
		} else {
			*(*SqlContainerSpecIndexingPolicy)(ptr) = SqlContainerSpecIndexingPolicy{}
		}
	default:
		iter.ReportError("decode SqlContainerSpecIndexingPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SqlDatabaseSpecAutoscaleSettingsCodec struct {
}

func (SqlDatabaseSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SqlDatabaseSpecAutoscaleSettings)(ptr) == nil
}

func (SqlDatabaseSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SqlDatabaseSpecAutoscaleSettings)(ptr)
	var objs []SqlDatabaseSpecAutoscaleSettings
	if obj != nil {
		objs = []SqlDatabaseSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlDatabaseSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SqlDatabaseSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SqlDatabaseSpecAutoscaleSettings)(ptr) = SqlDatabaseSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SqlDatabaseSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlDatabaseSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SqlDatabaseSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*SqlDatabaseSpecAutoscaleSettings)(ptr) = SqlDatabaseSpecAutoscaleSettings{}
			}
		} else {
			*(*SqlDatabaseSpecAutoscaleSettings)(ptr) = SqlDatabaseSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode SqlDatabaseSpecAutoscaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TableSpecAutoscaleSettingsCodec struct {
}

func (TableSpecAutoscaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TableSpecAutoscaleSettings)(ptr) == nil
}

func (TableSpecAutoscaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TableSpecAutoscaleSettings)(ptr)
	var objs []TableSpecAutoscaleSettings
	if obj != nil {
		objs = []TableSpecAutoscaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecAutoscaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TableSpecAutoscaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TableSpecAutoscaleSettings)(ptr) = TableSpecAutoscaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TableSpecAutoscaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecAutoscaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TableSpecAutoscaleSettings)(ptr) = objs[0]
			} else {
				*(*TableSpecAutoscaleSettings)(ptr) = TableSpecAutoscaleSettings{}
			}
		} else {
			*(*TableSpecAutoscaleSettings)(ptr) = TableSpecAutoscaleSettings{}
		}
	default:
		iter.ReportError("decode TableSpecAutoscaleSettings", "unexpected JSON type")
	}
}

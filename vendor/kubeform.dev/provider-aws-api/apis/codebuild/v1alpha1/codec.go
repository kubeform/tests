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
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecArtifacts{}).Type1()):                           ProjectSpecArtifactsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecCache{}).Type1()):                               ProjectSpecCacheCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironment{}).Type1()):                         ProjectSpecEnvironmentCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironmentRegistryCredential{}).Type1()):       ProjectSpecEnvironmentRegistryCredentialCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfig{}).Type1()):                          ProjectSpecLogsConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigCloudwatchLogs{}).Type1()):            ProjectSpecLogsConfigCloudwatchLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigS3Logs{}).Type1()):                    ProjectSpecLogsConfigS3LogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesAuth{}).Type1()):                ProjectSpecSecondarySourcesAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesGitSubmodulesConfig{}).Type1()): ProjectSpecSecondarySourcesGitSubmodulesConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSource{}).Type1()):                              ProjectSpecSourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceAuth{}).Type1()):                          ProjectSpecSourceAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceGitSubmodulesConfig{}).Type1()):           ProjectSpecSourceGitSubmodulesConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecVpcConfig{}).Type1()):                           ProjectSpecVpcConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfig{}).Type1()):                    ReportGroupSpecExportConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfigS3Destination{}).Type1()):       ReportGroupSpecExportConfigS3DestinationCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecArtifacts{}).Type1()):                           ProjectSpecArtifactsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecCache{}).Type1()):                               ProjectSpecCacheCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironment{}).Type1()):                         ProjectSpecEnvironmentCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironmentRegistryCredential{}).Type1()):       ProjectSpecEnvironmentRegistryCredentialCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfig{}).Type1()):                          ProjectSpecLogsConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigCloudwatchLogs{}).Type1()):            ProjectSpecLogsConfigCloudwatchLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigS3Logs{}).Type1()):                    ProjectSpecLogsConfigS3LogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesAuth{}).Type1()):                ProjectSpecSecondarySourcesAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesGitSubmodulesConfig{}).Type1()): ProjectSpecSecondarySourcesGitSubmodulesConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSource{}).Type1()):                              ProjectSpecSourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceAuth{}).Type1()):                          ProjectSpecSourceAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceGitSubmodulesConfig{}).Type1()):           ProjectSpecSourceGitSubmodulesConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecVpcConfig{}).Type1()):                           ProjectSpecVpcConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfig{}).Type1()):                    ReportGroupSpecExportConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfigS3Destination{}).Type1()):       ReportGroupSpecExportConfigS3DestinationCodec{},
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
type ProjectSpecArtifactsCodec struct {
}

func (ProjectSpecArtifactsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecArtifacts)(ptr) == nil
}

func (ProjectSpecArtifactsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecArtifacts)(ptr)
	var objs []ProjectSpecArtifacts
	if obj != nil {
		objs = []ProjectSpecArtifacts{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecArtifacts{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecArtifactsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecArtifacts)(ptr) = ProjectSpecArtifacts{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecArtifacts

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecArtifacts{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecArtifacts)(ptr) = objs[0]
			} else {
				*(*ProjectSpecArtifacts)(ptr) = ProjectSpecArtifacts{}
			}
		} else {
			*(*ProjectSpecArtifacts)(ptr) = ProjectSpecArtifacts{}
		}
	default:
		iter.ReportError("decode ProjectSpecArtifacts", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecCacheCodec struct {
}

func (ProjectSpecCacheCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecCache)(ptr) == nil
}

func (ProjectSpecCacheCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecCache)(ptr)
	var objs []ProjectSpecCache
	if obj != nil {
		objs = []ProjectSpecCache{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecCache{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecCacheCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecCache)(ptr) = ProjectSpecCache{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecCache

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecCache{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecCache)(ptr) = objs[0]
			} else {
				*(*ProjectSpecCache)(ptr) = ProjectSpecCache{}
			}
		} else {
			*(*ProjectSpecCache)(ptr) = ProjectSpecCache{}
		}
	default:
		iter.ReportError("decode ProjectSpecCache", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecEnvironmentCodec struct {
}

func (ProjectSpecEnvironmentCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecEnvironment)(ptr) == nil
}

func (ProjectSpecEnvironmentCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecEnvironment)(ptr)
	var objs []ProjectSpecEnvironment
	if obj != nil {
		objs = []ProjectSpecEnvironment{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironment{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecEnvironmentCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecEnvironment)(ptr) = ProjectSpecEnvironment{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecEnvironment

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironment{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecEnvironment)(ptr) = objs[0]
			} else {
				*(*ProjectSpecEnvironment)(ptr) = ProjectSpecEnvironment{}
			}
		} else {
			*(*ProjectSpecEnvironment)(ptr) = ProjectSpecEnvironment{}
		}
	default:
		iter.ReportError("decode ProjectSpecEnvironment", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecEnvironmentRegistryCredentialCodec struct {
}

func (ProjectSpecEnvironmentRegistryCredentialCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecEnvironmentRegistryCredential)(ptr) == nil
}

func (ProjectSpecEnvironmentRegistryCredentialCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecEnvironmentRegistryCredential)(ptr)
	var objs []ProjectSpecEnvironmentRegistryCredential
	if obj != nil {
		objs = []ProjectSpecEnvironmentRegistryCredential{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironmentRegistryCredential{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecEnvironmentRegistryCredentialCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecEnvironmentRegistryCredential)(ptr) = ProjectSpecEnvironmentRegistryCredential{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecEnvironmentRegistryCredential

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecEnvironmentRegistryCredential{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecEnvironmentRegistryCredential)(ptr) = objs[0]
			} else {
				*(*ProjectSpecEnvironmentRegistryCredential)(ptr) = ProjectSpecEnvironmentRegistryCredential{}
			}
		} else {
			*(*ProjectSpecEnvironmentRegistryCredential)(ptr) = ProjectSpecEnvironmentRegistryCredential{}
		}
	default:
		iter.ReportError("decode ProjectSpecEnvironmentRegistryCredential", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecLogsConfigCodec struct {
}

func (ProjectSpecLogsConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecLogsConfig)(ptr) == nil
}

func (ProjectSpecLogsConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecLogsConfig)(ptr)
	var objs []ProjectSpecLogsConfig
	if obj != nil {
		objs = []ProjectSpecLogsConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecLogsConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecLogsConfig)(ptr) = ProjectSpecLogsConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecLogsConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecLogsConfig)(ptr) = objs[0]
			} else {
				*(*ProjectSpecLogsConfig)(ptr) = ProjectSpecLogsConfig{}
			}
		} else {
			*(*ProjectSpecLogsConfig)(ptr) = ProjectSpecLogsConfig{}
		}
	default:
		iter.ReportError("decode ProjectSpecLogsConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecLogsConfigCloudwatchLogsCodec struct {
}

func (ProjectSpecLogsConfigCloudwatchLogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecLogsConfigCloudwatchLogs)(ptr) == nil
}

func (ProjectSpecLogsConfigCloudwatchLogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecLogsConfigCloudwatchLogs)(ptr)
	var objs []ProjectSpecLogsConfigCloudwatchLogs
	if obj != nil {
		objs = []ProjectSpecLogsConfigCloudwatchLogs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigCloudwatchLogs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecLogsConfigCloudwatchLogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecLogsConfigCloudwatchLogs)(ptr) = ProjectSpecLogsConfigCloudwatchLogs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecLogsConfigCloudwatchLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigCloudwatchLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecLogsConfigCloudwatchLogs)(ptr) = objs[0]
			} else {
				*(*ProjectSpecLogsConfigCloudwatchLogs)(ptr) = ProjectSpecLogsConfigCloudwatchLogs{}
			}
		} else {
			*(*ProjectSpecLogsConfigCloudwatchLogs)(ptr) = ProjectSpecLogsConfigCloudwatchLogs{}
		}
	default:
		iter.ReportError("decode ProjectSpecLogsConfigCloudwatchLogs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecLogsConfigS3LogsCodec struct {
}

func (ProjectSpecLogsConfigS3LogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecLogsConfigS3Logs)(ptr) == nil
}

func (ProjectSpecLogsConfigS3LogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecLogsConfigS3Logs)(ptr)
	var objs []ProjectSpecLogsConfigS3Logs
	if obj != nil {
		objs = []ProjectSpecLogsConfigS3Logs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigS3Logs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecLogsConfigS3LogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecLogsConfigS3Logs)(ptr) = ProjectSpecLogsConfigS3Logs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecLogsConfigS3Logs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecLogsConfigS3Logs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecLogsConfigS3Logs)(ptr) = objs[0]
			} else {
				*(*ProjectSpecLogsConfigS3Logs)(ptr) = ProjectSpecLogsConfigS3Logs{}
			}
		} else {
			*(*ProjectSpecLogsConfigS3Logs)(ptr) = ProjectSpecLogsConfigS3Logs{}
		}
	default:
		iter.ReportError("decode ProjectSpecLogsConfigS3Logs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecSecondarySourcesAuthCodec struct {
}

func (ProjectSpecSecondarySourcesAuthCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecSecondarySourcesAuth)(ptr) == nil
}

func (ProjectSpecSecondarySourcesAuthCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecSecondarySourcesAuth)(ptr)
	var objs []ProjectSpecSecondarySourcesAuth
	if obj != nil {
		objs = []ProjectSpecSecondarySourcesAuth{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesAuth{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecSecondarySourcesAuthCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecSecondarySourcesAuth)(ptr) = ProjectSpecSecondarySourcesAuth{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecSecondarySourcesAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecSecondarySourcesAuth)(ptr) = objs[0]
			} else {
				*(*ProjectSpecSecondarySourcesAuth)(ptr) = ProjectSpecSecondarySourcesAuth{}
			}
		} else {
			*(*ProjectSpecSecondarySourcesAuth)(ptr) = ProjectSpecSecondarySourcesAuth{}
		}
	default:
		iter.ReportError("decode ProjectSpecSecondarySourcesAuth", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecSecondarySourcesGitSubmodulesConfigCodec struct {
}

func (ProjectSpecSecondarySourcesGitSubmodulesConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecSecondarySourcesGitSubmodulesConfig)(ptr) == nil
}

func (ProjectSpecSecondarySourcesGitSubmodulesConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecSecondarySourcesGitSubmodulesConfig)(ptr)
	var objs []ProjectSpecSecondarySourcesGitSubmodulesConfig
	if obj != nil {
		objs = []ProjectSpecSecondarySourcesGitSubmodulesConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesGitSubmodulesConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecSecondarySourcesGitSubmodulesConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecSecondarySourcesGitSubmodulesConfig)(ptr) = ProjectSpecSecondarySourcesGitSubmodulesConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecSecondarySourcesGitSubmodulesConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSecondarySourcesGitSubmodulesConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecSecondarySourcesGitSubmodulesConfig)(ptr) = objs[0]
			} else {
				*(*ProjectSpecSecondarySourcesGitSubmodulesConfig)(ptr) = ProjectSpecSecondarySourcesGitSubmodulesConfig{}
			}
		} else {
			*(*ProjectSpecSecondarySourcesGitSubmodulesConfig)(ptr) = ProjectSpecSecondarySourcesGitSubmodulesConfig{}
		}
	default:
		iter.ReportError("decode ProjectSpecSecondarySourcesGitSubmodulesConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecSourceCodec struct {
}

func (ProjectSpecSourceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecSource)(ptr) == nil
}

func (ProjectSpecSourceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecSource)(ptr)
	var objs []ProjectSpecSource
	if obj != nil {
		objs = []ProjectSpecSource{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSource{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecSourceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecSource)(ptr) = ProjectSpecSource{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecSource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecSource)(ptr) = objs[0]
			} else {
				*(*ProjectSpecSource)(ptr) = ProjectSpecSource{}
			}
		} else {
			*(*ProjectSpecSource)(ptr) = ProjectSpecSource{}
		}
	default:
		iter.ReportError("decode ProjectSpecSource", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecSourceAuthCodec struct {
}

func (ProjectSpecSourceAuthCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecSourceAuth)(ptr) == nil
}

func (ProjectSpecSourceAuthCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecSourceAuth)(ptr)
	var objs []ProjectSpecSourceAuth
	if obj != nil {
		objs = []ProjectSpecSourceAuth{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceAuth{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecSourceAuthCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecSourceAuth)(ptr) = ProjectSpecSourceAuth{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecSourceAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecSourceAuth)(ptr) = objs[0]
			} else {
				*(*ProjectSpecSourceAuth)(ptr) = ProjectSpecSourceAuth{}
			}
		} else {
			*(*ProjectSpecSourceAuth)(ptr) = ProjectSpecSourceAuth{}
		}
	default:
		iter.ReportError("decode ProjectSpecSourceAuth", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecSourceGitSubmodulesConfigCodec struct {
}

func (ProjectSpecSourceGitSubmodulesConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecSourceGitSubmodulesConfig)(ptr) == nil
}

func (ProjectSpecSourceGitSubmodulesConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecSourceGitSubmodulesConfig)(ptr)
	var objs []ProjectSpecSourceGitSubmodulesConfig
	if obj != nil {
		objs = []ProjectSpecSourceGitSubmodulesConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceGitSubmodulesConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecSourceGitSubmodulesConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecSourceGitSubmodulesConfig)(ptr) = ProjectSpecSourceGitSubmodulesConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecSourceGitSubmodulesConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecSourceGitSubmodulesConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecSourceGitSubmodulesConfig)(ptr) = objs[0]
			} else {
				*(*ProjectSpecSourceGitSubmodulesConfig)(ptr) = ProjectSpecSourceGitSubmodulesConfig{}
			}
		} else {
			*(*ProjectSpecSourceGitSubmodulesConfig)(ptr) = ProjectSpecSourceGitSubmodulesConfig{}
		}
	default:
		iter.ReportError("decode ProjectSpecSourceGitSubmodulesConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSpecVpcConfigCodec struct {
}

func (ProjectSpecVpcConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSpecVpcConfig)(ptr) == nil
}

func (ProjectSpecVpcConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSpecVpcConfig)(ptr)
	var objs []ProjectSpecVpcConfig
	if obj != nil {
		objs = []ProjectSpecVpcConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecVpcConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSpecVpcConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSpecVpcConfig)(ptr) = ProjectSpecVpcConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSpecVpcConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSpecVpcConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSpecVpcConfig)(ptr) = objs[0]
			} else {
				*(*ProjectSpecVpcConfig)(ptr) = ProjectSpecVpcConfig{}
			}
		} else {
			*(*ProjectSpecVpcConfig)(ptr) = ProjectSpecVpcConfig{}
		}
	default:
		iter.ReportError("decode ProjectSpecVpcConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ReportGroupSpecExportConfigCodec struct {
}

func (ReportGroupSpecExportConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ReportGroupSpecExportConfig)(ptr) == nil
}

func (ReportGroupSpecExportConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ReportGroupSpecExportConfig)(ptr)
	var objs []ReportGroupSpecExportConfig
	if obj != nil {
		objs = []ReportGroupSpecExportConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ReportGroupSpecExportConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ReportGroupSpecExportConfig)(ptr) = ReportGroupSpecExportConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ReportGroupSpecExportConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ReportGroupSpecExportConfig)(ptr) = objs[0]
			} else {
				*(*ReportGroupSpecExportConfig)(ptr) = ReportGroupSpecExportConfig{}
			}
		} else {
			*(*ReportGroupSpecExportConfig)(ptr) = ReportGroupSpecExportConfig{}
		}
	default:
		iter.ReportError("decode ReportGroupSpecExportConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ReportGroupSpecExportConfigS3DestinationCodec struct {
}

func (ReportGroupSpecExportConfigS3DestinationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ReportGroupSpecExportConfigS3Destination)(ptr) == nil
}

func (ReportGroupSpecExportConfigS3DestinationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ReportGroupSpecExportConfigS3Destination)(ptr)
	var objs []ReportGroupSpecExportConfigS3Destination
	if obj != nil {
		objs = []ReportGroupSpecExportConfigS3Destination{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfigS3Destination{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ReportGroupSpecExportConfigS3DestinationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ReportGroupSpecExportConfigS3Destination)(ptr) = ReportGroupSpecExportConfigS3Destination{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ReportGroupSpecExportConfigS3Destination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ReportGroupSpecExportConfigS3Destination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ReportGroupSpecExportConfigS3Destination)(ptr) = objs[0]
			} else {
				*(*ReportGroupSpecExportConfigS3Destination)(ptr) = ReportGroupSpecExportConfigS3Destination{}
			}
		} else {
			*(*ReportGroupSpecExportConfigS3Destination)(ptr) = ReportGroupSpecExportConfigS3Destination{}
		}
	default:
		iter.ReportError("decode ReportGroupSpecExportConfigS3Destination", "unexpected JSON type")
	}
}

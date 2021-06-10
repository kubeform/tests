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
		jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProvider{}).Type1()):                      CapacityProviderSpecAutoScalingGroupProviderCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProviderManagedScaling{}).Type1()):        CapacityProviderSpecAutoScalingGroupProviderManagedScalingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentCircuitBreaker{}).Type1()):                               ServiceSpecDeploymentCircuitBreakerCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentController{}).Type1()):                                   ServiceSpecDeploymentControllerCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecNetworkConfiguration{}).Type1()):                                   ServiceSpecNetworkConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecServiceRegistries{}).Type1()):                                      ServiceSpecServiceRegistriesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecProxyConfiguration{}).Type1()):                              TaskDefinitionSpecProxyConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeDockerVolumeConfiguration{}).Type1()):                 TaskDefinitionSpecVolumeDockerVolumeConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfiguration{}).Type1()):                    TaskDefinitionSpecVolumeEfsVolumeConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{}).Type1()): TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProvider{}).Type1()):                      CapacityProviderSpecAutoScalingGroupProviderCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProviderManagedScaling{}).Type1()):        CapacityProviderSpecAutoScalingGroupProviderManagedScalingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentCircuitBreaker{}).Type1()):                               ServiceSpecDeploymentCircuitBreakerCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentController{}).Type1()):                                   ServiceSpecDeploymentControllerCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecNetworkConfiguration{}).Type1()):                                   ServiceSpecNetworkConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecServiceRegistries{}).Type1()):                                      ServiceSpecServiceRegistriesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecProxyConfiguration{}).Type1()):                              TaskDefinitionSpecProxyConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeDockerVolumeConfiguration{}).Type1()):                 TaskDefinitionSpecVolumeDockerVolumeConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfiguration{}).Type1()):                    TaskDefinitionSpecVolumeEfsVolumeConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{}).Type1()): TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfigCodec{},
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
type CapacityProviderSpecAutoScalingGroupProviderCodec struct {
}

func (CapacityProviderSpecAutoScalingGroupProviderCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CapacityProviderSpecAutoScalingGroupProvider)(ptr) == nil
}

func (CapacityProviderSpecAutoScalingGroupProviderCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CapacityProviderSpecAutoScalingGroupProvider)(ptr)
	var objs []CapacityProviderSpecAutoScalingGroupProvider
	if obj != nil {
		objs = []CapacityProviderSpecAutoScalingGroupProvider{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProvider{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CapacityProviderSpecAutoScalingGroupProviderCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CapacityProviderSpecAutoScalingGroupProvider)(ptr) = CapacityProviderSpecAutoScalingGroupProvider{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CapacityProviderSpecAutoScalingGroupProvider

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProvider{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CapacityProviderSpecAutoScalingGroupProvider)(ptr) = objs[0]
			} else {
				*(*CapacityProviderSpecAutoScalingGroupProvider)(ptr) = CapacityProviderSpecAutoScalingGroupProvider{}
			}
		} else {
			*(*CapacityProviderSpecAutoScalingGroupProvider)(ptr) = CapacityProviderSpecAutoScalingGroupProvider{}
		}
	default:
		iter.ReportError("decode CapacityProviderSpecAutoScalingGroupProvider", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CapacityProviderSpecAutoScalingGroupProviderManagedScalingCodec struct {
}

func (CapacityProviderSpecAutoScalingGroupProviderManagedScalingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CapacityProviderSpecAutoScalingGroupProviderManagedScaling)(ptr) == nil
}

func (CapacityProviderSpecAutoScalingGroupProviderManagedScalingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CapacityProviderSpecAutoScalingGroupProviderManagedScaling)(ptr)
	var objs []CapacityProviderSpecAutoScalingGroupProviderManagedScaling
	if obj != nil {
		objs = []CapacityProviderSpecAutoScalingGroupProviderManagedScaling{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProviderManagedScaling{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CapacityProviderSpecAutoScalingGroupProviderManagedScalingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CapacityProviderSpecAutoScalingGroupProviderManagedScaling)(ptr) = CapacityProviderSpecAutoScalingGroupProviderManagedScaling{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CapacityProviderSpecAutoScalingGroupProviderManagedScaling

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CapacityProviderSpecAutoScalingGroupProviderManagedScaling{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CapacityProviderSpecAutoScalingGroupProviderManagedScaling)(ptr) = objs[0]
			} else {
				*(*CapacityProviderSpecAutoScalingGroupProviderManagedScaling)(ptr) = CapacityProviderSpecAutoScalingGroupProviderManagedScaling{}
			}
		} else {
			*(*CapacityProviderSpecAutoScalingGroupProviderManagedScaling)(ptr) = CapacityProviderSpecAutoScalingGroupProviderManagedScaling{}
		}
	default:
		iter.ReportError("decode CapacityProviderSpecAutoScalingGroupProviderManagedScaling", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecDeploymentCircuitBreakerCodec struct {
}

func (ServiceSpecDeploymentCircuitBreakerCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecDeploymentCircuitBreaker)(ptr) == nil
}

func (ServiceSpecDeploymentCircuitBreakerCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecDeploymentCircuitBreaker)(ptr)
	var objs []ServiceSpecDeploymentCircuitBreaker
	if obj != nil {
		objs = []ServiceSpecDeploymentCircuitBreaker{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentCircuitBreaker{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecDeploymentCircuitBreakerCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecDeploymentCircuitBreaker)(ptr) = ServiceSpecDeploymentCircuitBreaker{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecDeploymentCircuitBreaker

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentCircuitBreaker{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecDeploymentCircuitBreaker)(ptr) = objs[0]
			} else {
				*(*ServiceSpecDeploymentCircuitBreaker)(ptr) = ServiceSpecDeploymentCircuitBreaker{}
			}
		} else {
			*(*ServiceSpecDeploymentCircuitBreaker)(ptr) = ServiceSpecDeploymentCircuitBreaker{}
		}
	default:
		iter.ReportError("decode ServiceSpecDeploymentCircuitBreaker", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecDeploymentControllerCodec struct {
}

func (ServiceSpecDeploymentControllerCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecDeploymentController)(ptr) == nil
}

func (ServiceSpecDeploymentControllerCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecDeploymentController)(ptr)
	var objs []ServiceSpecDeploymentController
	if obj != nil {
		objs = []ServiceSpecDeploymentController{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentController{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecDeploymentControllerCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecDeploymentController)(ptr) = ServiceSpecDeploymentController{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecDeploymentController

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecDeploymentController{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecDeploymentController)(ptr) = objs[0]
			} else {
				*(*ServiceSpecDeploymentController)(ptr) = ServiceSpecDeploymentController{}
			}
		} else {
			*(*ServiceSpecDeploymentController)(ptr) = ServiceSpecDeploymentController{}
		}
	default:
		iter.ReportError("decode ServiceSpecDeploymentController", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecNetworkConfigurationCodec struct {
}

func (ServiceSpecNetworkConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecNetworkConfiguration)(ptr) == nil
}

func (ServiceSpecNetworkConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecNetworkConfiguration)(ptr)
	var objs []ServiceSpecNetworkConfiguration
	if obj != nil {
		objs = []ServiceSpecNetworkConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecNetworkConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecNetworkConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecNetworkConfiguration)(ptr) = ServiceSpecNetworkConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecNetworkConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecNetworkConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecNetworkConfiguration)(ptr) = objs[0]
			} else {
				*(*ServiceSpecNetworkConfiguration)(ptr) = ServiceSpecNetworkConfiguration{}
			}
		} else {
			*(*ServiceSpecNetworkConfiguration)(ptr) = ServiceSpecNetworkConfiguration{}
		}
	default:
		iter.ReportError("decode ServiceSpecNetworkConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecServiceRegistriesCodec struct {
}

func (ServiceSpecServiceRegistriesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecServiceRegistries)(ptr) == nil
}

func (ServiceSpecServiceRegistriesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecServiceRegistries)(ptr)
	var objs []ServiceSpecServiceRegistries
	if obj != nil {
		objs = []ServiceSpecServiceRegistries{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecServiceRegistries{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecServiceRegistriesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecServiceRegistries)(ptr) = ServiceSpecServiceRegistries{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecServiceRegistries

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecServiceRegistries{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecServiceRegistries)(ptr) = objs[0]
			} else {
				*(*ServiceSpecServiceRegistries)(ptr) = ServiceSpecServiceRegistries{}
			}
		} else {
			*(*ServiceSpecServiceRegistries)(ptr) = ServiceSpecServiceRegistries{}
		}
	default:
		iter.ReportError("decode ServiceSpecServiceRegistries", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TaskDefinitionSpecProxyConfigurationCodec struct {
}

func (TaskDefinitionSpecProxyConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TaskDefinitionSpecProxyConfiguration)(ptr) == nil
}

func (TaskDefinitionSpecProxyConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TaskDefinitionSpecProxyConfiguration)(ptr)
	var objs []TaskDefinitionSpecProxyConfiguration
	if obj != nil {
		objs = []TaskDefinitionSpecProxyConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecProxyConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TaskDefinitionSpecProxyConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TaskDefinitionSpecProxyConfiguration)(ptr) = TaskDefinitionSpecProxyConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TaskDefinitionSpecProxyConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecProxyConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TaskDefinitionSpecProxyConfiguration)(ptr) = objs[0]
			} else {
				*(*TaskDefinitionSpecProxyConfiguration)(ptr) = TaskDefinitionSpecProxyConfiguration{}
			}
		} else {
			*(*TaskDefinitionSpecProxyConfiguration)(ptr) = TaskDefinitionSpecProxyConfiguration{}
		}
	default:
		iter.ReportError("decode TaskDefinitionSpecProxyConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TaskDefinitionSpecVolumeDockerVolumeConfigurationCodec struct {
}

func (TaskDefinitionSpecVolumeDockerVolumeConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TaskDefinitionSpecVolumeDockerVolumeConfiguration)(ptr) == nil
}

func (TaskDefinitionSpecVolumeDockerVolumeConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TaskDefinitionSpecVolumeDockerVolumeConfiguration)(ptr)
	var objs []TaskDefinitionSpecVolumeDockerVolumeConfiguration
	if obj != nil {
		objs = []TaskDefinitionSpecVolumeDockerVolumeConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeDockerVolumeConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TaskDefinitionSpecVolumeDockerVolumeConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TaskDefinitionSpecVolumeDockerVolumeConfiguration)(ptr) = TaskDefinitionSpecVolumeDockerVolumeConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TaskDefinitionSpecVolumeDockerVolumeConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeDockerVolumeConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TaskDefinitionSpecVolumeDockerVolumeConfiguration)(ptr) = objs[0]
			} else {
				*(*TaskDefinitionSpecVolumeDockerVolumeConfiguration)(ptr) = TaskDefinitionSpecVolumeDockerVolumeConfiguration{}
			}
		} else {
			*(*TaskDefinitionSpecVolumeDockerVolumeConfiguration)(ptr) = TaskDefinitionSpecVolumeDockerVolumeConfiguration{}
		}
	default:
		iter.ReportError("decode TaskDefinitionSpecVolumeDockerVolumeConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TaskDefinitionSpecVolumeEfsVolumeConfigurationCodec struct {
}

func (TaskDefinitionSpecVolumeEfsVolumeConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TaskDefinitionSpecVolumeEfsVolumeConfiguration)(ptr) == nil
}

func (TaskDefinitionSpecVolumeEfsVolumeConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TaskDefinitionSpecVolumeEfsVolumeConfiguration)(ptr)
	var objs []TaskDefinitionSpecVolumeEfsVolumeConfiguration
	if obj != nil {
		objs = []TaskDefinitionSpecVolumeEfsVolumeConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TaskDefinitionSpecVolumeEfsVolumeConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TaskDefinitionSpecVolumeEfsVolumeConfiguration)(ptr) = TaskDefinitionSpecVolumeEfsVolumeConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TaskDefinitionSpecVolumeEfsVolumeConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TaskDefinitionSpecVolumeEfsVolumeConfiguration)(ptr) = objs[0]
			} else {
				*(*TaskDefinitionSpecVolumeEfsVolumeConfiguration)(ptr) = TaskDefinitionSpecVolumeEfsVolumeConfiguration{}
			}
		} else {
			*(*TaskDefinitionSpecVolumeEfsVolumeConfiguration)(ptr) = TaskDefinitionSpecVolumeEfsVolumeConfiguration{}
		}
	default:
		iter.ReportError("decode TaskDefinitionSpecVolumeEfsVolumeConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfigCodec struct {
}

func (TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig)(ptr) == nil
}

func (TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig)(ptr)
	var objs []TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig
	if obj != nil {
		objs = []TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig)(ptr) = TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig)(ptr) = objs[0]
			} else {
				*(*TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig)(ptr) = TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{}
			}
		} else {
			*(*TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig)(ptr) = TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig{}
		}
	default:
		iter.ReportError("decode TaskDefinitionSpecVolumeEfsVolumeConfigurationAuthorizationConfig", "unexpected JSON type")
	}
}

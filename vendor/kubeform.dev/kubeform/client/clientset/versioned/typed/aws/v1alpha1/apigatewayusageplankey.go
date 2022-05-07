/*
Copyright The Kubeform Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApiGatewayUsagePlanKeysGetter has a method to return a ApiGatewayUsagePlanKeyInterface.
// A group's client should implement this interface.
type ApiGatewayUsagePlanKeysGetter interface {
	ApiGatewayUsagePlanKeys(namespace string) ApiGatewayUsagePlanKeyInterface
}

// ApiGatewayUsagePlanKeyInterface has methods to work with ApiGatewayUsagePlanKey resources.
type ApiGatewayUsagePlanKeyInterface interface {
	Create(ctx context.Context, apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey, opts v1.CreateOptions) (*v1alpha1.ApiGatewayUsagePlanKey, error)
	Update(ctx context.Context, apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayUsagePlanKey, error)
	UpdateStatus(ctx context.Context, apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayUsagePlanKey, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApiGatewayUsagePlanKey, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApiGatewayUsagePlanKeyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayUsagePlanKey, err error)
	ApiGatewayUsagePlanKeyExpansion
}

// apiGatewayUsagePlanKeys implements ApiGatewayUsagePlanKeyInterface
type apiGatewayUsagePlanKeys struct {
	client rest.Interface
	ns     string
}

// newApiGatewayUsagePlanKeys returns a ApiGatewayUsagePlanKeys
func newApiGatewayUsagePlanKeys(c *AwsV1alpha1Client, namespace string) *apiGatewayUsagePlanKeys {
	return &apiGatewayUsagePlanKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiGatewayUsagePlanKey, and returns the corresponding apiGatewayUsagePlanKey object, and an error if there is any.
func (c *apiGatewayUsagePlanKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.ApiGatewayUsagePlanKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiGatewayUsagePlanKeys that match those selectors.
func (c *apiGatewayUsagePlanKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiGatewayUsagePlanKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiGatewayUsagePlanKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiGatewayUsagePlanKeys.
func (c *apiGatewayUsagePlanKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apiGatewayUsagePlanKey and creates it.  Returns the server's representation of the apiGatewayUsagePlanKey, and an error, if there is any.
func (c *apiGatewayUsagePlanKeys) Create(ctx context.Context, apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey, opts v1.CreateOptions) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.ApiGatewayUsagePlanKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiGatewayUsagePlanKey).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apiGatewayUsagePlanKey and updates it. Returns the server's representation of the apiGatewayUsagePlanKey, and an error, if there is any.
func (c *apiGatewayUsagePlanKeys) Update(ctx context.Context, apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.ApiGatewayUsagePlanKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		Name(apiGatewayUsagePlanKey.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiGatewayUsagePlanKey).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *apiGatewayUsagePlanKeys) UpdateStatus(ctx context.Context, apiGatewayUsagePlanKey *v1alpha1.ApiGatewayUsagePlanKey, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.ApiGatewayUsagePlanKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		Name(apiGatewayUsagePlanKey.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiGatewayUsagePlanKey).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apiGatewayUsagePlanKey and deletes it. Returns an error if one occurs.
func (c *apiGatewayUsagePlanKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiGatewayUsagePlanKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apiGatewayUsagePlanKey.
func (c *apiGatewayUsagePlanKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayUsagePlanKey, err error) {
	result = &v1alpha1.ApiGatewayUsagePlanKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apigatewayusageplankeys").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
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

// ApiGatewayUsagePlansGetter has a method to return a ApiGatewayUsagePlanInterface.
// A group's client should implement this interface.
type ApiGatewayUsagePlansGetter interface {
	ApiGatewayUsagePlans(namespace string) ApiGatewayUsagePlanInterface
}

// ApiGatewayUsagePlanInterface has methods to work with ApiGatewayUsagePlan resources.
type ApiGatewayUsagePlanInterface interface {
	Create(ctx context.Context, apiGatewayUsagePlan *v1alpha1.ApiGatewayUsagePlan, opts v1.CreateOptions) (*v1alpha1.ApiGatewayUsagePlan, error)
	Update(ctx context.Context, apiGatewayUsagePlan *v1alpha1.ApiGatewayUsagePlan, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayUsagePlan, error)
	UpdateStatus(ctx context.Context, apiGatewayUsagePlan *v1alpha1.ApiGatewayUsagePlan, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayUsagePlan, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApiGatewayUsagePlan, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApiGatewayUsagePlanList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayUsagePlan, err error)
	ApiGatewayUsagePlanExpansion
}

// apiGatewayUsagePlans implements ApiGatewayUsagePlanInterface
type apiGatewayUsagePlans struct {
	client rest.Interface
	ns     string
}

// newApiGatewayUsagePlans returns a ApiGatewayUsagePlans
func newApiGatewayUsagePlans(c *AwsV1alpha1Client, namespace string) *apiGatewayUsagePlans {
	return &apiGatewayUsagePlans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiGatewayUsagePlan, and returns the corresponding apiGatewayUsagePlan object, and an error if there is any.
func (c *apiGatewayUsagePlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayUsagePlan, err error) {
	result = &v1alpha1.ApiGatewayUsagePlan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiGatewayUsagePlans that match those selectors.
func (c *apiGatewayUsagePlans) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiGatewayUsagePlanList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiGatewayUsagePlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiGatewayUsagePlans.
func (c *apiGatewayUsagePlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apiGatewayUsagePlan and creates it.  Returns the server's representation of the apiGatewayUsagePlan, and an error, if there is any.
func (c *apiGatewayUsagePlans) Create(ctx context.Context, apiGatewayUsagePlan *v1alpha1.ApiGatewayUsagePlan, opts v1.CreateOptions) (result *v1alpha1.ApiGatewayUsagePlan, err error) {
	result = &v1alpha1.ApiGatewayUsagePlan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiGatewayUsagePlan).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apiGatewayUsagePlan and updates it. Returns the server's representation of the apiGatewayUsagePlan, and an error, if there is any.
func (c *apiGatewayUsagePlans) Update(ctx context.Context, apiGatewayUsagePlan *v1alpha1.ApiGatewayUsagePlan, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayUsagePlan, err error) {
	result = &v1alpha1.ApiGatewayUsagePlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		Name(apiGatewayUsagePlan.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiGatewayUsagePlan).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *apiGatewayUsagePlans) UpdateStatus(ctx context.Context, apiGatewayUsagePlan *v1alpha1.ApiGatewayUsagePlan, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayUsagePlan, err error) {
	result = &v1alpha1.ApiGatewayUsagePlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		Name(apiGatewayUsagePlan.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiGatewayUsagePlan).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apiGatewayUsagePlan and deletes it. Returns an error if one occurs.
func (c *apiGatewayUsagePlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiGatewayUsagePlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apiGatewayUsagePlan.
func (c *apiGatewayUsagePlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayUsagePlan, err error) {
	result = &v1alpha1.ApiGatewayUsagePlan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apigatewayusageplans").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

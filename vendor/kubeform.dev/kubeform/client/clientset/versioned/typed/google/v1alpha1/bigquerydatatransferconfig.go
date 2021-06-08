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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BigqueryDataTransferConfigsGetter has a method to return a BigqueryDataTransferConfigInterface.
// A group's client should implement this interface.
type BigqueryDataTransferConfigsGetter interface {
	BigqueryDataTransferConfigs(namespace string) BigqueryDataTransferConfigInterface
}

// BigqueryDataTransferConfigInterface has methods to work with BigqueryDataTransferConfig resources.
type BigqueryDataTransferConfigInterface interface {
	Create(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.CreateOptions) (*v1alpha1.BigqueryDataTransferConfig, error)
	Update(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.UpdateOptions) (*v1alpha1.BigqueryDataTransferConfig, error)
	UpdateStatus(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.UpdateOptions) (*v1alpha1.BigqueryDataTransferConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BigqueryDataTransferConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BigqueryDataTransferConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BigqueryDataTransferConfig, err error)
	BigqueryDataTransferConfigExpansion
}

// bigqueryDataTransferConfigs implements BigqueryDataTransferConfigInterface
type bigqueryDataTransferConfigs struct {
	client rest.Interface
	ns     string
}

// newBigqueryDataTransferConfigs returns a BigqueryDataTransferConfigs
func newBigqueryDataTransferConfigs(c *GoogleV1alpha1Client, namespace string) *bigqueryDataTransferConfigs {
	return &bigqueryDataTransferConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bigqueryDataTransferConfig, and returns the corresponding bigqueryDataTransferConfig object, and an error if there is any.
func (c *bigqueryDataTransferConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	result = &v1alpha1.BigqueryDataTransferConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BigqueryDataTransferConfigs that match those selectors.
func (c *bigqueryDataTransferConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BigqueryDataTransferConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BigqueryDataTransferConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bigqueryDataTransferConfigs.
func (c *bigqueryDataTransferConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bigqueryDataTransferConfig and creates it.  Returns the server's representation of the bigqueryDataTransferConfig, and an error, if there is any.
func (c *bigqueryDataTransferConfigs) Create(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.CreateOptions) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	result = &v1alpha1.BigqueryDataTransferConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigqueryDataTransferConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bigqueryDataTransferConfig and updates it. Returns the server's representation of the bigqueryDataTransferConfig, and an error, if there is any.
func (c *bigqueryDataTransferConfigs) Update(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.UpdateOptions) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	result = &v1alpha1.BigqueryDataTransferConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		Name(bigqueryDataTransferConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigqueryDataTransferConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bigqueryDataTransferConfigs) UpdateStatus(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.UpdateOptions) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	result = &v1alpha1.BigqueryDataTransferConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		Name(bigqueryDataTransferConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigqueryDataTransferConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bigqueryDataTransferConfig and deletes it. Returns an error if one occurs.
func (c *bigqueryDataTransferConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bigqueryDataTransferConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bigqueryDataTransferConfig.
func (c *bigqueryDataTransferConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	result = &v1alpha1.BigqueryDataTransferConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bigquerydatatransferconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

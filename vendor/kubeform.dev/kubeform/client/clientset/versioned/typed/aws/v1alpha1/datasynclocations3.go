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

// DatasyncLocationS3sGetter has a method to return a DatasyncLocationS3Interface.
// A group's client should implement this interface.
type DatasyncLocationS3sGetter interface {
	DatasyncLocationS3s(namespace string) DatasyncLocationS3Interface
}

// DatasyncLocationS3Interface has methods to work with DatasyncLocationS3 resources.
type DatasyncLocationS3Interface interface {
	Create(ctx context.Context, datasyncLocationS3 *v1alpha1.DatasyncLocationS3, opts v1.CreateOptions) (*v1alpha1.DatasyncLocationS3, error)
	Update(ctx context.Context, datasyncLocationS3 *v1alpha1.DatasyncLocationS3, opts v1.UpdateOptions) (*v1alpha1.DatasyncLocationS3, error)
	UpdateStatus(ctx context.Context, datasyncLocationS3 *v1alpha1.DatasyncLocationS3, opts v1.UpdateOptions) (*v1alpha1.DatasyncLocationS3, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DatasyncLocationS3, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DatasyncLocationS3List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatasyncLocationS3, err error)
	DatasyncLocationS3Expansion
}

// datasyncLocationS3s implements DatasyncLocationS3Interface
type datasyncLocationS3s struct {
	client rest.Interface
	ns     string
}

// newDatasyncLocationS3s returns a DatasyncLocationS3s
func newDatasyncLocationS3s(c *AwsV1alpha1Client, namespace string) *datasyncLocationS3s {
	return &datasyncLocationS3s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the datasyncLocationS3, and returns the corresponding datasyncLocationS3 object, and an error if there is any.
func (c *datasyncLocationS3s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatasyncLocationS3, err error) {
	result = &v1alpha1.DatasyncLocationS3{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DatasyncLocationS3s that match those selectors.
func (c *datasyncLocationS3s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatasyncLocationS3List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatasyncLocationS3List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested datasyncLocationS3s.
func (c *datasyncLocationS3s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a datasyncLocationS3 and creates it.  Returns the server's representation of the datasyncLocationS3, and an error, if there is any.
func (c *datasyncLocationS3s) Create(ctx context.Context, datasyncLocationS3 *v1alpha1.DatasyncLocationS3, opts v1.CreateOptions) (result *v1alpha1.DatasyncLocationS3, err error) {
	result = &v1alpha1.DatasyncLocationS3{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(datasyncLocationS3).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a datasyncLocationS3 and updates it. Returns the server's representation of the datasyncLocationS3, and an error, if there is any.
func (c *datasyncLocationS3s) Update(ctx context.Context, datasyncLocationS3 *v1alpha1.DatasyncLocationS3, opts v1.UpdateOptions) (result *v1alpha1.DatasyncLocationS3, err error) {
	result = &v1alpha1.DatasyncLocationS3{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		Name(datasyncLocationS3.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(datasyncLocationS3).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *datasyncLocationS3s) UpdateStatus(ctx context.Context, datasyncLocationS3 *v1alpha1.DatasyncLocationS3, opts v1.UpdateOptions) (result *v1alpha1.DatasyncLocationS3, err error) {
	result = &v1alpha1.DatasyncLocationS3{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		Name(datasyncLocationS3.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(datasyncLocationS3).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the datasyncLocationS3 and deletes it. Returns an error if one occurs.
func (c *datasyncLocationS3s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *datasyncLocationS3s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasynclocations3s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched datasyncLocationS3.
func (c *datasyncLocationS3s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatasyncLocationS3, err error) {
	result = &v1alpha1.DatasyncLocationS3{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datasynclocations3s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
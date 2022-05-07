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

// ComputeInstanceTemplatesGetter has a method to return a ComputeInstanceTemplateInterface.
// A group's client should implement this interface.
type ComputeInstanceTemplatesGetter interface {
	ComputeInstanceTemplates(namespace string) ComputeInstanceTemplateInterface
}

// ComputeInstanceTemplateInterface has methods to work with ComputeInstanceTemplate resources.
type ComputeInstanceTemplateInterface interface {
	Create(ctx context.Context, computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate, opts v1.CreateOptions) (*v1alpha1.ComputeInstanceTemplate, error)
	Update(ctx context.Context, computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate, opts v1.UpdateOptions) (*v1alpha1.ComputeInstanceTemplate, error)
	UpdateStatus(ctx context.Context, computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate, opts v1.UpdateOptions) (*v1alpha1.ComputeInstanceTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ComputeInstanceTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ComputeInstanceTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeInstanceTemplate, err error)
	ComputeInstanceTemplateExpansion
}

// computeInstanceTemplates implements ComputeInstanceTemplateInterface
type computeInstanceTemplates struct {
	client rest.Interface
	ns     string
}

// newComputeInstanceTemplates returns a ComputeInstanceTemplates
func newComputeInstanceTemplates(c *GoogleV1alpha1Client, namespace string) *computeInstanceTemplates {
	return &computeInstanceTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeInstanceTemplate, and returns the corresponding computeInstanceTemplate object, and an error if there is any.
func (c *computeInstanceTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	result = &v1alpha1.ComputeInstanceTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeInstanceTemplates that match those selectors.
func (c *computeInstanceTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeInstanceTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeInstanceTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeInstanceTemplates.
func (c *computeInstanceTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeInstanceTemplate and creates it.  Returns the server's representation of the computeInstanceTemplate, and an error, if there is any.
func (c *computeInstanceTemplates) Create(ctx context.Context, computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate, opts v1.CreateOptions) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	result = &v1alpha1.ComputeInstanceTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeInstanceTemplate and updates it. Returns the server's representation of the computeInstanceTemplate, and an error, if there is any.
func (c *computeInstanceTemplates) Update(ctx context.Context, computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate, opts v1.UpdateOptions) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	result = &v1alpha1.ComputeInstanceTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		Name(computeInstanceTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *computeInstanceTemplates) UpdateStatus(ctx context.Context, computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate, opts v1.UpdateOptions) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	result = &v1alpha1.ComputeInstanceTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		Name(computeInstanceTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeInstanceTemplate and deletes it. Returns an error if one occurs.
func (c *computeInstanceTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeInstanceTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeInstanceTemplate.
func (c *computeInstanceTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	result = &v1alpha1.ComputeInstanceTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeinstancetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
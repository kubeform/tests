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

// RedshiftParameterGroupsGetter has a method to return a RedshiftParameterGroupInterface.
// A group's client should implement this interface.
type RedshiftParameterGroupsGetter interface {
	RedshiftParameterGroups(namespace string) RedshiftParameterGroupInterface
}

// RedshiftParameterGroupInterface has methods to work with RedshiftParameterGroup resources.
type RedshiftParameterGroupInterface interface {
	Create(ctx context.Context, redshiftParameterGroup *v1alpha1.RedshiftParameterGroup, opts v1.CreateOptions) (*v1alpha1.RedshiftParameterGroup, error)
	Update(ctx context.Context, redshiftParameterGroup *v1alpha1.RedshiftParameterGroup, opts v1.UpdateOptions) (*v1alpha1.RedshiftParameterGroup, error)
	UpdateStatus(ctx context.Context, redshiftParameterGroup *v1alpha1.RedshiftParameterGroup, opts v1.UpdateOptions) (*v1alpha1.RedshiftParameterGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RedshiftParameterGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RedshiftParameterGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RedshiftParameterGroup, err error)
	RedshiftParameterGroupExpansion
}

// redshiftParameterGroups implements RedshiftParameterGroupInterface
type redshiftParameterGroups struct {
	client rest.Interface
	ns     string
}

// newRedshiftParameterGroups returns a RedshiftParameterGroups
func newRedshiftParameterGroups(c *AwsV1alpha1Client, namespace string) *redshiftParameterGroups {
	return &redshiftParameterGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redshiftParameterGroup, and returns the corresponding redshiftParameterGroup object, and an error if there is any.
func (c *redshiftParameterGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RedshiftParameterGroup, err error) {
	result = &v1alpha1.RedshiftParameterGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedshiftParameterGroups that match those selectors.
func (c *redshiftParameterGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RedshiftParameterGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RedshiftParameterGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redshiftParameterGroups.
func (c *redshiftParameterGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a redshiftParameterGroup and creates it.  Returns the server's representation of the redshiftParameterGroup, and an error, if there is any.
func (c *redshiftParameterGroups) Create(ctx context.Context, redshiftParameterGroup *v1alpha1.RedshiftParameterGroup, opts v1.CreateOptions) (result *v1alpha1.RedshiftParameterGroup, err error) {
	result = &v1alpha1.RedshiftParameterGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redshiftParameterGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a redshiftParameterGroup and updates it. Returns the server's representation of the redshiftParameterGroup, and an error, if there is any.
func (c *redshiftParameterGroups) Update(ctx context.Context, redshiftParameterGroup *v1alpha1.RedshiftParameterGroup, opts v1.UpdateOptions) (result *v1alpha1.RedshiftParameterGroup, err error) {
	result = &v1alpha1.RedshiftParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		Name(redshiftParameterGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redshiftParameterGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *redshiftParameterGroups) UpdateStatus(ctx context.Context, redshiftParameterGroup *v1alpha1.RedshiftParameterGroup, opts v1.UpdateOptions) (result *v1alpha1.RedshiftParameterGroup, err error) {
	result = &v1alpha1.RedshiftParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		Name(redshiftParameterGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redshiftParameterGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the redshiftParameterGroup and deletes it. Returns an error if one occurs.
func (c *redshiftParameterGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redshiftParameterGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched redshiftParameterGroup.
func (c *redshiftParameterGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RedshiftParameterGroup, err error) {
	result = &v1alpha1.RedshiftParameterGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redshiftparametergroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

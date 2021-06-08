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

// EmrInstanceGroupsGetter has a method to return a EmrInstanceGroupInterface.
// A group's client should implement this interface.
type EmrInstanceGroupsGetter interface {
	EmrInstanceGroups(namespace string) EmrInstanceGroupInterface
}

// EmrInstanceGroupInterface has methods to work with EmrInstanceGroup resources.
type EmrInstanceGroupInterface interface {
	Create(ctx context.Context, emrInstanceGroup *v1alpha1.EmrInstanceGroup, opts v1.CreateOptions) (*v1alpha1.EmrInstanceGroup, error)
	Update(ctx context.Context, emrInstanceGroup *v1alpha1.EmrInstanceGroup, opts v1.UpdateOptions) (*v1alpha1.EmrInstanceGroup, error)
	UpdateStatus(ctx context.Context, emrInstanceGroup *v1alpha1.EmrInstanceGroup, opts v1.UpdateOptions) (*v1alpha1.EmrInstanceGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EmrInstanceGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EmrInstanceGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EmrInstanceGroup, err error)
	EmrInstanceGroupExpansion
}

// emrInstanceGroups implements EmrInstanceGroupInterface
type emrInstanceGroups struct {
	client rest.Interface
	ns     string
}

// newEmrInstanceGroups returns a EmrInstanceGroups
func newEmrInstanceGroups(c *AwsV1alpha1Client, namespace string) *emrInstanceGroups {
	return &emrInstanceGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the emrInstanceGroup, and returns the corresponding emrInstanceGroup object, and an error if there is any.
func (c *emrInstanceGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EmrInstanceGroup, err error) {
	result = &v1alpha1.EmrInstanceGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EmrInstanceGroups that match those selectors.
func (c *emrInstanceGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EmrInstanceGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EmrInstanceGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested emrInstanceGroups.
func (c *emrInstanceGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a emrInstanceGroup and creates it.  Returns the server's representation of the emrInstanceGroup, and an error, if there is any.
func (c *emrInstanceGroups) Create(ctx context.Context, emrInstanceGroup *v1alpha1.EmrInstanceGroup, opts v1.CreateOptions) (result *v1alpha1.EmrInstanceGroup, err error) {
	result = &v1alpha1.EmrInstanceGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(emrInstanceGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a emrInstanceGroup and updates it. Returns the server's representation of the emrInstanceGroup, and an error, if there is any.
func (c *emrInstanceGroups) Update(ctx context.Context, emrInstanceGroup *v1alpha1.EmrInstanceGroup, opts v1.UpdateOptions) (result *v1alpha1.EmrInstanceGroup, err error) {
	result = &v1alpha1.EmrInstanceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		Name(emrInstanceGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(emrInstanceGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *emrInstanceGroups) UpdateStatus(ctx context.Context, emrInstanceGroup *v1alpha1.EmrInstanceGroup, opts v1.UpdateOptions) (result *v1alpha1.EmrInstanceGroup, err error) {
	result = &v1alpha1.EmrInstanceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		Name(emrInstanceGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(emrInstanceGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the emrInstanceGroup and deletes it. Returns an error if one occurs.
func (c *emrInstanceGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *emrInstanceGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("emrinstancegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched emrInstanceGroup.
func (c *emrInstanceGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EmrInstanceGroup, err error) {
	result = &v1alpha1.EmrInstanceGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("emrinstancegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
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

// ConfigConfigurationRecorderStatus_sGetter has a method to return a ConfigConfigurationRecorderStatus_Interface.
// A group's client should implement this interface.
type ConfigConfigurationRecorderStatus_sGetter interface {
	ConfigConfigurationRecorderStatus_s(namespace string) ConfigConfigurationRecorderStatus_Interface
}

// ConfigConfigurationRecorderStatus_Interface has methods to work with ConfigConfigurationRecorderStatus_ resources.
type ConfigConfigurationRecorderStatus_Interface interface {
	Create(ctx context.Context, configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_, opts v1.CreateOptions) (*v1alpha1.ConfigConfigurationRecorderStatus_, error)
	Update(ctx context.Context, configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_, opts v1.UpdateOptions) (*v1alpha1.ConfigConfigurationRecorderStatus_, error)
	UpdateStatus(ctx context.Context, configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_, opts v1.UpdateOptions) (*v1alpha1.ConfigConfigurationRecorderStatus_, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ConfigConfigurationRecorderStatus_, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ConfigConfigurationRecorderStatus_List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error)
	ConfigConfigurationRecorderStatus_Expansion
}

// configConfigurationRecorderStatus_s implements ConfigConfigurationRecorderStatus_Interface
type configConfigurationRecorderStatus_s struct {
	client rest.Interface
	ns     string
}

// newConfigConfigurationRecorderStatus_s returns a ConfigConfigurationRecorderStatus_s
func newConfigConfigurationRecorderStatus_s(c *AwsV1alpha1Client, namespace string) *configConfigurationRecorderStatus_s {
	return &configConfigurationRecorderStatus_s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configConfigurationRecorderStatus_, and returns the corresponding configConfigurationRecorderStatus_ object, and an error if there is any.
func (c *configConfigurationRecorderStatus_s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigConfigurationRecorderStatus_{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigConfigurationRecorderStatus_s that match those selectors.
func (c *configConfigurationRecorderStatus_s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigConfigurationRecorderStatus_List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigConfigurationRecorderStatus_List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configConfigurationRecorderStatus_s.
func (c *configConfigurationRecorderStatus_s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a configConfigurationRecorderStatus_ and creates it.  Returns the server's representation of the configConfigurationRecorderStatus_, and an error, if there is any.
func (c *configConfigurationRecorderStatus_s) Create(ctx context.Context, configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_, opts v1.CreateOptions) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigConfigurationRecorderStatus_{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configConfigurationRecorderStatus_).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a configConfigurationRecorderStatus_ and updates it. Returns the server's representation of the configConfigurationRecorderStatus_, and an error, if there is any.
func (c *configConfigurationRecorderStatus_s) Update(ctx context.Context, configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_, opts v1.UpdateOptions) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigConfigurationRecorderStatus_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		Name(configConfigurationRecorderStatus_.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configConfigurationRecorderStatus_).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *configConfigurationRecorderStatus_s) UpdateStatus(ctx context.Context, configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_, opts v1.UpdateOptions) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigConfigurationRecorderStatus_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		Name(configConfigurationRecorderStatus_.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configConfigurationRecorderStatus_).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the configConfigurationRecorderStatus_ and deletes it. Returns an error if one occurs.
func (c *configConfigurationRecorderStatus_s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configConfigurationRecorderStatus_s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched configConfigurationRecorderStatus_.
func (c *configConfigurationRecorderStatus_s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigConfigurationRecorderStatus_{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configconfigurationrecorderstatus_s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
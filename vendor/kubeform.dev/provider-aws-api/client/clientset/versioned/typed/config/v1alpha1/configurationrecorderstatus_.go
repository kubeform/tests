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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-aws-api/apis/config/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigurationRecorderStatus_sGetter has a method to return a ConfigurationRecorderStatus_Interface.
// A group's client should implement this interface.
type ConfigurationRecorderStatus_sGetter interface {
	ConfigurationRecorderStatus_s(namespace string) ConfigurationRecorderStatus_Interface
}

// ConfigurationRecorderStatus_Interface has methods to work with ConfigurationRecorderStatus_ resources.
type ConfigurationRecorderStatus_Interface interface {
	Create(ctx context.Context, configurationRecorderStatus_ *v1alpha1.ConfigurationRecorderStatus_, opts v1.CreateOptions) (*v1alpha1.ConfigurationRecorderStatus_, error)
	Update(ctx context.Context, configurationRecorderStatus_ *v1alpha1.ConfigurationRecorderStatus_, opts v1.UpdateOptions) (*v1alpha1.ConfigurationRecorderStatus_, error)
	UpdateStatus(ctx context.Context, configurationRecorderStatus_ *v1alpha1.ConfigurationRecorderStatus_, opts v1.UpdateOptions) (*v1alpha1.ConfigurationRecorderStatus_, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ConfigurationRecorderStatus_, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ConfigurationRecorderStatus_List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigurationRecorderStatus_, err error)
	ConfigurationRecorderStatus_Expansion
}

// configurationRecorderStatus_s implements ConfigurationRecorderStatus_Interface
type configurationRecorderStatus_s struct {
	client rest.Interface
	ns     string
}

// newConfigurationRecorderStatus_s returns a ConfigurationRecorderStatus_s
func newConfigurationRecorderStatus_s(c *ConfigV1alpha1Client, namespace string) *configurationRecorderStatus_s {
	return &configurationRecorderStatus_s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configurationRecorderStatus_, and returns the corresponding configurationRecorderStatus_ object, and an error if there is any.
func (c *configurationRecorderStatus_s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigurationRecorderStatus_{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigurationRecorderStatus_s that match those selectors.
func (c *configurationRecorderStatus_s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigurationRecorderStatus_List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigurationRecorderStatus_List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configurationRecorderStatus_s.
func (c *configurationRecorderStatus_s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a configurationRecorderStatus_ and creates it.  Returns the server's representation of the configurationRecorderStatus_, and an error, if there is any.
func (c *configurationRecorderStatus_s) Create(ctx context.Context, configurationRecorderStatus_ *v1alpha1.ConfigurationRecorderStatus_, opts v1.CreateOptions) (result *v1alpha1.ConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigurationRecorderStatus_{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configurationRecorderStatus_).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a configurationRecorderStatus_ and updates it. Returns the server's representation of the configurationRecorderStatus_, and an error, if there is any.
func (c *configurationRecorderStatus_s) Update(ctx context.Context, configurationRecorderStatus_ *v1alpha1.ConfigurationRecorderStatus_, opts v1.UpdateOptions) (result *v1alpha1.ConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigurationRecorderStatus_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		Name(configurationRecorderStatus_.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configurationRecorderStatus_).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *configurationRecorderStatus_s) UpdateStatus(ctx context.Context, configurationRecorderStatus_ *v1alpha1.ConfigurationRecorderStatus_, opts v1.UpdateOptions) (result *v1alpha1.ConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigurationRecorderStatus_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		Name(configurationRecorderStatus_.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configurationRecorderStatus_).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the configurationRecorderStatus_ and deletes it. Returns an error if one occurs.
func (c *configurationRecorderStatus_s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configurationRecorderStatus_s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched configurationRecorderStatus_.
func (c *configurationRecorderStatus_s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigurationRecorderStatus_, err error) {
	result = &v1alpha1.ConfigurationRecorderStatus_{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configurationrecorderstatus_s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
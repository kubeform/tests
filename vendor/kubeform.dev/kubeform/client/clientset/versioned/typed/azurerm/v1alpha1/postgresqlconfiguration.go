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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PostgresqlConfigurationsGetter has a method to return a PostgresqlConfigurationInterface.
// A group's client should implement this interface.
type PostgresqlConfigurationsGetter interface {
	PostgresqlConfigurations(namespace string) PostgresqlConfigurationInterface
}

// PostgresqlConfigurationInterface has methods to work with PostgresqlConfiguration resources.
type PostgresqlConfigurationInterface interface {
	Create(ctx context.Context, postgresqlConfiguration *v1alpha1.PostgresqlConfiguration, opts v1.CreateOptions) (*v1alpha1.PostgresqlConfiguration, error)
	Update(ctx context.Context, postgresqlConfiguration *v1alpha1.PostgresqlConfiguration, opts v1.UpdateOptions) (*v1alpha1.PostgresqlConfiguration, error)
	UpdateStatus(ctx context.Context, postgresqlConfiguration *v1alpha1.PostgresqlConfiguration, opts v1.UpdateOptions) (*v1alpha1.PostgresqlConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PostgresqlConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PostgresqlConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PostgresqlConfiguration, err error)
	PostgresqlConfigurationExpansion
}

// postgresqlConfigurations implements PostgresqlConfigurationInterface
type postgresqlConfigurations struct {
	client rest.Interface
	ns     string
}

// newPostgresqlConfigurations returns a PostgresqlConfigurations
func newPostgresqlConfigurations(c *AzurermV1alpha1Client, namespace string) *postgresqlConfigurations {
	return &postgresqlConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the postgresqlConfiguration, and returns the corresponding postgresqlConfiguration object, and an error if there is any.
func (c *postgresqlConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PostgresqlConfiguration, err error) {
	result = &v1alpha1.PostgresqlConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PostgresqlConfigurations that match those selectors.
func (c *postgresqlConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PostgresqlConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PostgresqlConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested postgresqlConfigurations.
func (c *postgresqlConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a postgresqlConfiguration and creates it.  Returns the server's representation of the postgresqlConfiguration, and an error, if there is any.
func (c *postgresqlConfigurations) Create(ctx context.Context, postgresqlConfiguration *v1alpha1.PostgresqlConfiguration, opts v1.CreateOptions) (result *v1alpha1.PostgresqlConfiguration, err error) {
	result = &v1alpha1.PostgresqlConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(postgresqlConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a postgresqlConfiguration and updates it. Returns the server's representation of the postgresqlConfiguration, and an error, if there is any.
func (c *postgresqlConfigurations) Update(ctx context.Context, postgresqlConfiguration *v1alpha1.PostgresqlConfiguration, opts v1.UpdateOptions) (result *v1alpha1.PostgresqlConfiguration, err error) {
	result = &v1alpha1.PostgresqlConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		Name(postgresqlConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(postgresqlConfiguration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *postgresqlConfigurations) UpdateStatus(ctx context.Context, postgresqlConfiguration *v1alpha1.PostgresqlConfiguration, opts v1.UpdateOptions) (result *v1alpha1.PostgresqlConfiguration, err error) {
	result = &v1alpha1.PostgresqlConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		Name(postgresqlConfiguration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(postgresqlConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the postgresqlConfiguration and deletes it. Returns an error if one occurs.
func (c *postgresqlConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *postgresqlConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched postgresqlConfiguration.
func (c *postgresqlConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PostgresqlConfiguration, err error) {
	result = &v1alpha1.PostgresqlConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("postgresqlconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

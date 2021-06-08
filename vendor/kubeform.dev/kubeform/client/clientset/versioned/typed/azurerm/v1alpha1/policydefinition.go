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

// PolicyDefinitionsGetter has a method to return a PolicyDefinitionInterface.
// A group's client should implement this interface.
type PolicyDefinitionsGetter interface {
	PolicyDefinitions(namespace string) PolicyDefinitionInterface
}

// PolicyDefinitionInterface has methods to work with PolicyDefinition resources.
type PolicyDefinitionInterface interface {
	Create(ctx context.Context, policyDefinition *v1alpha1.PolicyDefinition, opts v1.CreateOptions) (*v1alpha1.PolicyDefinition, error)
	Update(ctx context.Context, policyDefinition *v1alpha1.PolicyDefinition, opts v1.UpdateOptions) (*v1alpha1.PolicyDefinition, error)
	UpdateStatus(ctx context.Context, policyDefinition *v1alpha1.PolicyDefinition, opts v1.UpdateOptions) (*v1alpha1.PolicyDefinition, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PolicyDefinition, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PolicyDefinitionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyDefinition, err error)
	PolicyDefinitionExpansion
}

// policyDefinitions implements PolicyDefinitionInterface
type policyDefinitions struct {
	client rest.Interface
	ns     string
}

// newPolicyDefinitions returns a PolicyDefinitions
func newPolicyDefinitions(c *AzurermV1alpha1Client, namespace string) *policyDefinitions {
	return &policyDefinitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the policyDefinition, and returns the corresponding policyDefinition object, and an error if there is any.
func (c *policyDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicyDefinition, err error) {
	result = &v1alpha1.PolicyDefinition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("policydefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PolicyDefinitions that match those selectors.
func (c *policyDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicyDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PolicyDefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("policydefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested policyDefinitions.
func (c *policyDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("policydefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a policyDefinition and creates it.  Returns the server's representation of the policyDefinition, and an error, if there is any.
func (c *policyDefinitions) Create(ctx context.Context, policyDefinition *v1alpha1.PolicyDefinition, opts v1.CreateOptions) (result *v1alpha1.PolicyDefinition, err error) {
	result = &v1alpha1.PolicyDefinition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("policydefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyDefinition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a policyDefinition and updates it. Returns the server's representation of the policyDefinition, and an error, if there is any.
func (c *policyDefinitions) Update(ctx context.Context, policyDefinition *v1alpha1.PolicyDefinition, opts v1.UpdateOptions) (result *v1alpha1.PolicyDefinition, err error) {
	result = &v1alpha1.PolicyDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("policydefinitions").
		Name(policyDefinition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyDefinition).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *policyDefinitions) UpdateStatus(ctx context.Context, policyDefinition *v1alpha1.PolicyDefinition, opts v1.UpdateOptions) (result *v1alpha1.PolicyDefinition, err error) {
	result = &v1alpha1.PolicyDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("policydefinitions").
		Name(policyDefinition.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyDefinition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the policyDefinition and deletes it. Returns an error if one occurs.
func (c *policyDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("policydefinitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *policyDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("policydefinitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched policyDefinition.
func (c *policyDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyDefinition, err error) {
	result = &v1alpha1.PolicyDefinition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("policydefinitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

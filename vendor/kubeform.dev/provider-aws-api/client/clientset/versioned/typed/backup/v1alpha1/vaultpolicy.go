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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/backup/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VaultPoliciesGetter has a method to return a VaultPolicyInterface.
// A group's client should implement this interface.
type VaultPoliciesGetter interface {
	VaultPolicies(namespace string) VaultPolicyInterface
}

// VaultPolicyInterface has methods to work with VaultPolicy resources.
type VaultPolicyInterface interface {
	Create(ctx context.Context, vaultPolicy *v1alpha1.VaultPolicy, opts v1.CreateOptions) (*v1alpha1.VaultPolicy, error)
	Update(ctx context.Context, vaultPolicy *v1alpha1.VaultPolicy, opts v1.UpdateOptions) (*v1alpha1.VaultPolicy, error)
	UpdateStatus(ctx context.Context, vaultPolicy *v1alpha1.VaultPolicy, opts v1.UpdateOptions) (*v1alpha1.VaultPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VaultPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VaultPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VaultPolicy, err error)
	VaultPolicyExpansion
}

// vaultPolicies implements VaultPolicyInterface
type vaultPolicies struct {
	client rest.Interface
	ns     string
}

// newVaultPolicies returns a VaultPolicies
func newVaultPolicies(c *BackupV1alpha1Client, namespace string) *vaultPolicies {
	return &vaultPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vaultPolicy, and returns the corresponding vaultPolicy object, and an error if there is any.
func (c *vaultPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VaultPolicy, err error) {
	result = &v1alpha1.VaultPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaultpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VaultPolicies that match those selectors.
func (c *vaultPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VaultPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VaultPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaultpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vaultPolicies.
func (c *vaultPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vaultpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vaultPolicy and creates it.  Returns the server's representation of the vaultPolicy, and an error, if there is any.
func (c *vaultPolicies) Create(ctx context.Context, vaultPolicy *v1alpha1.VaultPolicy, opts v1.CreateOptions) (result *v1alpha1.VaultPolicy, err error) {
	result = &v1alpha1.VaultPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vaultpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vaultPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vaultPolicy and updates it. Returns the server's representation of the vaultPolicy, and an error, if there is any.
func (c *vaultPolicies) Update(ctx context.Context, vaultPolicy *v1alpha1.VaultPolicy, opts v1.UpdateOptions) (result *v1alpha1.VaultPolicy, err error) {
	result = &v1alpha1.VaultPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vaultpolicies").
		Name(vaultPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vaultPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vaultPolicies) UpdateStatus(ctx context.Context, vaultPolicy *v1alpha1.VaultPolicy, opts v1.UpdateOptions) (result *v1alpha1.VaultPolicy, err error) {
	result = &v1alpha1.VaultPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vaultpolicies").
		Name(vaultPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vaultPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vaultPolicy and deletes it. Returns an error if one occurs.
func (c *vaultPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaultpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vaultPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaultpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vaultPolicy.
func (c *vaultPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VaultPolicy, err error) {
	result = &v1alpha1.VaultPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vaultpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

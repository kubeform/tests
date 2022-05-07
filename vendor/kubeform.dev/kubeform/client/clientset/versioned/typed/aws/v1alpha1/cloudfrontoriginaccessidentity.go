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

// CloudfrontOriginAccessIdentitiesGetter has a method to return a CloudfrontOriginAccessIdentityInterface.
// A group's client should implement this interface.
type CloudfrontOriginAccessIdentitiesGetter interface {
	CloudfrontOriginAccessIdentities(namespace string) CloudfrontOriginAccessIdentityInterface
}

// CloudfrontOriginAccessIdentityInterface has methods to work with CloudfrontOriginAccessIdentity resources.
type CloudfrontOriginAccessIdentityInterface interface {
	Create(ctx context.Context, cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity, opts v1.CreateOptions) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	Update(ctx context.Context, cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity, opts v1.UpdateOptions) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	UpdateStatus(ctx context.Context, cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity, opts v1.UpdateOptions) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CloudfrontOriginAccessIdentityList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error)
	CloudfrontOriginAccessIdentityExpansion
}

// cloudfrontOriginAccessIdentities implements CloudfrontOriginAccessIdentityInterface
type cloudfrontOriginAccessIdentities struct {
	client rest.Interface
	ns     string
}

// newCloudfrontOriginAccessIdentities returns a CloudfrontOriginAccessIdentities
func newCloudfrontOriginAccessIdentities(c *AwsV1alpha1Client, namespace string) *cloudfrontOriginAccessIdentities {
	return &cloudfrontOriginAccessIdentities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudfrontOriginAccessIdentity, and returns the corresponding cloudfrontOriginAccessIdentity object, and an error if there is any.
func (c *cloudfrontOriginAccessIdentities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudfrontOriginAccessIdentities that match those selectors.
func (c *cloudfrontOriginAccessIdentities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudfrontOriginAccessIdentityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudfrontOriginAccessIdentityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudfrontOriginAccessIdentities.
func (c *cloudfrontOriginAccessIdentities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloudfrontOriginAccessIdentity and creates it.  Returns the server's representation of the cloudfrontOriginAccessIdentity, and an error, if there is any.
func (c *cloudfrontOriginAccessIdentities) Create(ctx context.Context, cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity, opts v1.CreateOptions) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudfrontOriginAccessIdentity).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloudfrontOriginAccessIdentity and updates it. Returns the server's representation of the cloudfrontOriginAccessIdentity, and an error, if there is any.
func (c *cloudfrontOriginAccessIdentities) Update(ctx context.Context, cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity, opts v1.UpdateOptions) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(cloudfrontOriginAccessIdentity.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudfrontOriginAccessIdentity).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloudfrontOriginAccessIdentities) UpdateStatus(ctx context.Context, cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity, opts v1.UpdateOptions) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(cloudfrontOriginAccessIdentity.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudfrontOriginAccessIdentity).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloudfrontOriginAccessIdentity and deletes it. Returns an error if one occurs.
func (c *cloudfrontOriginAccessIdentities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudfrontOriginAccessIdentities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloudfrontOriginAccessIdentity.
func (c *cloudfrontOriginAccessIdentities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
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

// GameliftBuildsGetter has a method to return a GameliftBuildInterface.
// A group's client should implement this interface.
type GameliftBuildsGetter interface {
	GameliftBuilds(namespace string) GameliftBuildInterface
}

// GameliftBuildInterface has methods to work with GameliftBuild resources.
type GameliftBuildInterface interface {
	Create(ctx context.Context, gameliftBuild *v1alpha1.GameliftBuild, opts v1.CreateOptions) (*v1alpha1.GameliftBuild, error)
	Update(ctx context.Context, gameliftBuild *v1alpha1.GameliftBuild, opts v1.UpdateOptions) (*v1alpha1.GameliftBuild, error)
	UpdateStatus(ctx context.Context, gameliftBuild *v1alpha1.GameliftBuild, opts v1.UpdateOptions) (*v1alpha1.GameliftBuild, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GameliftBuild, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GameliftBuildList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GameliftBuild, err error)
	GameliftBuildExpansion
}

// gameliftBuilds implements GameliftBuildInterface
type gameliftBuilds struct {
	client rest.Interface
	ns     string
}

// newGameliftBuilds returns a GameliftBuilds
func newGameliftBuilds(c *AwsV1alpha1Client, namespace string) *gameliftBuilds {
	return &gameliftBuilds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gameliftBuild, and returns the corresponding gameliftBuild object, and an error if there is any.
func (c *gameliftBuilds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GameliftBuild, err error) {
	result = &v1alpha1.GameliftBuild{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GameliftBuilds that match those selectors.
func (c *gameliftBuilds) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GameliftBuildList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GameliftBuildList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gameliftBuilds.
func (c *gameliftBuilds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gameliftBuild and creates it.  Returns the server's representation of the gameliftBuild, and an error, if there is any.
func (c *gameliftBuilds) Create(ctx context.Context, gameliftBuild *v1alpha1.GameliftBuild, opts v1.CreateOptions) (result *v1alpha1.GameliftBuild, err error) {
	result = &v1alpha1.GameliftBuild{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameliftBuild).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gameliftBuild and updates it. Returns the server's representation of the gameliftBuild, and an error, if there is any.
func (c *gameliftBuilds) Update(ctx context.Context, gameliftBuild *v1alpha1.GameliftBuild, opts v1.UpdateOptions) (result *v1alpha1.GameliftBuild, err error) {
	result = &v1alpha1.GameliftBuild{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		Name(gameliftBuild.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameliftBuild).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gameliftBuilds) UpdateStatus(ctx context.Context, gameliftBuild *v1alpha1.GameliftBuild, opts v1.UpdateOptions) (result *v1alpha1.GameliftBuild, err error) {
	result = &v1alpha1.GameliftBuild{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		Name(gameliftBuild.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameliftBuild).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gameliftBuild and deletes it. Returns an error if one occurs.
func (c *gameliftBuilds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gameliftBuilds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameliftbuilds").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gameliftBuild.
func (c *gameliftBuilds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GameliftBuild, err error) {
	result = &v1alpha1.GameliftBuild{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gameliftbuilds").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

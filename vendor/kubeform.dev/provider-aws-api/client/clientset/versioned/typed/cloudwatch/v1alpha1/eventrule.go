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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudwatch/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EventRulesGetter has a method to return a EventRuleInterface.
// A group's client should implement this interface.
type EventRulesGetter interface {
	EventRules(namespace string) EventRuleInterface
}

// EventRuleInterface has methods to work with EventRule resources.
type EventRuleInterface interface {
	Create(ctx context.Context, eventRule *v1alpha1.EventRule, opts v1.CreateOptions) (*v1alpha1.EventRule, error)
	Update(ctx context.Context, eventRule *v1alpha1.EventRule, opts v1.UpdateOptions) (*v1alpha1.EventRule, error)
	UpdateStatus(ctx context.Context, eventRule *v1alpha1.EventRule, opts v1.UpdateOptions) (*v1alpha1.EventRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EventRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EventRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventRule, err error)
	EventRuleExpansion
}

// eventRules implements EventRuleInterface
type eventRules struct {
	client rest.Interface
	ns     string
}

// newEventRules returns a EventRules
func newEventRules(c *CloudwatchV1alpha1Client, namespace string) *eventRules {
	return &eventRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventRule, and returns the corresponding eventRule object, and an error if there is any.
func (c *eventRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventRule, err error) {
	result = &v1alpha1.EventRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventRules that match those selectors.
func (c *eventRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventRules.
func (c *eventRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a eventRule and creates it.  Returns the server's representation of the eventRule, and an error, if there is any.
func (c *eventRules) Create(ctx context.Context, eventRule *v1alpha1.EventRule, opts v1.CreateOptions) (result *v1alpha1.EventRule, err error) {
	result = &v1alpha1.EventRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a eventRule and updates it. Returns the server's representation of the eventRule, and an error, if there is any.
func (c *eventRules) Update(ctx context.Context, eventRule *v1alpha1.EventRule, opts v1.UpdateOptions) (result *v1alpha1.EventRule, err error) {
	result = &v1alpha1.EventRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventrules").
		Name(eventRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *eventRules) UpdateStatus(ctx context.Context, eventRule *v1alpha1.EventRule, opts v1.UpdateOptions) (result *v1alpha1.EventRule, err error) {
	result = &v1alpha1.EventRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventrules").
		Name(eventRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the eventRule and deletes it. Returns an error if one occurs.
func (c *eventRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched eventRule.
func (c *eventRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventRule, err error) {
	result = &v1alpha1.EventRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
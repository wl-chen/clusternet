/*
Copyright 2021 The Clusternet Authors.

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

	v1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	scheme "github.com/clusternet/clusternet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalizationsGetter has a method to return a GlobalizationInterface.
// A group's client should implement this interface.
type GlobalizationsGetter interface {
	Globalizations() GlobalizationInterface
}

// GlobalizationInterface has methods to work with Globalization resources.
type GlobalizationInterface interface {
	Create(ctx context.Context, globalization *v1alpha1.Globalization, opts v1.CreateOptions) (*v1alpha1.Globalization, error)
	Update(ctx context.Context, globalization *v1alpha1.Globalization, opts v1.UpdateOptions) (*v1alpha1.Globalization, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Globalization, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GlobalizationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Globalization, err error)
	GlobalizationExpansion
}

// globalizations implements GlobalizationInterface
type globalizations struct {
	client rest.Interface
}

// newGlobalizations returns a Globalizations
func newGlobalizations(c *AppsV1alpha1Client) *globalizations {
	return &globalizations{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalization, and returns the corresponding globalization object, and an error if there is any.
func (c *globalizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Globalization, err error) {
	result = &v1alpha1.Globalization{}
	err = c.client.Get().
		Resource("globalizations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Globalizations that match those selectors.
func (c *globalizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GlobalizationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GlobalizationList{}
	err = c.client.Get().
		Resource("globalizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalizations.
func (c *globalizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalization and creates it.  Returns the server's representation of the globalization, and an error, if there is any.
func (c *globalizations) Create(ctx context.Context, globalization *v1alpha1.Globalization, opts v1.CreateOptions) (result *v1alpha1.Globalization, err error) {
	result = &v1alpha1.Globalization{}
	err = c.client.Post().
		Resource("globalizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalization).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalization and updates it. Returns the server's representation of the globalization, and an error, if there is any.
func (c *globalizations) Update(ctx context.Context, globalization *v1alpha1.Globalization, opts v1.UpdateOptions) (result *v1alpha1.Globalization, err error) {
	result = &v1alpha1.Globalization{}
	err = c.client.Put().
		Resource("globalizations").
		Name(globalization.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalization).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalization and deletes it. Returns an error if one occurs.
func (c *globalizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalizations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalizations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalization.
func (c *globalizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Globalization, err error) {
	result = &v1alpha1.Globalization{}
	err = c.client.Patch(pt).
		Resource("globalizations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

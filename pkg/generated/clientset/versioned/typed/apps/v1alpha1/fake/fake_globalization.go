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

package fake

import (
	"context"

	v1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalizations implements GlobalizationInterface
type FakeGlobalizations struct {
	Fake *FakeAppsV1alpha1
}

var globalizationsResource = schema.GroupVersionResource{Group: "apps.clusternet.io", Version: "v1alpha1", Resource: "globalizations"}

var globalizationsKind = schema.GroupVersionKind{Group: "apps.clusternet.io", Version: "v1alpha1", Kind: "Globalization"}

// Get takes name of the globalization, and returns the corresponding globalization object, and an error if there is any.
func (c *FakeGlobalizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Globalization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(globalizationsResource, name), &v1alpha1.Globalization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Globalization), err
}

// List takes label and field selectors, and returns the list of Globalizations that match those selectors.
func (c *FakeGlobalizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GlobalizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(globalizationsResource, globalizationsKind, opts), &v1alpha1.GlobalizationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlobalizationList{ListMeta: obj.(*v1alpha1.GlobalizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlobalizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalizations.
func (c *FakeGlobalizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(globalizationsResource, opts))
}

// Create takes the representation of a globalization and creates it.  Returns the server's representation of the globalization, and an error, if there is any.
func (c *FakeGlobalizations) Create(ctx context.Context, globalization *v1alpha1.Globalization, opts v1.CreateOptions) (result *v1alpha1.Globalization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(globalizationsResource, globalization), &v1alpha1.Globalization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Globalization), err
}

// Update takes the representation of a globalization and updates it. Returns the server's representation of the globalization, and an error, if there is any.
func (c *FakeGlobalizations) Update(ctx context.Context, globalization *v1alpha1.Globalization, opts v1.UpdateOptions) (result *v1alpha1.Globalization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(globalizationsResource, globalization), &v1alpha1.Globalization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Globalization), err
}

// Delete takes name of the globalization and deletes it. Returns an error if one occurs.
func (c *FakeGlobalizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(globalizationsResource, name), &v1alpha1.Globalization{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(globalizationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlobalizationList{})
	return err
}

// Patch applies the patch and returns the patched globalization.
func (c *FakeGlobalizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Globalization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(globalizationsResource, name, pt, data, subresources...), &v1alpha1.Globalization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Globalization), err
}

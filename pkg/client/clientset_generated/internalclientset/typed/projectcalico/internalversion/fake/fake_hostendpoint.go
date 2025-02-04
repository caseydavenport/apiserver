// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
)

// FakeHostEndpoints implements HostEndpointInterface
type FakeHostEndpoints struct {
	Fake *FakeProjectcalico
}

var hostendpointsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "hostendpoints"}

var hostendpointsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "HostEndpoint"}

// Get takes name of the hostEndpoint, and returns the corresponding hostEndpoint object, and an error if there is any.
func (c *FakeHostEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(hostendpointsResource, name), &projectcalico.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.HostEndpoint), err
}

// List takes label and field selectors, and returns the list of HostEndpoints that match those selectors.
func (c *FakeHostEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.HostEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(hostendpointsResource, hostendpointsKind, opts), &projectcalico.HostEndpointList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.HostEndpointList{ListMeta: obj.(*projectcalico.HostEndpointList).ListMeta}
	for _, item := range obj.(*projectcalico.HostEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostEndpoints.
func (c *FakeHostEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(hostendpointsResource, opts))
}

// Create takes the representation of a hostEndpoint and creates it.  Returns the server's representation of the hostEndpoint, and an error, if there is any.
func (c *FakeHostEndpoints) Create(ctx context.Context, hostEndpoint *projectcalico.HostEndpoint, opts v1.CreateOptions) (result *projectcalico.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(hostendpointsResource, hostEndpoint), &projectcalico.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.HostEndpoint), err
}

// Update takes the representation of a hostEndpoint and updates it. Returns the server's representation of the hostEndpoint, and an error, if there is any.
func (c *FakeHostEndpoints) Update(ctx context.Context, hostEndpoint *projectcalico.HostEndpoint, opts v1.UpdateOptions) (result *projectcalico.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(hostendpointsResource, hostEndpoint), &projectcalico.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.HostEndpoint), err
}

// Delete takes name of the hostEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeHostEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(hostendpointsResource, name), &projectcalico.HostEndpoint{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(hostendpointsResource, listOpts)

	_, err := c.Fake.Invokes(action, &projectcalico.HostEndpointList{})
	return err
}

// Patch applies the patch and returns the patched hostEndpoint.
func (c *FakeHostEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.HostEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(hostendpointsResource, name, pt, data, subresources...), &projectcalico.HostEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.HostEndpoint), err
}

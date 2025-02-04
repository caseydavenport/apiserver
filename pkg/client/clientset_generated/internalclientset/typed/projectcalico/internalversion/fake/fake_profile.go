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

// FakeProfiles implements ProfileInterface
type FakeProfiles struct {
	Fake *FakeProjectcalico
}

var profilesResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "profiles"}

var profilesKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "Profile"}

// Get takes name of the profile, and returns the corresponding profile object, and an error if there is any.
func (c *FakeProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(profilesResource, name), &projectcalico.Profile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.Profile), err
}

// List takes label and field selectors, and returns the list of Profiles that match those selectors.
func (c *FakeProfiles) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.ProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(profilesResource, profilesKind, opts), &projectcalico.ProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.ProfileList{ListMeta: obj.(*projectcalico.ProfileList).ListMeta}
	for _, item := range obj.(*projectcalico.ProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested profiles.
func (c *FakeProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(profilesResource, opts))
}

// Create takes the representation of a profile and creates it.  Returns the server's representation of the profile, and an error, if there is any.
func (c *FakeProfiles) Create(ctx context.Context, profile *projectcalico.Profile, opts v1.CreateOptions) (result *projectcalico.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(profilesResource, profile), &projectcalico.Profile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.Profile), err
}

// Update takes the representation of a profile and updates it. Returns the server's representation of the profile, and an error, if there is any.
func (c *FakeProfiles) Update(ctx context.Context, profile *projectcalico.Profile, opts v1.UpdateOptions) (result *projectcalico.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(profilesResource, profile), &projectcalico.Profile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.Profile), err
}

// Delete takes name of the profile and deletes it. Returns an error if one occurs.
func (c *FakeProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(profilesResource, name), &projectcalico.Profile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(profilesResource, listOpts)

	_, err := c.Fake.Invokes(action, &projectcalico.ProfileList{})
	return err
}

// Patch applies the patch and returns the patched profile.
func (c *FakeProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(profilesResource, name, pt, data, subresources...), &projectcalico.Profile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.Profile), err
}

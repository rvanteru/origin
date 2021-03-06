// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	configv1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOAuths implements OAuthInterface
type FakeOAuths struct {
	Fake *FakeConfigV1
}

var oauthsResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "oauths"}

var oauthsKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "OAuth"}

// Get takes name of the oAuth, and returns the corresponding oAuth object, and an error if there is any.
func (c *FakeOAuths) Get(name string, options v1.GetOptions) (result *configv1.OAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(oauthsResource, name), &configv1.OAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OAuth), err
}

// List takes label and field selectors, and returns the list of OAuths that match those selectors.
func (c *FakeOAuths) List(opts v1.ListOptions) (result *configv1.OAuthList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(oauthsResource, oauthsKind, opts), &configv1.OAuthList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.OAuthList{ListMeta: obj.(*configv1.OAuthList).ListMeta}
	for _, item := range obj.(*configv1.OAuthList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oAuths.
func (c *FakeOAuths) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(oauthsResource, opts))
}

// Create takes the representation of a oAuth and creates it.  Returns the server's representation of the oAuth, and an error, if there is any.
func (c *FakeOAuths) Create(oAuth *configv1.OAuth) (result *configv1.OAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(oauthsResource, oAuth), &configv1.OAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OAuth), err
}

// Update takes the representation of a oAuth and updates it. Returns the server's representation of the oAuth, and an error, if there is any.
func (c *FakeOAuths) Update(oAuth *configv1.OAuth) (result *configv1.OAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(oauthsResource, oAuth), &configv1.OAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OAuth), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOAuths) UpdateStatus(oAuth *configv1.OAuth) (*configv1.OAuth, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(oauthsResource, "status", oAuth), &configv1.OAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OAuth), err
}

// Delete takes name of the oAuth and deletes it. Returns an error if one occurs.
func (c *FakeOAuths) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(oauthsResource, name), &configv1.OAuth{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOAuths) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(oauthsResource, listOptions)

	_, err := c.Fake.Invokes(action, &configv1.OAuthList{})
	return err
}

// Patch applies the patch and returns the patched oAuth.
func (c *FakeOAuths) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *configv1.OAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(oauthsResource, name, pt, data, subresources...), &configv1.OAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OAuth), err
}

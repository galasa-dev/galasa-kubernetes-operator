/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/galasa-dev/kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	scheme "github.com/galasa-dev/kubernetes-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GalasaMetricsComponentsGetter has a method to return a GalasaMetricsComponentInterface.
// A group's client should implement this interface.
type GalasaMetricsComponentsGetter interface {
	GalasaMetricsComponents(namespace string) GalasaMetricsComponentInterface
}

// GalasaMetricsComponentInterface has methods to work with GalasaMetricsComponent resources.
type GalasaMetricsComponentInterface interface {
	Create(ctx context.Context, galasaMetricsComponent *v2alpha1.GalasaMetricsComponent, opts v1.CreateOptions) (*v2alpha1.GalasaMetricsComponent, error)
	Update(ctx context.Context, galasaMetricsComponent *v2alpha1.GalasaMetricsComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaMetricsComponent, error)
	UpdateStatus(ctx context.Context, galasaMetricsComponent *v2alpha1.GalasaMetricsComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaMetricsComponent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaMetricsComponent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaMetricsComponentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaMetricsComponent, err error)
	GalasaMetricsComponentExpansion
}

// galasaMetricsComponents implements GalasaMetricsComponentInterface
type galasaMetricsComponents struct {
	client rest.Interface
	ns     string
}

// newGalasaMetricsComponents returns a GalasaMetricsComponents
func newGalasaMetricsComponents(c *GalasaV2alpha1Client, namespace string) *galasaMetricsComponents {
	return &galasaMetricsComponents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the galasaMetricsComponent, and returns the corresponding galasaMetricsComponent object, and an error if there is any.
func (c *galasaMetricsComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.GalasaMetricsComponent, err error) {
	result = &v2alpha1.GalasaMetricsComponent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GalasaMetricsComponents that match those selectors.
func (c *galasaMetricsComponents) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.GalasaMetricsComponentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.GalasaMetricsComponentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested galasaMetricsComponents.
func (c *galasaMetricsComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a galasaMetricsComponent and creates it.  Returns the server's representation of the galasaMetricsComponent, and an error, if there is any.
func (c *galasaMetricsComponents) Create(ctx context.Context, galasaMetricsComponent *v2alpha1.GalasaMetricsComponent, opts v1.CreateOptions) (result *v2alpha1.GalasaMetricsComponent, err error) {
	result = &v2alpha1.GalasaMetricsComponent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaMetricsComponent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a galasaMetricsComponent and updates it. Returns the server's representation of the galasaMetricsComponent, and an error, if there is any.
func (c *galasaMetricsComponents) Update(ctx context.Context, galasaMetricsComponent *v2alpha1.GalasaMetricsComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaMetricsComponent, err error) {
	result = &v2alpha1.GalasaMetricsComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		Name(galasaMetricsComponent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaMetricsComponent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *galasaMetricsComponents) UpdateStatus(ctx context.Context, galasaMetricsComponent *v2alpha1.GalasaMetricsComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaMetricsComponent, err error) {
	result = &v2alpha1.GalasaMetricsComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		Name(galasaMetricsComponent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaMetricsComponent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the galasaMetricsComponent and deletes it. Returns an error if one occurs.
func (c *galasaMetricsComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *galasaMetricsComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galasametricscomponents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched galasaMetricsComponent.
func (c *galasaMetricsComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaMetricsComponent, err error) {
	result = &v2alpha1.GalasaMetricsComponent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("galasametricscomponents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

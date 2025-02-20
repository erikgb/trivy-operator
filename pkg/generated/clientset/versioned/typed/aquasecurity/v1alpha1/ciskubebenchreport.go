// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"
	scheme "github.com/aquasecurity/trivy-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CISKubeBenchReportsGetter has a method to return a CISKubeBenchReportInterface.
// A group's client should implement this interface.
type CISKubeBenchReportsGetter interface {
	CISKubeBenchReports() CISKubeBenchReportInterface
}

// CISKubeBenchReportInterface has methods to work with CISKubeBenchReport resources.
type CISKubeBenchReportInterface interface {
	Create(ctx context.Context, cISKubeBenchReport *v1alpha1.CISKubeBenchReport, opts v1.CreateOptions) (*v1alpha1.CISKubeBenchReport, error)
	Update(ctx context.Context, cISKubeBenchReport *v1alpha1.CISKubeBenchReport, opts v1.UpdateOptions) (*v1alpha1.CISKubeBenchReport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CISKubeBenchReport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CISKubeBenchReportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CISKubeBenchReport, err error)
	CISKubeBenchReportExpansion
}

// cISKubeBenchReports implements CISKubeBenchReportInterface
type cISKubeBenchReports struct {
	client rest.Interface
}

// newCISKubeBenchReports returns a CISKubeBenchReports
func newCISKubeBenchReports(c *AquasecurityV1alpha1Client) *cISKubeBenchReports {
	return &cISKubeBenchReports{
		client: c.RESTClient(),
	}
}

// Get takes name of the cISKubeBenchReport, and returns the corresponding cISKubeBenchReport object, and an error if there is any.
func (c *cISKubeBenchReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CISKubeBenchReport, err error) {
	result = &v1alpha1.CISKubeBenchReport{}
	err = c.client.Get().
		Resource("ciskubebenchreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CISKubeBenchReports that match those selectors.
func (c *cISKubeBenchReports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CISKubeBenchReportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CISKubeBenchReportList{}
	err = c.client.Get().
		Resource("ciskubebenchreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cISKubeBenchReports.
func (c *cISKubeBenchReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciskubebenchreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cISKubeBenchReport and creates it.  Returns the server's representation of the cISKubeBenchReport, and an error, if there is any.
func (c *cISKubeBenchReports) Create(ctx context.Context, cISKubeBenchReport *v1alpha1.CISKubeBenchReport, opts v1.CreateOptions) (result *v1alpha1.CISKubeBenchReport, err error) {
	result = &v1alpha1.CISKubeBenchReport{}
	err = c.client.Post().
		Resource("ciskubebenchreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cISKubeBenchReport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cISKubeBenchReport and updates it. Returns the server's representation of the cISKubeBenchReport, and an error, if there is any.
func (c *cISKubeBenchReports) Update(ctx context.Context, cISKubeBenchReport *v1alpha1.CISKubeBenchReport, opts v1.UpdateOptions) (result *v1alpha1.CISKubeBenchReport, err error) {
	result = &v1alpha1.CISKubeBenchReport{}
	err = c.client.Put().
		Resource("ciskubebenchreports").
		Name(cISKubeBenchReport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cISKubeBenchReport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cISKubeBenchReport and deletes it. Returns an error if one occurs.
func (c *cISKubeBenchReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciskubebenchreports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cISKubeBenchReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciskubebenchreports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cISKubeBenchReport.
func (c *cISKubeBenchReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CISKubeBenchReport, err error) {
	result = &v1alpha1.CISKubeBenchReport{}
	err = c.client.Patch(pt).
		Resource("ciskubebenchreports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

/*
Copyright 2018 The Kong Authors.

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

package v1

import (
	"time"

	v1 "github.com/SergeRadinovich/kubernetes-ingress-controller/pkg/apis/configuration/v1"
	scheme "github.com/SergeRadinovich/kubernetes-ingress-controller/pkg/client/configuration/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KongCredentialsGetter has a method to return a KongCredentialInterface.
// A group's client should implement this interface.
type KongCredentialsGetter interface {
	KongCredentials(namespace string) KongCredentialInterface
}

// KongCredentialInterface has methods to work with KongCredential resources.
type KongCredentialInterface interface {
	Create(*v1.KongCredential) (*v1.KongCredential, error)
	Update(*v1.KongCredential) (*v1.KongCredential, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KongCredential, error)
	List(opts metav1.ListOptions) (*v1.KongCredentialList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KongCredential, err error)
	KongCredentialExpansion
}

// kongCredentials implements KongCredentialInterface
type kongCredentials struct {
	client rest.Interface
	ns     string
}

// newKongCredentials returns a KongCredentials
func newKongCredentials(c *ConfigurationV1Client, namespace string) *kongCredentials {
	return &kongCredentials{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kongCredential, and returns the corresponding kongCredential object, and an error if there is any.
func (c *kongCredentials) Get(name string, options metav1.GetOptions) (result *v1.KongCredential, err error) {
	result = &v1.KongCredential{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongcredentials").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KongCredentials that match those selectors.
func (c *kongCredentials) List(opts metav1.ListOptions) (result *v1.KongCredentialList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KongCredentialList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongcredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kongCredentials.
func (c *kongCredentials) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kongcredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kongCredential and creates it.  Returns the server's representation of the kongCredential, and an error, if there is any.
func (c *kongCredentials) Create(kongCredential *v1.KongCredential) (result *v1.KongCredential, err error) {
	result = &v1.KongCredential{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kongcredentials").
		Body(kongCredential).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kongCredential and updates it. Returns the server's representation of the kongCredential, and an error, if there is any.
func (c *kongCredentials) Update(kongCredential *v1.KongCredential) (result *v1.KongCredential, err error) {
	result = &v1.KongCredential{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kongcredentials").
		Name(kongCredential.Name).
		Body(kongCredential).
		Do().
		Into(result)
	return
}

// Delete takes name of the kongCredential and deletes it. Returns an error if one occurs.
func (c *kongCredentials) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongcredentials").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kongCredentials) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongcredentials").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kongCredential.
func (c *kongCredentials) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KongCredential, err error) {
	result = &v1.KongCredential{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kongcredentials").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

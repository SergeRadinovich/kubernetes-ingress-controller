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

// KongConsumersGetter has a method to return a KongConsumerInterface.
// A group's client should implement this interface.
type KongConsumersGetter interface {
	KongConsumers(namespace string) KongConsumerInterface
}

// KongConsumerInterface has methods to work with KongConsumer resources.
type KongConsumerInterface interface {
	Create(*v1.KongConsumer) (*v1.KongConsumer, error)
	Update(*v1.KongConsumer) (*v1.KongConsumer, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KongConsumer, error)
	List(opts metav1.ListOptions) (*v1.KongConsumerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KongConsumer, err error)
	KongConsumerExpansion
}

// kongConsumers implements KongConsumerInterface
type kongConsumers struct {
	client rest.Interface
	ns     string
}

// newKongConsumers returns a KongConsumers
func newKongConsumers(c *ConfigurationV1Client, namespace string) *kongConsumers {
	return &kongConsumers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kongConsumer, and returns the corresponding kongConsumer object, and an error if there is any.
func (c *kongConsumers) Get(name string, options metav1.GetOptions) (result *v1.KongConsumer, err error) {
	result = &v1.KongConsumer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongconsumers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KongConsumers that match those selectors.
func (c *kongConsumers) List(opts metav1.ListOptions) (result *v1.KongConsumerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KongConsumerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongconsumers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kongConsumers.
func (c *kongConsumers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kongconsumers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kongConsumer and creates it.  Returns the server's representation of the kongConsumer, and an error, if there is any.
func (c *kongConsumers) Create(kongConsumer *v1.KongConsumer) (result *v1.KongConsumer, err error) {
	result = &v1.KongConsumer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kongconsumers").
		Body(kongConsumer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kongConsumer and updates it. Returns the server's representation of the kongConsumer, and an error, if there is any.
func (c *kongConsumers) Update(kongConsumer *v1.KongConsumer) (result *v1.KongConsumer, err error) {
	result = &v1.KongConsumer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kongconsumers").
		Name(kongConsumer.Name).
		Body(kongConsumer).
		Do().
		Into(result)
	return
}

// Delete takes name of the kongConsumer and deletes it. Returns an error if one occurs.
func (c *kongConsumers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongconsumers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kongConsumers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongconsumers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kongConsumer.
func (c *kongConsumers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KongConsumer, err error) {
	result = &v1.KongConsumer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kongconsumers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/snowdrop/component-operator/pkg/apis/component/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CapabilityLister helps list Capabilities.
type CapabilityLister interface {
	// List lists all Capabilities in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.Capability, err error)
	// Capabilities returns an object that can list and get Capabilities.
	Capabilities(namespace string) CapabilityNamespaceLister
	CapabilityListerExpansion
}

// capabilityLister implements the CapabilityLister interface.
type capabilityLister struct {
	indexer cache.Indexer
}

// NewCapabilityLister returns a new CapabilityLister.
func NewCapabilityLister(indexer cache.Indexer) CapabilityLister {
	return &capabilityLister{indexer: indexer}
}

// List lists all Capabilities in the indexer.
func (s *capabilityLister) List(selector labels.Selector) (ret []*v1alpha2.Capability, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Capability))
	})
	return ret, err
}

// Capabilities returns an object that can list and get Capabilities.
func (s *capabilityLister) Capabilities(namespace string) CapabilityNamespaceLister {
	return capabilityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CapabilityNamespaceLister helps list and get Capabilities.
type CapabilityNamespaceLister interface {
	// List lists all Capabilities in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.Capability, err error)
	// Get retrieves the Capability from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.Capability, error)
	CapabilityNamespaceListerExpansion
}

// capabilityNamespaceLister implements the CapabilityNamespaceLister
// interface.
type capabilityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Capabilities in the indexer for a given namespace.
func (s capabilityNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.Capability, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Capability))
	})
	return ret, err
}

// Get retrieves the Capability from the indexer for a given namespace and name.
func (s capabilityNamespaceLister) Get(name string) (*v1alpha2.Capability, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("capability"), name)
	}
	return obj.(*v1alpha2.Capability), nil
}
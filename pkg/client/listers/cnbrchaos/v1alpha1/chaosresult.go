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

package v1alpha1

import (
	v1alpha1 "github.com/vossss/cnbrchaos/pkg/apis/cnbrchaos/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ChaosResultLister helps list ChaosResults.
type ChaosResultLister interface {
	// List lists all ChaosResults in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ChaosResult, err error)
	// ChaosResults returns an object that can list and get ChaosResults.
	ChaosResults(namespace string) ChaosResultNamespaceLister
	ChaosResultListerExpansion
}

// chaosResultLister implements the ChaosResultLister interface.
type chaosResultLister struct {
	indexer cache.Indexer
}

// NewChaosResultLister returns a new ChaosResultLister.
func NewChaosResultLister(indexer cache.Indexer) ChaosResultLister {
	return &chaosResultLister{indexer: indexer}
}

// List lists all ChaosResults in the indexer.
func (s *chaosResultLister) List(selector labels.Selector) (ret []*v1alpha1.ChaosResult, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChaosResult))
	})
	return ret, err
}

// ChaosResults returns an object that can list and get ChaosResults.
func (s *chaosResultLister) ChaosResults(namespace string) ChaosResultNamespaceLister {
	return chaosResultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ChaosResultNamespaceLister helps list and get ChaosResults.
type ChaosResultNamespaceLister interface {
	// List lists all ChaosResults in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ChaosResult, err error)
	// Get retrieves the ChaosResult from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ChaosResult, error)
	ChaosResultNamespaceListerExpansion
}

// chaosResultNamespaceLister implements the ChaosResultNamespaceLister
// interface.
type chaosResultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ChaosResults in the indexer for a given namespace.
func (s chaosResultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ChaosResult, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChaosResult))
	})
	return ret, err
}

// Get retrieves the ChaosResult from the indexer for a given namespace and name.
func (s chaosResultNamespaceLister) Get(name string) (*v1alpha1.ChaosResult, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("chaosresult"), name)
	}
	return obj.(*v1alpha1.ChaosResult), nil
}

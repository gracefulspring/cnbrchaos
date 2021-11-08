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
	v1alpha1 "github.com/vossss/cnbrchaos/chaos-operator/pkg/apis/cnbrchaos/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ChaosEngineLister helps list ChaosEngines.
type ChaosEngineLister interface {
	// List lists all ChaosEngines in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ChaosEngine, err error)
	// ChaosEngines returns an object that can list and get ChaosEngines.
	ChaosEngines(namespace string) ChaosEngineNamespaceLister
	ChaosEngineListerExpansion
}

// chaosEngineLister implements the ChaosEngineLister interface.
type chaosEngineLister struct {
	indexer cache.Indexer
}

// NewChaosEngineLister returns a new ChaosEngineLister.
func NewChaosEngineLister(indexer cache.Indexer) ChaosEngineLister {
	return &chaosEngineLister{indexer: indexer}
}

// List lists all ChaosEngines in the indexer.
func (s *chaosEngineLister) List(selector labels.Selector) (ret []*v1alpha1.ChaosEngine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChaosEngine))
	})
	return ret, err
}

// ChaosEngines returns an object that can list and get ChaosEngines.
func (s *chaosEngineLister) ChaosEngines(namespace string) ChaosEngineNamespaceLister {
	return chaosEngineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ChaosEngineNamespaceLister helps list and get ChaosEngines.
type ChaosEngineNamespaceLister interface {
	// List lists all ChaosEngines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ChaosEngine, err error)
	// Get retrieves the ChaosEngine from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ChaosEngine, error)
	ChaosEngineNamespaceListerExpansion
}

// chaosEngineNamespaceLister implements the ChaosEngineNamespaceLister
// interface.
type chaosEngineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ChaosEngines in the indexer for a given namespace.
func (s chaosEngineNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ChaosEngine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChaosEngine))
	})
	return ret, err
}

// Get retrieves the ChaosEngine from the indexer for a given namespace and name.
func (s chaosEngineNamespaceLister) Get(name string) (*v1alpha1.ChaosEngine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("chaosengine"), name)
	}
	return obj.(*v1alpha1.ChaosEngine), nil
}

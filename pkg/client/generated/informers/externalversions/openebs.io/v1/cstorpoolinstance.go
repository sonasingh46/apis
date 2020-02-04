/*
Copyright 2020 The OpenEBS Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	openebsiov1 "github.com/sonasingh46/apis/pkg/apis/openebs.io/v1"
	versioned "github.com/sonasingh46/apis/pkg/client/generated/clientset/versioned"
	internalinterfaces "github.com/sonasingh46/apis/pkg/client/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/sonasingh46/apis/pkg/client/generated/listers/openebs.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CStorPoolInstanceInformer provides access to a shared informer and lister for
// CStorPoolInstances.
type CStorPoolInstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CStorPoolInstanceLister
}

type cStorPoolInstanceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCStorPoolInstanceInformer constructs a new informer for CStorPoolInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCStorPoolInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCStorPoolInstanceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCStorPoolInstanceInformer constructs a new informer for CStorPoolInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCStorPoolInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1().CStorPoolInstances(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1().CStorPoolInstances(namespace).Watch(options)
			},
		},
		&openebsiov1.CStorPoolInstance{},
		resyncPeriod,
		indexers,
	)
}

func (f *cStorPoolInstanceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCStorPoolInstanceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cStorPoolInstanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openebsiov1.CStorPoolInstance{}, f.defaultInformer)
}

func (f *cStorPoolInstanceInformer) Lister() v1.CStorPoolInstanceLister {
	return v1.NewCStorPoolInstanceLister(f.Informer().GetIndexer())
}

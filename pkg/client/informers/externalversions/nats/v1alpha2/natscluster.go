// Copyright 2017-2018 The nats-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	time "time"

	natsv1alpha2 "github.com/nats-io/nats-operator/pkg/apis/nats/v1alpha2"
	versioned "github.com/nats-io/nats-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nats-io/nats-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/nats-io/nats-operator/pkg/client/listers/nats/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NatsClusterInformer provides access to a shared informer and lister for
// NatsClusters.
type NatsClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.NatsClusterLister
}

type natsClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNatsClusterInformer constructs a new informer for NatsCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNatsClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNatsClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNatsClusterInformer constructs a new informer for NatsCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNatsClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NatsV1alpha2().NatsClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NatsV1alpha2().NatsClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&natsv1alpha2.NatsCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *natsClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNatsClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *natsClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&natsv1alpha2.NatsCluster{}, f.defaultInformer)
}

func (f *natsClusterInformer) Lister() v1alpha2.NatsClusterLister {
	return v1alpha2.NewNatsClusterLister(f.Informer().GetIndexer())
}

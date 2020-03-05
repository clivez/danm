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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	danmv1 "github.com/nokia/danm/crd/apis/danm/v1"
	versioned "github.com/nokia/danm/crd/client/clientset/versioned"
	internalinterfaces "github.com/nokia/danm/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/nokia/danm/crd/client/listers/danm/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TenantConfigInformer provides access to a shared informer and lister for
// TenantConfigs.
type TenantConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TenantConfigLister
}

type tenantConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTenantConfigInformer constructs a new informer for TenantConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTenantConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTenantConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTenantConfigInformer constructs a new informer for TenantConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTenantConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DanmV1().TenantConfigs().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DanmV1().TenantConfigs().Watch(options)
			},
		},
		&danmv1.TenantConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *tenantConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTenantConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tenantConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&danmv1.TenantConfig{}, f.defaultInformer)
}

func (f *tenantConfigInformer) Lister() v1.TenantConfigLister {
	return v1.NewTenantConfigLister(f.Informer().GetIndexer())
}

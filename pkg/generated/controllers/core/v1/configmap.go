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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	"github.com/rancher/wrangler/pkg/generic"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

// ConfigMapController interface for managing ConfigMap resources.
type ConfigMapController interface {
	generic.ControllerMeta
	ConfigMapClient

	// OnChange runs the given handler when the controller detects a resource was changed.
	OnChange(ctx context.Context, name string, sync ConfigMapHandler)

	// OnRemove runs the given handler when the controller detects a resource was changed.
	OnRemove(ctx context.Context, name string, sync ConfigMapHandler)

	// Enqueue adds the resource with the given name to the worker queue of the controller.
	Enqueue(namespace, name string)

	// EnqueueAfter runs Enqueue after the provided duration.
	EnqueueAfter(namespace, name string, duration time.Duration)

	// Cache returns a cache for the resource type T.
	Cache() ConfigMapCache
}

// ConfigMapClient interface for managing ConfigMap resources in Kubernetes.
type ConfigMapClient interface {
	// Create creates a new object and return the newly created Object or an error.
	Create(*v1.ConfigMap) (*v1.ConfigMap, error)

	// Update updates the object and return the newly updated Object or an error.
	Update(*v1.ConfigMap) (*v1.ConfigMap, error)

	// Delete deletes the Object in the given name.
	Delete(namespace, name string, options *metav1.DeleteOptions) error

	// Get will attempt to retrieve the resource with the specified name.
	Get(namespace, name string, options metav1.GetOptions) (*v1.ConfigMap, error)

	// List will attempt to find multiple resources.
	List(namespace string, opts metav1.ListOptions) (*v1.ConfigMapList, error)

	// Watch will start watching resources.
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)

	// Patch will patch the resource with the matching name.
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ConfigMap, err error)
}

// ConfigMapCache interface for retrieving ConfigMap resources in memory.
type ConfigMapCache interface {
	// Get returns the resources with the specified name from the cache.
	Get(namespace, name string) (*v1.ConfigMap, error)

	// List will attempt to find resources from the Cache.
	List(namespace string, selector labels.Selector) ([]*v1.ConfigMap, error)

	// AddIndexer adds  a new Indexer to the cache with the provided name.
	// If you call this after you already have data in the store, the results are undefined.
	AddIndexer(indexName string, indexer ConfigMapIndexer)

	// GetByIndex returns the stored objects whose set of indexed values
	// for the named index includes the given indexed value.
	GetByIndex(indexName, key string) ([]*v1.ConfigMap, error)
}

// ConfigMapHandler is function for performing any potential modifications to a ConfigMap resource.
type ConfigMapHandler func(string, *v1.ConfigMap) (*v1.ConfigMap, error)

// ConfigMapIndexer computes a set of indexed values for the provided object.
type ConfigMapIndexer func(obj *v1.ConfigMap) ([]string, error)

// ConfigMapGenericController wraps wrangler/pkg/generic.Controller so that the function definitions adhere to ConfigMapController interface.
type ConfigMapGenericController struct {
	generic.ControllerInterface[*v1.ConfigMap, *v1.ConfigMapList]
}

// OnChange runs the given resource handler when the controller detects a resource was changed.
func (c *ConfigMapGenericController) OnChange(ctx context.Context, name string, sync ConfigMapHandler) {
	c.ControllerInterface.OnChange(ctx, name, generic.ObjectHandler[*v1.ConfigMap](sync))
}

// OnRemove runs the given object handler when the controller detects a resource was changed.
func (c *ConfigMapGenericController) OnRemove(ctx context.Context, name string, sync ConfigMapHandler) {
	c.ControllerInterface.OnRemove(ctx, name, generic.ObjectHandler[*v1.ConfigMap](sync))
}

// Cache returns a cache of resources in memory.
func (c *ConfigMapGenericController) Cache() ConfigMapCache {
	return &ConfigMapGenericCache{
		c.ControllerInterface.Cache(),
	}
}

// ConfigMapGenericCache wraps wrangler/pkg/generic.Cache so the function definitions adhere to ConfigMapCache interface.
type ConfigMapGenericCache struct {
	generic.CacheInterface[*v1.ConfigMap]
}

// AddIndexer adds  a new Indexer to the cache with the provided name.
// If you call this after you already have data in the store, the results are undefined.
func (c ConfigMapGenericCache) AddIndexer(indexName string, indexer ConfigMapIndexer) {
	c.CacheInterface.AddIndexer(indexName, generic.Indexer[*v1.ConfigMap](indexer))
}

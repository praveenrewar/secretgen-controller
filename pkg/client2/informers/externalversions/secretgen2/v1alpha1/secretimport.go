// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	secretgen2v1alpha1 "github.com/vmware-tanzu/carvel-secretgen-controller/pkg/apis/secretgen2/v1alpha1"
	versioned "github.com/vmware-tanzu/carvel-secretgen-controller/pkg/client2/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/carvel-secretgen-controller/pkg/client2/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vmware-tanzu/carvel-secretgen-controller/pkg/client2/listers/secretgen2/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecretImportInformer provides access to a shared informer and lister for
// SecretImports.
type SecretImportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SecretImportLister
}

type secretImportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecretImportInformer constructs a new informer for SecretImport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecretImportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecretImportInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecretImportInformer constructs a new informer for SecretImport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecretImportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretgenV1alpha1().SecretImports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretgenV1alpha1().SecretImports(namespace).Watch(context.TODO(), options)
			},
		},
		&secretgen2v1alpha1.SecretImport{},
		resyncPeriod,
		indexers,
	)
}

func (f *secretImportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecretImportInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *secretImportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&secretgen2v1alpha1.SecretImport{}, f.defaultInformer)
}

func (f *secretImportInformer) Lister() v1alpha1.SecretImportLister {
	return v1alpha1.NewSecretImportLister(f.Informer().GetIndexer())
}
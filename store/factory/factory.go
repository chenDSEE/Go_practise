package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

type factoryMap struct {
	sync.RWMutex
	implMap map[string]store.Store
}

var defaultFactoryMap = factoryMap{
	implMap: make(map[string]store.Store),
}

func Register(name string, impl store.Store) {
	if impl == nil {
		panic("Store: not to register a nil store.Store")
	}

	defaultFactoryMap.Lock()
	defer defaultFactoryMap.Unlock()

	if _, dup := defaultFactoryMap.implMap[name]; dup {
		panic("Store: register a dup store.Srote:" + name)
	}

	defaultFactoryMap.implMap[name] = impl
}

func New(name string) (store.Store, error) {
	defaultFactoryMap.RLock()
	defer defaultFactoryMap.RUnlock()

	impl, ok := defaultFactoryMap.implMap[name]
	if !ok {
		return nil, fmt.Errorf("Store: unknow store.Store[%s]", name)
	}

	return impl, nil
}
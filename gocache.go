package gocache

import (
	"errors"
	"sync"
)

type Gocache struct {
	mx         sync.RWMutex
	cacheItems map[string]CacheItems
}

type CacheItems struct {
	Value interface{}
}

// constructor
func New() *Gocache {
	return &Gocache{
		cacheItems: make(map[string]CacheItems),
	}
}

func (c *Gocache) Get(key string) (interface{}, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	item, ok := c.cacheItems[key]

	if !ok {
		return nil, errors.New("record not found")
	}

	return item.Value, nil
}

func (c *Gocache) Set(key string, val interface{}) {
	c.mx.Lock()
	c.cacheItems[key] = CacheItems{Value: val}
	c.mx.Unlock()
}

func (c *Gocache) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()

	if _, exist := c.cacheItems[key]; !exist {
		return errors.New("record not fund")
	}

	delete(c.cacheItems, key)

	return nil
}

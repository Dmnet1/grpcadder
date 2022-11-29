package cache

import (
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	cache "grpcadder/api/proto"
	"log"
	"sync"
)

type CacheServer struct {
	mx sync.RWMutex
	m  map[string]string
}

func (c *CacheServer) Get(_ context.Context, k *cache.Key) (*cache.Item, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	var a cache.Item

	if c.m[k.Key] != "" {
		a = cache.Item{Key: k.Key, Value: c.m[k.Key]}
		log.Printf("found Value for key %v", c.m[k.Key], a)
	}
	return &a, nil
}

func (c *CacheServer) Set(_ context.Context, i *cache.Item) (*empty.Empty, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	if c.m == nil {
		c.m = make(map[string]string)
		c.m[i.Key] = i.Value
	}

	out := new(empty.Empty)
	return out, nil
}
func (c *CacheServer) Delete(_ context.Context, k *cache.Key) (*empty.Empty, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	delete(c.m, k.Key)
	out := new(empty.Empty)
	return out, nil
}

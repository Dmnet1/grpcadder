package cache

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	cache "grpcadder/api/proto"
	"sync"
)

type CacheServer struct {
	mx sync.RWMutex
	m  map[string]string
}

func (c *CacheServer) Get(_ context.Context, in *cache.Key) (*cache.Item, error) {
	c.mx.RLock()
	_, found := c.m[in.Key]
	if !found {
		c.mx.RUnlock()
		err := errors.New("the cache has no values for the given key")
		return nil, err
	}

	getCache := cache.Item{Key: in.Key, Value: c.m[in.Key]}

	c.mx.RUnlock()
	return &getCache, nil
}

func (c *CacheServer) Set(_ context.Context, in *cache.Item) (*empty.Empty, error) {
	if c.m == nil {
		c.m = make(map[string]string)
	}

	c.m[in.Key] = in.Value
	out := new(empty.Empty)
	//var out *empty.Empty //Равна ли эта строка предыдущей?
	return out, nil

}

func (c *CacheServer) Delete(_ context.Context, in *cache.Key) (*empty.Empty, error) {
	delete(c.m, in.Key)
	out := new(empty.Empty)
	return out, nil
}

/*type CacheServer struct {
  memCache map[string]string
}

func (g *CacheServer) Get(ctx context.Context, in *cache.Key) (*cache.Item, error) {

  _, ok := g.memCache[in.Key]
  if ok != true {
    err := errors.New("the cache has no values for the given key")
    return nil, err
  }
  a := cache.Item{Key: in.Key, Value: g.memCache[in.Key]}
  return &a, nil
}

func (g *CacheServer) Set(ctx context.Context, in *cache.Item) (*empty.Empty, error) {
  if g.memCache == nil {
    g.memCache = make(map[string]string)
  }
  g.memCache[in.Key] = in.Value
  out := new(empty.Empty)
  return out, nil
}

func (g *CacheServer) Delete(ctx context.Context, in *cache.Key) (*empty.Empty, error) {
  delete(g.memCache, in.Key)
  out := new(empty.Empty)
  return out, nil
}*/

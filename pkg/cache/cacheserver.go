package cache

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	cache "grpcadder/api/proto"
	"sync"
)

type CacheServer struct {
	mx sync.RWMutex
	m  map[string]string
}

func (c *CacheServer) Get(ctx context.Context, key *cache.Key) (*cache.Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (c *CacheServer) Set(ctx context.Context, item *cache.Item) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (c *CacheServer) Delete(ctx context.Context, key *cache.Key) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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

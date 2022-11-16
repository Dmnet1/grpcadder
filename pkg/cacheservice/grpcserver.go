package cacheservice

import (
	"golang.org/x/net/context"
	"grpcadder/pkg/api/protobuf/cache"
)

type GRPCServer struct {
}

func (s *GRPCServer) Get(context.Context, *cache.Key) (*cache.Item, error) {

	return &cache.Item{
		Key:   "",
		Value: "",
	}, nil
}

func (s *GRPCServer) Set(context.Context, *cache.Item) (*cache.Item, error) {
	return &cache.Item{
		Key:   "",
		Value: "",
	}, nil
}

func (s *GRPCServer) Delete(context.Context, *cache.Key) (*cache.Item, error) {
	return &cache.Item{
		Key:   "",
		Value: "",
	}, nil
}

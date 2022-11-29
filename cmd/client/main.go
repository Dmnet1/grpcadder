package main

import (
	"flag"
	"google.golang.org/grpc"
	cache "grpcadder/api/proto"
	"log"
)

агтс

func main() {
	flag.Parse()
	if flag.NArg() < 5 {
		log.Panic("not enough arguments")
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := cache.NewCacheClient(conn)
	c.Set()
	c.Get()
	c.Delete()
}

package main

import (
	"google.golang.org/grpc"
	cache "grpcadder/api/proto"
	cache2 "grpcadder/pkg/cache"

	"log"
	"net"
)

func main() {
	s := grpc.NewServer()             // С помощью пакета grpc вызываем метод NewServer.
	srv := &cache2.CacheServer{}      // Переменная со структурой, которая реализует интерфейс сервера.
	cache.RegisterCacheServer(s, srv) // Регистрируем с помощью метода, который создан через protocol bufers.

	l, err := net.Listen("tcp", ":8080") // Создаем слушателя
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil { // Вызываем из сервера метод Serve, который принимает в качестве аргумента слушателя.
		log.Fatal(err)
	}

}

package main

import (
	"google.golang.org/grpc"
	"grpcadder/pkg/api/protobuf/cache"
	"grpcadder/pkg/cacheservice"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()             // С помощью пакета grpc вызываем метод NewServer.
	srv := &cacheservice.GRPCServer{} // Переменная со структурой, которая реализует интерфейс сервера.
	cache.RegisterCacheServer(s, srv) // Регистрируем с помощью метода, который создан через protocol bufers.

	l, err := net.Listen("tcp", ":8080") // Создаем слушателя
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil { // Вызываем из сервера метод Serve, который принимает в качестве аргумента слушателя.
		log.Fatal(err)
	}

}

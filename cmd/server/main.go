package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	cache "grpcadder/api/proto"
	cache2 "grpcadder/pkg/cache"
	"os"

	"log"
	"net"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	s := grpc.NewServer()             // С помощью пакета grpc вызываем метод NewServer.
	srv := &cache2.CacheServer{}      // Переменная со структурой, которая реализует интерфейс сервера.
	cache.RegisterCacheServer(s, srv) // Регистрируем с помощью метода, который создан через protocol bufers.

	l, err := net.Listen("tcp", os.Getenv("PORT")) // Создаем слушателя
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil { // Вызываем из сервера метод Serve, который принимает в качестве аргумента слушателя.
		log.Fatal(err)
	}

}

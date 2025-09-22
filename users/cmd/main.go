package main

import (
	"fmt"
	"log/slog"
	"msa_big_tech/users/internal/transport"
	users "msa_big_tech/users/pkg/proto/v1"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Инициализация контейнера зависимостей
	imp := transport.NewImplementation()

	// Создание нового gRPC-сервера
	grpcSrv := grpc.NewServer()

	// Регистрация сервиса регистрации V1 в gRPC-сервере
	users.RegisterUsersServiceServer(grpcSrv, imp)

	reflection.Register(grpcSrv)

	// Создание TCP-слушателя на порту, указанном в конфигурации
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50055"))
	if err != nil {
		slog.Error("failed to listen: %v", err)
		os.Exit(1)
	}
	slog.Info("Listening on port 50055")

	// Запуск gRPC-сервера для обработки входящих запросов
	err = grpcSrv.Serve(lis)
	if err != nil {
		slog.Error("failed to serve: %v", err)
		os.Exit(1)
	}

}

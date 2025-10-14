package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	container "msa_big_tech/auth/internal/di"
	auth "msa_big_tech/auth/pkg/proto/v1"

	"google.golang.org/grpc/reflection"
)

func main() {
	// Строим DI контейнер
	cont, err := container.BuildContainer()
	if err != nil {
		log.Fatalf("Failed to build container: %v", err)
	}

	// Получаем зависимости
	err = cont.Invoke(func(deps container.Dependencies) error {
		// Регистрация сервиса
		auth.RegisterAuthServiceServer(deps.GRPCServer, deps.AuthImpl)
		reflection.Register(deps.GRPCServer)

		// Создание TCP-слушателя
		lis, err := net.Listen("tcp", ":50054")
		if err != nil {
			return fmt.Errorf("failed to listen: %v", err)
		}

		slog.Info("Listening on port 50054")

		// Запуск сервера
		return deps.GRPCServer.Serve(lis)
	})

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

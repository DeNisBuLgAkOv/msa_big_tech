package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	delivery_grpc "msa_big_tech/chat/internal/delivery/grpc"
	"msa_big_tech/chat/internal/middleware"
	repo_chat "msa_big_tech/chat/internal/repository/postgress/chat"
	repo_message "msa_big_tech/chat/internal/repository/postgress/message"
	"msa_big_tech/chat/internal/usecase"
	chat "msa_big_tech/chat/pkg/proto/v1"

	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Инициализация контейнера зависимостей
	db := &sql.DB{}
	repoChat := repo_chat.NewChatRepository(db)
	repoMessage := repo_message.NewMessageRepository(db)
	srv := usecase.NewUseCase(*repoChat, *repoMessage)
	imp := delivery_grpc.NewImplementation(srv)

	// Создание нового gRPC-сервера
	grpcSrv := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthInterceptor()))

	// Регистрация сервиса регистрации V1 в gRPC-сервере
	chat.RegisterChatServiceServer(grpcSrv, imp)

	reflection.Register(grpcSrv)

	// Создание TCP-слушателя на порту, указанном в конфигурации
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50057"))
	if err != nil {
		slog.Error("failed to listen: %v", err)
		os.Exit(1)
	}
	slog.Info("Listening on port 50057")

	// Запуск gRPC-сервера для обработки входящих запросов
	err = grpcSrv.Serve(lis)
	if err != nil {
		slog.Error("failed to serve: %v", err)
		os.Exit(1)
	}

}

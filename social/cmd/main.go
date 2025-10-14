package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	social_grpc "msa_big_tech/social/internal/delivery/grpc"
	"msa_big_tech/social/internal/middleware"
	repos_friends "msa_big_tech/social/internal/repository/postgress/friends"
	repo_request "msa_big_tech/social/internal/repository/postgress/request"
	"msa_big_tech/social/internal/usecase"
	social "msa_big_tech/social/pkg/proto/v1"

	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db := &sql.DB{}
	repoRequest := repo_request.NewRequestsRepository(db)
	repoFriend := repos_friends.NewFriendsRepository(db)
	// Инициализация контейнера зависимостей
	srv := usecase.NewUsecase(*repoFriend, *repoRequest)
	imp := social_grpc.NewImplementation(srv)

	// Создание нового gRPC-сервера
	grpcSrv := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthInterceptor()))
	// Регистрация сервиса регистрации V1 в gRPC-сервере
	social.RegisterSocialServiceServer(grpcSrv, imp)

	reflection.Register(grpcSrv)

	// Создание TCP-слушателя на порту, указанном в конфигурации
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50056"))
	if err != nil {
		slog.Error("failed to listen: %v", err)
		os.Exit(1)
	}
	slog.Info("Listening on port 50056")

	// Запуск gRPC-сервера для обработки входящих запросов
	err = grpcSrv.Serve(lis)
	if err != nil {
		slog.Error("failed to serve: %v", err)
		os.Exit(1)
	}

}

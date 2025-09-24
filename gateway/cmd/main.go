package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbAuth "msa_big_tech/auth/pkg/proto/v1"
	pbChat "msa_big_tech/chat/pkg/proto/v1"
	pbSocial "msa_big_tech/social/pkg/proto/v1"
	pbUsers "msa_big_tech/users/pkg/proto/v1"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// mux для REST
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}))

	// Настройка подключения к gRPC-серверу
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрируем обработчики для gRPC-Gateway
	err := pbAuth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth:50054", opts)
	if err != nil {
		log.Fatalf("failed to connect to auth_service: %v", err)
	}
	err = pbChat.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "chat:50057", opts)
	if err != nil {
		log.Fatalf("failed to connect to chat_service: %v", err)
	}
	err = pbSocial.RegisterSocialServiceHandlerFromEndpoint(ctx, mux, "social:50056", opts)
	if err != nil {
		log.Fatalf("failed to connect to social_service: %v", err)
	}
	err = pbUsers.RegisterUsersServiceHandlerFromEndpoint(ctx, mux, "users:50055", opts)
	if err != nil {
		log.Fatalf("failed to connect to users_service: %v", err)
	}

	// Запускаем HTTP-сервер
	log.Println("gRPC-Gateway server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package delivery_grpc

import (
	"msa_big_tech/chat/internal/usecase"
	chat "msa_big_tech/chat/pkg/proto/v1"
)

type Implementation struct {
	chat.UnimplementedChatServiceServer
	srv usecase.ChatUsecase
}

func NewImplementation(srv usecase.ChatUsecase) *Implementation {
	return &Implementation{srv: srv}
}

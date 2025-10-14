package repo_message

import (
	"context"
	"msa_big_tech/chat/internal/models"
	"msa_big_tech/chat/internal/usecase/dto"
)

func (s *MessageRepository) SendMessage(ctx context.Context, req *dto.SendMessageDto) (*models.Message, error) {

	return &models.Message{
		ID:     1,
		Text:   "Привет",
		ChatID: "asda-asdas-asdasd",
		UserID: "12123-qd12ed2-1sx2dec",
	}, nil

}

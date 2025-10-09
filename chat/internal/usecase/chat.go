package usecase

import (
	repo_chat "msa_big_tech/chat/internal/repository/postgress/chat"
	repo_message "msa_big_tech/chat/internal/repository/postgress/message"
)

type UseCase struct {
	repo_chat     repo_chat.ChatRepository
	repos_message repo_message.MessageRepository
}

func NewUseCase(repo_chat repo_chat.ChatRepository, repos_message repo_message.MessageRepository) *UseCase {
	return &UseCase{repo_chat: repo_chat, repos_message: repos_message}
}

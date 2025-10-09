package models

import "errors"

var (
	// Чат не найден
	ErrChatNotFound = errors.New("chat not found")

	// Пользователь уже существует
	ErrChatAlreadyExists = errors.New("chat already exists")
)

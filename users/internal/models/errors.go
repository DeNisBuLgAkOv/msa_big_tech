package models

import "errors"

var (
	// Пользователь не найден
	ErrNotFound = errors.New("profile not found")

	// Пользователь уже существует
	ErrAlreadyExists = errors.New("user profile already exists")
)

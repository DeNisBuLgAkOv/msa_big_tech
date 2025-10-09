package models

import "errors"

var (
	ErrNotFound         = errors.New("user not found")
	ErrAlreadyExists    = errors.New("user already exists")
	ErrPermissionDenied = errors.New("permission denied")
	ErrRequestNotFound  = errors.New("request not found")
	ErrRequestExists    = errors.New("request already exists")
)

package models

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	// Пользователь уже зарегистрирован
	ErrUserAlreadyExists = errors.New("User already exists")
	// Пользователь не существует
	ErrNotExistedUser = errors.New("User does not exist")
	// Пароль неверный
	ErrInvalidPassword = errors.New("Incorrect password")
)

type User struct {
	ID       string
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}

// HashPassword хэширует пароль пользователя и сохраняет его в структуре
func (u *User) HashPassword() error {
	fmt.Println(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Println(string(hashedPassword))
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword проверяет, соответствует ли переданный пароль хэшированному паролю пользователя
func (u *User) CheckPassword(password string) bool {
	fmt.Println("новый", password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

// SetPassword устанавливает новый пароль (без хэширования)
// Обычно используется при создании пользователя перед вызовом HashPassword
func (u *User) SetPassword(password string) {
	u.Password = password
}

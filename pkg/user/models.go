package user

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email       string `json:"email";gorm:"email"`
	Name        string `json:"name";gorm:"name"`
	LastName    string `json:"lastName";gorm:"last_name"`
	Password    string `json:"password"`
	AccessToken string `json:"accessToken";gorm:"access_token"`
}

var (
	ErrUserExists              = errors.New("user_alredy_exists")
	ErrUserOrPasswordIncorrect = errors.New("user_or_password_is_incorrect")
	ErrInvalidToken            = errors.New("invalid_token")
)

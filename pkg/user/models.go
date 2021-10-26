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
	ErrUserNotExist            = errors.New("user_not_exist")
)

type Message struct {
	gorm.Model
	From int64  `gorm:"from"`
	To   int64  `gorm:"to"`
	Date int64  `gorm:"date"`
	Text string `gorm:"text"`
}

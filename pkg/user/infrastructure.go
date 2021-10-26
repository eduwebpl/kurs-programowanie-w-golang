package user

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type UserInfra interface {
	CreateNewUser(user User) error
	UpdateAccessToken(id uint, newAccessToken string) error
	GetUserInfo(accessToken string) (User, error)
	GetUserInfoByID(id uint) (User, error)
	GetUser(email string) (User, error)
	// Chat
	GetUsers(id []uint) ([]User, error)
	CreateMessage(message Message) error
	GetMessages(fromUserID uint, toUserID uint) ([]Message, error)
}

func DefaultUserInfra(db *gorm.DB) UserInfra {
	return &userInfra{
		db,
	}
}

type userInfra struct {
	db *gorm.DB
}

func (u *userInfra) CreateNewUser(user User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userInfra) UpdateAccessToken(id uint, newAccessToken string) error {
	result := u.db.Model(&User{}).Where("id = ?", id).Update("access_token", newAccessToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userInfra) GetUserInfo(accessToken string) (User, error) {
	user := User{}
	result := u.db.First(&user, "access_token = ?", accessToken)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (u *userInfra) GetUserInfoByID(id uint) (User, error) {
	user := User{}
	result := u.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (u *userInfra) GetUser(email string) (User, error) {
	user := User{}
	result := u.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

// Chat
func (u *userInfra) GetUsers(ids []uint) ([]User, error) {
	users := []User{}
	result := u.db.Find(&users, ids)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []User{}, nil
		}
		return []User{}, result.Error
	}

	return users, nil
}

func (u *userInfra) CreateMessage(message Message) error {
	result := u.db.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userInfra) GetMessages(fromUserID uint, toUserID uint) ([]Message, error) {
	messages := []Message{}
	result := u.db.Where(map[string]interface{}{"from": fromUserID, "to": toUserID}).Find(&messages)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []Message{}, nil
		}
		return []Message{}, result.Error
	}

	return messages, nil
}

package user

import (
	"19chat/pkg/helpers"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	CreateAccount(email string, name string, lastName string, password string) error
	Login(email string, password string) (string, error)
	Authorize(accessToken string) (uint, error)
	GetInfo(userId uint) (User, error)
}

func DefaultUserService(infra UserInfra, tokenSecret string) UserService {
	return &userService{
		infra,
		tokenSecret,
	}
}

type userService struct {
	infra       UserInfra
	tokenSecret string
}

func (u *userService) CreateAccount(email string, name string, lastName string, password string) error {
	user, _ := u.infra.GetUser(email)
	if user.Email != "" {
		return ErrUserExists
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return err
	}

	user = User{}
	user.Email = email
	user.Password = hashedPassword
	user.Name = name
	user.LastName = lastName

	err = u.infra.CreateNewUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Login(email string, password string) (string, error) {
	user, err := u.infra.GetUser(email)
	if err != nil {
		return "", err
	}
	if user.Email != email {
		return "", ErrUserOrPasswordIncorrect
	}

	if !helpers.ComparePasswords(password, user.Password) {
		return "", ErrUserOrPasswordIncorrect
	}

	token, err := helpers.GenerateJWTToken(user.ID, u.tokenSecret)
	if err != nil {
		return "", err
	}

	err = u.infra.UpdateAccessToken(user.ID, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userService) Authorize(accessToken string) (uint, error) {
	token, err := helpers.ValidateJWTToken(accessToken, u.tokenSecret)
	if err != nil {
		return 0, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	if !ok || claims.Valid() != nil {
		return 0, ErrInvalidToken
	}

	userID := int(claims["user_id"].(float64))
	fmt.Println(userID)
	if userID == 0 {
		return 0, ErrInvalidToken
	}

	user, err := u.infra.GetUserInfo(accessToken)
	if err != nil {
		return 0, ErrInvalidToken
	}

	if user.AccessToken != accessToken {
		return 0, ErrInvalidToken
	}

	return uint(userID), nil
}

func (u *userService) GetInfo(userId uint) (User, error) {
	user, err := u.infra.GetUserInfoByID(userId)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

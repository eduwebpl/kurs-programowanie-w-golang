package users

import (
	"encoding/json"
	"net/http"
)

const (
	usersEndpoint = "https://api.github.com/users"
	usersFileName = "users.txt"
)

type UsersService interface {
	GetAll() (Users, error)
}

func DefaultUsersService(client *http.Client) UsersService {
	service := &usersService{
		client,
	}
	return service
}

type usersService struct {
	client *http.Client
}

func (e *usersService) GetAll() (Users, error) {
	req, err := http.NewRequest("GET", usersEndpoint, nil)
	if err != nil {
		return Users{}, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Users{}, err
	}

	users := Users{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&users)

	return users, err
}

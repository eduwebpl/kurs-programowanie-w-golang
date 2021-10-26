package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	usersFromCache := e.readUsers()
	if usersFromCache != nil {
		return *usersFromCache, nil
	}

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

	e.saveUsers(users)
	return users, err
}

func (u *usersService) saveUsers(users Users) error {
	encodedUsers, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = os.WriteFile(usersFileName, encodedUsers, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u *usersService) readUsers() *Users {
	encodedUsers, err := os.ReadFile(usersFileName)
	if err != nil {
		return nil
	}
	decodedUsers := Users{}
	err = json.Unmarshal(encodedUsers, &decodedUsers)
	if err != nil {
		return nil
	}

	return &decodedUsers
}

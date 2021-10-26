package main

import (
	"10zapisywaniedopliku/internal/users"
	"fmt"
	"net/http"
)

func main() {

	usersService := users.DefaultUsersService(&http.Client{})

	users, err := usersService.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

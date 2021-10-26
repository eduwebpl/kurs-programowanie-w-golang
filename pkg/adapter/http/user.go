package http

import (
	"16kontouzytkownika/pkg/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAdapter interface {
	Login(c *gin.Context)
	Create(c *gin.Context)
	GetUserInfo(c *gin.Context)
	AuthorizationMiddleware(c *gin.Context)
}

func DefaultUserAdapter(userService user.UserService) UserAdapter {
	return &userAdapter{
		userService,
	}
}

type userAdapter struct {
	service user.UserService
}

func (u *userAdapter) Login(c *gin.Context) {
	loginUser := LoginUserRequest{}
	err := c.BindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	token, err := u.service.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		if err == user.ErrUserOrPasswordIncorrect {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	response := LoginUserResponse{}
	response.AccessToken = token

	c.JSON(http.StatusOK, response)
}

func (u *userAdapter) Create(c *gin.Context) {
	newUser := CreateUser{}
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = u.service.CreateAccount(newUser.Email, newUser.Name, newUser.LastName, newUser.Password)
	if err != nil {
		if err == user.ErrUserExists {
			c.JSON(http.StatusFound, nil)
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (u *userAdapter) GetUserInfo(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetHeader("user_id"))

	userInfo, err := u.service.GetInfo(uint(userID))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	info := UserInfo{}
	info.Email = userInfo.Email
	info.LastName = userInfo.LastName
	info.Name = userInfo.Name

	c.JSON(http.StatusOK, info)
}

func (u *userAdapter) AuthorizationMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	userID, err := u.service.Authorize(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	c.Request.Header.Set("user_id", strconv.Itoa(int(userID)))

	c.Next()
}

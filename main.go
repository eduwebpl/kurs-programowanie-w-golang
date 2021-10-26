package main

import (
	"16kontouzytkownika/pkg/adapter/http"
	"16kontouzytkownika/pkg/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const accessTokenSecret = "eduweb.pl"

func main() {
	os.Setenv("ACCESS_SECRET", accessTokenSecret)

	db, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal("Failed to open the SQLite database.")
	}

	db.AutoMigrate(&user.User{})

	userInfra := user.DefaultUserInfra(db)
	userService := user.DefaultUserService(userInfra, accessTokenSecret)

	server := gin.Default()

	adapter := http.DefaultUserAdapter(userService)

	userGroup := server.Group("/user")

	{
		userGroup.POST("/create", adapter.Create)
		userGroup.POST("/authorize", adapter.Login)

		authorizedGroup := userGroup.Group("", adapter.AuthorizationMiddleware)
		{
			authorizedGroup.GET("/info", adapter.GetUserInfo)
		}
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

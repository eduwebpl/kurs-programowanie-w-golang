package main

import (
	"17swaggo/pkg/adapter/http"
	"17swaggo/pkg/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	_ "17swaggo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const accessTokenSecret = "eduweb.pl"

// @title srv-user
// @version 0.0.0
// @host localhost:8080
// @BasePath /
// @license.name MIT License
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

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

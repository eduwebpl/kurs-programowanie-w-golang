package main

import (
	"19chat/pkg/adapter/http"
	"19chat/pkg/helpers"
	"19chat/pkg/user"
	"fmt"
	"log"
	"net"
	"os"

	userGrpc "19chat/pkg/adapter/grpc"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "19chat/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	accessTokenSecret = helpers.GetEnv("TOKEN_SECRET", "eduweb.pl")
	grpcPort          = helpers.GetEnv("GRPC_PORT", "8081")
	httpPort          = helpers.GetEnv("HTTP_PORT", "8080")
	postgresdsn       = helpers.GetEnv("POSTGRES_DSN", "")
)

// @title srv-user
// @version 0.0.0
// @host localhost:8080
// @BasePath /
// @license.name MIT License
func main() {
	fmt.Println("Chat is starting...")

	os.Setenv("ACCESS_SECRET", accessTokenSecret)

	db, err := gorm.Open("postgres", postgresdsn)

	if err != nil {
		log.Fatal("Failed to open the database.")
	}

	db.AutoMigrate(&user.User{}, &user.Message{})

	userInfra := user.DefaultUserInfra(db)
	userService := user.DefaultUserService(userInfra, accessTokenSecret)

	go startHTTPServer(userService)

	startGRPCServer(userService)
}

func startHTTPServer(userService user.UserService) {
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

	err := server.Run(":" + httpPort)
	if err != nil {
		log.Fatal(err)
	}
}

func startGRPCServer(userService user.UserService) {

	listener, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to start listening")
	}

	service := userGrpc.DefaultUserService(userService)

	grpcServer := grpc.NewServer()

	userGrpc.RegisterUserServiceServer(grpcServer, service)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Failed to start grpc server.")
	}
}

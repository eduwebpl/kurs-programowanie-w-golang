package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", rootPath)

	authorized := server.Group("/user/", authorizationMiddleware)
	{
		authorized.GET("", getUserInfo)
	}

	server.Run()
}

func rootPath(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"key": "value"})
}

func getUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"key": c.GetHeader("user")})
}

func authorizationMiddleware(c *gin.Context) {
	id := c.GetHeader("user_id")
	if id != "" {
		c.Request.Header.Add("user", "Kamil")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
	c.Next()
}

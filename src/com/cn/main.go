package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	f, _ := os.Create("/home/oufan/gowork/src/com/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.Use(cors())
	gin.SetMode(viper.GetString("mode"))
	user := r.Group("/api/user")
	{
		user.GET("/list", UserHandler.UserListHandler)
		user.GET("/info/:id", UserHandler.UserInfoHandler)
		user.POST("/add", UserHandler.AddUserHandler)
		user.POST("/edit", UserHandler.EditUserHandler)
		user.POST("/delete", UserHandler.DeleteUserHandler)

	}
	port := viper.GetString("port")
	r.Run(port)
}

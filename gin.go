package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-emqx/api"
)

var (
	router = gin.Default()
)

func Run() {

	router.Use(CORSMiddleware())

	getRoutes()

	err := router.Run(":8096")
	if err != nil {
		fmt.Println("Error] failed to start Gin server due to: ", err.Error())
		return
	}
}

func getRoutes() {

	v2 := router.Group("/v2")

	api.AddChatRoute(v2)
	api.AddMusicRoute(v2)
	api.AddDownloadRoutes(v2)

	api.AddUploadRoute(v2)
	api.AddUploadRoute2(v2)
	api.AddMqtt(v2)
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

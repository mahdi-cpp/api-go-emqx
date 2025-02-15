package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-emqx/repository"
	"net/http"
)

func AddMqtt(rg *gin.RouterGroup) {
	route := rg.Group("/mqtt")

	route.GET("/temp", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestChatV2())
	})

}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-emqx/repository"
	"net/http"
)

func AddMusicRoute(rg *gin.RouterGroup) {
	route := rg.Group("/chat")

	route.GET("/music", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.RestMusic())
	})

	route.GET("/subtitle", func(context *gin.Context) {
		repository.ReloadSubtitle()
		context.JSON(http.StatusOK, repository.RestSubtitle())
	})
}

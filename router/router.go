package router

import (
	"ginTemplateDemo/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler.Index)
	router.POST("/result", handler.SearchArticles)
	return router
}

package main

import (
	"log"

	"ginTemplateDemo/handler"
	"ginTemplateDemo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := handler.SetupDB(); err != nil {
		log.Fatalf("failed to connect database. error: %v", err)

	}
	gin.SetMode(gin.ReleaseMode)
	router := router.SetupRouter()
	router.Run(":8080")
	defer handler.CloseDB()
}

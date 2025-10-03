package routes

import (
	"github.com/alberto-debug/enrollmate/internal/handlers"
	"github.com/gin-gonic/gin"
)

func HomeRoutes(router *gin.Engine) {

	router.GET("/", handlers.Home)

}

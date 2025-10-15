package routes

import (
	"github.com/alberto-debug/enrollmate/internal/handlers"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine) {

	router.GET("/students", handlers.GetStudent)
	router.POST("/student", handlers.CreateStudent)

}

package routes

import (
	"github.com/alberto-debug/enrollmate/internal/handlers"
	"github.com/gin-gonic/gin"
)


func StudentRoutes(StudentRouter *gin.Engine){

	StudentRouter.GET("/students", handlers.GetStudents)
	StudentRouter.POST("/student", handlers.CreateStudent)
}


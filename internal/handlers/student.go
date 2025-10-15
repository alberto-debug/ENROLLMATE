package handlers

import (
	"net/http"

	"github.com/alberto-debug/enrollmate/config"
	"github.com/alberto-debug/enrollmate/internal/model"
	"github.com/gin-gonic/gin"
)

// Get /students
func GetStudents(c *gin.Context){
	var student []model.Student
	config.DB.Find(&student)
	c.JSON(http.StatusOK, student)
}

//Post /students


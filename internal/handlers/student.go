package handlers

import (
	"net/http"

	"github.com/alberto-debug/enrollmate/config"
	"github.com/alberto-debug/enrollmate/internal/model"
	"github.com/gin-gonic/gin"
)

// Get /students
func GetStudents(c *gin.Context) {

	var students []model.Student
	config.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

//Post /students
func CreateStudent(c *gin.Context){
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		
	}

	config.DB.Create(student)
	c.JSON(http.StatusCreated, student)
}

package main

import (
	"github.com/alberto-debug/enrollmate/config"
	"github.com/alberto-debug/enrollmate/internal/model"
	"github.com/alberto-debug/enrollmate/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//Connect Database and configure the auto migration of the scheme
	config.ConnectDb()
	config.DB.AutoMigrate(&model.Student{})

	// Setup router
	router := gin.Default()

	//Register the Routes
	routes.HomeRoutes(router)

	// Start server
	router.Run()
}

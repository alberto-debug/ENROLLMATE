package main

import (
	"fmt"
	"log"

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
	port := ":8080"
	fmt.Printf("🚀 Starting EnrollMate server on port %s...\n", port)
	fmt.Printf("🌐 Server will be available at: http://localhost%s\n", port)

	if err := router.Run(port); err != nil {
		log.Fatal("❌ Server failed to start: ", err)
	}

	//router.Run()
}

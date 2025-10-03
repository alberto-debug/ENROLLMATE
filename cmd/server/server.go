package main

import (
	"github.com/alberto-debug/enrollmate/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Setup router
	router := gin.Default()

	//Register the Routes
	routes.HomeRoutes(router)

	// Start server
	router.Run()
}

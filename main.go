package main

import (
	"log"

	"example.com/event-booking/db"
	"example.com/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	routes.RegisterRoutes(router)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

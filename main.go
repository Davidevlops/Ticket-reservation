package main

import (
	"log"
	"os"

	"go-backend-app/internal/connection"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables (optional in production)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database connection
	dbConn, err := connection.DatabasePoolConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	// Initialize Gin
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(func(c *gin.Context) {
		c.Set("db", dbConn)
		c.Next()
	})

	log.Printf("Server listening on port %s\n", port)

	// Start HTTP server (blocking)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("HTTP server failed:", err)
	}
}

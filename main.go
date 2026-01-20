package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-backend-app/internal/connection"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables (optional in production)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database connection
	dbConn, err := connection.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! Nice meeting you on %s", r.Host)
	})

	// Server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	log.Printf("Server listening on port %s\n", port)

	// Start HTTP server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("HTTP server failed:", err)
	}
}

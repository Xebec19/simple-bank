package main

import (
	"context"
	"log"
	"os"

	"github.com/Xebec19/simple-bank/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Create context
	ctx := context.Background()

	// Initialize server
	srv, err := server.NewServer(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	// Get server address from environment or use default
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	log.Printf("Starting server on %s", address)
	if err := srv.Start(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

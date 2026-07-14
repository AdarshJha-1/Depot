package main

import (
	"fmt"
	"log"

	"github.com/AdarshJha-1/Depot/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system environment")
	}
	server := server.New()
	fmt.Println("Server is running...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to run server", err)
	}
}

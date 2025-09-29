package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		// Log a fatal error if the .env file cannot be loaded
		// NOTE: If you deploy your app in a production environment
		// where variables are set by the OS/container, you might
		// want to handle this error differently (e.g., just log a warning).
		log.Fatalf("Error loading .env file: %v", err)
	}
	Host := os.Getenv("DB_HOST")
	fmt.Println(Host)
}

package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("We are loading %s and couldn't find it. Error loading .env file", key)
	}
	return os.Getenv(key)
}

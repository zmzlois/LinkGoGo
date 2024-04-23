package utils

import (
	"os"
)

// Config func to get env value
func Config(key string) string {
	// load .env file

	return os.Getenv(key)
}

package utils

import "os"

// GetPort returns port from environment variable "PORT", or default: 8080
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

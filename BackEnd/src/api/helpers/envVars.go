package envVars

import (
	"log"
	
	"github.com/joho/godotenv"
)

// Load environment vars from file
func GetEnvVars() {
	err := godotenv.Load("./config/cred.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
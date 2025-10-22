package config

import (
	"log"
	"os"
)

// EnvConfig stores the essential environment variables for the project
type EnvConfig struct {
	Port        string
	OpenAIKey   string
	OpenAIModel string
}

// LoadEnv loads and validates environment variables
func LoadEnv() *EnvConfig {
	return &EnvConfig{
		Port:        getPort(),
		OpenAIKey:   getRequiredEnv("OPENAI_API_KEY"),
		OpenAIModel: getEnvOrDefault("OPENAI_MODEL", "gpt-4"),
	}
}

// getPort returns the server port, defaulting to 8080 if not set
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// getRequiredEnv returns the value of a required environment variable
// or exits the program if it is not set
func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is not set in environment", key)
	}
	return value
}

// getEnvOrDefault returns the value of the environment variable or a default
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

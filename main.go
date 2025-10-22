package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jm-borges/llm-voice-agent-api/config"
	"github.com/jm-borges/llm-voice-agent-api/handlers"
	"github.com/joho/godotenv"
)

// main is the entry point of the LLM Voice Agent API server.
func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	env := config.LoadEnv()

	// Register HTTP routes
	registerRoutes()

	// Start HTTP server
	startServer(env.Port)
}

// registerRoutes sets up the HTTP routes and handlers.
func registerRoutes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/query", handlers.HandleQuery)
}

// rootHandler responds with a simple status message to indicate the server is running.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LLM Voice Agent API is running!")
}

// startServer starts the HTTP server on the specified port and logs any errors.
func startServer(port string) {
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

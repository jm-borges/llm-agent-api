package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Carrega .env se existir
    godotenv.Load()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "LLM Voice Agent API is running!")
    })

    fmt.Printf("Server listening on :%s\n", port)
    http.ListenAndServe(":"+port, nil)
}

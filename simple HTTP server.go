package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

// Response is a struct for JSON response
type Response struct {
    Message string `json:"message"`
    Time    string `json:"time"`
}

// handleRoot handles the root URL
func handleRoot(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "Welcome to the Go HTTP server!",
        Time:    time.Now().Format(time.RFC3339),
    }
    sendJSONResponse(w, http.StatusOK, response)
}

// handleGreet handles the /greet URL
func handleGreet(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }

    response := Response{
        Message: fmt.Sprintf("Hello, %s!", name),
        Time:    time.Now().Format(time.RFC3339),
    }
    sendJSONResponse(w, http.StatusOK, response)
}

// sendJSONResponse sends a JSON response with a given status code
func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func main() {
    // Define routes
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/greet", handleGreet)

    // Start the server
    port := ":8080"
    log.Printf("Starting server on port %s...", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}

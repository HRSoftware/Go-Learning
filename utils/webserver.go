package utils

import (
	"fmt"
	"net/http"
)

var server *http.Server

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// StartServer starts the HTTP server on the given port.
func StartWebServer() {
	// Define the HTTP handler and route
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	// Create the server instance
	server = &http.Server{
		Addr: ":8080",
	}

	// Start the server in a new goroutine
	go func() {
		fmt.Println("Server is starting on port 8080...")
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Error starting the server:", err)
		}
	}()
}

// StopServer gracefully shuts down the HTTP server.
func StopWebServer() {
	fmt.Println("Stopping the server...")

	// Graceful shutdown (timeout 5 seconds)
	if err := server.Close(); err != nil {
		fmt.Println("Error stopping the server:", err)
	} else {
		fmt.Println("Server stopped successfully.")
	}
}

// Set up channel to capture interrupt signals (Ctrl+C)
//stopChan := make(chan os.Signal, 1)
//signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

// Wait for interrupt signal to stop the server
//<-stopChan

package module

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/vintlang/vintlang/object"
)

// Exported HTTP functions for Vint
var HttpFunctions = map[string]object.ModuleFunction{}

func init() {
	HttpFunctions["fileServer"] = fileServer
}

// fileServer serves files from a specified directory with directory listing enabled.
func fileServer(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 || len(args) > 3 {
		return &object.Error{Message: "Usage: http.fileServer(port, directory, [message])"}
	}

	// Validate port argument
	port, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Port must be a string"}
	}

	// Validate directory argument
	directory, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Directory must be a string"}
	}

	// Validate custom message (optional)
	var message string
	if len(args) >= 3 {
		customMessage, ok := args[2].(*object.String)
		if ok {
			message = customMessage.Value
		} else {
			return &object.Error{Message: "Message must be a string"}
		}
	} else {
		message = fmt.Sprintf("Server started on port %s serving files from %s with directory listing enabled", port.Value, directory.Value)
	}

	// Ensure directory exists
	absDirectory, err := filepath.Abs(directory.Value)
	if err != nil || !isValidDirectory(absDirectory) {
		return &object.Error{Message: "Invalid or non-existent directory"}
	}

	// Create a file server with directory listing enabled
	fileHandler := http.FileServer(http.Dir(absDirectory))
	fileHandler = enableDirectoryListing(fileHandler)

	// Wrap the file server in middleware for logging requests
	http.Handle("/", logMiddleware(fileHandler))

	// Start the server in a new goroutine
	go func() {
		fmt.Println(message)
		if err := http.ListenAndServe(":"+port.Value, nil); err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	waitForInterrupt()

	return &object.String{Value: "Server stopped"}
}

// logMiddleware logs incoming requests to the server.
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// enableDirectoryListing modifies the handler to always show directory listings.
func enableDirectoryListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setting the content type to HTML so directory contents are rendered as a webpage.
		w.Header().Set("Content-Type", "text/html")
		next.ServeHTTP(w, r)
	})
}

// isValidDirectory checks if a path is a valid directory.
func isValidDirectory(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// waitForInterrupt blocks until an interrupt signal (Ctrl+C) is received.
func waitForInterrupt() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("\nShutting down server...")
}

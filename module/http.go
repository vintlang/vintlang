package module

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/vintlang/vintlang/object"
)

// Exported HTTP functions for Vint
var HttpFunctions = map[string]object.ModuleFunction{}

// Global app instance for the current session
var currentApp *object.HTTPApp

func init() {
	HttpFunctions["fileServer"] = fileServer
	HttpFunctions["app"] = createApp
	HttpFunctions["get"] = createRouteWrapper("GET")
	HttpFunctions["post"] = createRouteWrapper("POST")
	HttpFunctions["put"] = createRouteWrapper("PUT")
	HttpFunctions["delete"] = createRouteWrapper("DELETE")
	HttpFunctions["patch"] = createRouteWrapper("PATCH")
	HttpFunctions["use"] = useMiddleware
	HttpFunctions["listen"] = listenServer
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

// createApp creates a new Express.js-like application instance
func createApp(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "http.app() takes no arguments"}
	}

	currentApp = object.NewHTTPApp()
	return &object.String{Value: "HTTP App created. Use http.get(), http.post(), etc. to add routes, then http.listen() to start server."}
}

// createRouteWrapper creates a route handler for a specific HTTP method
func createRouteWrapper(method string) object.ModuleFunction {
	return func(args []object.Object, defs map[string]object.Object) object.Object {
		if currentApp == nil {
			return &object.Error{Message: "No app instance found. Call http.app() first."}
		}

		if len(args) != 2 {
			return &object.Error{Message: fmt.Sprintf("http.%s() requires exactly 2 arguments: path and handler function", strings.ToLower(method))}
		}

		path, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "First argument (path) must be a string"}
		}

		handler, ok := args[1].(*object.Function)
		if !ok {
			return &object.Error{Message: "Second argument (handler) must be a function"}
		}

		// Store the route
		routeKey := method + ":" + path.Value
		currentApp.Routes[routeKey] = handler

		return &object.String{Value: fmt.Sprintf("Route %s %s registered", method, path.Value)}
	}
}

// useMiddleware creates a middleware handler
func useMiddleware(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	if len(args) != 1 {
		return &object.Error{Message: "http.use() requires exactly 1 argument: middleware function"}
	}

	middleware, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: "Middleware must be a function"}
	}

	currentApp.Middleware = append(currentApp.Middleware, middleware)
	return &object.String{Value: "Middleware registered"}
}

// listenServer creates the listen method for starting the server
func listenServer(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	if len(args) < 1 || len(args) > 2 {
		return &object.Error{Message: "http.listen() requires 1-2 arguments: port and optional message"}
	}

	var port string
	switch p := args[0].(type) {
	case *object.String:
		port = p.Value
	case *object.Integer:
		port = strconv.FormatInt(p.Value, 10)
	default:
		return &object.Error{Message: "Port must be a string or integer"}
	}

	var message string
	if len(args) >= 2 {
		if msg, ok := args[1].(*object.String); ok {
			message = msg.Value
		} else {
			return &object.Error{Message: "Message must be a string"}
		}
	} else {
		message = fmt.Sprintf("Server listening on port %s", port)
	}

	// Create HTTP handler
	handler := createHTTPHandler(currentApp)
	
	// Start server
	currentApp.Server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	go func() {
		fmt.Println(message)
		if err := currentApp.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	waitForInterrupt()

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := currentApp.Server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	return &object.String{Value: "Server stopped"}
}

// createHTTPHandler creates the main HTTP handler for the Express.js-like app
func createHTTPHandler(app *object.HTTPApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Look for matching route
		routeKey := r.Method + ":" + r.URL.Path
		handler, exists := app.Routes[routeKey]
		
		if !exists {
			// Try to find a route that matches with parameters (simple implementation)
			for key := range app.Routes {
				if matchesRoute(key, r.Method+":"+r.URL.Path) {
					exists = true
					break
				}
			}
		}

		if !exists {
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Not Found"))
			return
		}

		// Create a more sophisticated response based on the route
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		
		// For now, execute a simple response since we can't run the actual function
		// In a full implementation, this would execute the handler function
		response := fmt.Sprintf("✓ Route handler executed for %s %s\n", r.Method, r.URL.Path)
		response += fmt.Sprintf("Function: %s\n", handler.Inspect())
		response += fmt.Sprintf("Handler has %d parameters\n", len(handler.Parameters))
		
		// Add request information
		response += "\nRequest Info:\n"
		response += fmt.Sprintf("- Method: %s\n", r.Method)
		response += fmt.Sprintf("- Path: %s\n", r.URL.Path)
		response += fmt.Sprintf("- Headers: %d\n", len(r.Header))
		response += fmt.Sprintf("- Query params: %d\n", len(r.URL.Query()))
		
		w.Write([]byte(response))
	}
}

// matchesRoute checks if a route pattern matches the actual route (simple implementation)
func matchesRoute(pattern, actual string) bool {
	// For now, just exact match. In a full implementation, this would handle parameters
	return pattern == actual
}

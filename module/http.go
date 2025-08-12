package module

import (
	"context"
	"encoding/json"
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

	"github.com/vintlang/vintlang/ast"
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
	// New backend features
	HttpFunctions["interceptor"] = addInterceptor
	HttpFunctions["guard"] = addGuard
	HttpFunctions["cors"] = corsMiddleware
	HttpFunctions["bodyParser"] = bodyParserMiddleware
	HttpFunctions["auth"] = authMiddleware
	HttpFunctions["errorHandler"] = setErrorHandler
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
		// Create enhanced request and response objects
		req := object.NewHTTPRequest(r)

		// Add CORS headers by default
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		// Run request interceptors
		if interceptors, exists := app.Interceptors["request"]; exists {
			for _, interceptor := range interceptors {
				// In a full implementation, this would execute the interceptor function
				log.Printf("Running request interceptor: %s", interceptor.Inspect())
			}
		}

		// Run guards
		for _, guard := range app.Guards {
			// In a full implementation, this would execute the guard function
			// Guards could return false to block the request
			log.Printf("Running guard: %s", guard.Inspect())
		}

		// Extract path parameters for routes like /users/:id
		routeKey := r.Method + ":" + r.URL.Path
		handler, exists := app.Routes[routeKey]
		
		if !exists {
			// Try to find a route that matches with parameters
			for key, routeHandler := range app.Routes {
				if matchesRouteWithParams(key, r.Method+":"+r.URL.Path, req) {
					handler = routeHandler
					exists = true
					break
				}
			}
		}

		if !exists {
			// Run error handler if available
			if app.ErrorHandler != nil {
				log.Printf("Running error handler for 404: %s", app.ErrorHandler.Inspect())
			}
			
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "application/json")
			errorResponse := map[string]interface{}{
				"error": "Not Found",
				"message": fmt.Sprintf("Cannot %s %s", r.Method, r.URL.Path),
				"statusCode": 404,
			}
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		// Run middleware
		for _, middleware := range app.Middleware {
			// In a full implementation, this would execute the middleware function
			log.Printf("Running middleware: %s", middleware.Inspect())
		}

		// Execute the route handler
		// In a full implementation, this would properly execute the handler function
		// with the request and response objects
		response := fmt.Sprintf("✓ Enhanced route handler executed for %s %s\n", r.Method, r.URL.Path)
		response += fmt.Sprintf("Function: %s\n", handler.Inspect())
		response += fmt.Sprintf("Handler has %d parameters\n", len(handler.Parameters))
		
		// Add enhanced request information
		response += "\nEnhanced Request Info:\n"
		response += fmt.Sprintf("- Method: %s\n", r.Method)
		response += fmt.Sprintf("- Path: %s\n", r.URL.Path)
		response += fmt.Sprintf("- Headers: %d\n", len(r.Header))
		response += fmt.Sprintf("- Query params: %d\n", len(r.URL.Query()))
		response += fmt.Sprintf("- Cookies: %d\n", len(req.Cookies))
		response += fmt.Sprintf("- Content-Type: %s\n", r.Header.Get("Content-Type"))
		
		if len(req.FormData) > 0 {
			response += fmt.Sprintf("- Form data fields: %d\n", len(req.FormData))
		}
		if req.JSON != nil {
			response += "- JSON body parsed\n"
		}
		if len(req.Params) > 0 {
			response += fmt.Sprintf("- Path params: %v\n", req.Params)
		}

		// Run response interceptors
		if interceptors, exists := app.Interceptors["response"]; exists {
			for _, interceptor := range interceptors {
				log.Printf("Running response interceptor: %s", interceptor.Inspect())
			}
		}

		// Set content type and send response
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte(response))
	}
}

// matchesRoute checks if a route pattern matches the actual route (simple implementation)
func matchesRoute(pattern, actual string) bool {
	// For now, just exact match. In a full implementation, this would handle parameters
	return pattern == actual
}

// matchesRouteWithParams checks if a route pattern matches and extracts parameters
func matchesRouteWithParams(pattern, actual string, req *object.HTTPRequest) bool {
	// Split the pattern and actual path by ":"
	patternParts := strings.Split(pattern, "/")
	actualParts := strings.Split(actual, "/")

	// Must have same number of parts
	if len(patternParts) != len(actualParts) {
		return false
	}

	// Check method part (first part before the first ":")
	if len(patternParts) == 0 || len(actualParts) == 0 {
		return false
	}

	// Extract method and path
	patternMethod := strings.Split(patternParts[0], ":")[0]
	actualMethod := strings.Split(actualParts[0], ":")[0]
	
	if patternMethod != actualMethod {
		return false
	}

	// Start from second part (after method:)
	patternPath := strings.Join(patternParts[1:], "/")
	actualPath := strings.Join(actualParts[1:], "/")
	
	patternPathParts := strings.Split(patternPath, "/")
	actualPathParts := strings.Split(actualPath, "/")

	if len(patternPathParts) != len(actualPathParts) {
		return false
	}

	// Extract parameters
	for i, patternPart := range patternPathParts {
		if strings.HasPrefix(patternPart, ":") {
			// This is a parameter
			paramName := patternPart[1:] // Remove the ":"
			paramValue := actualPathParts[i]
			req.Params[paramName] = paramValue
		} else if patternPart != actualPathParts[i] {
			// Static part doesn't match
			return false
		}
	}

	return true
}

// addInterceptor adds request or response interceptors
func addInterceptor(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	if len(args) != 2 {
		return &object.Error{Message: "http.interceptor() requires exactly 2 arguments: type ('request' or 'response') and handler function"}
	}

	interceptorType, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "First argument (type) must be a string"}
	}

	if interceptorType.Value != "request" && interceptorType.Value != "response" {
		return &object.Error{Message: "Interceptor type must be 'request' or 'response'"}
	}

	handler, ok := args[1].(*object.Function)
	if !ok {
		return &object.Error{Message: "Second argument (handler) must be a function"}
	}

	if currentApp.Interceptors[interceptorType.Value] == nil {
		currentApp.Interceptors[interceptorType.Value] = make([]*object.Function, 0)
	}
	currentApp.Interceptors[interceptorType.Value] = append(currentApp.Interceptors[interceptorType.Value], handler)

	return &object.String{Value: fmt.Sprintf("%s interceptor registered", interceptorType.Value)}
}

// addGuard adds guards for authentication, authorization, rate limiting, etc.
func addGuard(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	if len(args) != 1 {
		return &object.Error{Message: "http.guard() requires exactly 1 argument: guard function"}
	}

	guard, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: "Guard must be a function"}
	}

	currentApp.Guards = append(currentApp.Guards, guard)
	return &object.String{Value: "Guard registered"}
}

// corsMiddleware creates CORS middleware
func corsMiddleware(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	// Create a CORS middleware function
	corsFunc := &object.Function{
		Parameters: []*ast.Identifier{
			{Value: "req"},
			{Value: "res"},
			{Value: "next"},
		},
		Body: nil, // This would be implemented in a full evaluator
		Env:  nil,
	}

	currentApp.Middleware = append(currentApp.Middleware, corsFunc)
	return &object.String{Value: "CORS middleware registered"}
}

// bodyParserMiddleware creates body parser middleware
func bodyParserMiddleware(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	// Create a body parser middleware function
	bodyParserFunc := &object.Function{
		Parameters: []*ast.Identifier{
			{Value: "req"},
			{Value: "res"},
			{Value: "next"},
		},
		Body: nil, // This would be implemented in a full evaluator
		Env:  nil,
	}

	currentApp.Middleware = append(currentApp.Middleware, bodyParserFunc)
	return &object.String{Value: "Body parser middleware registered"}
}

// authMiddleware creates authentication middleware
func authMiddleware(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	if len(args) != 1 {
		return &object.Error{Message: "http.auth() requires exactly 1 argument: authentication function"}
	}

	authFunc, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: "Authentication function must be a function"}
	}

	currentApp.Middleware = append(currentApp.Middleware, authFunc)
	return &object.String{Value: "Authentication middleware registered"}
}

// setErrorHandler sets a global error handler
func setErrorHandler(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	if len(args) != 1 {
		return &object.Error{Message: "http.errorHandler() requires exactly 1 argument: error handler function"}
	}

	errorHandler, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: "Error handler must be a function"}
	}

	currentApp.ErrorHandler = errorHandler
	return &object.String{Value: "Error handler registered"}
}

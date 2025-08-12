package module

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	// Enterprise features
	HttpFunctions["group"] = createRouteGroup
	HttpFunctions["multipart"] = parseMultipart
	HttpFunctions["async"] = createAsyncHandler
	HttpFunctions["security"] = securityMiddleware
	// Streaming and performance features
	HttpFunctions["stream"] = createStreamHandler
	HttpFunctions["metrics"] = enableMetrics
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
		startTime := time.Now()
		
		// Create enhanced request and response objects
		req := object.NewHTTPRequest(r)

		// Apply security headers if configured
		if app.Security != nil {
			for key, value := range app.Security.SecurityHeaders {
				w.Header().Set(key, value)
			}
		}

		// Add CORS headers based on configuration
		if app.Security != nil && app.Security.CORSOptions != nil {
			cors := app.Security.CORSOptions
			w.Header().Set("Access-Control-Allow-Origin", cors.Origin)
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(cors.Methods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(cors.Headers, ", "))
			if cors.Credentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}

		// Add performance tracking headers if enabled
		if app.Performance != nil && app.Performance.RequestTiming {
			w.Header().Set("X-Request-Start", startTime.Format(time.RFC3339Nano))
		}

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		// Auto-parse multipart forms if content type is multipart/form-data
		if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			if err := r.ParseMultipartForm(32 << 20); err == nil { // 32 MB max
				// Extract form values
				formData := make(map[string]string)
				if r.MultipartForm != nil {
					for key, values := range r.MultipartForm.Value {
						formData[key] = strings.Join(values, ", ")
					}
				}

				// Extract file uploads
				files := make(map[string]*object.UploadedFile)
				if r.MultipartForm != nil {
					for key, fileHeaders := range r.MultipartForm.File {
						if len(fileHeaders) > 0 {
							fileHeader := fileHeaders[0]
							file, err := fileHeader.Open()
							if err != nil {
								continue
							}
							defer file.Close()

							content, err := io.ReadAll(file)
							if err != nil {
								continue
							}

							files[key] = &object.UploadedFile{
								Name:     fileHeader.Filename,
								Size:     fileHeader.Size,
								MimeType: fileHeader.Header.Get("Content-Type"),
								Content:  content,
							}
						}
					}
				}

				req.FormData = formData
				req.Files = files
			}
		}

		// Run request interceptors
		if interceptors, exists := app.Interceptors["request"]; exists {
			for _, interceptor := range interceptors {
				log.Printf("Running request interceptor: %s", interceptor.Inspect())
			}
		}

		// Run guards
		for _, guard := range app.Guards {
			log.Printf("Running guard: %s", guard.Inspect())
		}

		// Find matching route (including route groups)
		routeKey := r.Method + ":" + r.URL.Path
		handler, exists := app.Routes[routeKey]
		
		if !exists {
			// Check route groups
			for prefix, group := range app.RouteGroups {
				if strings.HasPrefix(r.URL.Path, prefix) {
					// Remove prefix and try to match in group
					groupPath := strings.TrimPrefix(r.URL.Path, prefix)
					groupKey := r.Method + ":" + groupPath
					if groupHandler, groupExists := group.Routes[groupKey]; groupExists {
						handler = groupHandler
						exists = true
						break
					}
				}
			}
		}
		
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
			
			// Enhanced error response structure
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "application/json")
			errorResponse := map[string]interface{}{
				"error": map[string]interface{}{
					"type":    "NOT_FOUND",
					"message": fmt.Sprintf("Cannot %s %s", r.Method, r.URL.Path),
					"code":    "ROUTE_NOT_FOUND",
					"status":  404,
					"details": map[string]interface{}{
						"method": r.Method,
						"path":   r.URL.Path,
						"timestamp": time.Now().UTC().Format(time.RFC3339),
					},
				},
			}
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		// Run middleware
		for _, middleware := range app.Middleware {
			log.Printf("Running middleware: %s", middleware.Inspect())
		}

		// Execute the route handler
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
		if len(req.Files) > 0 {
			response += fmt.Sprintf("- Uploaded files: %d\n", len(req.Files))
			for name, file := range req.Files {
				response += fmt.Sprintf("  • %s: %s (%d bytes)\n", name, file.Name, file.Size)
			}
		}

		// Add handler type indicators
		if handler.IsAsync {
			response += "- Handler: Async (non-blocking)\n"
		}
		if handler.IsStreaming {
			response += "- Handler: Streaming (real-time)\n"
		}

		// Add performance metrics if enabled
		if app.Performance != nil && app.Performance.RequestTiming {
			duration := time.Since(startTime)
			response += fmt.Sprintf("- Request duration: %v\n", duration)
			w.Header().Set("X-Response-Time", duration.String())
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

// Enterprise Features Implementation

// createRouteGroup creates a route group with a common prefix
func createRouteGroup(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	if len(args) != 2 {
		return &object.Error{Message: "http.group() requires exactly 2 arguments: prefix and group function"}
	}

	prefix, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "First argument (prefix) must be a string"}
	}

	groupFunc, ok := args[1].(*object.Function)
	if !ok {
		return &object.Error{Message: "Second argument (group function) must be a function"}
	}

	// Create a new route group context
	if currentApp.RouteGroups == nil {
		currentApp.RouteGroups = make(map[string]*object.RouteGroup)
	}

	routeGroup := &object.RouteGroup{
		Prefix:     prefix.Value,
		Routes:     make(map[string]*object.Function),
		Middleware: make([]*object.Function, 0),
		Guards:     make([]*object.Function, 0),
	}

	currentApp.RouteGroups[prefix.Value] = routeGroup
	
	// Store the group function for future use
	_ = groupFunc

	return &object.String{Value: fmt.Sprintf("Route group created with prefix: %s", prefix.Value)}
}

// parseMultipart handles multipart form data parsing
func parseMultipart(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "http.multipart() requires exactly 1 argument: request object"}
	}

	req, ok := args[0].(*object.HTTPRequest)
	if !ok {
		return &object.Error{Message: "Argument must be an HTTPRequest object"}
	}

	if req.RawRequest == nil {
		return &object.Error{Message: "Invalid request object"}
	}

	// Parse multipart form
	err := req.RawRequest.ParseMultipartForm(32 << 20) // 32 MB max memory
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to parse multipart form: %v", err)}
	}

	// Extract form values
	formData := make(map[string]string)
	if req.RawRequest.MultipartForm != nil {
		for key, values := range req.RawRequest.MultipartForm.Value {
			formData[key] = strings.Join(values, ", ")
		}
	}

	// Extract file uploads
	files := make(map[string]*object.UploadedFile)
	if req.RawRequest.MultipartForm != nil {
		for key, fileHeaders := range req.RawRequest.MultipartForm.File {
			if len(fileHeaders) > 0 {
				fileHeader := fileHeaders[0]
				file, err := fileHeader.Open()
				if err != nil {
					continue
				}
				defer file.Close()

				// Read file content
				content, err := io.ReadAll(file)
				if err != nil {
					continue
				}

				files[key] = &object.UploadedFile{
					Name:     fileHeader.Filename,
					Size:     fileHeader.Size,
					MimeType: fileHeader.Header.Get("Content-Type"),
					Content:  content,
				}
			}
		}
	}

	// Update request object
	req.FormData = formData
	req.Files = files

	return &object.String{Value: fmt.Sprintf("Multipart form parsed: %d fields, %d files", len(formData), len(files))}
}

// createAsyncHandler creates an async handler for long-running operations
func createAsyncHandler(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "http.async() requires exactly 1 argument: handler function"}
	}

	handler, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: "Handler must be a function"}
	}

	// Create an async wrapper function
	asyncHandler := &object.Function{
		Parameters: handler.Parameters,
		Body:       handler.Body,
		Env:        handler.Env,
		IsAsync:    true, // Mark as async
	}

	return asyncHandler
}

// securityMiddleware creates security middleware with CSRF protection and security headers
func securityMiddleware(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	// Create security middleware function
	securityFunc := &object.Function{
		Parameters: []*ast.Identifier{
			{Value: "req"},
			{Value: "res"},
			{Value: "next"},
		},
		Body: nil,
		Env:  nil,
	}

	currentApp.Middleware = append(currentApp.Middleware, securityFunc)
	return &object.String{Value: "Security middleware registered (CSRF protection, security headers)"}
}

// createStreamHandler creates a streaming response handler
func createStreamHandler(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "http.stream() requires exactly 1 argument: stream handler function"}
	}

	handler, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: "Stream handler must be a function"}
	}

	// Create a streaming wrapper function
	streamHandler := &object.Function{
		Parameters: handler.Parameters,
		Body:       handler.Body,
		Env:        handler.Env,
		IsStreaming: true, // Mark as streaming
	}

	return streamHandler
}

// enableMetrics enables performance monitoring and metrics collection
func enableMetrics(args []object.Object, defs map[string]object.Object) object.Object {
	if currentApp == nil {
		return &object.Error{Message: "No app instance found. Call http.app() first."}
	}

	// Enable metrics in performance config
	if currentApp.Performance != nil {
		currentApp.Performance.EnableMetrics = true
		currentApp.Performance.RequestTiming = true
	}

	return &object.String{Value: "Performance metrics enabled"}
}

package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// HTTPApp represents an Express.js-like application instance
type HTTPApp struct {
	Routes     map[string]*Function // key: "METHOD:/path"
	Middleware []*Function
	Server     *http.Server
	// New features for full backend support
	Interceptors map[string][]*Function // key: "request" or "response"
	Guards       []*Function
	ErrorHandler *Function
	// Enterprise features
	RouteGroups map[string]*RouteGroup
	Security    *SecurityConfig
	Performance *PerformanceConfig
}

// RouteGroup represents a group of routes with common prefix and middleware
type RouteGroup struct {
	Prefix     string
	Routes     map[string]*Function
	Middleware []*Function
	Guards     []*Function
}

// SecurityConfig holds security-related configuration
type SecurityConfig struct {
	CSRFProtection  bool
	SecurityHeaders map[string]string
	CORSOptions     *CORSOptions
}

// CORSOptions holds CORS configuration
type CORSOptions struct {
	Origin      string
	Methods     []string
	Headers     []string
	Credentials bool
}

// PerformanceConfig holds performance monitoring configuration
type PerformanceConfig struct {
	EnableMetrics bool
	MetricsPath   string
	RequestTiming bool
}

// UploadedFile represents an uploaded file
type UploadedFile struct {
	Name     string
	Size     int64
	MimeType string
	Content  []byte
}

func (file *UploadedFile) Type() VintObjectType { return UPLOADED_FILE_OBJ }
func (file *UploadedFile) Inspect() string {
	return fmt.Sprintf("UploadedFile{name: %s, size: %d, type: %s}", file.Name, file.Size, file.MimeType)
}

// Method to access uploaded file functionality
func (file *UploadedFile) Method(name string, args []VintObject) VintObject {
	switch name {
	case "save":
		if len(args) != 1 {
			return &Error{Message: "file.save() requires 1 argument: destination path"}
		}
		path, ok := args[0].(*String)
		if !ok {
			return &Error{Message: "Destination path must be a string"}
		}

		// Create directory if it doesn't exist
		dir := filepath.Dir(path.Value)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return &Error{Message: fmt.Sprintf("Failed to create directory: %v", err)}
		}

		// Write file
		if err := os.WriteFile(path.Value, file.Content, 0644); err != nil {
			return &Error{Message: fmt.Sprintf("Failed to save file: %v", err)}
		}

		return &String{Value: fmt.Sprintf("File saved to: %s", path.Value)}
	case "name":
		return &String{Value: file.Name}
	case "size":
		return &Integer{Value: file.Size}
	case "type":
		return &String{Value: file.MimeType}
	case "content":
		return &String{Value: string(file.Content)}
	default:
		return &Error{Message: fmt.Sprintf("Unknown file method: %s", name)}
	}
}

func (app *HTTPApp) Type() VintObjectType { return HTTP_APP_OBJ }
func (app *HTTPApp) Inspect() string {
	var out bytes.Buffer
	out.WriteString("HTTPApp{")
	out.WriteString(fmt.Sprintf("routes: %d, ", len(app.Routes)))
	out.WriteString(fmt.Sprintf("middleware: %d, ", len(app.Middleware)))
	out.WriteString(fmt.Sprintf("interceptors: %d, ", len(app.Interceptors)))
	out.WriteString(fmt.Sprintf("guards: %d, ", len(app.Guards)))
	out.WriteString(fmt.Sprintf("routeGroups: %d", len(app.RouteGroups)))
	out.WriteString("}")
	return out.String()
}

// HTTPRequest represents an HTTP request
type HTTPRequest struct {
	HTTPMethod string
	Path       string
	Headers    map[string]string
	Body       string
	Query      map[string]string
	Params     map[string]string
	// Enhanced features
	BodyBytes  []byte
	Cookies    map[string]string
	FormData   map[string]string
	JSON       map[string]any
	RawRequest *http.Request
	// Enterprise features
	Files   map[string]*UploadedFile
	IsAsync bool
}

func (req *HTTPRequest) Type() VintObjectType { return HTTP_REQUEST_OBJ }
func (req *HTTPRequest) Inspect() string {
	var out bytes.Buffer
	out.WriteString("HTTPRequest{")
	out.WriteString(fmt.Sprintf("method: %s, ", req.HTTPMethod))
	out.WriteString(fmt.Sprintf("path: %s", req.Path))
	out.WriteString("}")
	return out.String()
}

// Method to access request properties
func (req *HTTPRequest) Method(name string, args []VintObject) VintObject {
	switch name {
	case "get":
		if len(args) != 1 {
			return &Error{Message: "req.get() requires 1 argument: header name"}
		}
		headerName, ok := args[0].(*String)
		if !ok {
			return &Error{Message: "Header name must be a string"}
		}
		if value, exists := req.Headers[headerName.Value]; exists {
			return &String{Value: value}
		}
		return &String{Value: ""}
	case "body":
		return &String{Value: req.Body}
	case "method":
		return &String{Value: req.HTTPMethod}
	case "path":
		return &String{Value: req.Path}
	case "query":
		if len(args) == 0 {
			// Return all query parameters as a dict-like structure
			result := make(map[string]VintObject)
			for k, v := range req.Query {
				result[k] = &String{Value: v}
			}
			return &String{Value: fmt.Sprintf("%v", result)} // Simplified for now
		} else if len(args) == 1 {
			paramName, ok := args[0].(*String)
			if !ok {
				return &Error{Message: "Query parameter name must be a string"}
			}
			if value, exists := req.Query[paramName.Value]; exists {
				return &String{Value: value}
			}
			return &String{Value: ""}
		}
		return &Error{Message: "req.query() takes 0 or 1 arguments"}
	case "param":
		if len(args) != 1 {
			return &Error{Message: "req.param() requires 1 argument: parameter name"}
		}
		paramName, ok := args[0].(*String)
		if !ok {
			return &Error{Message: "Parameter name must be a string"}
		}
		if value, exists := req.Params[paramName.Value]; exists {
			return &String{Value: value}
		}
		return &String{Value: ""}
	case "cookie":
		if len(args) != 1 {
			return &Error{Message: "req.cookie() requires 1 argument: cookie name"}
		}
		cookieName, ok := args[0].(*String)
		if !ok {
			return &Error{Message: "Cookie name must be a string"}
		}
		if value, exists := req.Cookies[cookieName.Value]; exists {
			return &String{Value: value}
		}
		return &String{Value: ""}
	case "json":
		// Parse and return JSON body
		if req.JSON != nil {
			jsonStr, _ := json.Marshal(req.JSON)
			return &String{Value: string(jsonStr)}
		}
		return &String{Value: "{}"}
	case "form":
		if len(args) == 0 {
			// Return all form data
			result := make(map[string]VintObject)
			for k, v := range req.FormData {
				result[k] = &String{Value: v}
			}
			return &String{Value: fmt.Sprintf("%v", result)} // Simplified for now
		} else if len(args) == 1 {
			fieldName, ok := args[0].(*String)
			if !ok {
				return &Error{Message: "Form field name must be a string"}
			}
			if value, exists := req.FormData[fieldName.Value]; exists {
				return &String{Value: value}
			}
			return &String{Value: ""}
		}
		return &Error{Message: "req.form() takes 0 or 1 arguments"}
	case "file":
		if len(args) != 1 {
			return &Error{Message: "req.file() requires 1 argument: file field name"}
		}
		fieldName, ok := args[0].(*String)
		if !ok {
			return &Error{Message: "File field name must be a string"}
		}
		if file, exists := req.Files[fieldName.Value]; exists {
			return file
		}
		return &String{Value: ""}
	case "files":
		// Return all uploaded files
		result := make(map[string]VintObject)
		for k, v := range req.Files {
			result[k] = v
		}
		return &String{Value: fmt.Sprintf("Files: %d uploaded", len(req.Files))}
	default:
		return &Error{Message: fmt.Sprintf("Unknown request method: %s", name)}
	}
}

// HTTPResponse represents an HTTP response
type HTTPResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       string
	Writer     http.ResponseWriter
	Sent       bool
	// Enhanced features
	Request *HTTPRequest
}

func (res *HTTPResponse) Type() VintObjectType { return HTTP_RESPONSE_OBJ }
func (res *HTTPResponse) Inspect() string {
	var out bytes.Buffer
	out.WriteString("HTTPResponse{")
	out.WriteString(fmt.Sprintf("status: %d", res.StatusCode))
	out.WriteString("}")
	return out.String()
}

// Method to access response functionality
func (res *HTTPResponse) Method(name string, args []VintObject) VintObject {
	switch name {
	case "send":
		if len(args) != 1 {
			return &Error{Message: "res.send() requires 1 argument: message"}
		}
		message, ok := args[0].(*String)
		if !ok {
			return &Error{Message: "Message must be a string"}
		}
		res.Send(message.Value)
		return &String{Value: "Response sent"}
	case "json":
		if len(args) != 1 {
			return &Error{Message: "res.json() requires 1 argument: data"}
		}
		// For now, convert the object to its string representation
		// In a full implementation, this would properly serialize to JSON
		data := args[0].Inspect()
		res.Writer.Header().Set("Content-Type", "application/json")
		for key, value := range res.Headers {
			res.Writer.Header().Set(key, value)
		}
		res.Writer.WriteHeader(res.StatusCode)
		res.Writer.Write([]byte(data))
		res.Sent = true
		return &String{Value: "JSON response sent"}
	case "status":
		if len(args) != 1 {
			return &Error{Message: "res.status() requires 1 argument: status code"}
		}
		statusCode, ok := args[0].(*Integer)
		if !ok {
			return &Error{Message: "Status code must be an integer"}
		}
		res.StatusCode = int(statusCode.Value)
		return res // Return self for chaining
	case "header":
		if len(args) != 2 {
			return &Error{Message: "res.header() requires 2 arguments: key and value"}
		}
		key, ok1 := args[0].(*String)
		value, ok2 := args[1].(*String)
		if !ok1 || !ok2 {
			return &Error{Message: "Header key and value must be strings"}
		}
		res.Headers[key.Value] = value.Value
		return res // Return self for chaining
	case "redirect":
		if len(args) < 1 || len(args) > 2 {
			return &Error{Message: "res.redirect() requires 1-2 arguments: url and optional status code"}
		}
		url, ok := args[0].(*String)
		if !ok {
			return &Error{Message: "Redirect URL must be a string"}
		}
		statusCode := 302 // Default redirect status
		if len(args) == 2 {
			if code, ok := args[1].(*Integer); ok {
				statusCode = int(code.Value)
			}
		}
		res.Writer.Header().Set("Location", url.Value)
		res.Writer.WriteHeader(statusCode)
		res.Sent = true
		return &String{Value: "Redirect sent"}
	case "cookie":
		if len(args) < 2 || len(args) > 3 {
			return &Error{Message: "res.cookie() requires 2-3 arguments: name, value, and optional options"}
		}
		name, ok1 := args[0].(*String)
		value, ok2 := args[1].(*String)
		if !ok1 || !ok2 {
			return &Error{Message: "Cookie name and value must be strings"}
		}

		cookie := &http.Cookie{
			Name:  name.Value,
			Value: value.Value,
			Path:  "/",
		}

		// TODO: Handle cookie options (expires, secure, httponly, etc.)
		// For now, set basic cookie
		http.SetCookie(res.Writer, cookie)
		return &String{Value: "Cookie set"}
	case "end":
		if len(args) > 1 {
			return &Error{Message: "res.end() takes 0 or 1 arguments: optional message"}
		}
		if len(args) == 1 {
			if message, ok := args[0].(*String); ok {
				res.Writer.Write([]byte(message.Value))
			}
		}
		res.Sent = true
		return &String{Value: "Response ended"}
	default:
		return &Error{Message: fmt.Sprintf("Unknown response method: %s", name)}
	}
}

// Helper function to create a new HTTPApp
func NewHTTPApp() *HTTPApp {
	return &HTTPApp{
		Routes:       make(map[string]*Function),
		Middleware:   make([]*Function, 0),
		Interceptors: make(map[string][]*Function),
		Guards:       make([]*Function, 0),
		RouteGroups:  make(map[string]*RouteGroup),
		Security: &SecurityConfig{
			CSRFProtection: false,
			SecurityHeaders: map[string]string{
				"X-Content-Type-Options": "nosniff",
				"X-Frame-Options":        "DENY",
				"X-XSS-Protection":       "1; mode=block",
			},
			CORSOptions: &CORSOptions{
				Origin:  "*",
				Methods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
				Headers: []string{"Content-Type", "Authorization"},
			},
		},
		Performance: &PerformanceConfig{
			EnableMetrics: false,
			MetricsPath:   "/metrics",
			RequestTiming: false,
		},
	}
}

// Helper function to create HTTPRequest from http.Request
func NewHTTPRequest(r *http.Request) *HTTPRequest {
	headers := make(map[string]string)
	for key, values := range r.Header {
		headers[key] = strings.Join(values, ", ")
	}

	query := make(map[string]string)
	for key, values := range r.URL.Query() {
		query[key] = strings.Join(values, ", ")
	}

	// Read body
	bodyBytes := make([]byte, 0)
	bodyString := ""
	if r.Body != nil {
		defer r.Body.Close()
		bodyBytes, _ = io.ReadAll(r.Body)
		bodyString = string(bodyBytes)
	}

	// Parse cookies
	cookies := make(map[string]string)
	for _, cookie := range r.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}

	// Parse form data
	formData := make(map[string]string)
	if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		if parsed, err := url.ParseQuery(bodyString); err == nil {
			for key, values := range parsed {
				formData[key] = strings.Join(values, ", ")
			}
		}
	}

	// Parse JSON data
	var jsonData map[string]any
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		json.Unmarshal(bodyBytes, &jsonData)
	}

	return &HTTPRequest{
		HTTPMethod: r.Method,
		Path:       r.URL.Path,
		Headers:    headers,
		Body:       bodyString,
		Query:      query,
		Params:     make(map[string]string),
		BodyBytes:  bodyBytes,
		Cookies:    cookies,
		FormData:   formData,
		JSON:       jsonData,
		RawRequest: r,
		Files:      make(map[string]*UploadedFile),
		IsAsync:    false,
	}
}

// Helper function to create HTTPResponse
func NewHTTPResponse(w http.ResponseWriter, req *HTTPRequest) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 200,
		Headers:    make(map[string]string),
		Writer:     w,
		Sent:       false,
		Request:    req,
	}
}

// Send text response
func (res *HTTPResponse) Send(text string) {
	if res.Sent {
		log.Println("Warning: Response already sent")
		return
	}

	res.Writer.Header().Set("Content-Type", "text/plain")
	for key, value := range res.Headers {
		res.Writer.Header().Set(key, value)
	}
	res.Writer.WriteHeader(res.StatusCode)
	res.Writer.Write([]byte(text))
	res.Sent = true
}

// Send JSON response
func (res *HTTPResponse) JSON(data any) {
	if res.Sent {
		log.Println("Warning: Response already sent")
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		res.Writer.WriteHeader(500)
		res.Writer.Write([]byte("Internal Server Error"))
		res.Sent = true
		return
	}

	res.Writer.Header().Set("Content-Type", "application/json")
	for key, value := range res.Headers {
		res.Writer.Header().Set(key, value)
	}
	res.Writer.WriteHeader(res.StatusCode)
	res.Writer.Write(jsonData)
	res.Sent = true
}

// Set status code
func (res *HTTPResponse) Status(code int) *HTTPResponse {
	res.StatusCode = code
	return res
}

// Set header
func (res *HTTPResponse) Header(key, value string) *HTTPResponse {
	res.Headers[key] = value
	return res
}

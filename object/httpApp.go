package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// HTTPApp represents an Express.js-like application instance
type HTTPApp struct {
	Routes     map[string]*Function // key: "METHOD:/path"
	Middleware []*Function
	Server     *http.Server
}

func (app *HTTPApp) Type() ObjectType { return HTTP_APP_OBJ }
func (app *HTTPApp) Inspect() string {
	var out bytes.Buffer
	out.WriteString("HTTPApp{")
	out.WriteString(fmt.Sprintf("routes: %d, ", len(app.Routes)))
	out.WriteString(fmt.Sprintf("middleware: %d", len(app.Middleware)))
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
}

func (req *HTTPRequest) Type() ObjectType { return HTTP_REQUEST_OBJ }
func (req *HTTPRequest) Inspect() string {
	var out bytes.Buffer
	out.WriteString("HTTPRequest{")
	out.WriteString(fmt.Sprintf("method: %s, ", req.HTTPMethod))
	out.WriteString(fmt.Sprintf("path: %s", req.Path))
	out.WriteString("}")
	return out.String()
}

// Method to access request properties
func (req *HTTPRequest) Method(name string, args []Object) Object {
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
}

func (res *HTTPResponse) Type() ObjectType { return HTTP_RESPONSE_OBJ }
func (res *HTTPResponse) Inspect() string {
	var out bytes.Buffer
	out.WriteString("HTTPResponse{")
	out.WriteString(fmt.Sprintf("status: %d", res.StatusCode))
	out.WriteString("}")
	return out.String()
}

// Method to access response functionality
func (res *HTTPResponse) Method(name string, args []Object) Object {
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
	default:
		return &Error{Message: fmt.Sprintf("Unknown response method: %s", name)}
	}
}

// Helper function to create a new HTTPApp
func NewHTTPApp() *HTTPApp {
	return &HTTPApp{
		Routes:     make(map[string]*Function),
		Middleware: make([]*Function, 0),
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
	if r.Body != nil {
		defer r.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		bodyBytes = buf.Bytes()
	}

	return &HTTPRequest{
		HTTPMethod: r.Method,
		Path:       r.URL.Path,
		Headers:    headers,
		Body:       string(bodyBytes),
		Query:      query,
		Params:     make(map[string]string),
	}
}

// Helper function to create HTTPResponse
func NewHTTPResponse(w http.ResponseWriter) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 200,
		Headers:    make(map[string]string),
		Writer:     w,
		Sent:       false,
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
func (res *HTTPResponse) JSON(data interface{}) {
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
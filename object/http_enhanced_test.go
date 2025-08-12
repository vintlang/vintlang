package object

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPRequestEnhancements(t *testing.T) {
	// Create a mock HTTP request with various features
	reqBody := `{"name": "John", "email": "john@example.com"}`
	req := httptest.NewRequest("POST", "/users/123?sort=name&limit=10", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token123")
	req.AddCookie(&http.Cookie{Name: "session", Value: "abc123"})

	// Create enhanced HTTPRequest
	httpReq := NewHTTPRequest(req)

	// Test basic properties
	if httpReq.HTTPMethod != "POST" {
		t.Errorf("Expected method POST, got %s", httpReq.HTTPMethod)
	}

	if httpReq.Path != "/users/123" {
		t.Errorf("Expected path /users/123, got %s", httpReq.Path)
	}

	// Test headers
	if httpReq.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", httpReq.Headers["Content-Type"])
	}

	// Test query parameters
	if httpReq.Query["sort"] != "name" {
		t.Errorf("Expected query param sort=name, got %s", httpReq.Query["sort"])
	}

	if httpReq.Query["limit"] != "10" {
		t.Errorf("Expected query param limit=10, got %s", httpReq.Query["limit"])
	}

	// Test body
	if httpReq.Body != reqBody {
		t.Errorf("Expected body %s, got %s", reqBody, httpReq.Body)
	}

	// Test cookies
	if httpReq.Cookies["session"] != "abc123" {
		t.Errorf("Expected cookie session=abc123, got %s", httpReq.Cookies["session"])
	}

	// Test JSON parsing
	if httpReq.JSON == nil {
		t.Errorf("Expected JSON to be parsed")
	}

	if name, exists := httpReq.JSON["name"]; !exists || name != "John" {
		t.Errorf("Expected JSON name=John, got %v", name)
	}
}

func TestHTTPRequestMethods(t *testing.T) {
	// Create a test request
	req := httptest.NewRequest("GET", "/test?param=value", nil)
	req.Header.Set("User-Agent", "test-agent")
	req.AddCookie(&http.Cookie{Name: "test", Value: "cookie-value"})
	
	httpReq := NewHTTPRequest(req)
	httpReq.Params["id"] = "123" // Simulate path parameter

	// Test header method
	result := httpReq.Method("get", []Object{&String{Value: "User-Agent"}})
	if str, ok := result.(*String); !ok || str.Value != "test-agent" {
		t.Errorf("Expected User-Agent header, got %v", result)
	}

	// Test non-existent header
	result = httpReq.Method("get", []Object{&String{Value: "Non-Existent"}})
	if str, ok := result.(*String); !ok || str.Value != "" {
		t.Errorf("Expected empty string for non-existent header, got %v", result)
	}

	// Test query method
	result = httpReq.Method("query", []Object{&String{Value: "param"}})
	if str, ok := result.(*String); !ok || str.Value != "value" {
		t.Errorf("Expected query param value, got %v", result)
	}

	// Test param method
	result = httpReq.Method("param", []Object{&String{Value: "id"}})
	if str, ok := result.(*String); !ok || str.Value != "123" {
		t.Errorf("Expected path param 123, got %v", result)
	}

	// Test cookie method
	result = httpReq.Method("cookie", []Object{&String{Value: "test"}})
	if str, ok := result.(*String); !ok || str.Value != "cookie-value" {
		t.Errorf("Expected cookie value, got %v", result)
	}

	// Test method method
	result = httpReq.Method("method", []Object{})
	if str, ok := result.(*String); !ok || str.Value != "GET" {
		t.Errorf("Expected GET method, got %v", result)
	}

	// Test path method
	result = httpReq.Method("path", []Object{})
	if str, ok := result.(*String); !ok || str.Value != "/test" {
		t.Errorf("Expected /test path, got %v", result)
	}

	// Test body method
	result = httpReq.Method("body", []Object{})
	if _, ok := result.(*String); !ok {
		t.Errorf("Expected string body, got %v", result)
	}
}

func TestHTTPResponseEnhancements(t *testing.T) {
	// Create a test response writer
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	httpReq := NewHTTPRequest(req)
	
	// Create enhanced HTTPResponse
	httpRes := NewHTTPResponse(recorder, httpReq)

	// Test initial state
	if httpRes.StatusCode != 200 {
		t.Errorf("Expected default status code 200, got %d", httpRes.StatusCode)
	}

	if httpRes.Sent {
		t.Errorf("Expected response not to be sent initially")
	}

	if httpRes.Request != httpReq {
		t.Errorf("Expected response to have reference to request")
	}

	// Test headers initialization
	if httpRes.Headers == nil {
		t.Errorf("Expected headers map to be initialized")
	}
}

func TestHTTPResponseMethods(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	httpReq := NewHTTPRequest(req)
	httpRes := NewHTTPResponse(recorder, httpReq)

	// Test status method
	result := httpRes.Method("status", []Object{&Integer{Value: 404}})
	if httpRes.StatusCode != 404 {
		t.Errorf("Expected status code 404, got %d", httpRes.StatusCode)
	}

	// Should return self for chaining
	if result != httpRes {
		t.Errorf("Expected status method to return self for chaining")
	}

	// Test header method
	result = httpRes.Method("header", []Object{
		&String{Value: "X-Custom-Header"},
		&String{Value: "custom-value"},
	})
	if httpRes.Headers["X-Custom-Header"] != "custom-value" {
		t.Errorf("Expected custom header to be set")
	}

	// Should return self for chaining
	if result != httpRes {
		t.Errorf("Expected header method to return self for chaining")
	}

	// Test send method
	result = httpRes.Method("send", []Object{&String{Value: "Hello World"}})
	if !httpRes.Sent {
		t.Errorf("Expected response to be marked as sent")
	}

	// Test end method
	httpRes2 := NewHTTPResponse(httptest.NewRecorder(), httpReq)
	result = httpRes2.Method("end", []Object{&String{Value: "Goodbye"}})
	if !httpRes2.Sent {
		t.Errorf("Expected response to be marked as sent after end")
	}
}

func TestFormDataParsing(t *testing.T) {
	// Create a form data request
	formData := "name=John&email=john@example.com&age=30"
	req := httptest.NewRequest("POST", "/submit", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpReq := NewHTTPRequest(req)

	// Test form data parsing
	if httpReq.FormData["name"] != "John" {
		t.Errorf("Expected form data name=John, got %s", httpReq.FormData["name"])
	}

	if httpReq.FormData["email"] != "john@example.com" {
		t.Errorf("Expected form data email=john@example.com, got %s", httpReq.FormData["email"])
	}

	if httpReq.FormData["age"] != "30" {
		t.Errorf("Expected form data age=30, got %s", httpReq.FormData["age"])
	}

	// Test form method
	result := httpReq.Method("form", []Object{&String{Value: "name"}})
	if str, ok := result.(*String); !ok || str.Value != "John" {
		t.Errorf("Expected form field name=John, got %v", result)
	}

	// Test form method with all fields
	result = httpReq.Method("form", []Object{})
	if _, ok := result.(*String); !ok {
		t.Errorf("Expected string representation of all form fields, got %v", result)
	}
}

func TestErrorHandling(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	httpReq := NewHTTPRequest(req)

	// Test invalid method calls
	result := httpReq.Method("unknown", []Object{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for unknown method")
	}

	result = httpReq.Method("get", []Object{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for get method with wrong arguments")
	}

	result = httpReq.Method("query", []Object{&Integer{Value: 123}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for query method with wrong argument type")
	}

	// Test response error handling
	recorder := httptest.NewRecorder()
	httpRes := NewHTTPResponse(recorder, httpReq)

	result = httpRes.Method("unknown", []Object{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for unknown response method")
	}

	result = httpRes.Method("status", []Object{&String{Value: "not a number"}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for status method with wrong argument type")
	}
}
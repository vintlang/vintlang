package object

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPRequestIP(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.100:12345"

	httpReq := NewHTTPRequest(req)

	// Test remoteAddr field
	if httpReq.RemoteAddr != "192.168.1.100:12345" {
		t.Errorf("Expected RemoteAddr '192.168.1.100:12345', got '%s'", httpReq.RemoteAddr)
	}

	// Test remoteAddr method
	result := httpReq.Method("remoteAddr", []VintObject{})
	if str, ok := result.(*String); !ok || str.Value != "192.168.1.100:12345" {
		t.Errorf("Expected remoteAddr method to return '192.168.1.100:12345', got %v", result)
	}

	// Test ip method (should extract just the IP without port)
	result = httpReq.Method("ip", []VintObject{})
	if str, ok := result.(*String); !ok || str.Value != "192.168.1.100" {
		t.Errorf("Expected ip method to return '192.168.1.100', got %v", result)
	}
}

func TestHTTPRequestIPWithoutPort(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "10.0.0.1"

	httpReq := NewHTTPRequest(req)

	// Test ip method when no port is present
	result := httpReq.Method("ip", []VintObject{})
	if str, ok := result.(*String); !ok || str.Value != "10.0.0.1" {
		t.Errorf("Expected ip method to return '10.0.0.1', got %v", result)
	}
}

func TestHTTPRequestHeaders(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token123")

	httpReq := NewHTTPRequest(req)

	// Test headers method returns a Dict
	result := httpReq.Method("headers", []VintObject{})
	dict, ok := result.(*Dict)
	if !ok {
		t.Fatalf("Expected headers method to return a Dict, got %T", result)
	}

	// Verify we have headers
	if len(dict.Pairs) == 0 {
		t.Error("Expected headers dict to have entries")
	}

	// Check specific header values
	found := false
	for _, pair := range dict.Pairs {
		if key, ok := pair.Key.(*String); ok {
			if key.Value == "Content-Type" {
				if val, ok := pair.Value.(*String); ok && val.Value == "application/json" {
					found = true
				}
			}
		}
	}
	if !found {
		t.Error("Expected to find Content-Type: application/json in headers dict")
	}
}

func TestHTTPResponseMethodChaining(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	httpReq := NewHTTPRequest(req)
	httpRes := NewHTTPResponse(recorder, httpReq)

	// Test status + header chaining
	result := httpRes.Method("status", []VintObject{&Integer{Value: 201}})
	if result != httpRes {
		t.Error("Expected status method to return self for chaining")
	}
	if httpRes.StatusCode != 201 {
		t.Errorf("Expected status code 201, got %d", httpRes.StatusCode)
	}

	result = httpRes.Method("header", []VintObject{
		&String{Value: "X-Custom"},
		&String{Value: "test-value"},
	})
	if result != httpRes {
		t.Error("Expected header method to return self for chaining")
	}
	if httpRes.Headers["X-Custom"] != "test-value" {
		t.Error("Expected custom header to be set")
	}
}

func TestHTTPRequestBodyForwarding(t *testing.T) {
	bodyContent := `{"key": "value", "nested": {"a": 1}}`
	req := httptest.NewRequest("POST", "/proxy", strings.NewReader(bodyContent))
	req.Header.Set("Content-Type", "application/json")

	httpReq := NewHTTPRequest(req)

	// Body should be available as raw string
	result := httpReq.Method("body", []VintObject{})
	str, ok := result.(*String)
	if !ok {
		t.Fatalf("Expected body to be a string, got %T", result)
	}
	if str.Value != bodyContent {
		t.Errorf("Expected body '%s', got '%s'", bodyContent, str.Value)
	}
}

func TestCallbackRegistration(t *testing.T) {
	// Test that RegisterFuncCaller and CallFunction work correctly
	called := false
	RegisterFuncCaller(func(fn *Function, args []VintObject) VintObject {
		called = true
		return &String{Value: "callback result"}
	})

	fn := &Function{
		Parameters: nil,
		Body:       nil,
		Env:        nil,
	}

	result := CallFunction(fn, []VintObject{})
	if !called {
		t.Error("Expected callback to be called")
	}
	if str, ok := result.(*String); !ok || str.Value != "callback result" {
		t.Errorf("Expected 'callback result', got %v", result)
	}
}

func TestHTTPRequestMultipleHeaders(t *testing.T) {
	req := httptest.NewRequest("POST", "/api", strings.NewReader("test body"))
	req.Header.Set("X-Proxy-Secret", "my-secret")
	req.Header.Set("X-Custom-Header", "custom-value")
	req.Header.Set("Accept", "application/json")

	httpReq := NewHTTPRequest(req)

	// Test get method for each header
	tests := []struct {
		header   string
		expected string
	}{
		{"X-Proxy-Secret", "my-secret"},
		{"X-Custom-Header", "custom-value"},
		{"Accept", "application/json"},
	}

	for _, tt := range tests {
		result := httpReq.Method("get", []VintObject{&String{Value: tt.header}})
		str, ok := result.(*String)
		if !ok {
			t.Errorf("Expected string for header %s, got %T", tt.header, result)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("For header %s: expected '%s', got '%s'", tt.header, tt.expected, str.Value)
		}
	}
}

func TestHTTPAppType(t *testing.T) {
	app := NewHTTPApp()

	if app.Type() != HTTP_APP_OBJ {
		t.Errorf("Expected type HTTP_APP, got %s", app.Type())
	}

	if len(app.Routes) != 0 {
		t.Errorf("Expected 0 routes, got %d", len(app.Routes))
	}

	if app.Security == nil {
		t.Error("Expected security config to be initialized")
	}

	if app.Performance == nil {
		t.Error("Expected performance config to be initialized")
	}
}

func TestHTTPResponseSendWithHeaders(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	httpReq := NewHTTPRequest(req)
	httpRes := NewHTTPResponse(recorder, httpReq)

	// Set custom headers before sending
	httpRes.Method("header", []VintObject{
		&String{Value: "X-Response-Header"},
		&String{Value: "response-value"},
	})
	httpRes.Method("status", []VintObject{&Integer{Value: 403}})
	httpRes.Method("send", []VintObject{&String{Value: "forbidden"}})

	if !httpRes.Sent {
		t.Error("Expected response to be sent")
	}
	if recorder.Code != 403 {
		t.Errorf("Expected status code 403, got %d", recorder.Code)
	}
	if recorder.Body.String() != "forbidden" {
		t.Errorf("Expected body 'forbidden', got '%s'", recorder.Body.String())
	}
}

func TestHTTPRequestIPv6(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "[::1]:8080"

	httpReq := NewHTTPRequest(req)

	result := httpReq.Method("ip", []VintObject{})
	if str, ok := result.(*String); !ok || str.Value != "::1" {
		t.Errorf("Expected ip '::1', got %v", result)
	}
}

func TestHTTPResponseRedirect(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/old", nil)
	httpReq := NewHTTPRequest(req)
	httpRes := NewHTTPResponse(recorder, httpReq)

	result := httpRes.Method("redirect", []VintObject{&String{Value: "/new"}})
	if _, ok := result.(*String); !ok {
		t.Errorf("Expected string result from redirect, got %T", result)
	}
	if !httpRes.Sent {
		t.Error("Expected response to be sent after redirect")
	}
	if recorder.Header().Get("Location") != "/new" {
		t.Errorf("Expected Location header '/new', got '%s'", recorder.Header().Get("Location"))
	}
}

func TestHTTPRequestMethodAndPath(t *testing.T) {
	tests := []struct {
		method string
		path   string
	}{
		{"GET", "/health"},
		{"POST", "/api/users"},
		{"PUT", "/api/users/123"},
		{"DELETE", "/api/users/123"},
		{"PATCH", "/api/users/123"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.path, nil)
		httpReq := NewHTTPRequest(req)

		methodResult := httpReq.Method("method", []VintObject{})
		if str, ok := methodResult.(*String); !ok || str.Value != tt.method {
			t.Errorf("Expected method %s, got %v", tt.method, methodResult)
		}

		pathResult := httpReq.Method("path", []VintObject{})
		if str, ok := pathResult.(*String); !ok || str.Value != tt.path {
			t.Errorf("Expected path %s, got %v", tt.path, pathResult)
		}
	}
}

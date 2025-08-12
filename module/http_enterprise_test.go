package module

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

// Test enterprise HTTP features
func TestEnterpriseHTTPFeatures(t *testing.T) {
	// Reset global app for each test
	currentApp = nil

	// Test app creation
	_ = createApp([]object.Object{}, map[string]object.Object{})
	if currentApp == nil {
		t.Fatal("Expected currentApp to be created")
	}

	// Verify default security configuration
	if currentApp.Security == nil {
		t.Fatal("Expected security configuration to be initialized")
	}

	if currentApp.Security.SecurityHeaders["X-Content-Type-Options"] == "" {
		t.Error("Expected security headers to be set")
	}
}

func TestRouteGrouping(t *testing.T) {
	currentApp = object.NewHTTPApp()

	// Test route group creation
	prefix := &object.String{Value: "/api/v1"}
	groupFunc := &object.Function{
		Parameters: []*ast.Identifier{},
		Body:       nil,
		Env:        nil,
	}

	result := createRouteGroup([]object.Object{prefix, groupFunc}, map[string]object.Object{})
	
	if result.Type() != object.STRING_OBJ {
		t.Errorf("Expected string result, got %T", result)
	}

	if _, exists := currentApp.RouteGroups["/api/v1"]; !exists {
		t.Error("Expected route group to be created")
	}
}

func TestMultipartParsing(t *testing.T) {
	// Create a multipart form request
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	
	// Add form field
	writer.WriteField("username", "testuser")
	
	// Add file
	fileWriter, err := writer.CreateFormFile("avatar", "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	fileWriter.Write([]byte("test file content"))
	
	writer.Close()

	// Create HTTP request
	req := httptest.NewRequest("POST", "/upload", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Parse the multipart form first
	err = req.ParseMultipartForm(32 << 20)
	if err != nil {
		t.Fatal(err)
	}

	// Create HTTPRequest object
	httpReq := object.NewHTTPRequest(req)

	// Test multipart parsing
	result := parseMultipart([]object.Object{httpReq}, map[string]object.Object{})
	
	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Got error: %s", result.Inspect())
		return
	}

	if result.Type() != object.STRING_OBJ {
		t.Errorf("Expected string result, got %T", result)
	}

	// Verify form data was parsed
	if httpReq.FormData["username"] != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", httpReq.FormData["username"])
	}

	// Verify file was parsed
	if _, exists := httpReq.Files["avatar"]; !exists {
		t.Error("Expected file 'avatar' to be parsed")
	}
}

func TestAsyncHandler(t *testing.T) {
	handler := &object.Function{
		Parameters: []*ast.Identifier{
			{Value: "req"},
			{Value: "res"},
		},
		Body: nil,
		Env:  nil,
	}

	result := createAsyncHandler([]object.Object{handler}, map[string]object.Object{})
	
	asyncHandler, ok := result.(*object.Function)
	if !ok {
		t.Errorf("Expected Function result, got %T", result)
	}

	if !asyncHandler.IsAsync {
		t.Error("Expected handler to be marked as async")
	}
}

func TestSecurityMiddleware(t *testing.T) {
	currentApp = object.NewHTTPApp()

	result := securityMiddleware([]object.Object{}, map[string]object.Object{})
	
	if result.Type() != object.STRING_OBJ {
		t.Errorf("Expected string result, got %T", result)
	}

	// Verify middleware was added
	if len(currentApp.Middleware) == 0 {
		t.Error("Expected security middleware to be added")
	}
}

func TestEnhancedErrorHandling(t *testing.T) {
	currentApp = object.NewHTTPApp()

	// Create a test handler for 404 endpoint
	handler := createHTTPHandler(currentApp)

	// Test 404 response structure
	req := httptest.NewRequest("GET", "/nonexistent", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != 404 {
		t.Errorf("Expected status 404, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type 'application/json', got '%s'", contentType)
	}

	// Verify JSON structure contains enhanced error format
	body := w.Body.String()
	if !strings.Contains(body, "\"error\"") || !strings.Contains(body, "\"type\"") {
		t.Error("Expected enhanced error structure in response")
	}
}

func TestFileUploadMethods(t *testing.T) {
	// Create uploaded file object
	file := &object.UploadedFile{
		Name:     "test.txt",
		Size:     100,
		MimeType: "text/plain",
		Content:  []byte("test content"),
	}

	// Test name method
	nameResult := file.Method("name", []object.Object{})
	if nameStr, ok := nameResult.(*object.String); !ok || nameStr.Value != "test.txt" {
		t.Errorf("Expected name 'test.txt', got %v", nameResult)
	}

	// Test size method
	sizeResult := file.Method("size", []object.Object{})
	if sizeInt, ok := sizeResult.(*object.Integer); !ok || sizeInt.Value != 100 {
		t.Errorf("Expected size 100, got %v", sizeResult)
	}

	// Test type method
	typeResult := file.Method("type", []object.Object{})
	if typeStr, ok := typeResult.(*object.String); !ok || typeStr.Value != "text/plain" {
		t.Errorf("Expected type 'text/plain', got %v", typeResult)
	}
}

func TestRequestFileAccess(t *testing.T) {
	req := &object.HTTPRequest{
		Files: map[string]*object.UploadedFile{
			"avatar": {
				Name:     "avatar.jpg",
				Size:     1024,
				MimeType: "image/jpeg",
				Content:  []byte("fake image data"),
			},
		},
	}

	// Test file method
	fileName := &object.String{Value: "avatar"}
	result := req.Method("file", []object.Object{fileName})
	
	if file, ok := result.(*object.UploadedFile); !ok {
		t.Errorf("Expected UploadedFile, got %T", result)
	} else if file.Name != "avatar.jpg" {
		t.Errorf("Expected file name 'avatar.jpg', got '%s'", file.Name)
	}

	// Test files method
	filesResult := req.Method("files", []object.Object{})
	if filesResult.Type() != object.STRING_OBJ {
		t.Errorf("Expected string result for files method, got %T", filesResult)
	}
}

func TestSecurityHeaders(t *testing.T) {
	currentApp = object.NewHTTPApp()
	handler := createHTTPHandler(currentApp)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Add a test route with a simple function
	testHandler := &object.Function{
		Parameters: []*ast.Identifier{{Value: "req"}, {Value: "res"}},
		Body:       &ast.BlockStatement{Statements: []ast.Statement{}}, // Empty block
		Env:        nil,
	}
	currentApp.Routes["GET:/test"] = testHandler

	handler(w, req)

	// Check security headers
	headers := w.Header()
	
	expectedHeaders := map[string]string{
		"X-Content-Type-Options": "nosniff",
		"X-Frame-Options":        "DENY",
		"X-XSS-Protection":       "1; mode=block",
	}

	for key, expectedValue := range expectedHeaders {
		if actualValue := headers.Get(key); actualValue != expectedValue {
			t.Errorf("Expected header %s: %s, got %s", key, expectedValue, actualValue)
		}
	}
}

func TestCORSConfiguration(t *testing.T) {
	currentApp = object.NewHTTPApp()
	handler := createHTTPHandler(currentApp)

	req := httptest.NewRequest("OPTIONS", "/test", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Check CORS headers
	headers := w.Header()
	
	if origin := headers.Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Errorf("Expected CORS origin '*', got '%s'", origin)
	}

	if methods := headers.Get("Access-Control-Allow-Methods"); !strings.Contains(methods, "GET") {
		t.Errorf("Expected CORS methods to contain 'GET', got '%s'", methods)
	}
}

func TestStreamingHandler(t *testing.T) {
	handler := &object.Function{
		Parameters: []*ast.Identifier{
			{Value: "req"},
			{Value: "res"},
		},
		Body: nil,
		Env:  nil,
	}

	result := createStreamHandler([]object.Object{handler}, map[string]object.Object{})
	
	streamHandler, ok := result.(*object.Function)
	if !ok {
		t.Errorf("Expected Function result, got %T", result)
	}

	if !streamHandler.IsStreaming {
		t.Error("Expected handler to be marked as streaming")
	}
}

func TestMetricsEnable(t *testing.T) {
	currentApp = object.NewHTTPApp()

	result := enableMetrics([]object.Object{}, map[string]object.Object{})
	
	if result.Type() != object.STRING_OBJ {
		t.Errorf("Expected string result, got %T", result)
	}

	// Verify metrics were enabled
	if !currentApp.Performance.EnableMetrics {
		t.Error("Expected metrics to be enabled")
	}

	if !currentApp.Performance.RequestTiming {
		t.Error("Expected request timing to be enabled")
	}
}

func TestPerformanceHeaders(t *testing.T) {
	currentApp = object.NewHTTPApp()
	currentApp.Performance.RequestTiming = true
	handler := createHTTPHandler(currentApp)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Add a test route with a simple function
	testHandler := &object.Function{
		Parameters: []*ast.Identifier{{Value: "req"}, {Value: "res"}},
		Body:       &ast.BlockStatement{Statements: []ast.Statement{}},
		Env:        nil,
	}
	currentApp.Routes["GET:/test"] = testHandler

	handler(w, req)

	// Check performance headers
	headers := w.Header()
	
	if requestStart := headers.Get("X-Request-Start"); requestStart == "" {
		t.Error("Expected X-Request-Start header to be set")
	}

	if responseTime := headers.Get("X-Response-Time"); responseTime == "" {
		t.Error("Expected X-Response-Time header to be set")
	}
}
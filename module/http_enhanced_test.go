package module

import (
	"testing"

	"github.com/vintlang/vintlang/object"
)

func TestHTTPAppEnhancedFeatures(t *testing.T) {
	// Reset the current app
	currentApp = nil

	// Test app creation
	result := createApp([]object.VintObject{}, map[string]object.VintObject{})
	if _, ok := result.(*object.String); !ok {
		t.Errorf("Expected string result from createApp, got %T", result)
	}

	if currentApp == nil {
		t.Errorf("Expected currentApp to be set after createApp")
	}

	// Test that enhanced fields are initialized
	if currentApp.Interceptors == nil {
		t.Errorf("Expected Interceptors map to be initialized")
	}
	if currentApp.Guards == nil {
		t.Errorf("Expected Guards slice to be initialized")
	}
}

func TestInterceptors(t *testing.T) {
	// Reset and create app
	currentApp = nil
	createApp([]object.VintObject{}, map[string]object.VintObject{})

	// Test request interceptor
	interceptorFunc := &object.Function{}
	result := addInterceptor([]object.VintObject{
		&object.String{Value: "request"},
		interceptorFunc,
	}, map[string]object.VintObject{})

	if err, ok := result.(*object.Error); ok {
		t.Errorf("Expected successful interceptor registration, got error: %s", err.Message)
	}

	if len(currentApp.Interceptors["request"]) != 1 {
		t.Errorf("Expected 1 request interceptor, got %d", len(currentApp.Interceptors["request"]))
	}

	// Test response interceptor
	result = addInterceptor([]object.VintObject{
		&object.String{Value: "response"},
		interceptorFunc,
	}, map[string]object.VintObject{})

	if err, ok := result.(*object.Error); ok {
		t.Errorf("Expected successful interceptor registration, got error: %s", err.Message)
	}

	if len(currentApp.Interceptors["response"]) != 1 {
		t.Errorf("Expected 1 response interceptor, got %d", len(currentApp.Interceptors["response"]))
	}

	// Test invalid interceptor type
	result = addInterceptor([]object.VintObject{
		&object.String{Value: "invalid"},
		interceptorFunc,
	}, map[string]object.VintObject{})

	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for invalid interceptor type, got %T", result)
	}
}

func TestGuards(t *testing.T) {
	// Reset and create app
	currentApp = nil
	createApp([]object.VintObject{}, map[string]object.VintObject{})

	// Test guard registration
	guardFunc := &object.Function{}
	result := addGuard([]object.VintObject{guardFunc}, map[string]object.VintObject{})

	if err, ok := result.(*object.Error); ok {
		t.Errorf("Expected successful guard registration, got error: %s", err.Message)
	}

	if len(currentApp.Guards) != 1 {
		t.Errorf("Expected 1 guard, got %d", len(currentApp.Guards))
	}

	// Test multiple guards
	result = addGuard([]object.VintObject{guardFunc}, map[string]object.VintObject{})
	if len(currentApp.Guards) != 2 {
		t.Errorf("Expected 2 guards, got %d", len(currentApp.Guards))
	}
}

func TestCORSMiddleware(t *testing.T) {
	// Reset and create app
	currentApp = nil
	createApp([]object.VintObject{}, map[string]object.VintObject{})

	originalMiddlewareCount := len(currentApp.Middleware)

	// Test CORS middleware registration
	result := corsMiddleware([]object.VintObject{}, map[string]object.VintObject{})

	if err, ok := result.(*object.Error); ok {
		t.Errorf("Expected successful CORS middleware registration, got error: %s", err.Message)
	}

	if len(currentApp.Middleware) != originalMiddlewareCount+1 {
		t.Errorf("Expected middleware count to increase by 1, got %d", len(currentApp.Middleware))
	}
}

func TestAuthMiddleware(t *testing.T) {
	// Reset and create app
	currentApp = nil
	createApp([]object.VintObject{}, map[string]object.VintObject{})

	originalMiddlewareCount := len(currentApp.Middleware)

	// Test auth middleware registration
	authFunc := &object.Function{}
	result := authMiddleware([]object.VintObject{authFunc}, map[string]object.VintObject{})

	if err, ok := result.(*object.Error); ok {
		t.Errorf("Expected successful auth middleware registration, got error: %s", err.Message)
	}

	if len(currentApp.Middleware) != originalMiddlewareCount+1 {
		t.Errorf("Expected middleware count to increase by 1, got %d", len(currentApp.Middleware))
	}
}

func TestErrorHandler(t *testing.T) {
	// Reset and create app
	currentApp = nil
	createApp([]object.VintObject{}, map[string]object.VintObject{})

	// Test error handler registration
	errorHandlerFunc := &object.Function{}
	result := setErrorHandler([]object.VintObject{errorHandlerFunc}, map[string]object.VintObject{})

	if err, ok := result.(*object.Error); ok {
		t.Errorf("Expected successful error handler registration, got error: %s", err.Message)
	}

	if currentApp.ErrorHandler == nil {
		t.Errorf("Expected error handler to be set")
	}

	if currentApp.ErrorHandler != errorHandlerFunc {
		t.Errorf("Expected error handler to be the same function passed")
	}
}

func TestErrorConditions(t *testing.T) {
	// Test functions without app
	currentApp = nil

	// Test interceptor without app
	result := addInterceptor([]object.VintObject{
		&object.String{Value: "request"},
		&object.Function{},
	}, map[string]object.VintObject{})

	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error when calling interceptor without app")
	}

	// Test guard without app
	result = addGuard([]object.VintObject{&object.Function{}}, map[string]object.VintObject{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error when calling guard without app")
	}

	// Test CORS without app
	result = corsMiddleware([]object.VintObject{}, map[string]object.VintObject{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error when calling CORS without app")
	}

	// Test auth without app
	result = authMiddleware([]object.VintObject{&object.Function{}}, map[string]object.VintObject{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error when calling auth without app")
	}

	// Test error handler without app
	result = setErrorHandler([]object.VintObject{&object.Function{}}, map[string]object.VintObject{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error when calling errorHandler without app")
	}
}

func TestInvalidArguments(t *testing.T) {
	// Create app first
	currentApp = nil
	createApp([]object.VintObject{}, map[string]object.VintObject{})

	// Test interceptor with wrong number of arguments
	result := addInterceptor([]object.VintObject{
		&object.String{Value: "request"},
	}, map[string]object.VintObject{})

	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for wrong number of arguments to interceptor")
	}

	// Test interceptor with wrong argument types
	result = addInterceptor([]object.VintObject{
		&object.Integer{Value: 123},
		&object.Function{},
	}, map[string]object.VintObject{})

	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for wrong argument type to interceptor")
	}

	// Test guard with wrong argument type
	result = addGuard([]object.VintObject{&object.String{Value: "not a function"}}, map[string]object.VintObject{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for wrong argument type to guard")
	}

	// Test auth with wrong argument type
	result = authMiddleware([]object.VintObject{&object.String{Value: "not a function"}}, map[string]object.VintObject{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for wrong argument type to auth")
	}
}

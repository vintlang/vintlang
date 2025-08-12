# Enterprise HTTP Module Documentation

The Vint HTTP module has been enhanced with enterprise-level features to support building production-ready backend applications. This document covers the advanced features added to the original HTTP module.

## Table of Contents

1. [Overview](#overview)
2. [Route Grouping & API Versioning](#route-grouping--api-versioning)
3. [Multipart File Uploads](#multipart-file-uploads)
4. [Async Handlers](#async-handlers)
5. [Enhanced Security](#enhanced-security)
6. [Advanced Middleware](#advanced-middleware)
7. [Structured Error Handling](#structured-error-handling)
8. [Performance Monitoring](#performance-monitoring)
9. [Complete Examples](#complete-examples)

## Overview

The enterprise HTTP module extends the basic HTTP functionality with:

- **Route Grouping**: Organize routes with common prefixes and middleware
- **File Uploads**: Complete multipart/form-data support with file handling
- **Async Processing**: Non-blocking handlers for long-running operations
- **Security Features**: CSRF protection, security headers, enhanced CORS
- **Middleware Composition**: Advanced middleware stacking and composition
- **Error Handling**: Structured error responses with consistent format
- **Performance Hooks**: Request timing and metrics for APM integration

## Route Grouping & API Versioning

### Creating Route Groups

Route groups allow you to organize related routes under a common prefix:

```vint
import http

http.app()

// Create API v1 group
http.group("/api/v1", func() {
    // All routes in this group will be prefixed with /api/v1
})

// Create API v2 group  
http.group("/api/v2", func() {
    // All routes in this group will be prefixed with /api/v2
})
```

### Nested Route Groups

```vint
// Admin routes group
http.group("/admin", func() {
    // User management routes
    http.group("/users", func() {
        // /admin/users/* routes
    })
    
    // System routes
    http.group("/system", func() {
        // /admin/system/* routes
    })
})
```

## Multipart File Uploads

### Basic File Upload

```vint
http.post("/upload", func(req, res) {
    // Parse multipart form data
    http.multipart(req)
    
    // Access uploaded file
    let avatar = req.file("avatar")
    
    if avatar {
        // Get file information
        let name = avatar.name()
        let size = avatar.size()
        let type = avatar.type()
        
        // Save the file
        let saved = avatar.save("/uploads/" + name)
        
        res.json({
            "success": true,
            "file": {
                "name": name,
                "size": size,
                "type": type,
                "saved": saved
            }
        })
    } else {
        res.status(400).json({
            "error": "No file uploaded"
        })
    }
})
```

### Multiple File Upload

```vint
http.post("/multiple-upload", func(req, res) {
    http.multipart(req)
    
    let uploadedFiles = []
    let files = req.files()
    
    // Process each file
    for file in files {
        let savedPath = file.save("/uploads/" + file.name())
        uploadedFiles.push({
            "name": file.name(),
            "size": file.size(),
            "type": file.type(),
            "path": savedPath
        })
    }
    
    res.json({
        "success": true,
        "count": uploadedFiles.length,
        "files": uploadedFiles
    })
})
```

### Form Data with File Upload

```vint
http.post("/profile", func(req, res) {
    http.multipart(req)
    
    // Access form fields
    let username = req.form("username")
    let email = req.form("email")
    
    // Access uploaded file
    let avatar = req.file("avatar")
    
    if avatar {
        avatar.save("/avatars/" + username + "_" + avatar.name())
    }
    
    res.json({
        "message": "Profile updated",
        "user": {
            "username": username,
            "email": email,
            "avatar": avatar ? avatar.name() : null
        }
    })
})
```

## Async Handlers

Async handlers allow long-running operations without blocking other requests:

```vint
// Create async handler for heavy processing
let processDataAsync = http.async(func(req, res) {
    // This runs asynchronously and won't block other requests
    let data = req.json()
    
    // Simulate heavy processing
    // In real applications: database operations, API calls, image processing
    processLargeDataset(data)
    
    res.json({
        "message": "Processing started",
        "taskId": generateTaskId()
    })
})

http.post("/process", processDataAsync)

// Immediate response handler
http.post("/quick", func(req, res) {
    res.json({"message": "Quick response"})
})
```

## Enhanced Security

### Security Middleware

```vint
// Enable security features
http.security()

// This automatically adds:
// - X-Content-Type-Options: nosniff
// - X-Frame-Options: DENY  
// - X-XSS-Protection: 1; mode=block
// - CSRF protection (when enabled)
```

### Custom Security Headers

```vint
http.use(func(req, res, next) {
    res.header("Strict-Transport-Security", "max-age=31536000")
    res.header("Content-Security-Policy", "default-src 'self'")
    next()
})
```

### CORS Configuration

```vint
// Basic CORS
http.cors()

// Custom CORS (configuration would be enhanced in future)
http.use(func(req, res, next) {
    res.header("Access-Control-Allow-Origin", "https://myapp.com")
    res.header("Access-Control-Allow-Credentials", "true")
    next()
})
```

## Advanced Middleware

### Middleware Composition

```vint
// Authentication middleware
let authMiddleware = func(req, res, next) {
    let token = req.get("Authorization")
    if !token {
        res.status(401).json({"error": "Unauthorized"})
        return
    }
    // Validate token
    next()
}

// Logging middleware
let logMiddleware = func(req, res, next) {
    print("Request: " + req.method() + " " + req.path())
    next()
}

// Rate limiting middleware
let rateLimitMiddleware = func(req, res, next) {
    // Check rate limits
    next()
}

// Apply middleware in order
http.use(logMiddleware)
http.use(rateLimitMiddleware)
http.use(authMiddleware)
```

### Route-Specific Middleware

```vint
// Apply multiple middlewares to specific routes
http.post("/protected", [authMiddleware, rateLimitMiddleware], func(req, res) {
    res.json({"message": "Protected resource accessed"})
})
```

## Structured Error Handling

### Global Error Handler

```vint
http.errorHandler(func(err, req, res) {
    res.status(500).json({
        "error": {
            "type": "INTERNAL_SERVER_ERROR",
            "message": err.message,
            "code": "ERR_INTERNAL",
            "status": 500,
            "details": {
                "timestamp": Date.now(),
                "path": req.path(),
                "method": req.method(),
                "requestId": generateRequestId()
            }
        }
    })
})
```

### Custom Error Responses

```vint
http.get("/users/:id", func(req, res) {
    let userId = req.param("id")
    
    if !isValidId(userId) {
        res.status(400).json({
            "error": {
                "type": "VALIDATION_ERROR",
                "message": "Invalid user ID format",
                "code": "INVALID_USER_ID",
                "status": 400,
                "details": {
                    "field": "id",
                    "value": userId,
                    "expected": "numeric ID"
                }
            }
        })
        return
    }
    
    let user = findUser(userId)
    if !user {
        res.status(404).json({
            "error": {
                "type": "NOT_FOUND",
                "message": "User not found",
                "code": "USER_NOT_FOUND",
                "status": 404,
                "details": {
                    "userId": userId
                }
            }
        })
        return
    }
    
    res.json(user)
})
```

## Performance Monitoring

### Request Timing

```vint
// Middleware to track request timing
let timingMiddleware = func(req, res, next) {
    let startTime = Date.now()
    
    // Add custom response method to track timing
    let originalSend = res.send
    res.send = func(data) {
        let duration = Date.now() - startTime
        res.header("X-Response-Time", duration + "ms")
        originalSend(data)
    }
    
    next()
}

http.use(timingMiddleware)
```

### Metrics Endpoint

```vint
http.get("/metrics", func(req, res) {
    res.header("Content-Type", "text/plain")
    res.send(`
# HTTP Request Count
http_requests_total{method="GET"} 1234
http_requests_total{method="POST"} 567

# HTTP Request Duration
http_request_duration_seconds{quantile="0.5"} 0.05
http_request_duration_seconds{quantile="0.95"} 0.2

# System Metrics
memory_usage_bytes 104857600
cpu_usage_percent 15.5
`)
})
```

## Complete Examples

### Production-Ready API Server

```vint
import http

// Create application
http.app()

// Security setup
http.security()

// Global middleware
http.use(func(req, res, next) {
    print("Request: " + req.method() + " " + req.path())
    res.header("X-Powered-By", "VintLang")
    next()
})

// Authentication middleware
let authMiddleware = func(req, res, next) {
    let token = req.get("Authorization")
    if token && validateJWT(token) {
        next()
    } else {
        res.status(401).json({
            "error": {
                "type": "UNAUTHORIZED",
                "message": "Invalid or missing authentication token",
                "code": "AUTH_REQUIRED"
            }
        })
    }
}

// API v1 routes
http.group("/api/v1", func() {
    // Public routes
    http.post("/auth/login", func(req, res) {
        let credentials = req.json()
        let token = authenticateUser(credentials)
        
        if token {
            res.json({
                "token": token,
                "expires": Date.now() + 3600000
            })
        } else {
            res.status(401).json({
                "error": {
                    "type": "AUTHENTICATION_FAILED",
                    "message": "Invalid credentials"
                }
            })
        }
    })
    
    // Protected routes
    http.use(authMiddleware)
    
    http.get("/users", func(req, res) {
        let page = req.query("page") || "1"
        let users = getUsers(page)
        res.json(users)
    })
    
    http.post("/upload", func(req, res) {
        http.multipart(req)
        let file = req.file("document")
        
        if file {
            let savedPath = file.save("/secure-uploads/" + generateFileName())
            res.json({
                "message": "File uploaded successfully",
                "fileId": generateFileId(savedPath)
            })
        } else {
            res.status(400).json({
                "error": {
                    "type": "VALIDATION_ERROR",
                    "message": "No file provided"
                }
            })
        }
    })
})

// Health check
http.get("/health", func(req, res) {
    res.json({
        "status": "healthy",
        "timestamp": Date.now(),
        "version": "1.0.0"
    })
})

// Error handler
http.errorHandler(func(err, req, res) {
    print("Error: " + err.message)
    res.status(500).json({
        "error": {
            "type": "INTERNAL_SERVER_ERROR",
            "message": "An unexpected error occurred",
            "code": "ERR_INTERNAL"
        }
    })
})

// Start server
http.listen(3000, "Production API server running on port 3000")
```

### File Upload Service

```vint
import http

http.app()
http.security()

// File upload with validation
http.post("/files", func(req, res) {
    http.multipart(req)
    
    let files = req.files()
    let results = []
    
    for file in files {
        // Validate file type
        let allowedTypes = ["image/jpeg", "image/png", "application/pdf"]
        if !allowedTypes.includes(file.type()) {
            results.push({
                "name": file.name(),
                "error": "File type not allowed"
            })
            continue
        }
        
        // Validate file size (10MB max)
        if file.size() > 10485760 {
            results.push({
                "name": file.name(),
                "error": "File too large (max 10MB)"
            })
            continue
        }
        
        // Save file
        let fileName = Date.now() + "_" + file.name()
        let savedPath = file.save("/uploads/" + fileName)
        
        results.push({
            "name": file.name(),
            "fileName": fileName,
            "size": file.size(),
            "type": file.type(),
            "url": "/files/" + fileName,
            "status": "uploaded"
        })
    }
    
    res.json({
        "message": "File upload processed",
        "results": results
    })
})

// Serve uploaded files
http.get("/files/:filename", func(req, res) {
    let filename = req.param("filename")
    // In a real implementation, serve the file from storage
    res.send("File: " + filename)
})

http.listen(3000, "File upload service running on port 3000")
```

## Summary

The enterprise HTTP module provides all the features needed to build production-ready backend applications:

- **Scalable**: Route grouping and middleware composition for large applications
- **Secure**: Built-in security features and customizable protection
- **Performance**: Async handlers and monitoring capabilities  
- **Robust**: Structured error handling and comprehensive file upload support
- **Production-Ready**: All features needed for enterprise backend development

This makes VintLang suitable for building everything from simple APIs to complex enterprise applications with the same level of sophistication as modern frameworks like Express.js, FastAPI, or Spring Boot.
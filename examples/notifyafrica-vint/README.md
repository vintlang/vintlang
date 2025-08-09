# Enhanced NotifyAfrica SMS Pro

## Overview

This is an enhanced version of the NotifyAfrica SMS sender application built with VintLang, showcasing the comprehensive features and capabilities of the VintLang programming language.

## Features Demonstrated

### VintLang Modules Used

1. **notifyafrica_pkg** - Custom SMS API package
2. **os** - File operations and existence checks  
3. **time** - Time formatting and current time operations
4. **uuid** - Unique ID generation for request tracking
5. **dotenv** - Environment variable management
6. **json** - JSON encoding for data structures
7. **crypto** - MD5 hashing for security logging
8. **random** - Random number generation for verification codes
9. **sysinfo** - System information retrieval
10. **string** - String manipulation and type conversion

### Language Features Showcased

- ✅ Functions with parameters and return values
- ✅ Dictionaries and arrays for data structures
- ✅ For loops and while loops for iteration
- ✅ Conditional statements (if/else) for logic flow
- ✅ File I/O operations for logging and data persistence
- ✅ Error handling and validation
- ✅ User input and formatted output
- ✅ Type conversion and string operations

### Application Features

1. **Multi-Method SMS Sending**
   - Direct SMS to multiple recipients
   - Verification code SMS with auto-generated codes
   - Enhanced recipient input with validation

2. **Security & Logging**
   - Secure logging with MD5 hash verification
   - Comprehensive activity logging with timestamps
   - Error tracking and statistics

3. **Dashboard & Analytics**
   - Real-time SMS statistics tracking
   - System information display
   - Environment configuration status
   - Success rate calculations

4. **Enhanced User Experience**
   - Emoji-enhanced interface for better UX
   - Color-coded status messages
   - Progress indicators and request tracking
   - Interactive menus with clear options

5. **Environment Management**
   - Automatic environment configuration detection
   - API token validation
   - Base URL configuration display

## Usage

### Prerequisites

1. Set up your environment variables in `.env` file:
```bash
NOTIFYAFRICA_TOKEN=your_actual_api_token_here
NOTIFYAFRICA_BASEURL=https://notify.africa/api/v2/send-sms
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

### Running the Application

```bash
vint main.vint
```

### Menu Options

1. **📤 Send SMS** - Send SMS to multiple recipients
2. **🎲 Send SMS with Verification Code** - Generate and send verification codes
3. **📊 Dashboard** - View statistics and system information
4. **📋 View Logs** - Check application activity logs
5. **🚀 About VintLang Features** - Learn about features demonstrated
6. **🚪 Exit** - Close the application

## Code Highlights

### Security Features
```vint
// Enhanced logging with MD5 hash verification
let hash = crypto.hashMD5(logMessage + "security_salt")
let fullLog = "[" + timestamp + "] [" + level + "] " + logMessage + " (Hash: " + hash + ")\n"
```

### Random Code Generation
```vint
// Generate secure 6-digit verification codes
let generateVerificationCode = func() {
    return string(random.int(100000, 999999))
}
```

### System Information Display
```vint
// Real-time system information
print("   OS: " + sysinfo.os())
print("   Architecture: " + sysinfo.arch())
print("   Current Time: " + time.format(time.now(), "2006-01-02 15:04:05"))
```

### Environment Configuration
```vint
// Dynamic environment validation
let token = dotenv.get("NOTIFYAFRICA_TOKEN")
if (token != "" && token != "your_api_token_here") {
    print("   🔑 API Token: ✅ Configured")
} else {
    print("   🔑 API Token: ❌ Not configured")
}
```

## VintLang Features Showcase

This application serves as a comprehensive demonstration of VintLang's capabilities:

- **Module System**: Importing and using various built-in and custom modules
- **Error Handling**: Robust error checking and user feedback
- **Data Structures**: Working with dictionaries, arrays, and complex data
- **File Operations**: Reading, writing, and checking file existence
- **String Processing**: Advanced string manipulation and formatting
- **Time Operations**: Date/time formatting and operations
- **Cryptographic Functions**: Hashing for security
- **Random Operations**: Secure random number generation
- **System Integration**: Accessing system information
- **Environment Management**: Configuration through environment variables

## Technical Architecture

The application follows a modular design with clear separation of concerns:

- **Initialization Layer**: Environment setup and validation
- **Business Logic Layer**: SMS sending, verification, and logging
- **Presentation Layer**: User interface and interaction
- **Data Layer**: File operations and statistics
- **Security Layer**: Hashing and validation

This architecture demonstrates VintLang's capability to build well-structured, maintainable applications with enterprise-grade features.

## Output Examples

The application produces colorful, emoji-enhanced output:

```
🚀 Initializing NotifyAfrica SMS Pro...
✅ API token loaded successfully
✅ Initialization complete!

==================================================
📱 NotifyAfrica SMS Pro v2.0.0
==================================================
1. 📤 Send SMS
2. 🎲 Send SMS with Verification Code
3. 📊 Dashboard
4. 📋 View Logs
5. 🚀 About VintLang Features
6. 🚪 Exit
```

This enhanced NotifyAfrica SMS Pro application showcases VintLang as a powerful, feature-rich programming language capable of building professional-grade applications with modern UX and robust functionality.
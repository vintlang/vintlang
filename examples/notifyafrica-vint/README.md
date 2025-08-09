# Enhanced NotifyAfrica SMS Pro

## Overview

This is an enhanced version of the NotifyAfrica SMS sender application built with VintLang, showcasing the comprehensive features and capabilities of the VintLang programming language. The application uses a classic command-line interface for efficient automation and scripting.

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
11. **cli** - Command-line argument parsing

### Language Features Showcased

- ✅ Functions with parameters and return values
- ✅ Dictionaries and arrays for data structures
- ✅ For loops and conditional statements for logic flow
- ✅ File I/O operations for logging and data persistence
- ✅ Error handling and validation
- ✅ Command-line argument processing
- ✅ Type conversion and string operations

### Application Features

1. **Multi-Method SMS Sending**
   - Direct SMS to multiple recipients via command line
   - Verification code SMS with auto-generated codes
   - Enhanced recipient parsing with validation

2. **Security & Logging**
   - Secure logging with MD5 hash verification
   - Comprehensive activity logging with timestamps
   - Error tracking and persistent statistics

3. **Dashboard & Analytics**
   - Real-time SMS statistics tracking
   - System information display
   - Environment configuration status
   - Success rate calculations

4. **Command-Line Interface**
   - Professional CLI with comprehensive help
   - Multiple command modes for different operations
   - Automated scripting support
   - Clear error messages and usage instructions

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

### Command-Line Interface

The application provides a classic terminal command interface:

```bash
# Show help and available commands
vint main.vint help

# Send SMS to multiple recipients
vint main.vint send --sender-id "MySender" --message "Hello World" --recipients "+1234567890,+0987654321"

# Send verification SMS
vint main.vint verify --sender-id "MySender" --phone "+1234567890" --name "John Doe"

# View dashboard with statistics
vint main.vint dashboard

# View application logs
vint main.vint logs

# Show VintLang features demonstrated
vint main.vint features
```

### Command Reference

#### Send SMS
```bash
vint main.vint send [options]
  --sender-id     Sender ID (required)
  --message       SMS message (required)
  --recipients    Comma-separated phone numbers (required)
  --schedule      Schedule (optional, default: none)
```

#### Send Verification SMS
```bash
vint main.vint verify [options]
  --sender-id     Sender ID (required)
  --phone         Recipient phone number (required)
  --name          Recipient name (required)
```

#### Other Commands
- `dashboard` - Show system dashboard and statistics
- `logs` - View application logs  
- `features` - Show VintLang features demonstrated
- `help` - Show help message

## Code Highlights

### Command-Line Argument Processing
```vint
// Parse command-line arguments
let args = cli.getPositional()
let senderId = cli.getArgValue("--sender-id")
let message = cli.getArgValue("--message")
```

### Security Features
```vint
// Enhanced logging with MD5 hash verification
let hash = crypto.hashMD5(logMessage + "security_salt")
let fullLog = "[" + timestamp + "] [" + level + "] " + logMessage + " (Hash: " + hash + ")\n"
```

### Persistent Statistics
```vint
// Load and save statistics with JSON
let loadStats = func() {
    if (os.exists(stats_file)) {
        let statsData = os.readFile(stats_file)
        return json.decode(statsData)
    }
    return {"total_sent": 0, "total_failed": 0}
}
```

### Random Code Generation
```vint
// Generate secure 6-digit verification codes
let generateVerificationCode = func() {
    return string(random.int(100000, 999999))
}
```

## VintLang Features Showcase

This application serves as a comprehensive demonstration of VintLang's capabilities:

- **Module System**: Importing and using various built-in and custom modules
- **CLI Processing**: Professional command-line argument handling
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

- **CLI Layer**: Command-line argument parsing and validation
- **Business Logic Layer**: SMS sending, verification, and logging
- **Data Layer**: File operations and persistent statistics
- **Security Layer**: Hashing and validation
- **Environment Layer**: Configuration management

This architecture demonstrates VintLang's capability to build well-structured, maintainable CLI applications with enterprise-grade features.

## Output Examples

The application produces clean, professional command-line output:

```
NotifyAfrica SMS Pro v2.0.0
Command-line SMS sender showcasing VintLang features

Usage:
  vint main.vint [command] [options]

Commands:
  send              Send SMS
  verify            Send verification SMS
  dashboard         Show system dashboard and statistics
  logs              View application logs
  features          Show VintLang features demonstrated
  help              Show this help message

Examples:
  vint main.vint send --sender-id MySender --message "Hello World" --recipients "+1234567890,+0987654321"
  vint main.vint verify --sender-id MySender --phone "+1234567890" --name "John Doe"
  vint main.vint dashboard
```

This enhanced NotifyAfrica SMS Pro application showcases VintLang as a powerful, feature-rich programming language capable of building professional-grade CLI applications with robust functionality and excellent automation support.
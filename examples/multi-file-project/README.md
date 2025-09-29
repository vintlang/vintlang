# Multi-File Project Error Testing

This directory demonstrates the **filename-in-error-messages** feature across a complex Vint project structure.

## 📁 Project Structure

```
multi-file-project/
├── main.vint              # Entry point with syntax errors
├── config.vint            # Configuration with constant declaration errors  
├── utils/
│   ├── math.vint          # Math utilities with parameter and operator errors
│   └── string_helper.vint # String processing with unterminated string errors
└── test-errors.sh         # Comprehensive test script
```

## 🎯 Testing Filename Error Reporting

### Before Enhancement
```
Line 15: Expected next token to be =, got & instead
Line 21: Expected next token to be =, got STRING instead
```

### After Enhancement  
```
examples/multi-file-project/main.vint:15: Expected next token to be =, got & instead
examples/multi-file-project/utils/math.vint:9: Expected next token to be ), got IDENT instead
examples/multi-file-project/config.vint:15:17: Illegal character '$' - unexpected character
```

## 🏃‍♂️ Running Tests

### Individual File Testing
```bash
# Test main file
go run ../../main.go main.vint

# Test utility files  
go run ../../main.go utils/math.vint
go run ../../main.go utils/string_helper.vint

# Test configuration
go run ../../main.go config.vint
```

### Comprehensive Testing
```bash
# Run the automated test script
./test-errors.sh
```

## 🎨 Error Format Standards

| Error Type | Format | Example |
|------------|--------|---------|
| **Lexer Errors** | `filename:line:column: message` | `main.vint:15:20: Illegal character '$'` |
| **Parser Errors** | `filename:line: message` | `utils/math.vint:9: Expected next token to be )` |
| **String Errors** | `Line X: message` (legacy format) | `Line 32: Unterminated string literal` |

## ✅ Benefits Demonstrated

1. **🎯 Quick File Identification**: Developers immediately know which file contains errors
2. **🔍 Precise Location**: Line and column numbers pinpoint exact error locations  
3. **🏗️ Multi-File Project Support**: Essential for larger projects with multiple modules
4. **🛠️ IDE Integration**: Standard error format works with development tools
5. **📊 Error Categorization**: Clear distinction between lexer and parser errors

## 🔧 Implementation Details

The filename inclusion is achieved through:
- `lexer.NewWithFilename(content, filename)` constructor
- `parser.l.GetFilename()` method for accessing filename in parser
- Updated `repl.ReadWithFilename()` for file execution vs interactive mode
- Consistent error message formatting across all components

This enhancement significantly improves the debugging experience in multi-file Vint projects.
# VintLang Examples

Welcome to the VintLang examples directory! This collection contains 100 example files demonstrating various features and capabilities of the VintLang programming language.

## 📚 What's Inside

This directory contains comprehensive examples covering:

- **Core Language Features**: Variables, control flow, loops, functions
- **Data Structures**: Arrays, dictionaries, strings
- **Modules & Imports**: File organization, package system
- **Built-in Functions**: String manipulation, type conversion, utilities
- **File I/O**: Reading/writing files, directory operations
- **Databases**: SQLite, MySQL, PostgreSQL integration
- **Networking**: HTTP servers, API requests
- **Security**: Encryption, hashing, encoding
- **System Operations**: Shell commands, environment variables
- **Advanced Features**: Pointers, reflection, pattern matching, async operations

## 🚀 Getting Started

### Running Examples

To run any example, use the VintLang interpreter:

```bash
vint examples/example_name.vint
```

### Example Categories

#### Beginner-Friendly Examples
- `builtins.vint` - Built-in functions
- `functions.vint` - Function definitions and usage
- `switch.vint` - Switch-case statements
- `if_expression.vint` - If statements and expressions
- `test-for.vint` - For loops
- `repeat-keyword.vint` - Repeat loops

#### String & Data Manipulation
- `strings.vint` - String module functions
- `nativeStrings.vint` - Native string methods
- `json.vint` - JSON operations
- `csv.vint` - CSV file handling
- `encoding.vint` - Base64 encoding/decoding

#### File & System Operations
- `os.vint` - Operating system operations
- `path.vint` - Path manipulation
- `shell.vint` - Shell command execution

#### Database Examples
- `sqlite.vint` - SQLite database
- `mysql.vint` - MySQL database
- `postgres.vint` - PostgreSQL database

#### Networking & HTTP
- `http.vint` - HTTP file server
- `http_test.vint` - HTTP module testing
- `github-profile.vint` - HTTP requests

#### Advanced Features
- `pointers.vint` - Pointer operations
- `reflect.vint` - Runtime type inspection
- `defer_test.vint` - Defer statement
- `overloading_test.vint` - Function overloading
- `async_simple.vint` - Asynchronous operations

#### Security & Crypto
- `crypto.vint` - Hashing and encryption
- `dotenv.vint` - Environment variables from .env

#### AI/ML Integration
- `llm_openai.vint` - OpenAI GPT integration

## 📖 Documentation

All examples include comprehensive comments explaining:
- What the code does
- How VintLang features work
- Expected output
- Usage patterns

For a detailed analysis of all examples, see [TEST_RESULTS.md](TEST_RESULTS.md).

## ✅ Quality Assurance

All examples have been:
- ✅ Tested with VintLang v0.2.2
- ✅ Fixed for syntax errors
- ✅ Documented with explanatory comments
- ✅ Verified to compile successfully

## 🔧 Testing Examples

### Running Tests

Many examples include the word "test" in their names and demonstrate specific features:

```bash
# Test built-in functions
vint examples/builtins_test.vint

# Test array slicing
vint examples/array_slicing_test.vint

# Test function features
vint examples/function_test.vint
```

### Test Categories

- **builtins_test.vint** - Built-in function testing
- **declaratives_test.vint** - Declarative statements (info, debug, etc.)
- **function_test.vint** - Function default parameters
- **array_slicing_test.vint** - Array slicing operations
- **has_key_test.vint** - Dictionary key checking
- **overloading_test.vint** - Function overloading

## 📝 Notes

### External Dependencies

Some examples require external resources:
- **Database examples**: Need running database servers
- **dotenv.vint**: Needs a .env file
- **llm_openai.vint**: Requires OpenAI API key
- **Networking examples**: May require internet connectivity

These examples will fail expectedly without the required resources but demonstrate correct VintLang syntax.

### Module Availability

Some modules shown in examples may require additional setup or may be in development:
- **desktop** - Desktop GUI module (in development)
- **regex** - Regular expressions (in development)
- **package system** - Advanced package features (in development)

## 🤝 Contributing

When adding new examples:
1. Follow the naming convention (descriptive_name.vint)
2. Add comprehensive comments explaining the code
3. Test the example to ensure it works
4. Update this README if adding a new category
5. Add a description in TEST_RESULTS.md

## 📚 Learning Path

For beginners, we recommend following this learning path:

1. **Basics**: Start with `builtins.vint`, `functions.vint`, `switch.vint`
2. **Control Flow**: Try `if_expression.vint`, `test-for.vint`, `repeat-keyword.vint`
3. **Data Structures**: Explore `strings.vint`, `json.vint`, `array_slicing_test.vint`
4. **File Operations**: Learn from `os.vint`, `path.vint`
5. **Advanced**: Move to `pointers.vint`, `reflect.vint`, `overloading_test.vint`
6. **Modules**: Study various module examples for specific use cases

## 🎯 Quick Reference

| Category | Example Files | Count |
|----------|--------------|-------|
| Total Examples | All *.vint files | 100 |
| Test Files | *test*.vint | 29 |
| Showcase Files | *showcase*.vint | 5 |
| Database Files | sqlite, mysql, postgres | 3 |
| HTTP Files | http*.vint | 4 |

## 💡 Tips

- All examples use proper VintLang syntax as of v0.2.2
- Comments explain not just what code does, but why
- Most examples can be run independently
- Check example output to understand expected behavior
- Modify examples to experiment and learn

## 🔗 Resources

- [VintLang Documentation](https://vintlang.ekilie.com/docs)
- [VintLang GitHub Repository](https://github.com/vintlang/vintlang)
- [TEST_RESULTS.md](TEST_RESULTS.md) - Detailed test results and fixes

---

Happy coding with VintLang! 🎉

# VintLang CLI Development Guide

VintLang is now the ideal language for building simple, interactive terminal applications. This guide shows you how to leverage VintLang's enhanced terminal and CLI capabilities.

## Quick Start

Create a simple CLI tool in just a few lines:

```vintlang
import term
import cli

// Check for help
if (cli.hasArg("--help")) {
    cli.help("MyApp", "My awesome CLI application")
    exit(0)
}

// Display banner
let banner = term.banner("My CLI App")
term.println(banner)

// Interactive menu
let choice = term.select(["Option 1", "Option 2", "Exit"])
term.success("You selected: " + choice)
```

Run with: `vint myapp.vint --help`

## Terminal Module (`term`)

### Display Functions

#### Basic Output
- `term.print(message)` - Print without newline
- `term.println(message)` - Print with newline
- `term.clear()` - Clear terminal screen

#### Styled Output  
- `term.success(message)` - Green success message with checkmark
- `term.error(message)` - Red error message with X mark
- `term.warning(message)` - Yellow warning message with warning icon
- `term.info(message)` - Cyan info message with info icon
- `term.notify(message)` - Blue notification message

#### Visual Components
- `term.banner(text)` - Create bordered banner
- `term.box(text)` - Create boxed text
- `term.table(rows)` - Create formatted table
- `term.chart(data)` - Create simple bar chart

### Interactive Functions

#### Menu Systems
- `term.select(options)` - Single selection menu (numbered)
- `term.menu(items)` - Same as select, different name
- `term.radio(options)` - Radio button selection (numbered)
- `term.checkbox(options)` - Multiple selection (space-separated numbers)

#### User Input
- `term.input(prompt)` - Get text input from user
- `term.password(prompt)` - Get password input (currently visible)
- `term.confirm(question)` - Yes/no confirmation (y/n)

#### Feedback
- `term.loading(message)` - Show loading message
- `term.spinner(message)` - Show spinner with message

### Usage Examples

```vintlang
// Interactive menu
let choice = term.select([
    "Process files",
    "View reports", 
    "Settings",
    "Exit"
])

// Multiple selection
let features = term.checkbox([
    "Auto-save",
    "Dark theme", 
    "Notifications",
    "Advanced mode"
])
// User enters: "1 3" to select items 1 and 3

// Styled messages
term.success("Operation completed!")
term.warning("This action cannot be undone")
term.error("File not found")

// Data visualization
let data = [10, 25, 15, 30, 20]
let chart = term.chart(data)
term.println(chart)

// Table display
let table = term.table([
    ["Name", "Age", "City"],
    ["Alice", "25", "NYC"], 
    ["Bob", "30", "SF"],
    ["Carol", "28", "LA"]
])
term.println(table)
```

## CLI Module (`cli`)

### Argument Parsing

#### Basic Functions
- `cli.hasArg(flag)` - Check if flag exists
- `cli.getArgValue(flag)` - Get flag value
- `cli.getArgs()` - Get all arguments array
- `cli.getPositional()` - Get non-flag arguments

#### Utility Functions
- `cli.help(name, description)` - Generate help text
- `cli.version(name, version)` - Show version info
- `cli.prompt(message)` - Get user input
- `cli.confirm(question)` - Yes/no confirmation
- `cli.cliExit(code)` - Exit with status code

### Usage Examples

```vintlang
import cli

// Check for help/version
if (cli.hasArg("--help") || cli.hasArg("-h")) {
    cli.help("MyTool", "Process files and generate reports")
    exit(0)
}

if (cli.hasArg("--version")) {
    cli.version("MyTool", "2.1.0")
    exit(0)
}

// Get flag values
let inputFile = cli.getArgValue("--input")
let outputFile = cli.getArgValue("--output")
let verbose = cli.hasArg("--verbose")

// Handle different formats
let format = cli.getArgValue("--format")
if (!format) {
    format = "json"  // default
}

// Get positional arguments
let files = cli.getPositional()
if (len(files) == 0) {
    term.error("No input files specified")
    exit(1)
}
```

## Complete CLI Application Template

```vintlang
import term
import cli

// Application info
let APP_NAME = "MyTool"
let APP_VERSION = "1.0.0"
let APP_DESC = "A comprehensive CLI tool"

// Help and version handling
if (cli.hasArg("--help") || cli.hasArg("-h")) {
    cli.help(APP_NAME, APP_DESC)
    term.println("")
    term.info("Custom Options:")
    term.println("  --config FILE    Configuration file")
    term.println("  --dry-run        Show what would be done")
    term.println("  --format TYPE    Output format (json|csv|table)")
    exit(0)
}

if (cli.hasArg("--version") || cli.hasArg("-v")) {
    cli.version(APP_NAME, APP_VERSION)
    exit(0)
}

// Banner
let banner = term.banner(APP_NAME + " v" + APP_VERSION)
term.println(banner)

// Parse arguments
let verbose = cli.hasArg("--verbose")
let dryRun = cli.hasArg("--dry-run")
let configFile = cli.getArgValue("--config")
let format = cli.getArgValue("--format") || "table"

// Show configuration if verbose
if (verbose) {
    term.info("Configuration:")
    let config = term.table([
        ["Setting", "Value"],
        ["Verbose", verbose ? "Yes" : "No"],
        ["Dry Run", dryRun ? "Yes" : "No"],
        ["Config", configFile || "Default"],
        ["Format", format]
    ])
    term.println(config)
}

// Main application logic
term.info("What would you like to do?")
let action = term.select([
    "Process data",
    "Generate report",
    "View settings", 
    "Exit"
])

if (action == "Process data") {
    if (dryRun) {
        term.warning("DRY RUN: Would process data")
    } else {
        term.loading("Processing data...")
        term.success("Data processed successfully")
    }
    
} else if (action == "Generate report") {
    let reportType = term.radio([
        "Summary",
        "Detailed", 
        "Statistics"
    ])
    term.success("Generated " + reportType + " report")
    
} else if (action == "View settings") {
    let settings = term.table([
        ["Setting", "Value", "Source"],
        ["Theme", "Dark", "Default"],
        ["Language", "English", "System"],
        ["Timezone", "UTC", "Auto-detected"]
    ])
    term.println(settings)
}

// Cleanup
if (action != "Exit") {
    let cleanup = term.confirm("Clean up temporary files?")
    if (cleanup) {
        term.success("Cleanup completed")
    }
}

let message = term.box("Thank you for using " + APP_NAME + "!")
term.println(message)
```

## Best Practices

### 1. Always Provide Help
```vintlang
if (cli.hasArg("--help")) {
    cli.help("MyApp", "Brief description of what your app does")
    // Add custom usage info
    term.println("")
    term.info("Examples:")
    term.println("  vint myapp.vint --input data.txt")
    term.println("  vint myapp.vint --format json --verbose")
    exit(0)
}
```

### 2. Use Consistent Styling
```vintlang
// Success operations
term.success("File saved successfully")

// User warnings
term.warning("This will overwrite existing files")

// Error conditions  
term.error("Permission denied")

// Information
term.info("Processing 1,234 records")
```

### 3. Provide Interactive Fallbacks
```vintlang
let inputFile = cli.getArgValue("--input")
if (!inputFile) {
    inputFile = term.input("Enter input file path: ")
}
```

### 4. Validate User Input
```vintlang
let format = cli.getArgValue("--format")
if (format && format != "json" && format != "csv" && format != "xml") {
    term.error("Invalid format: " + format)
    term.info("Supported formats: json, csv, xml")
    exit(1)
}
```

### 5. Use Tables for Structured Data
```vintlang
let results = term.table([
    ["File", "Status", "Records"],
    ["data1.csv", "✓ Processed", "1,234"],
    ["data2.csv", "✗ Error", "0"],
    ["data3.csv", "✓ Processed", "856"]
])
term.println(results)
```

## Example Applications

The `examples/` directory contains complete CLI applications:

- `task_manager_cli.vint` - Task management with interactive menus
- `file_processor_cli.vint` - File processing with progress feedback  
- `system_monitor_cli.vint` - System monitoring with data visualization

Run any example with `--help` to see usage information:

```bash
vint examples/task_manager_cli.vint --help
vint examples/file_processor_cli.vint --version
vint examples/system_monitor_cli.vint --summary
```

## Cross-Platform Notes

VintLang CLI tools work across Linux, macOS, and Windows. The terminal styling uses standard ANSI codes and Unicode characters that work on modern terminals.

For maximum compatibility:
- Use `term.println()` instead of `print()` for better line ending handling
- Test interactive functions on your target platforms
- Provide command-line alternatives to interactive features

## Building and Distribution

To create a standalone executable:

```bash
# Build for current platform
go build -o mytool main.go

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o mytool-linux main.go
GOOS=windows GOARCH=amd64 go build -o mytool.exe main.go  
GOOS=darwin GOARCH=amd64 go build -o mytool-mac main.go
```

Users can then run your tool directly:
```bash
./mytool --help
./mytool --input data.txt --format json
```

## Troubleshooting

### Common Issues

1. **Parsing errors but functionality works**: This is a known issue that doesn't affect functionality
2. **Interactive functions not responding**: Ensure you're running in a proper terminal (not IDE output)
3. **Styling not appearing**: Your terminal may not support ANSI colors

### Debug Mode
Enable verbose output to see what's happening:
```bash
vint myapp.vint --verbose
```

---

**VintLang** - Making CLI development simple and powerful!
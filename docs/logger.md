# Logger Module in Vint

The Logger module in Vint provides structured logging functionality with different severity levels and timestamps. This module helps you log messages to stdout and stderr with proper formatting and timestamps.

---

## Importing the Logger Module

To use the Logger module, simply import it:
```js
import logger
```

---

## Functions and Examples

### 1. Log Info Messages (`info`)
The `info` function logs informational messages to stdout with a timestamp and INFO level indicator.

**Syntax**:
```js
info(message)
```

**Example**:
```js
import logger

logger.info("Application started successfully")
// Output: [2024-01-15 14:30:25] INFO: Application started successfully
```

---

### 2. Log Warning Messages (`warn`)
The `warn` function logs warning messages to stdout with a timestamp and WARN level indicator.

**Syntax**:
```js
warn(message)
```

**Example**:
```js
import logger

logger.warn("This is a warning message")
// Output: [2024-01-15 14:30:25] WARN: This is a warning message
```

---

### 3. Log Error Messages (`error`)
The `error` function logs error messages to stderr with a timestamp and ERROR level indicator.

**Syntax**:
```js
error(message)
```

**Example**:
```js
import logger

logger.error("Something went wrong")
// Output: [2024-01-15 14:30:25] ERROR: Something went wrong
```

---

### 4. Log Debug Messages (`debug`)
The `debug` function logs debug information to stdout with a timestamp and DEBUG level indicator.

**Syntax**:
```js
debug(message)
```

**Example**:
```js
import logger

logger.debug("Debug information for troubleshooting")
// Output: [2024-01-15 14:30:25] DEBUG: Debug information for troubleshooting
```

---

### 5. Log Fatal Messages (`fatal`)
The `fatal` function logs fatal error messages to stderr with a timestamp and FATAL level indicator.

**Syntax**:
```js
fatal(message)
```

**Example**:
```js
import logger

logger.fatal("Critical system failure")
// Output: [2024-01-15 14:30:25] FATAL: Critical system failure
```

---

## Usage Example

```js
import logger

print("=== Logger Module Example ===")

// Log application lifecycle
logger.info("Application starting...")
logger.debug("Loading configuration...")
logger.info("Configuration loaded successfully")

// Simulate some warnings and errors
logger.warn("Low disk space detected")
logger.error("Failed to connect to database")
logger.fatal("Unable to recover from critical error")
```

---

## Summary of Functions

| Function | Description                                           | Output Destination |
|----------|------------------------------------------------------|--------------------|
| `info`   | Logs informational messages with timestamp           | stdout             |
| `warn`   | Logs warning messages with timestamp                  | stdout             |
| `error`  | Logs error messages with timestamp                    | stderr             |
| `debug`  | Logs debug information with timestamp                 | stdout             |
| `fatal`  | Logs fatal error messages with timestamp              | stderr             |

All log messages include timestamps in the format `YYYY-MM-DD HH:MM:SS` for easy tracking and debugging.
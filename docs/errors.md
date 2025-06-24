# Errors Module

The `errors` module provides a way to create new errors.

## `errors.new(message)`

Creates a new error with the given message. This will stop the execution of the script.

### Parameters

- `message` (string): The error message.

### Example

```vint
import "errors"

errors.new("something went wrong")
# The script will stop here and print the error message
``` 
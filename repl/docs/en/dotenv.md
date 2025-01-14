# dotenv Module in Vint

The `dotenv` module in Vint is designed to load environment variables from a `.env` file into the application. This allows you to manage sensitive information such as API keys, database credentials, and other configuration values outside your codebase.

---

## Importing the dotenv Module

To use the `dotenv` module, import it as follows:

```vint
import dotenv
```

---

## Functions and Examples

### 1. Loading Environment Variables with `load()`
The `load` function loads environment variables from a `.env` file into the application's environment. This function should be called at the start of your application to ensure all the necessary environment variables are available.

**Syntax**:
```vint
load(filePath)
```
- `filePath`: The path to the `.env` file (relative or absolute).

**Example**:
```vint
import dotenv

dotenv.load(".env")
```
This loads the environment variables from the `.env` file located in the current directory.

---

### 2. Accessing an Environment Variable with `get()`
After loading the environment variables, you can access specific variables using the `get` function. This function retrieves the value of a given environment variable by its name.

**Syntax**:
```vint
get(variableName)
```
- `variableName`: The name of the environment variable to retrieve.

**Example**:
```vint
import dotenv

dotenv.load(".env")
apiKey = dotenv.get("API_KEY")
print(apiKey)  // Expected output: The value of the "API_KEY" from the .env file
```
In this example, the value of the `API_KEY` environment variable is retrieved and printed.

---

## `.env` File Format

The `.env` file should contain key-value pairs of environment variables. Each line in the file represents a separate environment variable.

**Example `.env` file**:
```
API_KEY=your_api_key_here
DB_HOST=localhost
DB_USER=root
DB_PASS=password123
```

In this case, you would retrieve the value of `API_KEY` as shown in the previous example.

---

## Summary of Functions

| Function           | Description                                             | Example Output                             |
|--------------------|---------------------------------------------------------|--------------------------------------------|
| `load(filePath)`    | Loads environment variables from the specified `.env` file. | No direct output, but environment variables are loaded. |
| `get(variableName)` | Retrieves the value of a specified environment variable.  | The value of the variable (e.g., `"your_api_key_here"`) |

---

The `dotenv` module is an essential tool for securely managing configuration settings in Vint applications. By keeping sensitive data in a `.env` file, you avoid hardcoding secrets into your source code, thus improving security and maintainability.
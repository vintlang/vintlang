# Sysinfo Module in Vint

The `sysinfo` module in Vint provides information about the system, such as the operating system and architecture. This can be useful for system diagnostics, logging, or adapting your program based on the system it’s running on.

---

## Importing the Sysinfo Module

To use the Sysinfo module, import it as follows:

```vint
import sysinfo
```

---

## Functions and Examples

### 1. Get the Operating System (`os`)
The `os` function returns the name of the operating system on which the Vint program is running.

**Syntax**:
```vint
os()
```

- Returns a string representing the operating system (e.g., `"Linux"`, `"Windows"`, `"macOS"`).

**Example**:
```vint
import sysinfo

os_name = sysinfo.os()
print("Operating System:", os_name)
// Output: "Operating System: Linux" (or whatever the actual OS is)
```

---

### 2. Get the System Architecture (`arch`)
The `arch` function returns the architecture of the system (e.g., 32-bit or 64-bit).

**Syntax**:
```vint
arch()
```

- Returns a string representing the architecture (e.g., `"x86_64"` for 64-bit or `"i386"` for 32-bit).

**Example**:
```vint
import sysinfo

architecture = sysinfo.arch()
print("System Architecture:", architecture)
// Output: "System Architecture: x86_64" (or whatever the actual architecture is)
```

---

### 3. Example Combining `os` and `arch`
You can combine both the `os()` and `arch()` functions to display comprehensive system information.

**Example**:
```vint
import sysinfo

os_name = sysinfo.os()
architecture = sysinfo.arch()

print("OS:", os_name, "Arch:", architecture)
// Output: "OS: Linux Arch: x86_64" (or whatever the actual OS and architecture are)
```

This will print out both the operating system and the architecture in a single output.

---

### Summary of Functions

| Function          | Description                                    | Example Output                              |
|-------------------|------------------------------------------------|---------------------------------------------|
| `os`              | Returns the operating system name.             | `"Linux"`, `"Windows"`, `"macOS"`           |
| `arch`            | Returns the system architecture.               | `"x86_64"`, `"i386"`                        |

---

The `sysinfo` module is useful for gathering basic information about the system running your Vint program. This can be used for system compatibility checks, logging system details, or tailoring the behavior of your application based on the underlying system.
# Sysinfo Module in Vint

The `sysinfo` module in Vint provides information about the system, such as the operating system and architecture. This can be useful for system diagnostics, logging, or adapting your program based on the system it’s running on.

---

## Importing the Sysinfo Module

To use the Sysinfo module, import it as follows:

```js
import sysinfo
```

---

## Functions and Examples

### 1. Get the Operating System (`os`)
The `os` function returns the name of the operating system on which the Vint program is running.

**Syntax**:
```js
os()
```

- Returns a string representing the operating system (e.g., `"Linux"`, `"Windows"`, `"macOS"`).

**Example**:
```js
import sysinfo

os_name = sysinfo.os()
print("Operating System:", os_name)
// Output: "Operating System: Linux" (or whatever the actual OS is)
```

---

### 2. Get the System Architecture (`arch`)
The `arch` function returns the architecture of the system (e.g., 32-bit or 64-bit).

**Syntax**:
```js
arch()
```

- Returns a string representing the architecture (e.g., `"x86_64"` for 64-bit or `"i386"` for 32-bit).

**Example**:
```js
import sysinfo

architecture = sysinfo.arch()
print("System Architecture:", architecture)
// Output: "System Architecture: x86_64" (or whatever the actual architecture is)
```

---

### 3. Example Combining `os` and `arch`
You can combine both the `os()` and `arch()` functions to display comprehensive system information.

**Example**:
```js
import sysinfo

os_name = sysinfo.os()
architecture = sysinfo.arch()

print("OS:", os_name, "Arch:", architecture)
// Output: "OS: Linux Arch: x86_64" (or whatever the actual OS and architecture are)
```

This will print out both the operating system and the architecture in a single output.

---

### 4. Get Memory Information (`memInfo`)
The `memInfo` function returns detailed information about system memory usage.

**Syntax**:
```js
memInfo()
```

- Returns a dictionary containing memory statistics in GB and usage percentage.

**Example**:
```js
import sysinfo

let memory = sysinfo.memInfo()
print("Total Memory:", memory["total"])
print("Available Memory:", memory["available"])
print("Used Memory:", memory["used"])
print("Free Memory:", memory["free"])
print("Usage Percentage:", memory["percent"] + "%")
// Output example:
// Total Memory: 15.62 GB
// Available Memory: 14.18 GB
// Used Memory: 1.08 GB
// Free Memory: 11.75 GB
// Usage Percentage: 6.89%
```

---

### 5. Get CPU Information (`cpuInfo`)
The `cpuInfo` function returns detailed information about the CPU.

**Syntax**:
```js
cpuInfo()
```

- Returns a dictionary containing CPU model, cores, frequency, and current usage.

**Example**:
```js
import sysinfo

let cpu = sysinfo.cpuInfo()
print("CPU Model:", cpu["model"])
print("CPU Cores:", cpu["cores"])
print("CPU Frequency:", cpu["frequency"])
print("CPU Usage:", cpu["usage"] + "%")
// Output example:
// CPU Model: AMD EPYC 7763 64-Core Processor
// CPU Cores: 1
// CPU Frequency: 3244.00 MHz
// CPU Usage: 33.33%
```

---

### 6. Get Disk Information (`diskInfo`)
The `diskInfo` function returns information about disk usage for the root filesystem.

**Syntax**:
```js
diskInfo()
```

- Returns a dictionary containing disk space information in GB and usage percentage.

**Example**:
```js
import sysinfo

let disk = sysinfo.diskInfo()
print("Total Disk Space:", disk["total"])
print("Used Disk Space:", disk["used"])
print("Free Disk Space:", disk["free"])
print("Disk Usage:", disk["percent"] + "%")
// Output example:
// Total Disk Space: 71.61 GB
// Used Disk Space: 49.81 GB
// Free Disk Space: 21.78 GB
// Disk Usage: 69.58%
```

---

### 7. Get Network Information (`netInfo`)
The `netInfo` function returns information about all network interfaces with addresses.

**Syntax**:
```js
netInfo()
```

- Returns an array of dictionaries, each containing interface name and addresses.

**Example**:
```js
import sysinfo

let interfaces = sysinfo.netInfo()
print("Number of interfaces:", len(interfaces))
for i = 0; i < len(interfaces); i = i + 1 {
    let iface = interfaces[i]
    print("Interface:", iface["name"])
    print("Addresses:", iface["addrs"])
}
// Output example:
// Number of interfaces: 3
// Interface: lo
// Addresses: [127.0.0.1/8, ::1/128]
// Interface: eth0
// Addresses: [10.1.0.75/20, fe80::7eed:8dff:fe4d:223/64]
// Interface: docker0
// Addresses: [172.17.0.1/16]
```

---

### 8. Comprehensive Example
Here's an example that demonstrates all available sysinfo functions:

**Example**:
```js
import sysinfo

print("=== System Information ===")

// Basic system info
print("OS:", sysinfo.os())
print("Architecture:", sysinfo.arch())

// Memory information
let mem = sysinfo.memInfo()
print("Memory - Total:", mem["total"], "Used:", mem["used"], "Usage:", mem["percent"] + "%")

// CPU information
let cpu = sysinfo.cpuInfo()
print("CPU:", cpu["model"], "(" + cpu["cores"] + " cores)")

// Disk information
let disk = sysinfo.diskInfo()
print("Disk - Total:", disk["total"], "Used:", disk["used"], "Usage:", disk["percent"] + "%")

// Network interfaces
let net = sysinfo.netInfo()
print("Network interfaces:", len(net))
```

---

### Summary of Functions

| Function          | Description                                    | Example Output                              |
|-------------------|------------------------------------------------|---------------------------------------------|
| `os`              | Returns the operating system name.             | `"Linux"`, `"Windows"`, `"macOS"`           |
| `arch`            | Returns the system architecture.               | `"x86_64"`, `"i386"`                        |
| `memInfo`         | Returns memory usage information.              | `{"total": "15.62 GB", "used": "1.08 GB", "percent": "6.89"}` |
| `cpuInfo`         | Returns CPU information and usage.             | `{"model": "AMD EPYC", "cores": "1", "usage": "33.33"}` |
| `diskInfo`        | Returns disk usage information.                | `{"total": "71.61 GB", "used": "49.81 GB", "percent": "69.58"}` |
| `netInfo`         | Returns network interface information.         | `[{"name": "eth0", "addrs": ["10.1.0.75/20"]}]` |

---

The `sysinfo` module is useful for comprehensive system monitoring and diagnostics in your Vint programs. It provides detailed information about memory, CPU, disk, and network resources, making it perfect for system administration tools, monitoring applications, and resource-aware programs.
# Overseer v2.0.0

**A developer-friendly CLI process manager for VintLang.**

Overseer manages long-running services on your machine using **systemd** (Linux)
or **launchd** (macOS). It provides a clean CLI for adding, starting, stopping,
monitoring, and configuring services — all from the terminal.

---

## Features

| Category | Capabilities |
|----------|-------------|
| **Service Management** | Add, start, stop, restart, enable, disable, remove services |
| **Monitoring** | Dashboard overview, per-service health checks via HTTP probes |
| **Environment** | Per-service environment variable management |
| **Tags** | Tag-based service discovery and batch operations |
| **Configuration** | JSON export/import, service descriptions, restart policies |
| **Cross-Platform** | systemd on Linux, launchd on macOS |
| **Build System** | Multi-platform cross-compilation via `build.vint` |

## VintLang Features Demonstrated

This project is a showcase of VintLang's language capabilities:

- **Packages** — Each module is a self-contained package (`package overseer_config { ... }`)
- **Structs** — `ServiceEntry`, `HealthResult`, `EnvVar` structs with methods
- **Enums** — `Platform` and `RestartPolicy` enums for type-safe constants
- **Switch/Case** — Command dispatching and status matching with multi-value cases
- **Modules** — `cli`, `term`, `json`, `os`, `path`, `shell`, `sysinfo`, `regex`, `xml`, `fmt`, `time`
- **First-class Functions** — Commands are functions stored in variables
- **Dictionaries & Arrays** — Service registry, tag lists, environment maps
- **String Methods** — `.trim()`, `.split()`, `.contains()`, `.substring()`
- **Error Handling** — Guard functions, validation, graceful error messages

## Project Structure

```
overseer/
├── main.vint              # CLI entry point & command dispatcher
├── overseer_config.vint   # Configuration, registry, enums, structs
├── overseer_display.vint  # Terminal UI & styled output
├── overseer_service.vint  # systemd/launchd service file generation
├── overseer_systemctl.vint# Service control backend (systemctl/launchctl)
├── overseer_health.vint   # HTTP health check monitoring
├── overseer_env.vint      # Per-service environment variable management
├── build.vint             # Multi-platform cross-compilation build system
└── README.md              # This file
```

## Quick Start

> **Note:** Use `vint main.vint` when running from source.
> After building with `vint build.vint build`, use the `overseer` binary directly.

```bash
# Show help
vint main.vint help

# Add a service
vint main.vint add api "./api --port 8080"

# Start it
vint main.vint start api

# Check status
vint main.vint status api

# View the dashboard
vint main.vint dashboard

# Set environment variables
vint main.vint env set api PORT 8080
vint main.vint env set api NODE_ENV production

# Add tags for grouping
vint main.vint tag add api backend
vint main.vint tag add api production

# Set a health endpoint
vint main.vint describe api "Main API server"

# Run health checks
vint main.vint health

# Export/import configuration
vint main.vint export backup.json
vint main.vint import backup.json

# View logs
vint main.vint logs api

# Stop and remove
vint main.vint stop api
vint main.vint remove api
```

## Commands Reference

### Service Management

| Command | Description |
|---------|-------------|
| `add <name> <command>` | Add and install a new managed service |
| `start <name>` | Start a service |
| `stop <name>` | Stop a service |
| `restart <name>` | Restart a service |
| `enable <name>` | Enable auto-start at login |
| `disable <name>` | Disable auto-start at login |
| `status <name>` | Show detailed service status |
| `logs <name>` | View recent logs (last 50 lines) |
| `list` | List all managed services |
| `remove <name>` | Stop, disable, and remove a service |

### Monitoring & Environment

| Command | Description |
|---------|-------------|
| `dashboard` | Show service overview with status counts |
| `health [name]` | Run health checks (all or specific service) |
| `env list <name>` | List environment variables |
| `env set <name> <KEY> <VALUE>` | Set an environment variable |
| `env rm <name> <KEY>` | Remove an environment variable |
| `tag add <name> <tag>` | Add a tag to a service |
| `tag rm <name> <tag>` | Remove a tag from a service |
| `tag find <tag>` | Find services with a given tag |

### Configuration

| Command | Description |
|---------|-------------|
| `describe <name> [text]` | Set or show service description |
| `export [file]` | Export configuration to JSON (stdout or file) |
| `import <file>` | Import configuration from a JSON file |
| `info` | Show overseer paths and system information |

## Building Standalone Binaries

The `build.vint` script uses VintLang's bundler for cross-compilation:

```bash
# Build for current platform
vint build.vint build

# Build for all platforms
vint build.vint build-all

# Platform-specific builds
vint build.vint build-darwin
vint build.vint build-linux
vint build.vint build-windows

# Create a GitHub release
vint build.vint release
```

## Architecture

### Data Model

Services are stored in `~/.config/overseer/registry.json` with the following
structure per entry:

```json
{
  "name": "api",
  "command": "./api --port 8080",
  "platform": "linux",
  "created_at": "2025-01-15 10:30:00",
  "tags": ["backend", "production"],
  "env": { "PORT": "8080", "NODE_ENV": "production" },
  "health_endpoint": "http://localhost:8080/health",
  "restart_policy": "always",
  "max_retries": 3,
  "description": "Main API server"
}
```

### Enums

```
Platform:       LINUX = "linux",  DARWIN = "darwin"
RestartPolicy:  ALWAYS = "always", ON_FAILURE = "on-failure", NEVER = "never"
```

### Health Checks

When a service has a `health_endpoint` configured, the `health` command sends
an HTTP GET request (via `curl`) with a 5-second timeout. Status codes 200, 201,
and 204 are considered healthy.

## Requirements

- **VintLang** — Install from [github.com/vintlang/vintlang](https://github.com/vintlang/vintlang)
- **Linux**: systemd with user session support
- **macOS**: launchd (built-in)
- **For building**: Go compiler + VintLang bundler

## License

Part of the VintLang examples collection. See the main repository for license details.

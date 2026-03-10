# Overseer - Process Manager

**Overseer** is a developer-friendly CLI process manager built on top of systemd. It wraps the complexity of writing systemd service files and running `systemctl` commands into a clean, simple interface.

---

## Overview

Instead of writing a service file like this:

```ini
# /etc/systemd/system/api.service
[Unit]
Description=My API

[Service]
ExecStart=/home/user/api
Restart=always
User=user

[Install]
WantedBy=multi-user.target
```

You use overseer:

```sh
vint main.vint add api "./api --port 8080"
vint main.vint start api
vint main.vint logs api
```

Under the hood, overseer generates the service file, installs it as a **systemd user unit**, and uses `systemctl --user` to manage it.

---

## Requirements

- A Linux system with **systemd** installed
- A user session with systemd support (standard on modern Linux distros)
- VintLang interpreter (`vint`)

### Enabling Services at Boot

To keep services running after you log out, enable lingering for your user:

```sh
loginctl enable-linger $USER
```

---

## Installation

Overseer is a multi-file VintLang project located in `examples/overseer/`. Run it directly with `vint`:

```sh
cd examples/overseer
vint main.vint help
```

---

## Commands

### `add <name> <command>`

Add and install a new managed service.

```sh
vint main.vint add api "./api --port 8080"
vint main.vint add worker "python3 worker.py --queue default"
vint main.vint add nginx "nginx -g 'daemon off;'"
```

This:

1. Generates a systemd unit file at `~/.config/overseer/services/<name>.service`
2. Copies it to `~/.config/systemd/user/<name>.service`
3. Runs `systemctl --user daemon-reload`
4. Enables the service (`systemctl --user enable <name>`)
5. Registers the service in the overseer registry (`~/.config/overseer/registry.json`)

---

### `start <name>`

Start a managed service.

```sh
vint main.vint start api
```

---

### `stop <name>`

Stop a running service.

```sh
vint main.vint stop api
```

---

### `restart <name>`

Restart a service (stop then start).

```sh
vint main.vint restart api
```

---

### `status <name>`

Show the current status and recent events for a service.

```sh
vint main.vint status api
```

---

### `logs <name>`

View the last 50 log lines for a service using `journalctl`.

```sh
vint main.vint logs api
```

---

### `list`

List all services managed by overseer, along with their current systemd status.

```sh
vint main.vint list
```

---

### `enable <name>`

Enable a service to start automatically at login.

```sh
vint main.vint enable api
```

---

### `disable <name>`

Disable a service from starting at login.

```sh
vint main.vint disable api
```

---

### `remove <name>`

Stop, disable, and completely remove a service from overseer and systemd.

```sh
vint main.vint remove api
```

This:

1. Stops the service
2. Disables the service (removes from auto-start)
3. Removes the unit files
4. Reloads systemd
5. Unregisters the service from the overseer registry

---

### `info`

Show overseer configuration paths and system info.

```sh
vint main.vint info
```

---

### `--version`

Show the overseer version.

```sh
vint main.vint --version
```

---

## Service Name Validation

Service names must:

- Start with a letter or digit
- Contain only letters, digits, hyphens (`-`), and underscores (`_`)

Invalid names (e.g., names containing spaces, slashes, or special characters) are rejected to prevent command injection and ensure compatibility with systemd.

---

## Project Structure

```
examples/overseer/
├── main.vint                 # CLI entry point - argument parsing and command dispatch
├── overseer_config.vint      # Service registry management (package)
├── overseer_display.vint     # Terminal output and formatting utilities (package)
├── overseer_service.vint     # Systemd unit file generation and installation (package)
└── overseer_systemctl.vint   # systemctl / journalctl wrappers (package)
```

### VintLang Features Used

| Feature              | Where Used                                                           |
| -------------------- | -------------------------------------------------------------------- |
| `package`            | All 4 sub-modules encapsulate their logic cleanly                    |
| `const`              | Version, colors, config constants                                    |
| `switch`             | Clean command dispatch in `main.vint`                                |
| `import` (built-in)  | `cli`, `os`, `shell`, `json`, `term`, `time`, `path`, `regex`, `fmt` |
| `import` (user pkg)  | The 4 overseer packages imported by `main.vint`                      |
| Package `init`       | Auto-run on import to set up config paths                            |
| Private members      | `_` prefix for internal-only package state                           |
| `path.join()`        | Safe cross-platform path construction                                |
| `regex.match()`      | Service name validation against injection                            |
| `fmt.padRight()`     | Aligned key-value output in status display                           |
| File I/O (`os`)      | Reading/writing service and registry files                           |
| Shell (`shell`)      | Running `systemctl` and `journalctl` commands                        |
| JSON (`json`)        | Persisting the service registry                                      |
| Terminal (`term`)    | Styled output, tables, banners, messages                             |
| CLI (`cli`)          | Argument parsing                                                     |
| Dicts & arrays       | Service registry, table rows, command routing                        |
| Dict methods         | `.has()`, `.keys()`, `.remove()` for registry ops                    |
| Functions / closures | Command handlers, guards, validators                                 |
| String methods       | `.trim()`, `.contains()`, `.split()`, `.substring()`                 |
| Error handling       | Guard functions with early exit                                      |

---

## Service File Format

Generated service files use the systemd **user unit** format:

```ini
[Unit]
Description=Overseer managed service: <name>
After=network.target

[Service]
Type=simple
ExecStart=<your command>
WorkingDirectory=/home/<user>
Environment=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/home/<user>/.local/bin
Restart=on-failure
RestartSec=5
StartLimitIntervalSec=60
StartLimitBurst=3
StandardOutput=journal
StandardError=journal
SyslogIdentifier=<name>

[Install]
WantedBy=default.target
```

Service files are stored in:

- **Overseer copy**: `~/.config/overseer/services/<name>.service`
- **Systemd unit**: `~/.config/systemd/user/<name>.service`

---

## Registry

The service registry is stored as JSON at `~/.config/overseer/registry.json`:

```json
{
  "services": {
    "api": {
      "name": "api",
      "command": "./api --port 8080",
      "created_at": "2024-01-15 10:30:00"
    }
  }
}
```

---

## Examples

### Full Lifecycle

```sh
# Navigate to the overseer directory
cd examples/overseer

# Add a service
vint main.vint add api "./api --port 8080"

# Start it
vint main.vint start api

# Check it's running
vint main.vint status api

# View logs
vint main.vint logs api

# Restart after a code update
vint main.vint restart api

# Control auto-start
vint main.vint enable api
vint main.vint disable api

# List all services
vint main.vint list

# Show config paths
vint main.vint info

# Remove it
vint main.vint remove api
```

### Running From Any Directory

Thanks to VintLang's script-relative import resolution, you can also run overseer from any directory by providing the full path:

```sh
vint /path/to/examples/overseer/main.vint list
```

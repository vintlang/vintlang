
# VintSocket Module (Experimental)

The `vintSocket` module provides functions for working with WebSockets. It allows you to create WebSocket servers and connect to WebSocket servers as a client. This module is experimental and its API may change in the future.

## Functions

### `createServer(port)`

Creates a WebSocket server on the specified port.

- `port` (string): The port number to listen on.

**Usage:**

```vint
import vintSocket

vintSocket.createServer("8080")
println("WebSocket server started on port 8080")
```

### `connect(url)`

Connects to a WebSocket server.

- `url` (string): The URL of the WebSocket server (e.g., `"ws://localhost:8080"`).

**Usage:**

```vint
import vintSocket

vintSocket.connect("ws://localhost:8080")
```

### `sendMessage(clientIndex, message)`

Sends a message to a specific connected client.

- `clientIndex` (integer): The index of the client in the list of connections.
- `message` (string): The message to send.

**Usage:**

```vint
import vintSocket

// Assuming a client is connected at index 0
vintSocket.sendMessage(0, "Hello, client!")
```

### `broadcast(message)`

Sends a message to all connected clients.

- `message` (string): The message to send.

**Usage:**

```vint
import vintSocket

vintSocket.broadcast("Hello, everyone!")
```

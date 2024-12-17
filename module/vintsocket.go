package module  //THIS IS AN EXPERIMENTAL MODULE

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"

	"github.com/ekilie/vint-lang/object"
)

var VintSocketFunctions = map[string]object.ModuleFunction{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = struct {
	sync.Mutex
	clients []*websocket.Conn
}{}

func init() {
	VintSocketFunctions["createServer"] = createServer
	VintSocketFunctions["connect"] = connect
	VintSocketFunctions["sendMessage"] = sendMessage
	VintSocketFunctions["broadcast"] = broadcast
}

// Create WebSocket server
func createServer(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "createServer requires one argument: port"}
	}

	port := args[0].Inspect()
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				return
			}
			connections.Lock()
			connections.clients = append(connections.clients, conn)
			connections.Unlock()

			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					break
				}
				conn.WriteMessage(websocket.TextMessage, msg)
			}
		})
		http.ListenAndServe(":"+port, nil)
	}()

	return &object.String{Value: "WebSocket server started on port " + port}
}

// Connect to WebSocket server
func connect(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "connect requires one argument: URL"}
	}

	url := args[0].Inspect()
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return &object.Error{Message: "Failed to connect: " + err.Error()}
	}

	connections.Lock()
	connections.clients = append(connections.clients, conn)
	connections.Unlock()

	return &object.String{Value: "Connected to " + url}
}

// Send message to a specific client
func sendMessage(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "sendMessage requires two arguments: client index and message"}
	}

	index := int(args[0].(*object.Integer).Value)
	message := args[1].Inspect()

	connections.Lock()
	if index >= len(connections.clients) {
		connections.Unlock()
		return &object.Error{Message: "Client index out of range"}
	}
	conn := connections.clients[index]
	connections.Unlock()

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return &object.Error{Message: "Failed to send message: " + err.Error()}
	}

	return &object.Boolean{Value: true}
}

// Broadcast message to all clients
func broadcast(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "broadcast requires one argument: message"}
	}

	message := args[0].Inspect()

	connections.Lock()
	for _, conn := range connections.clients {
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
	connections.Unlock()

	return &object.Boolean{Value: true}
}

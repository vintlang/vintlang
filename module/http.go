package module

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ekilie/vint-lang/object"
)

var HttpFunctions = map[string]object.ModuleFunction{}

func init() {
	HttpFunctions["createServer"] = createServer
	HttpFunctions["listen"] = listen
}

type HttpServer struct {
	routes map[string]map[string]*object.Function // Method -> Path -> Handler
}

// Global map of created servers
var servers = map[string]*HttpServer{}

func createServer(args []object.Object, env *object.Environment) object.Object {
	serverID := "default"
	if len(args) > 0 {
		id, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "Server ID must be a string"}
		}
		serverID = id.Value
	}

	if _, exists := servers[serverID]; exists {
		return &object.Error{Message: "Server already exists with ID: " + serverID}
	}

	servers[serverID] = &HttpServer{routes: make(map[string]map[string]*object.Function)}
	return &object.String{Value: serverID}
}

func listen(args []object.Object, env *object.Environment) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Usage: listen(serverID, port)"}
	}

	serverID, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Server ID must be a string"}
	}

	port, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Port must be a string"}
	}

	server, exists := servers[serverID.Value]
	if !exists {
		return &object.Error{Message: "No server found with ID: " + serverID.Value}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method := strings.ToUpper(r.Method)
		path := r.URL.Path

		if server.routes[method] == nil || server.routes[method][path] == nil {
			http.NotFound(w, r)
			return
		}

		handler := server.routes[method][path]

		// Prepare arguments for the handler function
		params := map[string]object.Object{
			"method": &object.String{Value: method},
			"path":   &object.String{Value: path},
		}

		// Parse JSON body if provided
		if r.Body != nil && r.Method == "POST" {
			defer r.Body.Close()
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err == nil {
				params["body"] = convertGoMapToDict(body)
			}
		}

		handlerEnv := object.NewEnclosedEnvironment(handler.Env)
		for key, value := range params {
			handlerEnv.Set(key, value)
		}

		// Call the handler function
		result := Eval(handler.Body, handlerEnv)
		if errObj, ok := result.(*object.Error); ok {
			http.Error(w, errObj.Message, http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result.Inspect()))
		}
	})

	go func() {
		if err := http.ListenAndServe(":"+port.Value, nil); err != nil {
			panic(err)
		}
	}()

	return &object.String{Value: "Server listening on port " + port.Value}
}

func (s *HttpServer) AddRoute(method, path string, handler *object.Function) object.Object {
	if s.routes[method] == nil {
		s.routes[method] = make(map[string]*object.Function)
	}
	s.routes[method][path] = handler
	return &object.Boolean{Value: true}
}

// Converts Go map to object.Dict
func convertGoMapToDict(data map[string]interface{}) *object.Dict {
	dict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	for key, value := range data {
		strKey := &object.String{Value: key}
		var valObj object.Object

		switch v := value.(type) {
		case string:
			valObj = &object.String{Value: v}
		case float64:
			valObj = &object.Float{Value: v}
		case bool:
			valObj = &object.Boolean{Value: v}
		case map[string]interface{}:
			valObj = convertGoMapToDict(v)
		default:
			valObj = &object.Null{}
		}

		hashKey := strKey.HashKey()
		dict.Pairs[hashKey] = object.DictPair{Key: strKey, Value: valObj}
	}
	return dict
}

package module

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekilie/vint-lang/object"
)

// Exported HTTP functions for Vint
var HttpFunctions = map[string]object.ModuleFunction{}

func init() {
	HttpFunctions["fileServer"] = fileServer
}

// fileServer serves files from a specified directory on a given port.
func fileServer(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 || len(args) > 3 {
		return &object.Error{Message: "Usage: http.fileServer(port, directory, [message])"}
	}

	// Validates port argument
	port, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Port must be a string"}
	}

	// Validates directory argument
	directory, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Directory must be a string"}
	}

	// Determines custom or default message
	var message string
	if len(args) == 3 {
		customMessage, ok := args[2].(*object.String)
		if ok {
			message = customMessage.Value
		} else {
			return &object.Error{Message: "Message must be a string"}
		}
	} else {
		message = fmt.Sprintf("Server started on port %s serving files from %s", port.Value, directory.Value)
	}

	// Creates a file server
	fs := http.FileServer(http.Dir(directory.Value))

	// Sets up the HTTP handler
	http.Handle("/", fs)

	// Starts the server in a new goroutine
	go func() {
		fmt.Println(message)
		if err := http.ListenAndServe(":"+port.Value, nil); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}
	}()

	// Waits for interrupt signal to gracefully stop the program
	waitForInterrupt()

	return &object.String{Value: "Server stopped on port " + port.Value}
}

// waitForInterrupt blocks until an interrupt signal (Ctrl+C) is received.
func waitForInterrupt() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("\nShutting down server...")
}

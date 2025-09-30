package builtins

import "github.com/vintlang/vintlang/object"

func init() {
	registerChannelBuiltins()
}

func registerChannelBuiltins() {
	RegisterBuiltin("send", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("send() takes exactly 2 arguments (channel, value), got %d", len(args))
			}

			ch, ok := args[0].(*object.Channel)
			if !ok {
				return newError("first argument to send() must be a channel, got %T", args[0])
			}

			if err := ch.Send(args[1]); err != nil {
				return newError("send error: %s", err.Error())
			}

			return NULL
		},
	})

	RegisterBuiltin("receive", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("receive() takes exactly 1 argument (channel), got %d", len(args))
			}

			ch, ok := args[0].(*object.Channel)
			if !ok {
				return newError("argument to receive() must be a channel, got %T", args[0])
			}

			value, ok := ch.Receive()
			if !ok {
				return NULL // Channel closed
			}

			return value
		},
	})

	RegisterBuiltin("close", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("close() takes exactly 1 argument (channel), got %d", len(args))
			}

			ch, ok := args[0].(*object.Channel)
			if !ok {
				return newError("argument to close() must be a channel, got %T", args[0])
			}

			ch.Close()
			return NULL
		},
	})
}
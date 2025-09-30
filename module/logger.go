package module

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vintlang/vintlang/object"
)

var LoggerFunctions = map[string]object.ModuleFunction{}

func init() {
	LoggerFunctions["info"] = logInfo
	LoggerFunctions["warn"] = logWarn
	LoggerFunctions["error"] = logError
	LoggerFunctions["debug"] = logDebug
	LoggerFunctions["fatal"] = logFatal
}

func logInfo(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"logger", "info",
			"1 argument: message (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`logger.info("Application started") -> logs info message`,
		)
	}

	message := args[0]
	if message.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"logger", "info",
			"string message",
			string(message.Type()),
			`logger.info("Application started") -> logs info message`,
		)
	}

	msg := message.(*object.String).Value
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] INFO: %s", timestamp, msg)

	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.Print(logMsg)

	return &object.String{Value: logMsg}
}

func logWarn(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"logger", "warn",
			"1 argument: message (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`logger.warn("Warning message") -> logs warning message`,
		)
	}

	message := args[0]
	if message.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"logger", "warn",
			"string message",
			string(message.Type()),
			`logger.warn("Warning message") -> logs warning message`,
		)
	}

	msg := message.(*object.String).Value
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] WARN: %s", timestamp, msg)

	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.Print(logMsg)

	return &object.String{Value: logMsg}
}

func logError(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"logger", "error",
			"1 argument: message (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`logger.error("Error occurred") -> logs error message`,
		)
	}

	message := args[0]
	if message.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"logger", "error",
			"string message",
			string(message.Type()),
			`logger.error("Error occurred") -> logs error message`,
		)
	}

	msg := message.(*object.String).Value
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] ERROR: %s", timestamp, msg)

	log.SetOutput(os.Stderr)
	log.SetFlags(0)
	log.Print(logMsg)

	return &object.String{Value: logMsg}
}

func logDebug(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"logger", "debug",
			"1 argument: message (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`logger.debug("Debug info") -> logs debug message`,
		)
	}

	message := args[0]
	if message.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"logger", "debug",
			"string message",
			string(message.Type()),
			`logger.debug("Debug info") -> logs debug message`,
		)
	}

	msg := message.(*object.String).Value
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] DEBUG: %s", timestamp, msg)

	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.Print(logMsg)

	return &object.String{Value: logMsg}
}

func logFatal(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"logger", "fatal",
			"1 argument: message (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`logger.fatal("Fatal error") -> logs fatal message and exits`,
		)
	}

	message := args[0]
	if message.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"logger", "fatal",
			"string message",
			string(message.Type()),
			`logger.fatal("Fatal error") -> logs fatal message and exits`,
		)
	}

	msg := message.(*object.String).Value
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] FATAL: %s", timestamp, msg)

	log.SetOutput(os.Stderr)
	log.SetFlags(0)
	log.Print(logMsg)

	// Note: In a real implementation, you might want to handle fatal differently
	// For now, we'll just return the message without exiting
	return &object.String{Value: logMsg}
}

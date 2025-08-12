package module

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var EmailFunctions = map[string]object.ModuleFunction{}

func init() {
	EmailFunctions["validate"] = emailValidate
	EmailFunctions["extractDomain"] = emailExtractDomain
	EmailFunctions["extractUsername"] = emailExtractUsername
	EmailFunctions["normalize"] = emailNormalize
}

func emailValidate(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"email", "validate",
			"1 argument: email address (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`email.validate("user@example.com") -> true/false`,
		)
	}

	emailStr := args[0]
	if emailStr.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"email", "validate",
			"string email address",
			string(emailStr.Type()),
			`email.validate("user@example.com") -> true/false`,
		)
	}

	email := emailStr.(*object.String).Value
	
	// Use Go's mail package for validation
	_, err := mail.ParseAddress(email)
	if err != nil {
		return &object.Boolean{Value: false}
	}
	
	// Additional validation with regex for common patterns
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	isValid := emailRegex.MatchString(email)
	
	return &object.Boolean{Value: isValid}
}

func emailExtractDomain(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"email", "extractDomain",
			"1 argument: email address (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`email.extractDomain("user@example.com") -> "example.com"`,
		)
	}

	emailStr := args[0]
	if emailStr.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"email", "extractDomain",
			"string email address",
			string(emailStr.Type()),
			`email.extractDomain("user@example.com") -> "example.com"`,
		)
	}

	email := emailStr.(*object.String).Value
	
	// Find the @ symbol and extract domain
	atIndex := strings.LastIndex(email, "@")
	if atIndex == -1 || atIndex == len(email)-1 {
		return &object.String{Value: ""}
	}
	
	domain := email[atIndex+1:]
	return &object.String{Value: domain}
}

func emailExtractUsername(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"email", "extractUsername",
			"1 argument: email address (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`email.extractUsername("user@example.com") -> "user"`,
		)
	}

	emailStr := args[0]
	if emailStr.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"email", "extractUsername",
			"string email address",
			string(emailStr.Type()),
			`email.extractUsername("user@example.com") -> "user"`,
		)
	}

	email := emailStr.(*object.String).Value
	
	// Find the @ symbol and extract username
	atIndex := strings.LastIndex(email, "@")
	if atIndex == -1 || atIndex == 0 {
		return &object.String{Value: ""}
	}
	
	username := email[:atIndex]
	return &object.String{Value: username}
}

func emailNormalize(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"email", "normalize",
			"1 argument: email address (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`email.normalize("User@EXAMPLE.COM") -> "user@example.com"`,
		)
	}

	emailStr := args[0]
	if emailStr.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"email", "normalize",
			"string email address",
			string(emailStr.Type()),
			`email.normalize("User@EXAMPLE.COM") -> "user@example.com"`,
		)
	}

	email := emailStr.(*object.String).Value
	
	// Convert to lowercase and trim whitespace
	normalized := strings.ToLower(strings.TrimSpace(email))
	
	// Validate that it's still a valid email after normalization
	_, err := mail.ParseAddress(normalized)
	if err != nil {
		return &object.String{Value: email} // Return original if normalization makes it invalid
	}
	
	return &object.String{Value: normalized}
}
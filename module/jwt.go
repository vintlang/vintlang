package module

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vintlang/vintlang/object"
)

// JwtFunctions is a map that holds the available functions in the JWT module.
var JwtFunctions = map[string]object.ModuleFunction{
	"create":     createJWT,
	"verify":     verifyJWT,
	"decode":     decodeJWT,
	"createHS256": createJWTHS256,
	"verifyHS256": verifyJWTHS256,
}

// createJWT creates a JWT token with the provided payload and secret
// Uses HS256 signing method by default
func createJWT(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"jwt", "create",
			"2 arguments: payload (hash), secret (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`jwt.create({"user": "john", "exp": 1234567890}, "secret") -> "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`,
		)
	}

	if args[0].Type() != object.DICT_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"jwt", "create",
			"payload (dict), secret (string)",
			fmt.Sprintf("got %s, %s", args[0].Type(), args[1].Type()),
			`jwt.create({"user": "john", "exp": 1234567890}, "secret") -> "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`,
		)
	}

	payload := args[0].(*object.Dict)
	secret := args[1].(*object.String).Value

	// Convert payload to jwt.MapClaims
	claims := jwt.MapClaims{}
	for _, pair := range payload.Pairs {
		key := pair.Key.Inspect()
		value := convertToGoValue(pair.Value)
		claims[key] = value
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to create JWT: %s", err.Error())}
	}

	return &object.String{Value: tokenString}
}

// createJWTHS256 creates a JWT token with HS256 signing method explicitly
func createJWTHS256(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || len(args) > 3 {
		return ErrorMessage(
			"jwt", "createHS256",
			"2-3 arguments: payload (hash), secret (string), [expiration_hours (number)]",
			fmt.Sprintf("%d arguments", len(args)),
			`jwt.createHS256({"user": "john"}, "secret", 24) -> "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`,
		)
	}

	if args[0].Type() != object.DICT_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"jwt", "createHS256",
			"payload (dict), secret (string), [expiration_hours (number)]",
			fmt.Sprintf("got %s, %s", args[0].Type(), args[1].Type()),
			`jwt.createHS256({"user": "john"}, "secret", 24) -> "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`,
		)
	}

	payload := args[0].(*object.Dict)
	secret := args[1].(*object.String).Value

	// Convert payload to jwt.MapClaims
	claims := jwt.MapClaims{}
	for _, pair := range payload.Pairs {
		key := pair.Key.Inspect()
		value := convertToGoValue(pair.Value)
		claims[key] = value
	}

	// Add expiration if provided
	if len(args) == 3 {
		if args[2].Type() != object.INTEGER_OBJ && args[2].Type() != object.FLOAT_OBJ {
			return ErrorMessage(
				"jwt", "createHS256",
				"expiration_hours as number",
				fmt.Sprintf("got %s", args[2].Type()),
				`jwt.createHS256({"user": "john"}, "secret", 24) -> "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`,
			)
		}

		var hours float64
		if args[2].Type() == object.INTEGER_OBJ {
			hours = float64(args[2].(*object.Integer).Value)
		} else {
			hours = args[2].(*object.Float).Value
		}

		claims["exp"] = time.Now().Add(time.Duration(hours) * time.Hour).Unix()
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to create JWT: %s", err.Error())}
	}

	return &object.String{Value: tokenString}
}

// verifyJWT verifies a JWT token and returns the payload if valid
func verifyJWT(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"jwt", "verify",
			"2 arguments: token (string), secret (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`jwt.verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", "secret") -> {"user": "john", "exp": 1234567890}`,
		)
	}

	if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"jwt", "verify",
			"token (string), secret (string)",
			fmt.Sprintf("got %s, %s", args[0].Type(), args[1].Type()),
			`jwt.verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", "secret") -> {"user": "john", "exp": 1234567890}`,
		)
	}

	tokenString := args[0].(*object.String).Value
	secret := args[1].(*object.String).Value

	// Parse and verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to verify JWT: %s", err.Error())}
	}

	if !token.Valid {
		return &object.Error{Message: "Invalid JWT token"}
	}

	// Convert claims to VintLang hash
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &object.Error{Message: "Failed to extract claims from JWT"}
	}

	return convertClaimsToHash(claims)
}

// verifyJWTHS256 verifies a JWT token with HS256 signing method explicitly
func verifyJWTHS256(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"jwt", "verifyHS256",
			"2 arguments: token (string), secret (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`jwt.verifyHS256("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", "secret") -> {"user": "john", "exp": 1234567890}`,
		)
	}

	if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"jwt", "verifyHS256",
			"token (string), secret (string)",
			fmt.Sprintf("got %s, %s", args[0].Type(), args[1].Type()),
			`jwt.verifyHS256("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", "secret") -> {"user": "john", "exp": 1234567890}`,
		)
	}

	tokenString := args[0].(*object.String).Value
	secret := args[1].(*object.String).Value

	// Parse and verify the token with explicit HS256 method
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: expected HS256, got %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to verify JWT: %s", err.Error())}
	}

	if !token.Valid {
		return &object.Error{Message: "Invalid JWT token"}
	}

	// Convert claims to VintLang hash
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &object.Error{Message: "Failed to extract claims from JWT"}
	}

	return convertClaimsToHash(claims)
}

// decodeJWT decodes a JWT token without verification (useful for inspecting headers/payload)
func decodeJWT(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"jwt", "decode",
			"1 argument: token (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`jwt.decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...") -> {"header": {...}, "payload": {...}}`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"jwt", "decode",
			"token (string)",
			fmt.Sprintf("got %s", args[0].Type()),
			`jwt.decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...") -> {"header": {...}, "payload": {...}}`,
		)
	}

	tokenString := args[0].(*object.String).Value

	// Parse without verification
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to decode JWT: %s", err.Error())}
	}

	// Create result dict with header and payload
	pairs := make(map[object.HashKey]object.DictPair)

	// Add header
	headerHash := convertMapToHash(token.Header)
	headerKey := (&object.String{Value: "header"}).HashKey()
	pairs[headerKey] = object.DictPair{Key: &object.String{Value: "header"}, Value: headerHash}

	// Add payload
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		payloadHash := convertClaimsToHash(claims)
		payloadKey := (&object.String{Value: "payload"}).HashKey()
		pairs[payloadKey] = object.DictPair{Key: &object.String{Value: "payload"}, Value: payloadHash}
	}

	return &object.Dict{Pairs: pairs}
}

// Helper function to convert VintLang object to Go value
func convertToGoValue(obj object.VintObject) interface{} {
	switch o := obj.(type) {
	case *object.Integer:
		return o.Value
	case *object.Float:
		return o.Value
	case *object.String:
		return o.Value
	case *object.Boolean:
		return o.Value
	default:
		return o.Inspect()
	}
}

// Helper function to convert jwt.MapClaims to VintLang Dict
func convertClaimsToHash(claims jwt.MapClaims) *object.Dict {
	pairs := make(map[object.HashKey]object.DictPair)

	for key, value := range claims {
		keyObj := &object.String{Value: key}
		var valueObj object.VintObject

		switch v := value.(type) {
		case string:
			valueObj = &object.String{Value: v}
		case float64:
			// Check if it's an integer
			if v == float64(int64(v)) {
				valueObj = &object.Integer{Value: int64(v)}
			} else {
				valueObj = &object.Float{Value: v}
			}
		case bool:
			valueObj = &object.Boolean{Value: v}
		case int64:
			valueObj = &object.Integer{Value: v}
		default:
			valueObj = &object.String{Value: fmt.Sprintf("%v", v)}
		}

		hashKey := keyObj.HashKey()
		pairs[hashKey] = object.DictPair{Key: keyObj, Value: valueObj}
	}

	return &object.Dict{Pairs: pairs}
}

// Helper function to convert map[string]interface{} to VintLang Dict
func convertMapToHash(m map[string]interface{}) *object.Dict {
	pairs := make(map[object.HashKey]object.DictPair)

	for key, value := range m {
		keyObj := &object.String{Value: key}
		var valueObj object.VintObject

		switch v := value.(type) {
		case string:
			valueObj = &object.String{Value: v}
		case float64:
			if v == float64(int64(v)) {
				valueObj = &object.Integer{Value: int64(v)}
			} else {
				valueObj = &object.Float{Value: v}
			}
		case bool:
			valueObj = &object.Boolean{Value: v}
		case int64:
			valueObj = &object.Integer{Value: v}
		default:
			valueObj = &object.String{Value: fmt.Sprintf("%v", v)}
		}

		hashKey := keyObj.HashKey()
		pairs[hashKey] = object.DictPair{Key: keyObj, Value: valueObj}
	}

	return &object.Dict{Pairs: pairs}
}
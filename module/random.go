package module

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	mathrand "math/rand"
	"time"

	"github.com/vintlang/vintlang/object"
)

var RandomFunctions = map[string]object.ModuleFunction{}

func init() {
	mathrand.Seed(time.Now().UnixNano())
	RandomFunctions["int"] = randomInt
	RandomFunctions["float"] = randomFloat
	RandomFunctions["string"] = randomString
	RandomFunctions["choice"] = randomChoice
	RandomFunctions["otp"] = randomOTP
	RandomFunctions["token"] = randomToken
	RandomFunctions["password"] = randomPassword
	RandomFunctions["shuffle"] = randomShuffle
	RandomFunctions["sample"] = randomSample
	RandomFunctions["bool"] = randomBool
	RandomFunctions["range"] = randomRange
}

func randomInt(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.INTEGER_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"random",
			"int",
			"2 integer arguments (min, max)",
			formatArgs(args),
			`random.int(1, 10) -> 7`,
		)
	}
	min := int(args[0].(*object.Integer).Value)
	max := int(args[1].(*object.Integer).Value)
	if min > max {
		return ErrorMessage(
			"random",
			"int",
			"min <= max",
			fmt.Sprintf("min=%d, max=%d", min, max),
			`random.int(1, 10) -> 7`,
		)
	}
	return &object.Integer{Value: int64(mathrand.Intn(max-min+1) + min)}
}

func randomFloat(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"random",
			"float",
			"no arguments",
			formatArgs(args),
			`random.float() -> 0.527391`,
		)
	}
	return &object.Float{Value: mathrand.Float64()}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digitBytes = "0123456789"
const specialBytes = "!@#$%^&*()-_=+[]{}|;:,.<>?"
const passwordBytes = letterBytes + digitBytes + specialBytes

func randomString(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"random",
			"string",
			"1 integer argument (length)",
			formatArgs(args),
			`random.string(8) -> "aZxRtQwe"`,
		)
	}
	n := int(args[0].(*object.Integer).Value)
	if n < 0 {
		return ErrorMessage(
			"random",
			"string",
			"a non-negative integer",
			fmt.Sprintf("%d", n),
			`random.string(8) -> "aZxRtQwe"`,
		)
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[mathrand.Intn(len(letterBytes))]
	}
	return &object.String{Value: string(b)}
}

func randomChoice(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.ARRAY_OBJ {
		return ErrorMessage(
			"random",
			"choice",
			"1 array argument",
			formatArgs(args),
			`random.choice(["apple", "banana", "cherry"]) -> "banana"`,
		)
	}
	arr := args[0].(*object.Array)
	if len(arr.Elements) == 0 {
		return &object.Null{}
	}
	index := mathrand.Intn(len(arr.Elements))
	return arr.Elements[index]
}

// randomOTP generates a numeric OTP of specified length using crypto/rand
func randomOTP(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"random",
			"otp",
			"1 integer argument (length)",
			formatArgs(args),
			`random.otp(6) -> "482031"`,
		)
	}
	n := int(args[0].(*object.Integer).Value)
	if n <= 0 {
		return ErrorMessage(
			"random",
			"otp",
			"a positive integer",
			fmt.Sprintf("%d", n),
			`random.otp(6) -> "482031"`,
		)
	}
	b := make([]byte, n)
	for i := range b {
		randomByte := make([]byte, 1)
		if _, err := rand.Read(randomByte); err != nil {
			return &object.Error{Message: "failed to generate secure random OTP"}
		}
		b[i] = digitBytes[int(randomByte[0])%len(digitBytes)]
	}
	return &object.String{Value: string(b)}
}

// randomToken generates a cryptographically secure hex token
func randomToken(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"random",
			"token",
			"1 integer argument (byte length)",
			formatArgs(args),
			`random.token(16) -> "a3f1b2c4d5e6f7089012abcd3456ef78"`,
		)
	}
	n := int(args[0].(*object.Integer).Value)
	if n <= 0 {
		return ErrorMessage(
			"random",
			"token",
			"a positive integer",
			fmt.Sprintf("%d", n),
			`random.token(16) -> "a3f1b2c4d5e6f7089012abcd3456ef78"`,
		)
	}
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return &object.Error{Message: "failed to generate secure random token"}
	}
	return &object.String{Value: hex.EncodeToString(b)}
}

// randomPassword generates a random password with mixed character types
func randomPassword(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"random",
			"password",
			"1 integer argument (length)",
			formatArgs(args),
			`random.password(16) -> "aB3!xK9@mP2#nQ5&"`,
		)
	}
	n := int(args[0].(*object.Integer).Value)
	if n < 4 {
		return ErrorMessage(
			"random",
			"password",
			"an integer >= 4 (to include all character types)",
			fmt.Sprintf("%d", n),
			`random.password(16) -> "aB3!xK9@mP2#nQ5&"`,
		)
	}

	// Ensure at least one of each character type using crypto/rand
	password := make([]byte, n)
	charSets := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		digitBytes,
		specialBytes,
	}
	// Place one char from each required set
	for i := 0; i < 4; i++ {
		randomByte := make([]byte, 1)
		if _, err := rand.Read(randomByte); err != nil {
			return &object.Error{Message: "failed to generate secure random password"}
		}
		password[i] = charSets[i][int(randomByte[0])%len(charSets[i])]
	}
	// Fill remaining with random chars from all sets
	for i := 4; i < n; i++ {
		randomByte := make([]byte, 1)
		if _, err := rand.Read(randomByte); err != nil {
			return &object.Error{Message: "failed to generate secure random password"}
		}
		password[i] = passwordBytes[int(randomByte[0])%len(passwordBytes)]
	}
	// Shuffle to avoid predictable positions
	for i := n - 1; i > 0; i-- {
		randomByte := make([]byte, 1)
		if _, err := rand.Read(randomByte); err != nil {
			return &object.Error{Message: "failed to generate secure random password"}
		}
		j := int(randomByte[0]) % (i + 1)
		password[i], password[j] = password[j], password[i]
	}
	return &object.String{Value: string(password)}
}

// randomShuffle returns a new shuffled copy of the array
func randomShuffle(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.ARRAY_OBJ {
		return ErrorMessage(
			"random",
			"shuffle",
			"1 array argument",
			formatArgs(args),
			`random.shuffle([1, 2, 3, 4, 5]) -> [3, 1, 5, 2, 4]`,
		)
	}
	arr := args[0].(*object.Array)
	elements := make([]object.VintObject, len(arr.Elements))
	copy(elements, arr.Elements)
	mathrand.Shuffle(len(elements), func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return &object.Array{Elements: elements}
}

// randomSample returns N unique random elements from an array
func randomSample(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.ARRAY_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"random",
			"sample",
			"1 array argument and 1 integer argument (count)",
			formatArgs(args),
			`random.sample([1, 2, 3, 4, 5], 3) -> [2, 5, 1]`,
		)
	}
	arr := args[0].(*object.Array)
	count := int(args[1].(*object.Integer).Value)
	if count < 0 || count > len(arr.Elements) {
		return ErrorMessage(
			"random",
			"sample",
			fmt.Sprintf("count between 0 and %d (array length)", len(arr.Elements)),
			fmt.Sprintf("%d", count),
			`random.sample([1, 2, 3, 4, 5], 3) -> [2, 5, 1]`,
		)
	}
	// Fisher-Yates partial shuffle
	elements := make([]object.VintObject, len(arr.Elements))
	copy(elements, arr.Elements)
	for i := 0; i < count; i++ {
		j := i + mathrand.Intn(len(elements)-i)
		elements[i], elements[j] = elements[j], elements[i]
	}
	return &object.Array{Elements: elements[:count]}
}

// randomBool returns a random boolean
func randomBool(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"random",
			"bool",
			"no arguments",
			formatArgs(args),
			`random.bool() -> true`,
		)
	}
	return &object.Boolean{Value: mathrand.Intn(2) == 1}
}

// randomRange returns an array of random integers in range [min, max]
func randomRange(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.INTEGER_OBJ || args[1].Type() != object.INTEGER_OBJ || args[2].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"random",
			"range",
			"3 integer arguments (min, max, count)",
			formatArgs(args),
			`random.range(1, 100, 5) -> [42, 17, 88, 3, 56]`,
		)
	}
	min := int(args[0].(*object.Integer).Value)
	max := int(args[1].(*object.Integer).Value)
	count := int(args[2].(*object.Integer).Value)
	if min > max {
		return ErrorMessage(
			"random",
			"range",
			"min <= max",
			fmt.Sprintf("min=%d, max=%d", min, max),
			`random.range(1, 100, 5) -> [42, 17, 88, 3, 56]`,
		)
	}
	if count < 0 {
		return ErrorMessage(
			"random",
			"range",
			"a non-negative count",
			fmt.Sprintf("%d", count),
			`random.range(1, 100, 5) -> [42, 17, 88, 3, 56]`,
		)
	}
	elements := make([]object.VintObject, count)
	for i := 0; i < count; i++ {
		elements[i] = &object.Integer{Value: int64(mathrand.Intn(max-min+1) + min)}
	}
	return &object.Array{Elements: elements}
}

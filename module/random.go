package module

import (
	"math/rand"
	"time"

	"github.com/vintlang/vintlang/object"
)

var RandomFunctions = map[string]object.ModuleFunction{}

func init() {
	rand.Seed(time.Now().UnixNano())
	RandomFunctions["int"] = randomInt
	RandomFunctions["float"] = randomFloat
	RandomFunctions["string"] = randomString
	RandomFunctions["choice"] = randomChoice
}

func randomInt(args []object.Object, defs map[string]object.Object) object.Object {
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
	return &object.Integer{Value: int64(rand.Intn(max-min+1) + min)}
}

func randomFloat(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return ErrorMessage(
			"random",
			"float",
			"no arguments",
			formatArgs(args),
			`random.float() -> 0.527391`,
		)
	}
	return &object.Float{Value: rand.Float64()}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(args []object.Object, defs map[string]object.Object) object.Object {
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
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return &object.String{Value: string(b)}
}

func randomChoice(args []object.Object, defs map[string]object.Object) object.Object {
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
	index := rand.Intn(len(arr.Elements))
	return arr.Elements[index]
}

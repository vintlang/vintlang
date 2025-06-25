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
	min, max := 0, 1
	if len(args) == 2 {
		minArg, ok1 := args[0].(*object.Integer)
		maxArg, ok2 := args[1].(*object.Integer)
		if !ok1 || !ok2 {
			return &object.Error{Message: "int() expects two integer arguments: min and max"}
		}
		min = int(minArg.Value)
		max = int(maxArg.Value)
	}
	return &object.Integer{Value: int64(rand.Intn(max-min+1) + min)}
}

func randomFloat(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return &object.Error{Message: "float() expects no arguments"}
	}
	return &object.Float{Value: rand.Float64()}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(args []object.Object, defs map[string]object.Object) object.Object {
	n := 10
	if len(args) == 1 {
		lenArg, ok := args[0].(*object.Integer)
		if !ok {
			return &object.Error{Message: "string() expects an integer argument for length"}
		}
		n = int(lenArg.Value)
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return &object.String{Value: string(b)}
}

func randomChoice(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || args[0].Type() != object.ARRAY_OBJ {
		return &object.Error{Message: "choice() expects a single array argument"}
	}
	arr := args[0].(*object.Array)
	if len(arr.Elements) == 0 {
		return &object.Null{}
	}
	index := rand.Intn(len(arr.Elements))
	return arr.Elements[index]
}

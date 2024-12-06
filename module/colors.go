package module

import (
	"fmt"

	"github.com/ekilie/vint-lang/object"
)

var ColorsFunctions = map[string]object.ModuleFunction{}

func init() {
	ColorsFunctions["rgbToHex"] = rgbToHex
}

func rgbToHex(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "rgbToHex requires three arguments: R, G, B values"}
	}

	r, g, b := args[0].Inspect(), args[1].Inspect(), args[2].Inspect()
	hex := fmt.Sprintf("#%02x%02x%02x", r, g, b)
	return &object.String{Value: hex}
}

package module

import (
    "fmt"
    "github.com/vintlang/vintlang/object"
)

var ColorsFunctions = map[string]object.ModuleFunction{}

func init() {
    ColorsFunctions["rgbToHex"] = rgbToHex
}

func rgbToHex(args []object.Object, defs map[string]object.Object) object.Object {
    if len(args) != 3 {
        return ErrorMessage(
            "colors", "rgbToHex",
            "3 arguments: RED (int), GREEN (int), BLUE (int)",
            fmt.Sprintf("%d arguments", len(args)),
            `colors.rgbToHex(255, 0, 128) -> "#FF0080"`,
        )
    }

    for i, arg := range args {
        if arg.Type() != object.INTEGER_OBJ {
            return ErrorMessage(
                "colors", "rgbToHex",
                fmt.Sprintf("integer values for all channels (argument %d)", i+1),
                string(arg.Type()),
                `colors.rgbToHex(255, 0, 128) -> "#FF0080"`,
            )
        }
    }

    r := args[0].(*object.Integer).Value
    g := args[1].(*object.Integer).Value
    b := args[2].(*object.Integer).Value

    if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
        return &object.Error{
            Message: "\033[1;31mError in colors.rgbToHex()\033[0m:\n" +
                "  RGB values must be in the range 0-255.\n" +
                "  Usage: colors.rgbToHex(255, 0, 128) -> \"#FF0080\"\n",
        }
    }

    hex := fmt.Sprintf("#%02X%02X%02X", r, g, b)
    return &object.String{Value: hex}
}

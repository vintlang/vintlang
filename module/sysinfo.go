package module

import (
	"runtime"

	"github.com/ekilie/vint-lang/object"
)

var SysInfoFunctions = map[string]object.ModuleFunction{}

func init() {
	SysInfoFunctions["os"] = getOS
	SysInfoFunctions["arch"] = getArch
}

func getOS(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.String{Value: runtime.GOOS}
}

func getArch(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.String{Value: runtime.GOARCH}
}

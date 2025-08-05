package module

import (
	"fmt"
	"runtime"

	"github.com/vintlang/vintlang/object"
)

var SysInfoFunctions = map[string]object.ModuleFunction{}

func init() {
	SysInfoFunctions["os"] = getOS
	SysInfoFunctions["arch"] = getArch
}

func getOS(args []object.Object, defs map[string]object.Object) object.Object {
    if len(args) != 0 {
        return ErrorMessage(
            "sysinfo", "os",
            "No arguments",
            fmt.Sprintf("%d arguments", len(args)),
            `sysinfo.os() -> "linux"`,
        )
    }
    return &object.String{Value: runtime.GOOS}
}
	

func getArch(args []object.Object, defs map[string]object.Object) object.Object {
    if len(args) != 0 {
        return ErrorMessage(
            "sysinfo", "arch",
            "No arguments",
            fmt.Sprintf("%d arguments", len(args)),
            `sysinfo.arch() -> "amd64"`,
        )
    }
    return &object.String{Value: runtime.GOARCH}
}

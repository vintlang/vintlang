package module

import "github.com/NuruProgramming/Nuru/object"

var Mapper = map[string]*object.Module{}

func init() {
	Mapper["os"] = &object.Module{Name: "os", Functions: OsFunctions}
	Mapper["time"] = &object.Module{Name: "time", Functions: TimeFunctions}
	Mapper["net"] = &object.Module{Name: "net", Functions: NetFunctions}
	Mapper["json"] = &object.Module{Name: "json", Functions: JsonFunctions}
	Mapper["math"] = &object.Module{Name: "hisabati", Functions: MathFunctions}
}

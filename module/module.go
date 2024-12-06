package module

import "github.com/ekilie/vint-lang/object"

var Mapper = map[string]*object.Module{}

func init() {
	Mapper["os"] = &object.Module{Name: "os", Functions: OsFunctions}
	Mapper["time"] = &object.Module{Name: "time", Functions: TimeFunctions}
	Mapper["net"] = &object.Module{Name: "net", Functions: NetFunctions}
	Mapper["json"] = &object.Module{Name: "json", Functions: JsonFunctions}
	Mapper["math"] = &object.Module{Name: "math", Functions: MathFunctions}
	Mapper["cli"] = &object.Module{Name: "cli", Functions: CliFunctions}
	Mapper["uuid"] = &object.Module{Name: "uuid", Functions: UuidFunctions}
	Mapper["string"] = &object.Module{Name: "string", Functions: StringFunctions}
	Mapper["crypto"] = &object.Module{Name: "crypto", Functions: CryptoFunctions}
	Mapper["regex"] = &object.Module{Name: "regex", Functions: RegexFunctions}
	Mapper["shell"] = &object.Module{Name: "shell", Functions: RegexFunctions}
	Mapper["dotenv"] = &object.Module{Name: "dotenv", Functions: RegexFunctions}
	Mapper["sysinfo"] = &object.Module{Name: "sysinfo", Functions: RegexFunctions}
	Mapper["encoding"] = &object.Module{Name: "encoding", Functions: RegexFunctions}
	Mapper["colors"] = &object.Module{Name: "colors", Functions: RegexFunctions}

}

package module

import "github.com/vintlang/vintlang/object"

var Mapper = map[string]*object.Module{}

func init() {
	Mapper["os"] = &object.Module{Name: "os", Functions: OsFunctions}
	Mapper["time"] = &object.Module{Name: "time", Functions: TimeFunctions}
	Mapper["net"] = &object.Module{Name: "net", Functions: NetFunctions}
	Mapper["http"] = &object.Module{Name: "http", Functions: HttpFunctions}
	Mapper["json"] = &object.Module{Name: "json", Functions: JsonFunctions}
	Mapper["math"] = &object.Module{Name: "math", Functions: MathFunctions}
	Mapper["cli"] = &object.Module{Name: "cli", Functions: CliFunctions}
	Mapper["term"] = &object.Module{Name: "term", Functions: TermFunctions}
	Mapper["uuid"] = &object.Module{Name: "uuid", Functions: UuidFunctions}
	Mapper["string"] = &object.Module{Name: "string", Functions: StringFunctions}
	Mapper["crypto"] = &object.Module{Name: "crypto", Functions: CryptoFunctions}
	Mapper["regex"] = &object.Module{Name: "regex", Functions: RegexFunctions}
	Mapper["shell"] = &object.Module{Name: "shell", Functions: ShellFunctions}
	Mapper["dotenv"] = &object.Module{Name: "dotenv", Functions: DotenvFunctions}
	Mapper["sysinfo"] = &object.Module{Name: "sysinfo", Functions: SysInfoFunctions}
	Mapper["sqlite"] = &object.Module{Name: "sqlite", Functions: SQLiteFunctions}
	Mapper["mysql"] = &object.Module{Name: "mysql", Functions: MySQLFunctions}
	Mapper["postgres"] = &object.Module{Name: "postgres", Functions: PostGresFunctions}
	Mapper["path"] = &object.Module{Name: "path", Functions: PathFunctions}
	Mapper["random"] = &object.Module{Name: "random", Functions: RandomFunctions}
	Mapper["csv"] = &object.Module{Name: "csv", Functions: CsvFunctions}
	Mapper["encoding"] = &object.Module{Name: "encoding", Functions: EncodingFunctions}
	Mapper["errors"] = &object.Module{Name: "errors", Functions: ErrorFunctions}
	Mapper["colors"] = &object.Module{Name: "colors", Functions: ColorsFunctions}
	Mapper["vintSocket"] = &object.Module{Name: "vintSocket", Functions: VintSocketFunctions}
	Mapper["vintChart"] = &object.Module{Name: "vintChart", Functions: VintChartFunctions}
	Mapper["llm"] = &object.Module{Name: "llm", Functions: LLMFunctions}
	Mapper["openai"] = &object.Module{Name: "openai", Functions: OpenAIFunctions}
}

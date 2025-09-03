package module

import (
	"fmt"

	"github.com/vintlang/vintlang/object"
)

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
	Mapper["schedule"] = &object.Module{Name: "schedule", Functions: ScheduleFunctions}
	Mapper["logger"] = &object.Module{Name: "logger", Functions: LoggerFunctions}
	Mapper["hash"] = &object.Module{Name: "hash", Functions: HashFunctions}
	Mapper["xml"] = &object.Module{Name: "xml", Functions: XMLFunctions}
	Mapper["url"] = &object.Module{Name: "url", Functions: URLFunctions}
	Mapper["email"] = &object.Module{Name: "email", Functions: EmailFunctions}
	Mapper["reflect"] = &object.Module{Name: "reflect", Functions: ReflectFunctions}
}

// ErrorMessage formats an error message for module functions
// It provides a consistent error format for incorrect usage of module functions
// including the module name, function name, expected arguments, received arguments, and usage instructions.
func ErrorMessage(module, function, expected, received, usage string) *object.Error {
    return &object.Error{
        Message: fmt.Sprintf(
            "\033[1; -> %s.%s()\033[0m:\n"+
                "  Expected: %s\n"+
                "  Received: %s\n"+
                "  Usage: %s\n"+
                "  See documentation for details. https://vintlang.ekilie.com/docs\n",
            module, function, expected, received, usage,
        ),
    }
}

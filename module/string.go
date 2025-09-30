package module

import (
	"regexp"
	"strings"

	"github.com/vintlang/vintlang/object"
	"github.com/xrash/smetrics" // string metrics, like Levenshtein
)

var StringFunctions = map[string]object.ModuleFunction{}

func init() {
	StringFunctions["trim"] = trim
	StringFunctions["contains"] = contains
	StringFunctions["toUpper"] = toUpper
	StringFunctions["toLower"] = toLower
	StringFunctions["replace"] = replace
	StringFunctions["split"] = split
	StringFunctions["join"] = join
	StringFunctions["substring"] = substring
	StringFunctions["length"] = length
	StringFunctions["indexOf"] = indexOf
	StringFunctions["similarity"] = similarity
	StringFunctions["slug"] = slug
	StringFunctions["startsWith"] = startsWith
	StringFunctions["endsWith"] = endsWith
	StringFunctions["chr"] = chr
	StringFunctions["ord"] = ord
}

// slug creates a URL-friendly slug from a normal string
func slug(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "slug",
			"1 string argument (text to convert to slug)",
			formatArgs(args),
			`string.slug("Hello World!") -> "hello-world"`,
		)
	}

	input := strings.ToLower(args[0].(*object.String).Value)
	re := regexp.MustCompile(`[^a-z0-9\s-]+`)
	input = re.ReplaceAllString(input, "")
	re = regexp.MustCompile(`[\s-]+`)
	input = re.ReplaceAllString(input, "-")
	input = strings.Trim(input, "-")

	return &object.String{Value: input}
}

// similarity computes a similarity score between two strings
func similarity(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "similarity",
			"2 string arguments (string1, string2)",
			formatArgs(args),
			`string.similarity("hello", "hallo") -> 0.8`,
		)
	}

	str1 := args[0].(*object.String).Value
	str2 := args[1].(*object.String).Value
	distance := smetrics.WagnerFischer(str1, str2, 1, 1, 2)
	maxLen := len(str1)
	if len(str2) > maxLen {
		maxLen = len(str2)
	}
	if maxLen == 0 {
		return &object.Float{Value: 1.0}
	}
	return &object.Float{Value: 1.0 - float64(distance)/float64(maxLen)}
}

func trim(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "trim",
			"1 string argument",
			formatArgs(args),
			`string.trim("  hi  ") -> "hi"`,
		)
	}
	return &object.String{Value: strings.TrimSpace(args[0].(*object.String).Value)}
}

func contains(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "contains",
			"2 string arguments (string, substring)",
			formatArgs(args),
			`string.contains("hello world", "world") -> true`,
		)
	}
	return &object.Boolean{Value: strings.Contains(args[0].(*object.String).Value, args[1].(*object.String).Value)}
}

func toUpper(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "toUpper",
			"1 string argument",
			formatArgs(args),
			`string.toUpper("hello") -> "HELLO"`,
		)
	}
	return &object.String{Value: strings.ToUpper(args[0].(*object.String).Value)}
}

func toLower(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "toLower",
			"1 string argument",
			formatArgs(args),
			`string.toLower("HELLO") -> "hello"`,
		)
	}
	return &object.String{Value: strings.ToLower(args[0].(*object.String).Value)}
}

func replace(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "replace",
			"3 string arguments (original, old, new)",
			formatArgs(args),
			`string.replace("hello world", "world", "gophers") -> "hello gophers"`,
		)
	}
	return &object.String{Value: strings.ReplaceAll(args[0].(*object.String).Value, args[1].(*object.String).Value, args[2].(*object.String).Value)}
}

func split(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "split",
			"2 string arguments (string, delimiter)",
			formatArgs(args),
			`string.split("a,b,c", ",") -> ["a","b","c"]`,
		)
	}
	parts := strings.Split(args[0].(*object.String).Value, args[1].(*object.String).Value)
	elements := make([]object.VintObject, len(parts))
	for i, part := range parts {
		elements[i] = &object.String{Value: part}
	}
	return &object.Array{Elements: elements}
}

func join(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 || args[0].Type() != object.ARRAY_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "join",
			"array of strings and delimiter",
			formatArgs(args),
			`string.join(["a","b"], ",") -> "a,b"`,
		)
	}
	array := args[0].(*object.Array)
	delim := args[1].(*object.String).Value
	var parts []string
	for _, elem := range array.Elements {
		if elem.Type() != object.STRING_OBJ {
			return &object.Error{Message: "join expects an array of strings"}
		}
		parts = append(parts, elem.(*object.String).Value)
	}
	return &object.String{Value: strings.Join(parts, delim)}
}

func substring(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.INTEGER_OBJ || args[2].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"string", "substring",
			"string, int start, int end",
			formatArgs(args),
			`string.substring("hello", 0, 4) -> "hell"`,
		)
	}
	str := args[0].(*object.String).Value
	start := int(args[1].(*object.Integer).Value)
	end := int(args[2].(*object.Integer).Value)
	if start < 0 || end > len(str) || start >= end {
		return &object.Error{Message: "Invalid start or end index"}
	}
	return &object.String{Value: str[start:end]}
}

func length(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "length",
			"1 string argument",
			formatArgs(args),
			`string.length("hello") -> 5`,
		)
	}
	return &object.Integer{Value: int64(len(args[0].(*object.String).Value))}
}

func indexOf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "indexOf",
			"string, substring",
			formatArgs(args),
			`string.indexOf("hello", "e") -> 1`,
		)
	}
	index := strings.Index(args[0].(*object.String).Value, args[1].(*object.String).Value)
	if index == -1 {
		return &object.Error{Message: "Substring not found"}
	}
	return &object.Integer{Value: int64(index)}
}

func startsWith(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "startsWith",
			"string, prefix",
			formatArgs(args),
			`string.startsWith("hello", "he") -> true`,
		)
	}

	str := args[0].(*object.String).Value
	prefix := args[1].(*object.String).Value
	return &object.Boolean{Value: strings.HasPrefix(str, prefix)}
}

func endsWith(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "endsWith",
			"string, suffix",
			formatArgs(args),
			`string.endsWith("hello", "lo") -> true`,
		)
	}

	str := args[0].(*object.String).Value
	suffix := args[1].(*object.String).Value
	return &object.Boolean{Value: strings.HasSuffix(str, suffix)}
}

func chr(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"string", "chr",
			"integer (ASCII/Unicode code)",
			formatArgs(args),
			`string.chr(65) -> "A"`,
		)
	}

	code := args[0].(*object.Integer).Value
	return &object.String{Value: string(rune(code))}
}

func ord(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"string", "ord",
			"single character string",
			formatArgs(args),
			`string.ord("A") -> 65`,
		)
	}

	s := args[0].(*object.String).Value
	if len(s) != 1 {
		return ErrorMessage(
			"string", "ord",
			"single character string",
			formatArgs(args),
			`string.ord("A") -> 65`,
		)
	}
	return &object.Integer{Value: int64(s[0])}
}

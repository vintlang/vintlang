package module

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var URLFunctions = map[string]object.ModuleFunction{}

func init() {
	URLFunctions["parse"] = urlParse
	URLFunctions["encode"] = urlEncode
	URLFunctions["decode"] = urlDecode
	URLFunctions["join"] = urlJoin
	URLFunctions["isValid"] = urlIsValid
	URLFunctions["build"] = urlBuild
}

func urlParse(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"url", "parse",
			"1 argument: URL string (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`url.parse("https://example.com/path?query=value") -> returns URL components`,
		)
	}

	urlStr := args[0]
	if urlStr.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"url", "parse",
			"string URL",
			string(urlStr.Type()),
			`url.parse("https://example.com/path?query=value") -> returns URL components`,
		)
	}

	input := urlStr.(*object.String).Value
	parsedURL, err := url.Parse(input)
	if err != nil {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> url.parse()\033[0m:\n"+
				"  Invalid URL: %s\n"+
				"  Usage: url.parse(\"https://example.com/path?query=value\")\n", err.Error()),
		}
	}

	// Return a formatted string with URL components
	result := fmt.Sprintf("scheme:%s host:%s path:%s query:%s fragment:%s",
		parsedURL.Scheme, parsedURL.Host, parsedURL.Path, parsedURL.RawQuery, parsedURL.Fragment)
	
	return &object.String{Value: result}
}

func urlEncode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"url", "encode",
			"1 argument: text (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`url.encode("hello world!") -> "hello%20world%21"`,
		)
	}

	text := args[0]
	if text.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"url", "encode",
			"string text",
			string(text.Type()),
			`url.encode("hello world!") -> "hello%20world%21"`,
		)
	}

	input := text.(*object.String).Value
	encoded := url.QueryEscape(input)
	
	return &object.String{Value: encoded}
}

func urlDecode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"url", "decode",
			"1 argument: encoded text (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`url.decode("hello%20world%21") -> "hello world!"`,
		)
	}

	encoded := args[0]
	if encoded.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"url", "decode",
			"string encoded text",
			string(encoded.Type()),
			`url.decode("hello%20world%21") -> "hello world!"`,
		)
	}

	input := encoded.(*object.String).Value
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> url.decode()\033[0m:\n"+
				"  Invalid URL encoding: %s\n"+
				"  Usage: url.decode(\"hello%%20world%%21\") -> \"hello world!\"\n", err.Error()),
		}
	}
	
	return &object.String{Value: decoded}
}

func urlJoin(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"url", "join",
			"2 arguments: base URL (string), path (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`url.join("https://example.com", "/path/to/resource") -> "https://example.com/path/to/resource"`,
		)
	}

	baseURL := args[0]
	path := args[1]
	
	if baseURL.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"url", "join",
			"string base URL for first argument",
			string(baseURL.Type()),
			`url.join("https://example.com", "/path/to/resource") -> "https://example.com/path/to/resource"`,
		)
	}
	
	if path.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"url", "join",
			"string path for second argument",
			string(path.Type()),
			`url.join("https://example.com", "/path/to/resource") -> "https://example.com/path/to/resource"`,
		)
	}

	base := baseURL.(*object.String).Value
	relativePath := path.(*object.String).Value
	
	parsedBase, err := url.Parse(base)
	if err != nil {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> url.join()\033[0m:\n"+
				"  Invalid base URL: %s\n"+
				"  Usage: url.join(\"https://example.com\", \"/path/to/resource\")\n", err.Error()),
		}
	}

	parsedPath, err := url.Parse(relativePath)
	if err != nil {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> url.join()\033[0m:\n"+
				"  Invalid path: %s\n"+
				"  Usage: url.join(\"https://example.com\", \"/path/to/resource\")\n", err.Error()),
		}
	}

	result := parsedBase.ResolveReference(parsedPath)
	return &object.String{Value: result.String()}
}

func urlIsValid(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"url", "isValid",
			"1 argument: URL string (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`url.isValid("https://example.com") -> true/false`,
		)
	}

	urlStr := args[0]
	if urlStr.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"url", "isValid",
			"string URL",
			string(urlStr.Type()),
			`url.isValid("https://example.com") -> true/false`,
		)
	}

	input := urlStr.(*object.String).Value
	_, err := url.Parse(input)
	
	// Additional validation - must have a scheme for a complete URL
	if err != nil || !strings.Contains(input, "://") {
		return &object.Boolean{Value: false}
	}
	
	return &object.Boolean{Value: true}
}

func urlBuild(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"url", "build",
			"1 argument: components dictionary (dict)",
			fmt.Sprintf("%d arguments", len(args)),
			`url.build({"scheme": "https", "host": "example.com", "path": "/api"}) -> "https://example.com/api"`,
		)
	}

	components := args[0]
	if components.Type() != object.DICT_OBJ {
		return ErrorMessage(
			"url", "build",
			"dictionary of URL components",
			string(components.Type()),
			`url.build({"scheme": "https", "host": "example.com", "path": "/api"}) -> "https://example.com/api"`,
		)
	}

	dict := components.(*object.Dict)
	u := &url.URL{}

	// Extract components from dictionary
	for _, pair := range dict.Pairs {
		key := pair.Key.Inspect()
		value := pair.Value

		if value.Type() != object.STRING_OBJ {
			return &object.Error{
				Message: fmt.Sprintf("\033[1;31m -> url.build()\033[0m:\n"+
					"  Component '%s' must be a string, got %s\n"+
					"  Usage: url.build({\"scheme\": \"https\", \"host\": \"example.com\"})\n", 
					key, string(value.Type())),
			}
		}

		valueStr := value.(*object.String).Value

		switch strings.ToLower(key) {
		case "scheme":
			u.Scheme = valueStr
		case "host":
			u.Host = valueStr
		case "path":
			u.Path = valueStr
		case "query":
			u.RawQuery = valueStr
		case "fragment":
			u.Fragment = valueStr
		case "user":
			if userInfo, err := url.Parse("http://" + valueStr + "@example.com"); err == nil {
				u.User = userInfo.User
			}
		case "port":
			if u.Host != "" && !strings.Contains(u.Host, ":") {
				u.Host = u.Host + ":" + valueStr
			}
		default:
			return &object.Error{
				Message: fmt.Sprintf("\033[1;31m -> url.build()\033[0m:\n"+
					"  Unknown URL component: '%s'\n"+
					"  Valid components: scheme, host, path, query, fragment, user, port\n", key),
			}
		}
	}

	// Basic validation - at least scheme or host should be provided for a valid URL
	if u.Scheme == "" && u.Host == "" {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> url.build()\033[0m:\n"+
				"  URL must have at least a scheme or host component\n"+
				"  Usage: url.build({\"scheme\": \"https\", \"host\": \"example.com\"})\n"),
		}
	}

	return &object.String{Value: u.String()}
}
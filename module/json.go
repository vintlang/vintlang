package module

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var JsonFunctions = map[string]object.ModuleFunction{}

func init() {
	JsonFunctions["decode"] = decode
	JsonFunctions["encode"] = encode
	JsonFunctions["stringify"] = encode //Experimental
	JsonFunctions["pretty"] = pretty
	JsonFunctions["merge"] = merge
	JsonFunctions["get"] = get
}

func decode(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return ErrorMessage(
			"json", "decode",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			`json.decode("{\"key\": \"value\"}")`,
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"json", "decode",
			"1 string argument (JSON string)",
			fmt.Sprintf("%d arguments", len(args)),
			`json.decode("{\"key\": \"value\"}")`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"json", "decode",
			"string argument containing JSON",
			string(args[0].Type()),
			`json.decode("{\"key\": \"value\"}")`,
		)
	}

	var i any

	input := args[0].(*object.String).Value
	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		return &object.Error{Message: "This data is not valid JSON"}
	}

	return convertWhateverToObject(i)
}

func convertWhateverToObject(i any) object.VintObject {
	switch v := i.(type) {
	case map[string]any:
		dict := &object.Dict{}
		dict.Pairs = make(map[object.HashKey]object.DictPair)

		for k, v := range v {
			pair := object.DictPair{
				Key:   &object.String{Value: k},
				Value: convertWhateverToObject(v),
			}
			dict.Pairs[pair.Key.(object.Hashable).HashKey()] = pair
		}

		return dict
	case []any:
		list := &object.Array{}
		for _, e := range v {
			list.Elements = append(list.Elements, convertWhateverToObject(e))
		}

		return list
	case string:
		return &object.String{Value: v}
	case int64:
		return &object.Integer{Value: v}
	case float64:
		return &object.Float{Value: v}
	case bool:
		if v {
			return &object.Boolean{Value: true}
		} else {
			return &object.Boolean{Value: false}
		}
	}
	return &object.Null{}
}

func encode(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return ErrorMessage(
			"json", "encode",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			`json.encode(data) or json.encode(data, 2)`,
		)
	}

	if len(args) < 1 || len(args) > 2 {
		return ErrorMessage(
			"json", "encode",
			"1 or 2 arguments (data and optional indent)",
			fmt.Sprintf("%d arguments", len(args)),
			`json.encode(data) or json.encode(data, 2)`,
		)
	}

	input := args[0]
	i := convertObjectToWhatever(input)

	// Default to no indentation
	indent := ""
	if len(args) == 2 {
		if args[1].Type() != object.INTEGER_OBJ {
			return ErrorMessage(
				"json", "encode",
				"integer argument for indent",
				string(args[1].Type()),
				`json.encode(data, 2) - indent with 2 spaces`,
			)
		}
		spaces := int(args[1].(*object.Integer).Value)
		indent = strings.Repeat(" ", spaces)
	}

	var data []byte
	var err error

	if indent != "" {
		data, err = json.MarshalIndent(i, "", indent)
	} else {
		data, err = json.Marshal(i)
	}

	if err != nil {
		return &object.Error{Message: "Unable to convert data to JSON"}
	}

	return &object.String{Value: string(data)}
}

func convertObjectToWhatever(obj object.VintObject) any {
	switch v := obj.(type) {
	case *object.Dict:
		m := make(map[string]any)
		for _, pair := range v.Pairs {
			key := pair.Key.(*object.String).Value
			m[key] = convertObjectToWhatever(pair.Value)
		}
		return m
	case *object.Array:
		list := make([]any, len(v.Elements))
		for i, e := range v.Elements {
			list[i] = convertObjectToWhatever(e)
		}
		return list
	case *object.String:
		return v.Value
	case *object.Integer:
		return v.Value
	case *object.Float:
		return v.Value
	case *object.Boolean:
		return v.Value
	case *object.Null:
		return nil
	}
	return nil
}

// pretty formats JSON with indentation for better readability
func pretty(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Expect a single string argument"}
	}

	var i any
	input := args[0].(*object.String).Value
	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		return &object.Error{Message: "Invalid JSON input"}
	}

	prettyJSON, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return &object.Error{Message: "Unable to format JSON"}
	}

	return &object.String{Value: string(prettyJSON)}
}

// merge combines two JSON objects into one
func merge(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 {
		return &object.Error{Message: "Expect exactly two arguments"}
	}

	obj1 := args[0]
	obj2 := args[1]
	map1, ok1 := convertObjectToWhatever(obj1).(map[string]any)
	map2, ok2 := convertObjectToWhatever(obj2).(map[string]any)
	if !ok1 || !ok2 {
		return &object.Error{Message: "Arguments must be JSON objects"}
	}

	// Merging maps
	for k, v := range map2 {
		map1[k] = v
	}

	mergedJSON, err := json.Marshal(map1)
	if err != nil {
		return &object.Error{Message: "Unable to merge JSON"}
	}

	return &object.String{Value: string(mergedJSON)}
}

// get retrieves a value from a JSON object by key
func get(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 2 {
		return &object.Error{Message: "Expect two arguments: JSON object and key"}
	}

	obj := args[0]
	key := args[1]
	if key.Type() != object.STRING_OBJ {
		return &object.Error{Message: "Key must be a string"}
	}

	mapObj, ok := convertObjectToWhatever(obj).(map[string]any)
	if !ok {
		return &object.Error{Message: "First argument must be a JSON object"}
	}

	val, exists := mapObj[key.(*object.String).Value]
	if !exists {
		return &object.Null{}
	}

	return convertWhateverToObject(val)
}

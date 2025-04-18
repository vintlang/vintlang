package module

import (
	"encoding/json"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var JsonFunctions = map[string]object.ModuleFunction{}

func init() {
	JsonFunctions["decode"] = decode
	JsonFunctions["encode"] = encode
	JsonFunctions["stringify"] = encode  //Experimental
	JsonFunctions["pretty"] = pretty
	JsonFunctions["merge"] = merge
	JsonFunctions["get"] = get
} 

func decode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "This argument is not allowed"}
	}
	if len(args) != 1 {
		return &object.Error{Message: "We only need one argument"}
	}

	if args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "The argument must be a string"}
	}

	var i interface{}

	input := args[0].(*object.String).Value
	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		return &object.Error{Message: "This data is not valid JSON"}
	}

	return convertWhateverToObject(i)
}

func convertWhateverToObject(i interface{}) object.Object {
	switch v := i.(type) {
	case map[string]interface{}:
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
	case []interface{}:
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

func encode(args []object.Object, defs map[string]object.Object) object.Object {
    if len(defs) != 0 {
        return &object.Error{Message: "This argument is not allowed"}
    }

    if len(args) < 1 || len(args) > 2 {
        return &object.Error{Message: "Expect one or two arguments: data and optional indent"}
    }

    input := args[0]
    i := convertObjectToWhatever(input)

    // Default to no indentation
    indent := ""
    if len(args) == 2 {
        if args[1].Type() != object.INTEGER_OBJ {
            return &object.Error{Message: "Indent must be an integer"}
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


func convertObjectToWhatever(obj object.Object) interface{} {
	switch v := obj.(type) {
	case *object.Dict:
		m := make(map[string]interface{})
		for _, pair := range v.Pairs {
			key := pair.Key.(*object.String).Value
			m[key] = convertObjectToWhatever(pair.Value)
		}
		return m
	case *object.Array:
		list := make([]interface{}, len(v.Elements))
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
func pretty(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Expect a single string argument"}
	}

	var i interface{}
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
func merge(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 2 {
		return &object.Error{Message: "Expect exactly two arguments"}
	}

	obj1 := args[0]
	obj2 := args[1]
	map1, ok1 := convertObjectToWhatever(obj1).(map[string]interface{})
	map2, ok2 := convertObjectToWhatever(obj2).(map[string]interface{})
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
func get(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 2 {
		return &object.Error{Message: "Expect two arguments: JSON object and key"}
	}

	obj := args[0]
	key := args[1]
	if key.Type() != object.STRING_OBJ {
		return &object.Error{Message: "Key must be a string"}
	}

	mapObj, ok := convertObjectToWhatever(obj).(map[string]interface{})
	if !ok {
		return &object.Error{Message: "First argument must be a JSON object"}
	}

	val, exists := mapObj[key.(*object.String).Value]
	if !exists {
		return &object.Null{}
	}

	return convertWhateverToObject(val)
}
package module

import (
	"fmt"

	"github.com/vintlang/vintlang/object"
	"gopkg.in/yaml.v3"
)

var YAMLFunctions = map[string]object.ModuleFunction{}

func init() {
	YAMLFunctions["decode"] = yamlDecode
	YAMLFunctions["encode"] = yamlEncode
	YAMLFunctions["merge"] = yamlMerge
	YAMLFunctions["get"] = yamlGet
}

// yamlDecode parses a YAML string and returns a Vint object
func yamlDecode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"yaml", "decode",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			`yaml.decode("key: value")`,
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"yaml", "decode",
			"1 string argument (YAML string)",
			fmt.Sprintf("%d arguments", len(args)),
			`yaml.decode("key: value")`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"yaml", "decode",
			"string argument containing YAML",
			string(args[0].Type()),
			`yaml.decode("key: value")`,
		)
	}

	var i interface{}
	input := args[0].(*object.String).Value
	err := yaml.Unmarshal([]byte(input), &i)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Invalid YAML: %s", err.Error())}
	}

	return convertYAMLToObject(i)
}

// yamlEncode converts a Vint object to a YAML string
func yamlEncode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"yaml", "encode",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			`yaml.encode(data)`,
		)
	}

	if len(args) != 1 {
		return ErrorMessage(
			"yaml", "encode",
			"1 argument (data to encode)",
			fmt.Sprintf("%d arguments", len(args)),
			`yaml.encode(data)`,
		)
	}

	input := args[0]
	i := convertObjectToYAML(input)

	data, err := yaml.Marshal(i)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Unable to convert data to YAML: %s", err.Error())}
	}

	return &object.String{Value: string(data)}
}

// yamlMerge combines two YAML-compatible objects
func yamlMerge(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"yaml", "merge",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			`yaml.merge(obj1, obj2)`,
		)
	}
	if len(args) != 2 {
		return ErrorMessage(
			"yaml", "merge",
			"2 arguments (objects to merge)",
			fmt.Sprintf("%d arguments", len(args)),
			`yaml.merge(obj1, obj2)`,
		)
	}

	obj1 := args[0]
	obj2 := args[1]

	// Convert objects to interface{} for merging
	map1, ok1 := convertObjectToYAML(obj1).(map[string]interface{})
	map2, ok2 := convertObjectToYAML(obj2).(map[string]interface{})

	if !ok1 || !ok2 {
		return &object.Error{Message: "Arguments must be dictionary-like objects"}
	}

	// Create a new map and merge
	merged := make(map[string]interface{})
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}

	return convertYAMLToObject(merged)
}

// yamlGet retrieves a value from a YAML-compatible object by key
func yamlGet(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"yaml", "get",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			`yaml.get(obj, "key")`,
		)
	}
	if len(args) != 2 {
		return ErrorMessage(
			"yaml", "get",
			"2 arguments (object and key)",
			fmt.Sprintf("%d arguments", len(args)),
			`yaml.get(obj, "key")`,
		)
	}

	obj := args[0]
	key := args[1]

	if key.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"yaml", "get",
			"string key",
			string(key.Type()),
			`yaml.get(obj, "key")`,
		)
	}

	mapObj, ok := convertObjectToYAML(obj).(map[string]interface{})
	if !ok {
		return &object.Error{Message: "First argument must be a dictionary-like object"}
	}

	val, exists := mapObj[key.(*object.String).Value]
	if !exists {
		return &object.Null{}
	}

	return convertYAMLToObject(val)
}

// convertYAMLToObject converts a Go interface{} from YAML parsing to a Vint object
func convertYAMLToObject(i interface{}) object.Object {
	switch v := i.(type) {
	case map[string]interface{}:
		dict := &object.Dict{}
		dict.Pairs = make(map[object.HashKey]object.DictPair)

		for k, v := range v {
			pair := object.DictPair{
				Key:   &object.String{Value: k},
				Value: convertYAMLToObject(v),
			}
			dict.Pairs[pair.Key.(object.Hashable).HashKey()] = pair
		}
		return dict

	case map[interface{}]interface{}:
		// Handle YAML's tendency to use interface{} keys
		dict := &object.Dict{}
		dict.Pairs = make(map[object.HashKey]object.DictPair)

		for k, v := range v {
			var key object.Object
			switch kv := k.(type) {
			case string:
				key = &object.String{Value: kv}
			case int:
				key = &object.String{Value: fmt.Sprintf("%d", kv)}
			case int64:
				key = &object.String{Value: fmt.Sprintf("%d", kv)}
			case float64:
				key = &object.String{Value: fmt.Sprintf("%g", kv)}
			default:
				key = &object.String{Value: fmt.Sprintf("%v", kv)}
			}

			pair := object.DictPair{
				Key:   key,
				Value: convertYAMLToObject(v),
			}
			dict.Pairs[pair.Key.(object.Hashable).HashKey()] = pair
		}
		return dict

	case []interface{}:
		list := &object.Array{}
		for _, e := range v {
			list.Elements = append(list.Elements, convertYAMLToObject(e))
		}
		return list

	case string:
		return &object.String{Value: v}

	case int:
		return &object.Integer{Value: int64(v)}

	case int64:
		return &object.Integer{Value: v}

	case float64:
		return &object.Float{Value: v}

	case bool:
		return &object.Boolean{Value: v}

	case nil:
		return &object.Null{}
	}

	// Fallback for unknown types
	return &object.String{Value: fmt.Sprintf("%v", i)}
}

// convertObjectToYAML converts a Vint object to a Go interface{} for YAML marshaling
func convertObjectToYAML(obj object.Object) interface{} {
	switch v := obj.(type) {
	case *object.Dict:
		m := make(map[string]interface{})
		for _, pair := range v.Pairs {
			key := pair.Key.(*object.String).Value
			m[key] = convertObjectToYAML(pair.Value)
		}
		return m

	case *object.Array:
		list := make([]interface{}, len(v.Elements))
		for i, e := range v.Elements {
			list[i] = convertObjectToYAML(e)
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

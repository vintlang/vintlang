package module

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vintlang/vintlang/object"
)

var PostGresFunctions = map[string]object.ModuleFunction{}

func init() {
	PostGresFunctions["open"] = openPqConnection
	PostGresFunctions["close"] = closePqConnection
	PostGresFunctions["execute"] = executePqQuery
	PostGresFunctions["fetchAll"] = fetchAllPq
	PostGresFunctions["fetchOne"] = fetchOnePq
}

type PostGresConnection struct {
	db *sql.DB
}

func openPqConnection(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Invalid arguments: Expected 'open(connectionString)' where 'connectionString' is a string"}
	}

	connStr := args[0].(*object.String).Value
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to open database: %s", err)}
	}

	err = db.Ping()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to connect to database: %s", err)}
	}

	conn := &PostGresConnection{db: db}
	return &object.NativeObject{
		Value: conn,
	}
}

func closePqConnection(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Invalid arguments: Expected 'close(conn)'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*PostGresConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	err := conn.Value.(*PostGresConnection).db.Close()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to close database connection: %s", err)}
	}

	return &object.Null{}
}

func executePqQuery(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Invalid arguments: Expected 'execute(conn, query, [params...])'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*PostGresConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Query must be a string"}
	}

	params := convertObjectsToPqParams(args[2:])
	_, err := conn.Value.(*PostGresConnection).db.Exec(query.Value, params...)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Query execution failed: %s", err)}
	}

	return &object.Null{}
}

func fetchAllPq(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Invalid arguments: Expected 'fetchAll(conn, query, [params...])'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*PostGresConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Query must be a string"}
	}

	params := convertObjectsToPqParams(args[2:])
	rows, err := conn.Value.(*PostGresConnection).db.Query(query.Value, params...)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Query execution failed: %s", err)}
	}
	defer rows.Close()

	result := make([]object.Object, 0)
	cols, _ := rows.Columns()
	for rows.Next() {
		values := make([]interface{}, len(cols))
		scanArgs := make([]interface{}, len(cols))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return &object.Error{Message: fmt.Sprintf("Failed to scan row: %s", err)}
		}

		row := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
		for i, col := range cols {
			key := &object.String{Value: col}
			value := convertPqToObject(values[i])
			row.Pairs[key.HashKey()] = object.DictPair{Key: key, Value: value}
		}

		result = append(result, row)
	}

	return &object.Array{Elements: result}
}

func fetchOnePq(args []object.Object, defs map[string]object.Object) object.Object {
	result := fetchAllPq(args, defs)
	if result.Type() == object.ARRAY_OBJ {
		array := result.(*object.Array)
		if len(array.Elements) > 0 {
			return array.Elements[0]
		}
	}
	return &object.Null{}
}

func convertPqToObject(val interface{}) object.Object {
	switch v := val.(type) {
	case int64:
		return &object.Integer{Value: v}
	case float64:
		return &object.Float{Value: v}
	case []byte:
		return &object.String{Value: string(v)}
	case string:
		return &object.String{Value: v}
	case bool:
		return &object.Boolean{Value: v}
	case nil:
		return &object.Null{}
	default:
		return &object.String{Value: fmt.Sprintf("%v", v)}
	}
}

func convertObjectsToPqParams(objects []object.Object) []interface{} {
	params := make([]interface{}, len(objects))
	for i, obj := range objects {
		switch v := obj.(type) {
		case *object.String:
			params[i] = v.Value
		case *object.Integer:
			params[i] = v.Value
		case *object.Float:
			params[i] = v.Value
		case *object.Boolean:
			params[i] = v.Value
		default:
			params[i] = nil
		}
	}
	return params
}

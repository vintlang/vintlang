package module

import (
	"database/sql"
	// "errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3" 
	"github.com/vintlang/vintlang/object"
)

var SQLiteFunctions = map[string]object.ModuleFunction{}

func init() {
	SQLiteFunctions["open"] = openDatabase
	SQLiteFunctions["execute"] = executeQuery
	SQLiteFunctions["fetchAll"] = fetchAll
	SQLiteFunctions["fetchOne"] = fetchOne
}

type SQLiteConnection struct {
	db *sql.DB
}

// Open a SQLite database
func openDatabase(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Usage: open(path)"}
	}

	dbPath := args[0].(*object.String).Value
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to open database: %s", err)}
	}

	conn := &SQLiteConnection{db: db}
	return &object.NativeObject{
		Value: conn,
	}
}

// Execute a query (INSERT, UPDATE, DELETE)
func executeQuery(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Usage: execute(conn, query, [params...])"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*SQLiteConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	query := args[1].(*object.String).Value
	params := convertObjectsToParams(args[2:])

	_, err := conn.Value.(*SQLiteConnection).db.Exec(query, params...)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Query failed: %s", err)}
	}

	return &object.Null{}
}

// Fetch all rows (SELECT)
func fetchAll(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Usage: fetchAll(conn, query, [params...])"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*SQLiteConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	query := args[1].(*object.String).Value
	params := convertObjectsToParams(args[2:])

	rows, err := conn.Value.(*SQLiteConnection).db.Query(query, params...)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Query failed: %s", err)}
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
			return &object.Error{Message: fmt.Sprintf("Row scan failed: %s", err)}
		}

		row := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
		for i, col := range cols {
			key := &object.String{Value: col}
			value := convertToObject(values[i])
			row.Pairs[key.HashKey()] = object.DictPair{Key: key, Value: value}
		}

		result = append(result, row)
	}

	return &object.Array{Elements: result}
}

// Fetch a single row
func fetchOne(args []object.Object, defs map[string]object.Object) object.Object {
	result := fetchAll(args, defs)
	if result.Type() == object.ARRAY_OBJ {
		array := result.(*object.Array)
		if len(array.Elements) > 0 {
			return array.Elements[0]
		}
	}
	return &object.Null{}
}

// Utility to convert Go values to Vint objects
func convertToObject(val interface{}) object.Object {
	switch v := val.(type) {
	case int64:
		return &object.Integer{Value: v}
	case float64:
		return &object.Float{Value: v}
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

// Utility to convert Vint objects to Go parameters
func convertObjectsToParams(objects []object.Object) []interface{} {
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

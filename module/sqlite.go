package module

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vintlang/vintlang/object"
)

var SQLiteFunctions = map[string]object.ModuleFunction{}

func init() {
	SQLiteFunctions["open"] = openDatabase
	SQLiteFunctions["close"] = closeDatabase
	SQLiteFunctions["execute"] = executeQuery
	SQLiteFunctions["fetchAll"] = fetchAll
	SQLiteFunctions["fetchOne"] = fetchOne
	SQLiteFunctions["createTable"] = createTable
	SQLiteFunctions["dropTable"] = dropTable
}

type SQLiteConnection struct {
	db *sql.DB
}

// Open a SQLite database
func openDatabase(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Invalid arguments: Expected 'open(path)' where 'path' is a string"}
	}

	dbPath := args[0].(*object.String).Value
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to open database at '%s': %s", dbPath, err)}
	}

	conn := &SQLiteConnection{db: db}
	return &object.NativeObject{
		Value: conn,
	}
}

// Close the database connection
func closeDatabase(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Invalid arguments: Expected 'close(conn)'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*SQLiteConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	err := conn.Value.(*SQLiteConnection).db.Close()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to close database connection: %s", err)}
	}

	return &object.Null{}
}

// Execute a query (INSERT, UPDATE, DELETE)
func executeQuery(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Invalid arguments: Expected 'execute(conn, query, [params...])'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*SQLiteConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Query must be a string"}
	}

	params := convertObjectsToParams(args[2:])
	_, err := conn.Value.(*SQLiteConnection).db.Exec(query.Value, params...)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Query execution failed: %s", err)}
	}

	return &object.Null{}
}

// Fetch all rows (SELECT)
func fetchAll(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "Invalid arguments: Expected 'fetchAll(conn, query, [params...])'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*SQLiteConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Query must be a string"}
	}

	params := convertObjectsToParams(args[2:])
	rows, err := conn.Value.(*SQLiteConnection).db.Query(query.Value, params...)
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

// Create a table
func createTable(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "Invalid arguments: Expected 'createTable(conn, query)'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*SQLiteConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Query must be a string"}
	}

	_, err := conn.Value.(*SQLiteConnection).db.Exec(query.Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to create table: %s", err)}
	}

	return &object.Null{}
}

// Drop a table
func dropTable(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "Invalid arguments: Expected 'dropTable(conn, tableName)'"}
	}

	conn, ok := args[0].(*object.NativeObject)
	if !ok || conn.Value.(*SQLiteConnection).db == nil {
		return &object.Error{Message: "Invalid database connection"}
	}

	tableName, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Table name must be a string"}
	}

	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName.Value)
	_, err := conn.Value.(*SQLiteConnection).db.Exec(query)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to drop table '%s': %s", tableName.Value, err)}
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

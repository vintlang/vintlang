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

func openDatabase(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"sqlite",
			"open",
			"1 string argument (path)",
			formatArgs(args),
			`sqlite.open("my.db") -> connection`,
		)
	}
	dbPath := args[0].(*object.String).Value
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to open database at '%s': %s", dbPath, err)}
	}
	return &object.NativeObject{Value: &SQLiteConnection{db: db}}
}

func closeDatabase(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.NATIVE_OBJ {
		return ErrorMessage(
			"sqlite",
			"close",
			"1 connection argument",
			formatArgs(args),
			`sqlite.close(conn) -> null`,
		)
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

func executeQuery(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"sqlite",
			"execute",
			"connection, query, [params...]",
			formatArgs(args),
			`sqlite.execute(conn, "INSERT INTO users(name) VALUES(?)", "John") -> null`,
		)
	}
	conn := args[0].(*object.NativeObject).Value.(*SQLiteConnection)
	params := convertObjectsToParams(args[2:])
	_, err := conn.db.Exec(args[1].(*object.String).Value, params...)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Query execution failed: %s", err)}
	}
	return &object.Null{}
}

func fetchAll(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"sqlite",
			"fetchAll",
			"connection, query, [params...]",
			formatArgs(args),
			`sqlite.fetchAll(conn, "SELECT * FROM users") -> [{...}]`,
		)
	}
	conn := args[0].(*object.NativeObject).Value.(*SQLiteConnection)
	rows, err := conn.db.Query(args[1].(*object.String).Value, convertObjectsToParams(args[2:])...)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Query execution failed: %s", err)}
	}
	defer rows.Close()

	result := make([]object.VintObject, 0)
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

func fetchOne(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	result := fetchAll(args, defs)
	if result.Type() == object.ARRAY_OBJ {
		array := result.(*object.Array)
		if len(array.Elements) > 0 {
			return array.Elements[0]
		}
	}
	return &object.Null{}
}

func createTable(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"sqlite",
			"createTable",
			"connection, query",
			formatArgs(args),
			`sqlite.createTable(conn, "CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, name TEXT)") -> null`,
		)
	}
	conn := args[0].(*object.NativeObject).Value.(*SQLiteConnection)
	_, err := conn.db.Exec(args[1].(*object.String).Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to create table: %s", err)}
	}
	return &object.Null{}
}

func dropTable(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.NATIVE_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"sqlite",
			"dropTable",
			"connection, tableName",
			formatArgs(args),
			`sqlite.dropTable(conn, "users") -> null`,
		)
	}
	conn := args[0].(*object.NativeObject).Value.(*SQLiteConnection)
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", args[1].(*object.String).Value)
	_, err := conn.db.Exec(query)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to drop table '%s': %s", args[1].(*object.String).Value, err)}
	}
	return &object.Null{}
}

func convertToObject(val interface{}) object.VintObject {
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

func convertObjectsToParams(objects []object.VintObject) []interface{} {
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

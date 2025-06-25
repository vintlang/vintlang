package module

import (
	"encoding/csv"
	"os"

	"github.com/vintlang/vintlang/object"
)

var CsvFunctions = map[string]object.ModuleFunction{}

func init() {
	CsvFunctions["read"] = readCsv
	CsvFunctions["write"] = writeCsv
}

func readCsv(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "read() expects a single string argument for the file path"}
	}
	filePath := args[0].(*object.String).Value

	file, err := os.Open(filePath)
	if err != nil {
		return &object.Error{Message: "Error opening file: " + err.Error()}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return &object.Error{Message: "Error reading CSV: " + err.Error()}
	}

	var rows []object.Object
	for _, record := range records {
		var row []object.Object
		for _, value := range record {
			row = append(row, &object.String{Value: value})
		}
		rows = append(rows, &object.Array{Elements: row})
	}

	return &object.Array{Elements: rows}
}

func writeCsv(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.ARRAY_OBJ {
		return &object.Error{Message: "write() expects a file path (string) and data (array of arrays)"}
	}
	filePath := args[0].(*object.String).Value
	data := args[1].(*object.Array)

	file, err := os.Create(filePath)
	if err != nil {
		return &object.Error{Message: "Error creating file: " + err.Error()}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, rowObj := range data.Elements {
		if rowObj.Type() != object.ARRAY_OBJ {
			return &object.Error{Message: "All elements of data must be arrays"}
		}
		rowArr := rowObj.(*object.Array)
		var record []string
		for _, valueObj := range rowArr.Elements {
			if valueObj.Type() != object.STRING_OBJ {
				return &object.Error{Message: "All cell values must be strings"}
			}
			record = append(record, valueObj.(*object.String).Value)
		}
		if err := writer.Write(record); err != nil {
			return &object.Error{Message: "Error writing record to CSV: " + err.Error()}
		}
	}

	return &object.Null{}
}

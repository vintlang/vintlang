package module

import (
	"encoding/csv"
	"fmt"
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
		return ErrorMessage(
			"csv", "read",
			"1 string argument (file path to read)",
			fmt.Sprintf("%d arguments", len(args)),
			`csv.read("data.csv") -> returns an array of arrays with CSV data`,
		)
	}

	filePath := args[0].(*object.String).Value

	file, err := os.Open(filePath)
	if err != nil {
		return ErrorMessage(
			"csv", "read",
			"valid file path",
			fmt.Sprintf("error: %v", err),
			`csv.read("data.csv") -> returns an array of arrays with CSV data`,
		)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return ErrorMessage(
			"csv", "read",
			"valid CSV file",
			fmt.Sprintf("error: %v", err),
			`csv.read("data.csv") -> returns an array of arrays with CSV data`,
		)
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
		return ErrorMessage(
			"csv", "write",
			"1 string argument (file path) and 1 array argument (data)",
			fmt.Sprintf("%d arguments", len(args)),
			`csv.write("data.csv", [["header1", "header2"], ["value1", "value2"]]) -> writes CSV data to a file`,
		)
	}
	filePath := args[0].(*object.String).Value
	data := args[1].(*object.Array)

	file, err := os.Create(filePath)
	if err != nil {
		return ErrorMessage(
			"csv", "write",
			"valid file path",
			fmt.Sprintf("error: %v", err),
			`csv.write("data.csv", [["header1", "header2"], ["value1", "value2"]]) -> writes CSV data to a file`,
		)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, rowObj := range data.Elements {
		if rowObj.Type() != object.ARRAY_OBJ {
			return ErrorMessage(
				"csv", "write",
				"valid array of arrays",
				fmt.Sprintf("error: %v", err),
				`csv.write("data.csv", [["header1", "header2"], ["value1", "value2"]]) -> writes CSV data to a file`,
			)
		}
		rowArr := rowObj.(*object.Array)
		var record []string
		for _, valueObj := range rowArr.Elements {
			if valueObj.Type() != object.STRING_OBJ {
				return ErrorMessage(
					"csv", "write",
					"valid string values",
					fmt.Sprintf("error: %v", err),
					`csv.write("data.csv", [["header1", "header2"], ["value1", "value2"]]) -> writes CSV data to a file`,
				)
			}
			record = append(record, valueObj.(*object.String).Value)
		}
		if err := writer.Write(record); err != nil {
			return ErrorMessage(
				"csv", "write",
				"valid CSV file",
				fmt.Sprintf("error: %v", err),
				`csv.write("data.csv", [["header1", "header2"], ["value1", "value2"]]) -> writes CSV data to a file`,
			)
		}
	}

	return &object.Null{}
}

package module

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/vintlang/vintlang/object"
	"github.com/xuri/excelize/v2"
)

var ExcelFunctions = map[string]object.ModuleFunction{}

func init() {
	// File operations
	ExcelFunctions["create"] = excelCreate
	ExcelFunctions["open"] = excelOpen
	ExcelFunctions["openWithPassword"] = excelOpenWithPassword
	ExcelFunctions["save"] = excelSave
	ExcelFunctions["saveAs"] = excelSaveAs
	ExcelFunctions["close"] = excelClose
	ExcelFunctions["setPassword"] = excelSetPassword

	// Sheet operations
	ExcelFunctions["getSheets"] = excelGetSheets
	ExcelFunctions["addSheet"] = excelAddSheet
	ExcelFunctions["deleteSheet"] = excelDeleteSheet
	ExcelFunctions["renameSheet"] = excelRenameSheet
	ExcelFunctions["setActiveSheet"] = excelSetActiveSheet
	ExcelFunctions["getActiveSheet"] = excelGetActiveSheet

	// Cell operations
	ExcelFunctions["getCell"] = excelGetCell
	ExcelFunctions["setCell"] = excelSetCell
	ExcelFunctions["getCellFormula"] = excelGetCellFormula
	ExcelFunctions["setCellFormula"] = excelSetCellFormula
	ExcelFunctions["getCellStyle"] = excelGetCellStyle
	ExcelFunctions["setCellStyle"] = excelSetCellStyle

	// Range operations
	ExcelFunctions["getRange"] = excelGetRange
	ExcelFunctions["setRange"] = excelSetRange
	ExcelFunctions["copyRange"] = excelCopyRange
	ExcelFunctions["clearRange"] = excelClearRange

	// Row/Column operations
	ExcelFunctions["insertRow"] = excelInsertRow
	ExcelFunctions["insertColumn"] = excelInsertColumn
	ExcelFunctions["deleteRow"] = excelDeleteRow
	ExcelFunctions["deleteColumn"] = excelDeleteColumn
	ExcelFunctions["setRowHeight"] = excelSetRowHeight
	ExcelFunctions["setColumnWidth"] = excelSetColumnWidth
	ExcelFunctions["getRowHeight"] = excelGetRowHeight
	ExcelFunctions["getColumnWidth"] = excelGetColumnWidth

	// Formatting and styling
	ExcelFunctions["mergeCells"] = excelMergeCells
	ExcelFunctions["unmergeCells"] = excelUnmergeCells
	ExcelFunctions["setFont"] = excelSetFont
	ExcelFunctions["setBorder"] = excelSetBorder
	ExcelFunctions["setFill"] = excelSetFill
	ExcelFunctions["setAlignment"] = excelSetAlignment
	ExcelFunctions["setNumberFormat"] = excelSetNumberFormat

	// Data operations
	ExcelFunctions["addTable"] = excelAddTable
	ExcelFunctions["addChart"] = excelAddChart
	ExcelFunctions["addPicture"] = excelAddPicture
	ExcelFunctions["addComment"] = excelAddComment

	// Export/Import operations
	ExcelFunctions["toCSV"] = excelToCSV
	ExcelFunctions["fromCSV"] = excelFromCSV
	ExcelFunctions["toJSON"] = excelToJSON
	ExcelFunctions["fromJSON"] = excelFromJSON

	// Utility functions
	ExcelFunctions["getFileInfo"] = excelGetFileInfo
	ExcelFunctions["searchText"] = excelSearchText
	ExcelFunctions["replaceText"] = excelReplaceText
	ExcelFunctions["calculateFormulas"] = excelCalculateFormulas
}

// Global registry to keep track of open files
var openFiles = make(map[string]*excelize.File)

// File operations
func excelCreate(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) > 1 {
		return ErrorMessage(
			"excel", "create",
			"0-1 arguments (optional file path)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.create() or excel.create(\"workbook.xlsx\")",
		)
	}

	f := excelize.NewFile()
	defer f.Close()

	// If file path is provided, save the file
	if len(args) == 1 {
		if args[0].Type() != object.STRING_OBJ {
			return ErrorMessage(
				"excel", "create",
				"string argument (file path)",
				fmt.Sprintf("argument type %T", args[0]),
				"excel.create(\"workbook.xlsx\")",
			)
		}

		filePath := args[0].(*object.String).Value
		if err := f.SaveAs(filePath); err != nil {
			return &object.Error{Message: fmt.Sprintf("Failed to create Excel file: %v", err)}
		}

		// Store in registry for later use
		openFiles[filePath] = f

		return &object.String{Value: filePath}
	}

	// Return a temporary file identifier
	tempID := fmt.Sprintf("temp_%d", time.Now().UnixNano())
	openFiles[tempID] = f
	return &object.String{Value: tempID}
}

func excelOpen(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "open",
			"1 string argument (file path)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.open(\"workbook.xlsx\")",
		)
	}

	filePath := args[0].(*object.String).Value

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to open Excel file: %v", err)}
	}

	openFiles[filePath] = f
	return &object.String{Value: filePath}
}

func excelOpenWithPassword(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "openWithPassword",
			"2 string arguments (file path, password)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.openWithPassword(\"workbook.xlsx\", \"password123\")",
		)
	}

	filePath := args[0].(*object.String).Value
	password := args[1].(*object.String).Value

	f, err := excelize.OpenFile(filePath, excelize.Options{Password: password})
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to open password-protected Excel file: %v", err)}
	}

	openFiles[filePath] = f
	return &object.String{Value: filePath}
}

func excelSave(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "save",
			"1 string argument (file identifier)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.save(file_id)",
		)
	}

	fileID := args[0].(*object.String).Value
	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.Save(); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to save Excel file: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelSaveAs(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "saveAs",
			"2 string arguments (file identifier, new file path)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.saveAs(file_id, \"new_workbook.xlsx\")",
		)
	}

	fileID := args[0].(*object.String).Value
	newPath := args[1].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.SaveAs(newPath); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to save Excel file as %s: %v", newPath, err)}
	}

	// Update registry with new path
	openFiles[newPath] = f
	delete(openFiles, fileID)

	return &object.String{Value: newPath}
}

func excelClose(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "close",
			"1 string argument (file identifier)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.close(file_id)",
		)
	}

	fileID := args[0].(*object.String).Value
	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.Close(); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to close Excel file: %v", err)}
	}

	delete(openFiles, fileID)
	return &object.Boolean{Value: true}
}

func excelSetPassword(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "setPassword",
			"2 string arguments (file identifier, password)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.setPassword(file_id, \"password123\")",
		)
	}

	// Note: Password protection for workbooks is not directly available in excelize v2.8.0
	// This is a placeholder implementation
	return &object.Error{Message: "Password protection for workbooks is not yet implemented in this version of excelize"}
}

// Sheet operations
func excelGetSheets(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "getSheets",
			"1 string argument (file identifier)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.getSheets(file_id)",
		)
	}

	fileID := args[0].(*object.String).Value
	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	sheets := f.GetSheetList()
	result := make([]object.VintObject, len(sheets))
	for i, sheet := range sheets {
		result[i] = &object.String{Value: sheet}
	}

	return &object.Array{Elements: result}
}

func excelAddSheet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "addSheet",
			"2 string arguments (file identifier, sheet name)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.addSheet(file_id, \"Sheet2\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	index, err := f.NewSheet(sheetName)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to add sheet: %v", err)}
	}

	return &object.Integer{Value: int64(index)}
}

func excelDeleteSheet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "deleteSheet",
			"2 string arguments (file identifier, sheet name)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.deleteSheet(file_id, \"Sheet2\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.DeleteSheet(sheetName); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to delete sheet: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelRenameSheet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "renameSheet",
			"3 string arguments (file identifier, old name, new name)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.renameSheet(file_id, \"Sheet1\", \"Data\")",
		)
	}

	fileID := args[0].(*object.String).Value
	oldName := args[1].(*object.String).Value
	newName := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.SetSheetName(oldName, newName); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to rename sheet: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelSetActiveSheet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "setActiveSheet",
			"2 arguments (file identifier, sheet index as integer)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.setActiveSheet(file_id, 0)",
		)
	}

	fileID := args[0].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	var index int
	switch v := args[1].(type) {
	case *object.Integer:
		index = int(v.Value)
	case *object.String:
		// Try to convert string to int
		if i, err := strconv.Atoi(v.Value); err == nil {
			index = i
		} else {
			return &object.Error{Message: "Sheet index must be an integer"}
		}
	default:
		return &object.Error{Message: "Sheet index must be an integer"}
	}

	f.SetActiveSheet(index)
	return &object.Boolean{Value: true}
}

func excelGetActiveSheet(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "getActiveSheet",
			"1 string argument (file identifier)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.getActiveSheet(file_id)",
		)
	}

	fileID := args[0].(*object.String).Value
	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	index := f.GetActiveSheetIndex()
	return &object.Integer{Value: int64(index)}
}

// Cell operations
func excelGetCell(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "getCell",
			"3 string arguments (file identifier, sheet name, cell reference)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.getCell(file_id, \"Sheet1\", \"A1\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	cellRef := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	value, err := f.GetCellValue(sheetName, cellRef)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get cell value: %v", err)}
	}

	return &object.String{Value: value}
}

func excelSetCell(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 4 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "setCell",
			"4 arguments (file identifier, sheet name, cell reference, value)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.setCell(file_id, \"Sheet1\", \"A1\", \"Hello World\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	cellRef := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	var value any
	switch v := args[3].(type) {
	case *object.String:
		value = v.Value
	case *object.Integer:
		value = v.Value
	case *object.Float:
		value = v.Value
	case *object.Boolean:
		value = v.Value
	default:
		value = fmt.Sprintf("%v", v)
	}

	if err := f.SetCellValue(sheetName, cellRef, value); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to set cell value: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelGetCellFormula(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "getCellFormula",
			"3 string arguments (file identifier, sheet name, cell reference)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.getCellFormula(file_id, \"Sheet1\", \"A1\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	cellRef := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	formula, err := f.GetCellFormula(sheetName, cellRef)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get cell formula: %v", err)}
	}

	return &object.String{Value: formula}
}

func excelSetCellFormula(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 4 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ || args[3].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "setCellFormula",
			"4 string arguments (file identifier, sheet name, cell reference, formula)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.setCellFormula(file_id, \"Sheet1\", \"A1\", \"=SUM(B1:B10)\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	cellRef := args[2].(*object.String).Value
	formula := args[3].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.SetCellFormula(sheetName, cellRef, formula); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to set cell formula: %v", err)}
	}

	return &object.Boolean{Value: true}
}

// Range operations
func excelGetRange(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "getRange",
			"3 string arguments (file identifier, sheet name, range reference)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.getRange(file_id, \"Sheet1\", \"A1:C3\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	rangeRef := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	// For now, get all rows (TODO: implement proper range parsing)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to get range %s: %v", rangeRef, err)}
	}

	// Return all rows as nested arrays (simplified implementation)
	result := make([]object.VintObject, len(rows))
	for i, row := range rows {
		cells := make([]object.VintObject, len(row))
		for j, cell := range row {
			cells[j] = &object.String{Value: cell}
		}
		result[i] = &object.Array{Elements: cells}
	}

	return &object.Array{Elements: result}
}

func excelSetRange(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 4 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ || args[3].Type() != object.ARRAY_OBJ {
		return ErrorMessage(
			"excel", "setRange",
			"4 arguments (file identifier, sheet name, range reference, 2D array data)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.setRange(file_id, \"Sheet1\", \"A1:C3\", [[\"A1\", \"B1\", \"C1\"]])",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	rangeRef := args[2].(*object.String).Value
	data := args[3].(*object.Array)

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	// Parse the range to get starting cell
	parts := strings.Split(rangeRef, ":")
	if len(parts) < 1 {
		return &object.Error{Message: "Invalid range format"}
	}

	startCell := parts[0]

	// Convert data to any slice
	values := make([][]any, len(data.Elements))
	for i, row := range data.Elements {
		if rowArray, ok := row.(*object.Array); ok {
			values[i] = make([]any, len(rowArray.Elements))
			for j, cell := range rowArray.Elements {
				switch v := cell.(type) {
				case *object.String:
					values[i][j] = v.Value
				case *object.Integer:
					values[i][j] = v.Value
				case *object.Float:
					values[i][j] = v.Value
				case *object.Boolean:
					values[i][j] = v.Value
				default:
					values[i][j] = fmt.Sprintf("%v", v)
				}
			}
		}
	}

	if err := f.SetSheetRow(sheetName, startCell, &values[0]); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to set range: %v", err)}
	}

	return &object.Boolean{Value: true}
}

// Row/Column operations
func excelInsertRow(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "insertRow",
			"3 arguments (file identifier, sheet name, row number as integer)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.insertRow(file_id, \"Sheet1\", 2)",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	var row int
	switch v := args[2].(type) {
	case *object.Integer:
		row = int(v.Value)
	default:
		return &object.Error{Message: "Row number must be an integer"}
	}

	if err := f.InsertRows(sheetName, row, 1); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to insert row: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelInsertColumn(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "insertColumn",
			"3 string arguments (file identifier, sheet name, column letter)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.insertColumn(file_id, \"Sheet1\", \"B\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	column := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.InsertCols(sheetName, column, 1); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to insert column: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelDeleteRow(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "deleteRow",
			"3 arguments (file identifier, sheet name, row number as integer)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.deleteRow(file_id, \"Sheet1\", 2)",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	var row int
	switch v := args[2].(type) {
	case *object.Integer:
		row = int(v.Value)
	default:
		return &object.Error{Message: "Row number must be an integer"}
	}

	if err := f.RemoveRow(sheetName, row); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to delete row: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelDeleteColumn(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "deleteColumn",
			"3 string arguments (file identifier, sheet name, column letter)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.deleteColumn(file_id, \"Sheet1\", \"B\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	column := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	if err := f.RemoveCol(sheetName, column); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to delete column: %v", err)}
	}

	return &object.Boolean{Value: true}
}

// Formatting functions (simplified implementations)
func excelMergeCells(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "mergeCells",
			"3 string arguments (file identifier, sheet name, range reference)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.mergeCells(file_id, \"Sheet1\", \"A1:C1\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	rangeRef := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	// Parse range to get start and end cells
	parts := strings.Split(rangeRef, ":")
	if len(parts) != 2 {
		return &object.Error{Message: "Invalid range format, expected format like A1:C1"}
	}

	if err := f.MergeCell(sheetName, parts[0], parts[1]); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to merge cells: %v", err)}
	}

	return &object.Boolean{Value: true}
}

func excelUnmergeCells(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "unmergeCells",
			"3 string arguments (file identifier, sheet name, range reference)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.unmergeCells(file_id, \"Sheet1\", \"A1:C1\")",
		)
	}

	fileID := args[0].(*object.String).Value
	sheetName := args[1].(*object.String).Value
	rangeRef := args[2].(*object.String).Value

	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	// Parse range to get start and end cells
	parts := strings.Split(rangeRef, ":")
	if len(parts) != 2 {
		return &object.Error{Message: "Invalid range format, expected format like A1:C1"}
	}

	if err := f.UnmergeCell(sheetName, parts[0], parts[1]); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to unmerge cells: %v", err)}
	}

	return &object.Boolean{Value: true}
}

// Utility functions
func excelGetFileInfo(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"excel", "getFileInfo",
			"1 string argument (file identifier)",
			fmt.Sprintf("%d arguments", len(args)),
			"excel.getFileInfo(file_id)",
		)
	}

	fileID := args[0].(*object.String).Value
	f, exists := openFiles[fileID]
	if !exists {
		return &object.Error{Message: "File not found in registry"}
	}

	sheets := f.GetSheetList()
	activeSheet := f.GetActiveSheetIndex()

	result := make(map[string]object.VintObject)
	result["sheets"] = &object.Array{Elements: func() []object.VintObject {
		elements := make([]object.VintObject, len(sheets))
		for i, sheet := range sheets {
			elements[i] = &object.String{Value: sheet}
		}
		return elements
	}()}
	result["activeSheet"] = &object.Integer{Value: int64(activeSheet)}
	result["sheetCount"] = &object.Integer{Value: int64(len(sheets))}

	// Convert map to Dict
	pairs := make(map[object.HashKey]object.DictPair)
	for k, v := range result {
		key := &object.String{Value: k}
		hashKey := key.HashKey()
		pairs[hashKey] = object.DictPair{Key: key, Value: v}
	}
	return &object.Dict{Pairs: pairs}
}

// Stub implementations for remaining functions (can be expanded)
func excelGetCellStyle(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSetCellStyle(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelCopyRange(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelClearRange(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSetRowHeight(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSetColumnWidth(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelGetRowHeight(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: 15.0} // Placeholder
}

func excelGetColumnWidth(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: 8.43} // Placeholder
}

func excelSetFont(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSetBorder(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSetFill(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSetAlignment(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSetNumberFormat(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelAddTable(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelAddChart(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelAddPicture(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelAddComment(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelToCSV(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelFromCSV(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelToJSON(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelFromJSON(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelSearchText(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Array{Elements: []object.VintObject{}} // Placeholder
}

func excelReplaceText(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

func excelCalculateFormulas(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Boolean{Value: true} // Placeholder
}

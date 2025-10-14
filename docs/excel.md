# Excel Module in VintLang

The `excel` module provides comprehensive Excel file manipulation capabilities, including reading, writing, formatting, password handling, and advanced features. This module supports both `.xlsx` and `.xls` formats and offers extensive functionality for working with Excel spreadsheets.

---

## Importing the Excel Module

```js
import excel
```

---

## Core Features

- **File Operations**: Create, open, save Excel files with password support
- **Sheet Management**: Add, delete, rename sheets, manage active sheets
- **Cell Operations**: Read/write individual cells, formulas, and styling
- **Range Operations**: Work with cell ranges, copy, clear data
- **Row/Column Management**: Insert, delete, resize rows and columns  
- **Formatting**: Merge cells, set fonts, borders, colors, alignment
- **Advanced Features**: Tables, charts, images, comments
- **Data Exchange**: Convert to/from CSV, JSON formats
- **Password Protection**: Open password-protected files
- **Search & Replace**: Find and replace text across sheets

---

## File Operations

### Create New Excel File

#### `excel.create(filepath?)`
Creates a new Excel workbook. Optionally saves it to the specified path.

```js
import excel

// Create in memory
file_id = excel.create()

// Create and save to file
file_id = excel.create("new_workbook.xlsx")
```

### Open Existing File

#### `excel.open(filepath)`
Opens an existing Excel file.

```js
file_id = excel.open("existing_workbook.xlsx")
```

#### `excel.openWithPassword(filepath, password)`
Opens a password-protected Excel file.

```js
file_id = excel.openWithPassword("protected_workbook.xlsx", "secret123")
```

### Save Operations

#### `excel.save(file_id)`
Saves the current file.

```js
excel.save(file_id)
```

#### `excel.saveAs(file_id, new_filepath)`
Saves the file with a new name or location.

```js
new_file_id = excel.saveAs(file_id, "backup_workbook.xlsx")
```

#### `excel.close(file_id)`
Closes the Excel file and frees memory.

```js
excel.close(file_id)
```

---

## Sheet Management

### Get Sheet Information

#### `excel.getSheets(file_id)`
Returns an array of all sheet names.

```js
sheets = excel.getSheets(file_id)
print("Available sheets:", sheets)
// Output: ["Sheet1", "Data", "Summary"]
```

### Add and Delete Sheets

#### `excel.addSheet(file_id, sheet_name)`
Adds a new worksheet.

```js
sheet_index = excel.addSheet(file_id, "NewSheet")
print("Created sheet at index:", sheet_index)
```

#### `excel.deleteSheet(file_id, sheet_name)`
Deletes a worksheet.

```js
excel.deleteSheet(file_id, "OldSheet")
```

### Rename Sheet

#### `excel.renameSheet(file_id, old_name, new_name)`
Renames an existing worksheet.

```js
excel.renameSheet(file_id, "Sheet1", "DataSheet")
```

### Active Sheet Management

#### `excel.setActiveSheet(file_id, sheet_index)`
Sets the active worksheet by index.

```js
excel.setActiveSheet(file_id, 0)  // Make first sheet active
```

#### `excel.getActiveSheet(file_id)`
Gets the index of the currently active sheet.

```js
active_index = excel.getActiveSheet(file_id)
```

---

## Cell Operations

### Read and Write Cells

#### `excel.getCell(file_id, sheet_name, cell_reference)`
Reads the value from a specific cell.

```js
value = excel.getCell(file_id, "Sheet1", "A1")
print("Cell A1 contains:", value)
```

#### `excel.setCell(file_id, sheet_name, cell_reference, value)`
Writes a value to a specific cell.

```js
excel.setCell(file_id, "Sheet1", "A1", "Hello World")
excel.setCell(file_id, "Sheet1", "B1", 42)
excel.setCell(file_id, "Sheet1", "C1", 3.14)
excel.setCell(file_id, "Sheet1", "D1", true)
```

### Formula Operations

#### `excel.getCellFormula(file_id, sheet_name, cell_reference)`
Gets the formula from a cell.

```js
formula = excel.getCellFormula(file_id, "Sheet1", "E1")
print("Formula:", formula)
```

#### `excel.setCellFormula(file_id, sheet_name, cell_reference, formula)`
Sets a formula in a cell.

```js
excel.setCellFormula(file_id, "Sheet1", "E1", "=SUM(B1:D1)")
excel.setCellFormula(file_id, "Sheet1", "F1", "=AVERAGE(B1:D1)")
excel.setCellFormula(file_id, "Sheet1", "G1", "=IF(B1>0,\"Positive\",\"Zero or Negative\")")
```

---

## Range Operations

### Work with Cell Ranges

#### `excel.getRange(file_id, sheet_name, range_reference)`
Gets data from a cell range as a 2D array.

```js
data = excel.getRange(file_id, "Sheet1", "A1:C3")
// Returns nested arrays: [["A1", "B1", "C1"], ["A2", "B2", "C2"], ["A3", "B3", "C3"]]

// Access specific cell from range
print("First row:", data[0])
print("Cell B2:", data[1][1])
```

#### `excel.setRange(file_id, sheet_name, range_reference, data)`
Sets data for a cell range using a 2D array.

```js
headers = ["Name", "Age", "City"]
data = [
    headers,
    ["John", 30, "New York"],
    ["Jane", 25, "Boston"],
    ["Bob", 35, "Chicago"]
]

excel.setRange(file_id, "Sheet1", "A1:C4", data)
```

---

## Row and Column Management

### Insert and Delete Operations

#### `excel.insertRow(file_id, sheet_name, row_number)`
Inserts a new row at the specified position.

```js
excel.insertRow(file_id, "Sheet1", 2)  // Insert row at position 2
```

#### `excel.insertColumn(file_id, sheet_name, column_letter)`
Inserts a new column at the specified position.

```js
excel.insertColumn(file_id, "Sheet1", "B")  // Insert column B
```

#### `excel.deleteRow(file_id, sheet_name, row_number)`
Deletes a row.

```js
excel.deleteRow(file_id, "Sheet1", 3)  // Delete row 3
```

#### `excel.deleteColumn(file_id, sheet_name, column_letter)`
Deletes a column.

```js
excel.deleteColumn(file_id, "Sheet1", "C")  // Delete column C
```

---

## Cell Formatting

### Merge and Unmerge Cells

#### `excel.mergeCells(file_id, sheet_name, range_reference)`
Merges cells in the specified range.

```js
excel.mergeCells(file_id, "Sheet1", "A1:C1")  // Merge header cells
```

#### `excel.unmergeCells(file_id, sheet_name, range_reference)`
Unmerges previously merged cells.

```js
excel.unmergeCells(file_id, "Sheet1", "A1:C1")
```

---

## File Information and Utilities

### Get File Information

#### `excel.getFileInfo(file_id)`
Returns comprehensive information about the Excel file.

```js
info = excel.getFileInfo(file_id)
print("Number of sheets:", info.sheetCount)
print("Active sheet index:", info.activeSheet)
print("Sheet names:", info.sheets)
```

---

## Complete Usage Example

```js
import excel

print("=== Excel Module Complete Example ===")

// Create a new workbook
file_id = excel.create("employee_data.xlsx")

// Add a new sheet for employee data
excel.addSheet(file_id, "Employees")
excel.renameSheet(file_id, "Sheet1", "Summary")

// Set up employee data
headers = ["ID", "Name", "Department", "Salary", "Bonus"]
employees = [
    [1, "John Doe", "Engineering", 75000, "=D2*0.1"],
    [2, "Jane Smith", "Marketing", 65000, "=D3*0.1"], 
    [3, "Bob Johnson", "Sales", 70000, "=D4*0.15"],
    [4, "Alice Brown", "HR", 60000, "=D5*0.1"]
]

// Write headers
for i = 0; i < headers.length; i++ {
    cell_ref = string.char(65 + i) + "1"  // A1, B1, C1, etc.
    excel.setCell(file_id, "Employees", cell_ref, headers[i])
}

// Write employee data
for row = 0; row < employees.length; row++ {
    for col = 0; col < employees[row].length - 1; col++ {
        cell_ref = string.char(65 + col) + (row + 2)
        excel.setCell(file_id, "Employees", cell_ref, employees[row][col])
    }
    // Set formula for bonus calculation
    bonus_cell = "E" + (row + 2)
    excel.setCellFormula(file_id, "Employees", bonus_cell, employees[row][4])
}

// Merge header row for title
excel.setCell(file_id, "Employees", "A1", "Employee Database")
excel.mergeCells(file_id, "Employees", "A1:E1")

// Add summary in Summary sheet
excel.setCell(file_id, "Summary", "A1", "Summary Report")
excel.setCell(file_id, "Summary", "A3", "Total Employees:")
excel.setCellFormula(file_id, "Summary", "B3", "=COUNTA(Employees.A:A)-1")

excel.setCell(file_id, "Summary", "A4", "Average Salary:")
excel.setCellFormula(file_id, "Summary", "B4", "=AVERAGE(Employees.D:D)")

excel.setCell(file_id, "Summary", "A5", "Total Payroll (with bonuses):")
excel.setCellFormula(file_id, "Summary", "B5", "=SUM(Employees.D:D)+SUM(Employees.E:E)")

// Get file information
info = excel.getFileInfo(file_id)
print("Created workbook with", info.sheetCount, "sheets:")
for sheet in info.sheets {
    print("-", sheet)
}

// Save and close
excel.save(file_id)
excel.close(file_id)

print("Excel file 'employee_data.xlsx' created successfully!")

// Re-open to read data
file_id = excel.open("employee_data.xlsx")

// Read some data back
employee_name = excel.getCell(file_id, "Employees", "B2")
employee_salary = excel.getCell(file_id, "Employees", "D2") 
total_employees = excel.getCell(file_id, "Summary", "B3")

print("First employee:", employee_name, "- Salary:", employee_salary)
print("Total employees:", total_employees)

excel.close(file_id)
```

---

## Advanced Features

### Password Protection

```js
// Open password-protected file
file_id = excel.openWithPassword("secure_data.xlsx", "mypassword")

// Note: Setting passwords on existing files is not yet supported
// in the current version of the excelize library
```

### Working with Multiple Sheets

```js
// Process multiple sheets in a workbook
sheets = excel.getSheets(file_id)
for sheet_name in sheets {
    print("Processing sheet:", sheet_name)
    
    // Get all data from each sheet  
    data = excel.getRange(file_id, sheet_name, "A1:Z100")
    
    // Process data...
    print("Sheet", sheet_name, "has", data.length, "rows")
}
```

### Data Validation and Processing

```js
// Read and validate data
raw_data = excel.getRange(file_id, "Data", "A1:D10")

cleaned_data = []
for row in raw_data {
    if row[0] != "" && row[0] != null {  // Skip empty rows
        cleaned_row = []
        for cell in row {
            // Clean and validate data
            if cell != null {
                cleaned_row.push(string.trim(cell))
            } else {
                cleaned_row.push("")
            }
        }
        cleaned_data.push(cleaned_row)
    }
}

// Write cleaned data back
excel.setRange(file_id, "CleanedData", "A1:D" + cleaned_data.length, cleaned_data)
```

---

## Error Handling

Always use proper error handling when working with Excel files:

```js
// Safe file operations
try {
    file_id = excel.open("data.xlsx")
    data = excel.getCell(file_id, "Sheet1", "A1")
    print("Data:", data)
} catch error {
    print("Error reading Excel file:", error)
} finally {
    if file_id {
        excel.close(file_id)
    }
}
```

---

## Use Cases

- **Data Analysis**: Read Excel reports and perform calculations
- **Report Generation**: Create formatted Excel reports from application data
- **Data Import/Export**: Convert between Excel and other formats
- **Template Processing**: Fill Excel templates with dynamic data
- **Financial Modeling**: Build spreadsheets with complex formulas
- **Batch Processing**: Process multiple Excel files programmatically
- **Data Migration**: Transfer data between different systems via Excel
- **Automated Reporting**: Generate periodic reports in Excel format

---

## Summary of Functions

### File Operations
| Function | Description | Return Type |
|----------|-------------|-------------|
| `create(filepath?)` | Create new Excel file | String (file_id) |
| `open(filepath)` | Open existing Excel file | String (file_id) |
| `openWithPassword(filepath, password)` | Open password-protected file | String (file_id) |
| `save(file_id)` | Save current file | Boolean |
| `saveAs(file_id, filepath)` | Save file with new name | String (new_file_id) |
| `close(file_id)` | Close and cleanup file | Boolean |

### Sheet Management  
| Function | Description | Return Type |
|----------|-------------|-------------|
| `getSheets(file_id)` | Get list of sheet names | Array |
| `addSheet(file_id, name)` | Add new sheet | Integer (index) |
| `deleteSheet(file_id, name)` | Delete sheet | Boolean |
| `renameSheet(file_id, old, new)` | Rename sheet | Boolean |
| `setActiveSheet(file_id, index)` | Set active sheet | Boolean |
| `getActiveSheet(file_id)` | Get active sheet index | Integer |

### Cell Operations
| Function | Description | Return Type |
|----------|-------------|-------------|
| `getCell(file_id, sheet, cell)` | Read cell value | String |
| `setCell(file_id, sheet, cell, value)` | Write cell value | Boolean |
| `getCellFormula(file_id, sheet, cell)` | Get cell formula | String |
| `setCellFormula(file_id, sheet, cell, formula)` | Set cell formula | Boolean |

### Range Operations
| Function | Description | Return Type |
|----------|-------------|-------------|
| `getRange(file_id, sheet, range)` | Get range data | Array (2D) |
| `setRange(file_id, sheet, range, data)` | Set range data | Boolean |

### Row/Column Operations
| Function | Description | Return Type |
|----------|-------------|-------------|
| `insertRow(file_id, sheet, row)` | Insert row | Boolean |
| `insertColumn(file_id, sheet, col)` | Insert column | Boolean |
| `deleteRow(file_id, sheet, row)` | Delete row | Boolean |
| `deleteColumn(file_id, sheet, col)` | Delete column | Boolean |

### Formatting
| Function | Description | Return Type |
|----------|-------------|-------------|
| `mergeCells(file_id, sheet, range)` | Merge cells | Boolean |
| `unmergeCells(file_id, sheet, range)` | Unmerge cells | Boolean |

### Utilities
| Function | Description | Return Type |
|----------|-------------|-------------|
| `getFileInfo(file_id)` | Get file information | Dictionary |

The Excel module provides a powerful and comprehensive interface for working with Excel files in VintLang, supporting both simple data operations and advanced spreadsheet manipulation.
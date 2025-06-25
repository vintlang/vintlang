# CSV Module in VintLang

The `csv` module provides functions to read from and write to CSV (Comma-Separated Values) files.

## Functions

### `csv.read(filePath)`
Reads a CSV file and returns its contents as an array of arrays.

```vint
// Assuming 'data.csv' contains:
// name,age
// alice,30
// bob,25

data = csv.read("data.csv")
print(data) 
// Outputs: [["name", "age"], ["alice", "30"], ["bob", "25"]]
```

### `csv.write(filePath, data)`
Writes a 2D array to a CSV file. The `data` argument must be an array of arrays, and all cell values must be strings.

```vint
users = [
    ["name", "email"],
    ["John Doe", "john.doe@example.com"],
    ["Jane Smith", "jane.smith@example.com"]
]

csv.write("users.csv", users)
// This will create 'users.csv' with the provided data.
``` 
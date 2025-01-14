# Files in Vint

The `files` module in Vint provides basic functionality for working with files. Currently, it allows you to open and read the contents of files.

---

## Opening a File

You can open a file using the `open` keyword. This will return an object of type `FAILI`, which represents the file.

### Syntax

```vint
fileObject = open("filename.txt")
```

### Example

```vint
myFile = open("file.txt")

aina(myFile) // Output: FAILI
```

---

## Reading a File

After opening a file, you can read its contents using the `read()` method. This method retrieves the entire content of the file as a string.

### Syntax

```vint
fileObject.read()
```

### Example

```vint
myFile = open("file.txt")

contents = myFile.read()
print(contents)
```

---

## Notes

- Ensure that the file you are trying to open exists and is accessible from your program.
- The current implementation does not support writing to files or advanced file manipulation.
- File paths should be specified relative to the current working directory or as an absolute path.

---

## Example Usage

### Reading and Printing File Contents

```vint
fileName = "example.txt"
file = open(fileName)

contents = file.read()
print("File Contents:")
print(contents)
```

---
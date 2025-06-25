# Path Module in VintLang

The `path` module provides functions for working with file system paths.

## Functions

### `path.join([...paths])`
Joins one or more path components intelligently.

```vint
p = path.join("/users", "alice", "docs", "file.txt")
print(p) // Outputs: /users/alice/docs/file.txt
```

### `path.basename(path)`
Returns the last portion of a path.

```vint
p = path.basename("/users/alice/docs/file.txt")
print(p) // Outputs: file.txt
```

### `path.dirname(path)`
Returns the directory name of a path.

```vint
p = path.dirname("/users/alice/docs/file.txt")
print(p) // Outputs: /users/alice/docs
```

### `path.ext(path)`
Returns the file extension of the path.

```vint
p = path.ext("/users/alice/docs/file.txt")
print(p) // Outputs: .txt
```

### `path.isAbs(path)`
Returns `true` if the path is absolute.

```vint
p = path.isAbs("/users/alice/docs/file.txt")
print(p) // Outputs: true

p2 = path.isAbs("docs/file.txt")
print(p2) // Outputs: false
``` 
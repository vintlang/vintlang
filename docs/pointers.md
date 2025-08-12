 # Pointers in VintLang

VintLang now supports basic pointer operations, allowing you to reference and dereference values in your programs. This feature enables more advanced data manipulation and can be useful for certain algorithms and data structures.

## Syntax

- **Address-of:** Use `&` to get a pointer to a value.
- **Dereference:** Use `*` to access the value pointed to by a pointer.

## Usage

### Creating a Pointer
```js
let x = 42
let p = &x  # p is now a pointer to the value of x
```

### Dereferencing a Pointer
```js
print(*p)  # prints 42
```

### Printing a Pointer
```js
print(p)  # prints something like Pointer(42) or Pointer(addr=0x..., value=42)
```

## Limitations
- **Pointers in VintLang are pointers to values, not to variables.**
  - If you change the value of `x` after creating `p = &x`, the pointer `p` will still point to the original value, not the updated value of `x`.
- You cannot assign through a pointer (e.g., `*p = 100` is not supported).
- Pointers to literals (e.g., `let p = &42`) are allowed, but they are just pointers to the value at the time of creation.

## Example
```js
let x = 10
let p = &x
print(p)    # Pointer(10)
print(*p)   # 10
x = 20
print(*p)   # Still 10, because p points to the original value
```

## Error Handling
- Dereferencing a non-pointer or a nil pointer will result in a runtime error.

## Summary
- Use `&` to create pointers to values.
- Use `*` to dereference pointers.
- Pointers are useful for referencing values, but do not provide full variable reference semantics.

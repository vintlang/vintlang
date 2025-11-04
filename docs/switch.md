# Switch Statements in Vint

Switch statements in **Vint** allow you to execute different code blocks based on the value of a given expression. This guide covers the basics of switch statements and their usage.

## Basic Syntax

A switch statement starts with the `switch` keyword, followed by the expression inside parentheses `()`, and all cases enclosed within curly braces `{}`.

Each case uses the keyword `case`, followed by a value to check. Multiple values in a case can be separated by commas `,`. The block of code to execute if the condition is met is placed within curly braces `{}`.

### Example:

```js
let a = 2

switch (a) {
	case 3 {
		print("a is three")
	}
	case 2 {
		print("a is two")
	}
}
```

## Multiple Values in a Case

A single `case` can handle multiple possible values. These values are separated by commas `,`.

### Example:

```js
switch (a) {
	case 1, 2, 3 {
		print("a is one, two, or three")
	}
	case 4 {
		print("a is four")
	}
}
```

## Default Case (`default`)

The `default` statement is executed when none of the specified cases match. It is represented by the `default` keyword.

### Example:

```js
let z = 20

switch(z) {
	case 10 {
		print("ten")
	}
	case 30 {
		print("thirty")
	}
	default {
		print("twenty")
	}
}
```

## Nested Switch Statements

Switch statements can be nested to handle more complex conditions.

### Example:

```js
let x = 1
let y = 2

switch (x) {
	case 1 {
		switch (y) {
			case 2 {
				print("x is one and y is two")
			}
			case 3 {
				print("x is one and y is three")
			}
		}
	}
	case 2 {
		print("x is two")
	}
}
```

## Logical Conditions in Cases

Cases can also be used with logical conditions.

### Example:

```js
let isTrue = true
let isFalse = false

switch (isTrue) {
	case true {
		print("isTrue is true")
	}
	case isFalse {
		print("isFalse is true")
	}
	default {
		print("Neither condition is true")
	}
}
```

## Guard Conditions (Advanced)

Switch statements now support **guard conditions** using the `if` keyword. This allows you to bind the switch value to a variable and add additional conditions.

### Variable Binding with Guards

You can bind the switch value to a variable and use it in guard conditions:

```js
let number = 15

switch (number) {
    case x if x > 0 && x < 10 {
        print("Small positive number:", x)
    }
    case x if x >= 10 && x < 100 {
        print("Medium positive number:", x)
    }
    case x if x >= 100 {
        print("Large positive number:", x)
    }
    case x if x < 0 {
        print("Negative number:", x)
    }
    case 0 {
        print("Zero")
    }
    default {
        print("Unknown number")
    }
}
// Output: Medium positive number: 15
```

### Type-Based Switch Cases

Guard conditions enable type checking in switch statements:

```js
let value = "hello world"

switch (value) {
    case x if type(x) == "STRING" && len(x) > 5 {
        print("Long string:", x)
    }
    case x if type(x) == "STRING" {
        print("Short string:", x)
    }
    case x if type(x) == "INTEGER" && x > 0 {
        print("Positive integer:", x)
    }
    case x if type(x) == "BOOLEAN" {
        print("Boolean value:", x)
    }
    default {
        print("Other type:", type(value))
    }
}
// Output: Long string: hello world
```

### Combining Regular Cases with Guard Cases

You can mix regular value-based cases with guard condition cases:

```js
let input = 42

switch (input) {
    case 0 {
        print("Exactly zero")
    }
    case 1 {
        print("Exactly one")
    }
    case x if x > 1 && x <= 10 {
        print("Small number:", x)
    }
    case x if x > 10 && x <= 100 {
        print("Medium number:", x)
    }
    case x if x > 100 {
        print("Large number:", x)
    }
    default {
        print("Negative or unknown")
    }
}
// Output: Medium number: 42
```

By mastering switch statements in **Vint**, you can write clean, structured, and efficient code that efficiently handles complex branching logic with powerful guard conditions and variable binding.

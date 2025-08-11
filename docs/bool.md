# Working with Booleans in vint

Boolean objects in vint are truthy, meaning that any value is true, except tupu and false. They are used to evaluate expressions that return true or false values.

## Evaluating Boolean Expressions

### Evaluating Simple Expressions

In vint, you can evaluate simple expressions that return a boolean value:

```s
print(1 > 2) // Output: `false`

print(1 + 3 < 10) // Output: `true`
```

### Evaluating Complex Expressions

In vint, you can use boolean operators to evaluate complex expressions:

```s
a = 5
b = 10
c = 15

result = (a < b) && (b < c)

if (result) {
    print("Both conditions are true")
} else {
    print("At least one condition is false")
}
// Output: "Both conditions are true"
```

Here, we create three variables a, b, and c. We then evaluate the expression (a < b) && (b < c). Since both conditions are true, the output will be "Both conditions are true".

## Boolean Operators

vint has several boolean operators that you can use to evaluate expressions:

### The && Operator

The && operator evaluates to true only if both operands are true. Here's an example:

```s
print(true && true) // Output: `true`

print(true && false) // Output: `false`
```

### The and() function

```s
print(and(true,true)) // Output: `true`

print(and(true,false)) // Output: `false`
```

### The || Operator

The || operator evaluates to true if at least one of the operands is true. Here's an example:

```s
print(true || false) // Output: `true`

print(false || false) // Output: `false`
```
### The or() Function
```s
print(or(true,false)) // Output: `true`

print(or(false,false)) // Output: `false`
```

### The ! Operator

The ! operator negates the value of the operand. Here's an example:

```s
print(!true) // Output: `false`

print(!false) // Output: `true`
```

### The not() function

```s
print(not(true)) // Output: `false`

print(not(false)) // Output: `true`
```

## Working with Boolean Values in Loops

In vint, you can use boolean expressions in loops to control their behavior. Here's an example:

```s
num = [1, 2, 3, 4, 5]

for v in num {
    if (v % 2 == 0) {
        print(v, "is even")
    } else {
        print(v, "is odd")
    }
}
// Output:
// 1 is odd
// 2 is even
// 3 is odd
// 4 is even
// 5 is odd
```

## Boolean Methods

Boolean values in vint come with several useful built-in methods for conversion and logical operations:

### to_string()

Converts the boolean value to a string representation:

```s
let flag = true
print(flag.to_string())     // "true"

let disabled = false
print(disabled.to_string()) // "false"
```

### to_int()

Converts the boolean value to an integer (1 for true, 0 for false):

```s
let enabled = true
print(enabled.to_int())     // 1

let disabled = false
print(disabled.to_int())    // 0
```

### negate()

Returns the logical negation of the boolean value:

```s
let flag = true
print(flag.negate())        // false

let condition = false
print(condition.negate())   // true
```

### and()

Performs logical AND operation with another boolean:

```s
let a = true
let b = false
print(a.and(b))            // false
print(a.and(true))         // true
```

### or()

Performs logical OR operation with another boolean:

```s
let a = true
let b = false
print(a.or(b))             // true
print(b.or(false))         // false
```

### xor()

Performs logical XOR (exclusive OR) operation:

```s
let a = true
let b = false
print(a.xor(b))            // true
print(a.xor(true))         // false
```

### implies()

Performs logical implication (if A then B):

```s
let premise = true
let conclusion = false
print(premise.implies(conclusion))  // false
print(false.implies(false))         // true
```

### equivalent()

Checks if two boolean values are logically equivalent:

```s
let a = true
let b = true
print(a.equivalent(b))     // true
print(a.equivalent(false)) // false
```

### nor()

Performs logical NOR operation (NOT OR):

```s
let a = false
let b = false
print(a.nor(b))            // true
print(a.nor(true))         // false
```

### nand()

Performs logical NAND operation (NOT AND):

```s
let a = true
let b = true
print(a.nand(b))           // false
print(a.nand(false))       // true
```

## Practical Boolean Examples

Here are some practical examples using boolean methods:

```s
// Feature flags system
let features = {
    "dark_mode": true,
    "notifications": false,
    "beta_features": true
}

// Convert to configuration strings
for key, value in features {
    config_string = key + "=" + value.to_string()
    print(config_string)
}
// Output:
// dark_mode=true
// notifications=false
// beta_features=true

// Permission system using logical operations
let is_admin = true
let is_owner = false
let can_read = true

// Complex permission checks
let can_write = is_admin.or(is_owner)
let can_delete = is_admin.and(is_owner.negate())
let has_access = can_read.and(can_write.or(is_owner))

print("Can write:", can_write.to_string())     // true
print("Can delete:", can_delete.to_string())   // true
print("Has access:", has_access.to_string())   // true

// State machine logic
let door_open = false
let key_inserted = true
let button_pressed = true

// Door can be opened if key is inserted XOR button is pressed (but not both)
let can_open = key_inserted.xor(button_pressed).and(door_open.negate())
print("Can open door:", can_open.to_string())  // false

// Validation logic using implications
let form_valid = true
let submit_enabled = true

// If form is valid, then submit should be enabled
let validation_check = form_valid.implies(submit_enabled)
print("Validation passes:", validation_check.to_string())  // true
```

## Boolean Method Chaining

Boolean methods support method chaining for complex logical operations:

```s
// Complex boolean logic with chaining
let user_active = true
let subscription_valid = false
let trial_period = true

// Chain multiple operations
let has_access = user_active
    .and(subscription_valid.or(trial_period))
    .and(false.negate())

print("User has access:", has_access.to_string())  // true

// Truth table generation
conditions = [true, false]
for a in conditions {
    for b in conditions {
        print("A:", a.to_string(), "B:", b.to_string())
        print("  AND:", a.and(b).to_string())
        print("  OR:", a.or(b).to_string())
        print("  XOR:", a.xor(b).to_string())
        print("  NAND:", a.nand(b).to_string())
        print("---")
    }
}
```


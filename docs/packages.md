# Packages in VintLang

Packages in VintLang provide a powerful way to organize, encapsulate, and reuse your code. They allow you to group related functions, variables, and state into a single, importable unit, similar to modules or libraries in other languages.

---

## Defining a Package

You can define a package using the `package` keyword, followed by the package name and a block of code enclosed in curly braces `{}`.

A single `.vint` file can contain one package definition. The name of the file does not need to match the package name, but it is good practice to keep them related.

**Syntax:**
```js
package MyPackage {
    // ... package members ...
}
```

### Package Members

Inside a package block, you can define:
- **Variables**: To hold the package's state.
- **Functions**: To provide the package's functionality.

All members defined with `let` inside a package are public and can be accessed after the package is imported.

---

## The Automatic `init` Function

VintLang's package system includes a special feature for initialization. If you define a function named `init` inside your package, the Vint interpreter will **automatically execute it** when the package is first loaded.

This is useful for setting up initial state, connecting to services, or performing any other setup work the package needs before it can be used.

**Example:**
```js
package Counter {
    let count = 0

    // This function will run automatically
    let init = func() {
        print("Counter package has been initialized!")
        @.count = 100 // Set initial state
    }

    let getCount = func() {
        return @.count
    }
}
```

---

## The `@` Operator: Self-Reference

Inside a package, you may need to refer to the package's own members or state. VintLang provides the special `@` operator for this purpose. The `@` operator is a reference to the package's own scope.

This is similar to `this` or `self` in other object-oriented languages.

You use it with dot notation to access other members within the same package.

**Example:**
```js
package Greeter {
    let greeting = "Hello"

    let setGreeting = func(newGreeting) {
        // Use @ to access the 'greeting' variable
        @.greeting = newGreeting
    }

    let sayHello = func(name) {
        // Use @ to access the 'greeting' variable
        print(@.greeting + ", " + name + "!")
    }
}
```
Using `@` is necessary to distinguish between a package-level variable and a local variable with the same name.

---

## Importing and Using Packages

To use a package, you import the file that contains its definition. The package object is then assigned to a variable with the same name as the package.

1.  **Create your package file** (e.g., `utils.vint`).
2.  **Import it in another file** (e.g., `main.vint`).
3.  **Access its members** using dot notation.

If `utils.vint` contains `package utils { ... }`, you would use it like this:

```js
// main.vint

// Import the file containing the package
import "utils"

// Now you can use the 'utils' package
utils.doSomething()
```

---
For a complete, runnable example, see the files in the `examples/packages_example/` directory.

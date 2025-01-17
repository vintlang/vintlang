# Vint Extension for VS Code

Welcome to the **Vint** extension for Visual Studio Code! This extension enhances your coding experience by providing a collection of snippets tailored for the Vint programming language. With these snippets, you can write code faster and more efficiently, allowing you to focus on building amazing applications.

## Features

- **Predefined Snippets**: Easily insert commonly used code patterns, such as functions, loops, conditionals, and more.
- **Syntax Assistance**: Quickly scaffold structures like if-else statements, loops, and switch cases.
- **Customization**: Tailor snippets to suit your coding style and project needs.

### Example

Quickly create a function with the `func` snippet:

```js

// Main logic to split and print characters of a string
let name = "VintLang"
s = name.split("") 
for i in s { 
    print(i)
}

// Demonstrating type conversion and conditional statements
age = "10"
convert(age, "INTEGER")  // Convert age string to integer
print(type(age))          // Uncomment to check the type of ageInInt

// Conditional statements to compare the age variable
if (age == 20) {
    print(age)
} else if (age == 10) {
    print("Age is " + age)
} else {
    print((age == "20"))
}

// Working with height variable
height = "6.0" // Height in feet
print("My name is " + name)

// Define a function to print details
let printDetails = func(name, age, height) {
    print("My name is " + name + ", I am " + age + " years old, and my height is " + height + " feet.")
}

// Calling the printDetails function with initial values
printDetails(name, age, height)

```

> Tip: Explore the full list of snippets in the snippet documentation or directly in the VS Code IntelliSense suggestions.

## Requirements

This extension does not have specific dependencies. However, ensure that you have:

1. Visual Studio Code installed.
2. The Vint language environment properly configured in your system (if applicable).

## Extension Settings

This extension contributes the following settings:

- `vint.snippets.enable`: Enable or disable the Vint snippets.
- `vint.snippets.custom`: Add your custom snippets for Vint.

To configure these settings, navigate to `File > Preferences > Settings` (or `Code > Preferences > Settings` on macOS) and search for "vint".

## Known Issues

- Snippets may not appear in non-Vint files.
- Some advanced code patterns might require manual adjustments.

If you encounter issues, please report them in the [GitHub issues](https://github.com/your-repo/vint/issues) section of the project.

## Release Notes

### 1.0.0

- Initial release with basic snippets for functions, loops, conditionals, and more.

### 1.1.0

- Added new snippets for switch cases, try-catch blocks, and variable declarations.
- Improved snippet descriptions for better IntelliSense integration.

### 1.2.0

- Enhanced customization options for user-defined snippets.
- Bug fixes and performance improvements.

---

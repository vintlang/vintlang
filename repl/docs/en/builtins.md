# Built-in Functions in vint

vint has several built-in functions that perform specific tasks.

## The print() Function

The print() function is used to print out messages to the console. It can take zero or more arguments, and the arguments will be printed out with a space in between them. Additionally, print() supports basic formatting such as /n for a new line, /t for a tab space, and \\ for a backslash. Here's an example:

```s
print(1, 2, 3) // Output: "1 2 3"
```

## The input() Function

The input() function is used to get input from the user. It can take zero or one argument, which is a string that will be used as a prompt for the user. Here's an example:

```s
let salamu = func() {
    let jina = input("Unaitwa nani? ")
    print("Mambo vipi", jina)
}

salamu()
```

In this example, we define a function `salamu()` that prompts the user to enter their name using the `input()` function. We then use the `print()` function to print out a message that includes the user's name.

## The type() Function

The `type()` function is used to determine the type of an object. It accepts one argument, and the return value will be a string indicating the type of the object. Here's an example:

```s
type(2) // Output: "NAMBA"
type("vint") // Output: "NENO"
```

## The open() Function

The `open()` function is used to open a file. It accepts one argument, which is the path to the file that you want to open. Here's an example:

```s
file = open("data.txt")
```

In this example, we use the `open()` function to open a file named "data.txt". The variable file will contain a reference to the opened file.
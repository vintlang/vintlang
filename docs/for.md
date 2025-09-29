# For Loops in vint

For loops are a fundamental control structure in vint, used for iterating over iterable objects such as strings, arrays, and dictionaries. This page covers the syntax and usage of for loops in vint, including key-value pair iteration, and the use of break and continue statements.

## Basic Syntax

To create a for loop, use the for keyword followed by a temporary identifier (such as i or v) and the iterable object. Enclose the loop body in curly braces {}. Here's an example with a string:

```s
name = "hello"

for i in name {
    print(i)
}
```
Output:

```s
h
e
l
l
o
```

## Iterating Over Key-Value Pairs

### Dictionaries

vint allows you to iterate over both the value or the key-value pair of an iterable. To iterate over just the values, use one temporary identifier:

```s
dict = {"a": "apple", "b": "banana"}

for v in dict {
    print(v)
}
```

Output:

```s
apple
banana
```

To iterate over both the keys and the values, use two temporary identifiers:

```s
for k, v in dict {
    print(k + " is " + v)
}
```

Output:

```s
a is apple
b is banana
```

### Strings

To iterate over just the values in a string, use one temporary identifier:

```s
for v in "mojo" {
    print(v)
}
```

Output:
```s
m
o
j
o
```
To iterate over both the keys and the values in a string, use two temporary identifiers:

```s
for i, v in "mojo" {
    print(i, "->", v)
}
```
Output:
```s
0 -> m
1 -> o
2 -> j
3 -> o
```

### Lists

To iterate over just the values in a list, use one temporary identifier:

```s
names = ["alice", "bob", "charlie"]

for v in names {
    print(v)
}
```

Output:

```s
alice
bob
charlie
```

To iterate over both the keys and the values in a list, use two temporary identifiers:

```s
for i, v in names {
    print(i, "-", v)
}
```

Output:

```s
0 - alice
1 - bob
2 - charlie
```

## Break and Continue

### Break

Use the break keyword to terminate a loop:

```s
for i, v in "hello" {
    if (i == 2) {
        print("breaking loop")
        break
    }
    print(v)
}
```

Output:

```s
h
e
breaking loop
```

### Continue

Use the continue keyword to skip a specific iteration:

```s
for i, v in "hello" {
    if (i == 2) {
        print("skipping iteration")
        continue
    }
    print(v)
}
```

Output:

```s
h
e
skipping iteration
l
o
```
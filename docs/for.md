# For Loops in vint

For loops are a ffuncmental control structure in vint, used for iterating over iterable objects such as strings, arrays, and dictionaries. This page covers the syntax and usage of for loops in vint, including key-value pair iteration, and the use of break and continue statements.

## Basic Syntax
To create a for loop, use the for keyword followed by a temporary identifier (such as i or v) and the iterable object. Enclose the loop body in curly braces {}. Here's an example with a string:

```s
jina = "lugano"

for i in jina {
    print(i)
}
```
Output:

```s
l
u
g
a
n
o
```

## Iterating Over Key-Value Pairs

### Dictionaries

vint allows you to iterate over both the value or the key-value pair of an iterable. To iterate over just the values, use one temporary identifier:

```s
kamusi = {"a": "andaa", "b": "baba"}

for v in kamusi {
    print(v)
}
```

Output:

```s
andaa
baba
```
To iterate over both the keys and the values, use two temporary identifiers:

```s

for k, v in kamusi {
    print(k + " ni " + v)
}
```
Output:

```s
a ni andaa
b ni baba
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
majina = ["juma", "asha", "haruna"]

for v in majina {
    print(v)
}
```

Output:

```s
juma
asha
haruna
```

To iterate over both the keys and the values in a list, use two temporary identifiers:

```s
for i, v in majina {
    print(i, "-", v)
}
```

Output:

```s
0 - juma
1 - asha
2 - haruna
```

## Break  and Continue 

### Break 

Use the break keyword to terminate a loop:

```s

for i, v in "mojo" {
    if (i == 2) {
        print("break")
        break
    }
    print(v)
}
```

Output:

```s
m
o
nimebreak
```

### Continue (Endelea)

Use the continue keyword to skip a specific iteration:

```s
for i, v in "mojo" {
    if (i == 2) {
        print("nimeruka")
        continue
    }
    print(v)
}
```

Output:

```s
m
o
nimeruka
o
```
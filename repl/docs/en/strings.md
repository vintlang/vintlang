# Strings in vint

Strings are a sequence of characters that can represent text in the vint programming language. This page covers the basics of strings, their manipulation, and some built-in methods.

## Basic Syntax

Strings can be enclosed in either single quotes '' or double quotes "":

```s
print("mambo") // mambo

fanya a = 'niaje'

print("mambo", a) // mambo niaje
```

## Concatenating Strings

Strings can be concatenated using the + operator:

```s
fanya a = "habari" + " " + "yako"

print(a) // habari yako

fanya b = "habari"

b += " yako"

// habari yako
```

You can also repeat a string n number of times using the * operator:

```s
print("mambo " * 4)

// mambo mambo mambo mambo

fanya a = "habari"

a *= 4

// habarihabarihabarihabari
```

## Looping over a String

You can loop through a string using the kwa keyword:

```s
fanya jina = "avicenna"

kwa i ktk jina {print(i)}
```
Output
```s 
a
v
i
c
e
n
n
a  
```

And for key-value pairs:

```s
kwa i, v ktk jina {
	print(i, "=>", v)
}
```
Output
```s
0 => a
1 => v
2 => i
3 => c
4 => e
5 => n
6 => n
7 => a
```

## Comparing Strings

You can compare two strings using the == operator:

```s
fanya a = "vint"

print(a == "vint") // kweli

print(a == "mambo") // sikweli
```

## String Methods

### idadi()

You can find the length of a string using the idadi method. It does not accept any parameters.

```s
fanya a = "mambo"
a.idadi() // 5
```

### herufikubwa()

This method converts a string to uppercase. It does not accept any parameters.

```s
fanya a = "vint"
a.herufikubwa() // vint
```

### herufindogo

This method converts a string to lowercase. It does not accept any parameters.

```s
fanya a = "vint"
a.herufindogo() // vint
```

### gawa

The gawa method splits a string into an array based on a specified delimiter. If no argument is provided, it will split the string according to whitespace.

Example without a parameter:

```s
fanya a = "vint mambo habari"
fanya b = a.gawa()
print(b) // ["vint", "mambo", "habari"]
```

Example with a parameter:

```s
fanya a = "vint,mambo,habari"
fanya b = a.gawa(",")
print(b) // ["vint", "mambo", "habari"]
```

By understanding strings and their manipulation in vint, you can effectively work with text data in your programs.
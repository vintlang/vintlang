# OPERATORS
Operators are the fofunction of any programming language, allowing you to perform various operations on variables and values. This page covers the syntax and usage of operators in vint, including assignment, arithmetic, comparison, member, and logic operators.

## ASSIGNMENT

Assuming `i` and `v` are predefined variables, vint supports the following assignment operators:

- `i = v`: which is the regular assignment operator
- `i += v`: which is the equivalent of `i = i + v`
- `i -= v`: which is the equivalent of `i = i - v`
- `i *= v`: which is the equivalent of `i = i * v`
- `i /= v`: which is the equivalent of `i = i / v`
- `i += v`: which is the equivalent of `i = i + v`

For `strings`, `arrays` and `dictionaries`, the `+=` sign operator is permissible. Example:
```
list1 += list2 // this is equivalent to list1 = list1 + list2
```

## ARITHMETIC OPERATORS

vint supports the following arithmetic operators:

- `+`: Additon
- `-`: Subtraction
- `*`: Multiplication
- `/`: Division
- `%`: Modulo (ie the remainder of a division)
- `**`: Exponential power (eg: `2**3 = 8`)

## COMPARISON OPERATORS

vint supports the following comparison operators:

- `==`: Equal to
- `!=`: Not equal to
- `>`: Greater than
- `>=`: Greater than or equal to
- `<`: Less than
- `<=`: Less than or equal to

## MEMBER OPERATOR

The member operator in vint is `in`. It will check if an object exists in another object:
```go
let majina = ['juma', 'asha', 'haruna']

"haruna" in majina // true
"halima" in majina // false
```

## LOGIC OPERATORS

vint supports the following logic operators:

- `&&`: Logical `AND`. It will evaluate to true if both are true, otherwise it will evaluate to false.
- `||`: Logical `OR`. It will evaluate to false if both are false, otherwise it will evaluate to true.
- `!`: Logical `NOT`. It will evaluate to the opposite of a given expression.

## PRECEDENCE OF OPERATORS

Operators have the following precedence, starting from the highest priority to the lowest:

- `()` : Items in paranthesis have the highest priority
- `!`: Negation
- `%`: Modulo
- `**`: Exponential power
- `/, *`: Division and Multiplication
- `+, +=, -, -=`: Addition and Subtraction
- `>, >=, <, <=`: Comparison operators
- `==, !=`: Equal or Not Equal to
- `=`: Assignment Operator
- `in`: Member Operator
- `&&, ||`: Logical AND and OR

Understanding operators in vint allows you to create complex expressions, perform calculations, and make decisions based on the values of variables.

# Arrays in vint

Arrays in vint are versatile data structures that can hold multiple items, including different types such as numbers, strings, booleans, functions, and null values. This page covers various aspects of arrays, including how to create, manipulate, and iterate over them using vint's built-in keywords and methods.

## Creating Arrays

To create an array, use square brackets [] and separate items with commas:

```s
list = [1, "second", true]
```
## Accessing and Modifying Array Elements

Arrays in vint are zero-indexed. To access an element, use the element's index in square brackets:

```s
num = [10, 20, 30]
n = num[1]  // n is 20
```

You can reassign an element in an array using its index:

```s
num[1] = 25
```

## Concatenating Arrays

To concatenate two or more arrays, use the + operator:

```s
a = [1, 2, 3]
b = [4, 5, 6]
c = a + b
// c is now [1, 2, 3, 4, 5, 6]
```

## Checking for Array Membership

Use the `in` keyword to check if an item exists in an array:

```s
num = [10, 20, 30]
print(20 in num)  // will print true
```

## Looping Over Arrays

You can use the for and in keywords to loop over array elements. To loop over just the values, use the following syntax:

```
num = [1, 2, 3, 4, 5]

for value in num {
    print(value)
}
```

To loop over both index and value pairs, use this syntax:

```s
man = ["Juma", "Asha", "Haruna"]

for idx, n in man {
    print(idx, "-", n)
}
```

## Array Methods

Arrays in vint have several built-in methods:

### length()

length() returns the length of an array:

```s
a = [1, 2, 3]
urefu = a.length()
print(urefu)  // will print 3
```

### push()

push() adds one or more items to the end of an array:

```s
a = [1, 2, 3]
a.push("s", "g")
print(a)  // will print [1, 2, 3, "s", "g"]
```

### last()

last() returns the last item in an array, or tupu if the array is empty:

```s
a = [1, 2, 3]
last_el = a.last()
print(last_el)  // will print 3

b = []
last_el = b.last()
print(last_el)  // will print tupu
```

### map()

map() goes through every element in the array and applies the passed function to each element. It will then return a new array with the updated elements:
```s
a = [1, 2, 3]

b = a.map(func(x){return x*2})

print(b) // [2, 4, 6]
```

### filter()

filter() will go through every single element of an array and checks if that element returns true or false when passed into a function. It will return a new array with elements that returned true:
```s
a = [1, 2, 3, 4]

b = a.filter(func(x){
    if (x % 2 == 0) 
        {return true}
    return false
    })

print(b) // [2, 4]
```

With this information, you can now effectively work with arrays in vint, making it easy to manipulate collections of data in your programs.

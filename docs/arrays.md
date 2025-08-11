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
man = ["Tach", "ekilie", "Tachera Sasi"]

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

last() returns the last item in an array, or null if the array is empty:

```s
a = [1, 2, 3]
last_el = a.last()
print(last_el)  // will print 3

b = []
last_el = b.last()
print(last_el)  // will print tupu
```

### pop()

pop() removes and returns the last item in the array. If the array is empty, it returns null:

```s
a = [1, 2, 3]
last = a.pop()
print(last)  // will print 3
print(a)     // will print [1, 2]
```

### shift()

shift() removes and returns the first item in the array. If the array is empty, it returns null:

```s
a = [1, 2, 3]
first = a.shift()
print(first)  // will print 1
print(a)      // will print [2, 3]
```

### unshift()

unshift() adds one or more items to the beginning of the array:

```s
a = [3, 4]
a.unshift(1, 2)
print(a)  // will print [1, 2, 3, 4]
```

### reverse()

reverse() reverses the array in place:

```s
a = [1, 2, 3]
a.reverse()
print(a)  // will print [3, 2, 1]
```

### sort()

sort() sorts the array in place. It only works for arrays of integers or strings:

```s
a = [3, 1, 2]
a.sort()
print(a)  // will print [1, 2, 3]

b = ["banana", "apple", "cherry"]
b.sort()
print(b)  // will print ["apple", "banana", "cherry"]
```

### map()

map() goes through every element in the array and applies the passed function to each element. It returns a new array with the updated elements:

```s
a = [1, 2, 3]
b = a.map(func(x){ return x * 2 })
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

### slice()

slice() extracts a section of an array and returns a new array:

```s
a = [1, 2, 3, 4, 5]
sliced = a.slice(1, 3)
print(sliced)  // [2, 3]

// With just start index
sliced2 = a.slice(2)
print(sliced2)  // [3, 4, 5]
```

### concat()

concat() merges two or more arrays into a new array:

```s
a = [1, 2]
b = [3, 4]
c = [5, 6]
combined = a.concat(b, c)
print(combined)  // [1, 2, 3, 4, 5, 6]
```

### includes()

includes() checks if an array contains a specific element:

```s
numbers = [1, 2, 3, 4, 5]
print(numbers.includes(3))  // true
print(numbers.includes(10)) // false
```

### every()

every() tests whether all elements pass a test function:

```s
numbers = [2, 4, 6, 8]
allEven = numbers.every(func(x){ return x % 2 == 0 })
print(allEven)  // true
```

### some()

some() tests whether at least one element passes a test function:

```s
numbers = [1, 3, 5, 8]
hasEven = numbers.some(func(x){ return x % 2 == 0 })
print(hasEven)  // true
```

### reduce()

reduce() reduces the array to a single value using an accumulator function:

```s
numbers = [1, 2, 3, 4]
sum = numbers.reduce(func(acc, val){ return acc + val }, 0)
print(sum)  // 10

// Without initial value
product = numbers.reduce(func(acc, val){ return acc * val })
print(product)  // 24
```

### flatten()

flatten() flattens nested arrays into a single array:

```s
nested = [[1, 2], [3, 4], [5]]
flat = nested.flatten()
print(flat)  // [1, 2, 3, 4, 5]

// With depth limit
deep = [[[1, 2]], [3, 4]]
flatOne = deep.flatten(1)
print(flatOne)  // [[1, 2], 3, 4]
```

### unique()

unique() returns a new array with duplicate elements removed:

```s
numbers = [1, 2, 2, 3, 3, 4]
uniqueNumbers = numbers.unique()
print(uniqueNumbers)  // [1, 2, 3, 4]
```

### fill()

fill() fills all elements of an array with a static value:

```s
arr = [1, 2, 3, 4]
arr.fill(0)
print(arr)  // [0, 0, 0, 0]

// Fill with start and end positions
arr2 = [1, 2, 3, 4, 5]
arr2.fill(9, 1, 3)
print(arr2)  // [1, 9, 9, 4, 5]
```

### lastIndexOf()

lastIndexOf() returns the last index at which a given element can be found:

```s
numbers = [1, 2, 3, 2, 4]
lastIndex = numbers.lastIndexOf(2)
print(lastIndex)  // 3

// Element not found
notFound = numbers.lastIndexOf(10)
print(notFound)  // -1
```

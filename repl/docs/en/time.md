# Time in Vint

## Importing Time

To use time-related functionalities in Vint, you first need to import the `time` module as follows:
```vint
import time
```

## Time Methods

### `now()`

To get the current time, use the `time.now()` method. This will return the current time as a `time` object:
```vint
import time

current_time = time.now()
```

### `since()`

Use this method to get the total time since in seconds. It accepts a time object or a string in the format `HH:mm:ss dd-MM-YYYY`:

```vint
import time

now = time.now()

time.since(now) // returns the since time

// alternatively:

now.since("00:00:00 01-01-1900") // returns the since time in seconds since that date
```

### `sleep()`

Use `sleep()` if you want your program to pause or "sleep." It accepts one argument, which is the total time to sleep in seconds:

```vint
time.sleep(10) // will pause the program for ten seconds
```

### `add()`

Use the `add()` method to add to the current time, explained with an example:

```vint
import time

now = time.now()

tomorrow = now.add(days=1)
next_hour = now.add(hours=24)
next_year = now.add(years=1)
three_months_later = now.add(months=3)
next_week = now.add(days=7)
custom_time = now.add(days=3, hours=4, minutes=50, seconds=3)
```

It will return a `time` object with the specified time added.

## Example Usage

### Print the current timestamp
```vint
print(time.now())
```

### Function to greet a user based on the time of the day
```vint
let greet = func(name) {
    let current_time = time.now()  // Get the current time
    print(current_time)            // Print the current time
    if (current_time.hour < 12) {  // Check if it's before noon
        print("Good morning, " + name + "!")
    } else {
        print("Good evening, " + name + "!")
    }
}
```

### Time-related operations
```vint
year = 2024
print("Is", year, "Leap year:", time.isLeapYear(year))
print(time.format(time.now(), "02-01-2006 15:04:05"))
print(time.add(time.now(), "1h"))
print(time.subtract(time.now(), "2h30m45s"))
```
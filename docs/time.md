# Time in Vint

## Importing Time

To use time-related functionalities in Vint, you first need to import the `time` module as follows:
```js
import time
```

## Time Methods

### `now()`

To get the current time, use the `time.now()` method. This will return the current time as a `time` object:
```js
import time

current_time = time.now()
```

### `since()`

Use this method to get the total time since in seconds. It accepts a time object or a string in the format `HH:mm:ss dd-MM-YYYY`:

```js
import time

now = time.now()

time.since(now) // returns the since time

// alternatively:

now.since("00:00:00 01-01-1900") // returns the since time in seconds since that date
```

### `sleep()`

Use `sleep()` if you want your program to pause or "sleep." It accepts one argument, which is the total time to sleep in seconds:

```js
time.sleep(10) // will pause the program for ten seconds
```

### `add()`

Use the `add()` method to add to the current time, explained with an example:

```js
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
```js
print(time.now())
```

### Function to greet a user based on the time of the day
```js
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
```js
year = 2024
print("Is", year, "Leap year:", time.isLeapYear(year))
print(time.format(time.now(), "02-01-2006 15:04:05"))
print(time.add(time.now(), "1h"))
print(time.subtract(time.now(), "2h30m45s"))
```

## Time Object Methods

Time objects in Vint have several powerful built-in methods for manipulation and extraction of time components:

### format()

Format the time object using a custom format string:

```js
import time

now = time.now()
formatted = now.format("2006-01-02 15:04:05")  // Standard format
print(formatted)  // 2024-08-11 15:30:45

// Custom formats
print(now.format("02-01-2006"))           // 11-08-2024
print(now.format("15:04"))                // 15:30
print(now.format("Monday, January 2, 2006"))  // Sunday, August 11, 2024
```

### year()

Get the year component of the time:

```js
import time

now = time.now()
current_year = now.year()
print("Current year:", current_year)  // Current year: 2024
```

### month()

Get the month component of the time (1-12):

```js
import time

now = time.now()
current_month = now.month()
print("Current month:", current_month)  // Current month: 8
```

### day()

Get the day component of the time (1-31):

```js
import time

now = time.now()
current_day = now.day()
print("Current day:", current_day)  // Current day: 11
```

### hour()

Get the hour component of the time (0-23):

```js
import time

now = time.now()
current_hour = now.hour()
print("Current hour:", current_hour)  // Current hour: 15
```

### minute()

Get the minute component of the time (0-59):

```js
import time

now = time.now()
current_minute = now.minute()
print("Current minute:", current_minute)  // Current minute: 30
```

### second()

Get the second component of the time (0-59):

```js
import time

now = time.now()
current_second = now.second()
print("Current second:", current_second)  // Current second: 45
```

### weekday()

Get the weekday name of the time:

```js
import time

now = time.now()
day_name = now.weekday()
print("Today is:", day_name)  // Today is: Sunday
```

## Practical Time Examples

Here are some practical examples using time methods:

```js
import time

// Create a timestamp logger
let log_with_timestamp = func(message) {
    let now = time.now()
    let timestamp = now.format("2006-01-02 15:04:05")
    print("[" + timestamp + "] " + message)
}

log_with_timestamp("Application started")
// Output: [2024-08-11 15:30:45] Application started

// Build a custom date display
let display_date = func() {
    let now = time.now()
    let weekday = now.weekday()
    let day = now.day()
    let month = now.month()
    let year = now.year()
    
    let months = ["", "January", "February", "March", "April", "May", "June",
                  "July", "August", "September", "October", "November", "December"]
    
    let formatted = weekday + ", " + months[month] + " " + day.to_string() + ", " + year.to_string()
    print(formatted)
}

display_date()
// Output: Sunday, August 11, 2024

// Time-based conditional logic
let get_greeting = func() {
    let now = time.now()
    let hour = now.hour()
    
    if (hour < 12) {
        return "Good morning!"
    } else if (hour < 18) {
        return "Good afternoon!"
    } else {
        return "Good evening!"
    }
}

print(get_greeting())

// Schedule checker
let is_business_hours = func() {
    let now = time.now()
    let hour = now.hour()
    let weekday = now.weekday()
    
    // Check if it's a weekday (Monday-Friday) and between 9 AM and 5 PM
    let is_weekday = weekday != "Saturday" && weekday != "Sunday"
    let is_work_time = hour >= 9 && hour < 17
    
    return is_weekday && is_work_time
}

if (is_business_hours()) {
    print("Office is open!")
} else {
    print("Office is closed!")
}

// Age calculator
let calculate_age = func(birth_year) {
    let now = time.now()
    let current_year = now.year()
    return current_year - birth_year
}

let age = calculate_age(1990)
print("Age:", age)

// Deadline checker
let check_deadline = func(deadline_date) {
    let now = time.now()
    let deadline = time.parse(deadline_date)  // Assuming we have a parse method
    
    let days_left = deadline.since(now) / (24 * 60 * 60)  // Convert seconds to days
    
    if (days_left > 0) {
        print("Deadline in", days_left.floor(), "days")
    } else {
        print("Deadline has passed!")
    }
}
```

## Method Chaining with Time

Time methods can be used in combination for complex operations:

```js
import time

// Get a formatted timestamp for a specific time
let birthday = time.now().add(days=30)
let birthday_info = "Birthday: " + birthday.weekday() + ", " + 
                   birthday.format("January 2, 2006") + " at " +
                   birthday.format("15:04")

print(birthday_info)
// Output: Birthday: Tuesday, September 10, 2024 at 15:30

// Create time-based file naming
let create_backup_filename = func(base_name) {
    let now = time.now()
    let timestamp = now.year().to_string() + 
                   now.month().to_string().padStart(2, "0") +
                   now.day().to_string().padStart(2, "0") + "_" +
                   now.hour().to_string().padStart(2, "0") +
                   now.minute().to_string().padStart(2, "0")
    
    return base_name + "_" + timestamp + ".backup"
}

let filename = create_backup_filename("database")
print(filename)  // database_20240811_1530.backup
```
# DateTime Module in VintLang

The `datetime` module provides comprehensive date and time manipulation capabilities with timezone support, duration handling, and advanced datetime operations.

## Importing DateTime

```js
import datetime
```

## Basic Functions

### `datetime.now([timezone])`

Get the current date and time, optionally in a specific timezone.

```js
let current = datetime.now()
print(current)  // 10:15:32 26-09-2025

let ny_time = datetime.now("America/New_York") 
print(ny_time)  // 06:15:32 26-09-2025
```

### `datetime.utcNow()`

Get the current UTC time.

```js
let utc = datetime.utcNow()
print(utc)  // 10:15:32 26-09-2025
```

### `datetime.parse(datetime_string, [format], [timezone])`

Parse a datetime string into a Time object.

```js
let parsed = datetime.parse("2024-12-25 15:30:00", "2006-01-02 15:04:05")
print(parsed)  // 15:30:00 25-12-2024

let with_tz = datetime.parse("2024-01-01 00:00:00", "2006-01-02 15:04:05", "America/New_York")
```

### `datetime.fromTimestamp(timestamp, [timezone])`

Create a Time object from a Unix timestamp.

```js
let time_from_ts = datetime.fromTimestamp(1704063000)
print(time_from_ts)  // Unix timestamp converted to local time
```

## Duration Functions

### `datetime.duration(string | keyword_args)`

Create a Duration object from a string or keyword arguments.

```js
// From string
let dur1 = datetime.duration("2h30m15s")

// From keyword arguments
let dur2 = datetime.duration(hours=2, minutes=30, seconds=15, days=1, weeks=1)

// Supported units: nanoseconds, microseconds, milliseconds, seconds, minutes, hours, days, weeks
```

### `datetime.sleep(duration)`

Sleep for a specified duration.

```js
datetime.sleep(datetime.duration("2s"))  // Sleep for 2 seconds
datetime.sleep(5)  // Sleep for 5 seconds (integer)
datetime.sleep("1m30s")  // Sleep for 1 minute 30 seconds (string)
```

## Time Utility Functions

### `datetime.since(time)`

Get the duration since a specific time.

```js
let past_time = datetime.parse("2024-01-01 00:00:00", "2006-01-02 15:04:05")
let duration_since = datetime.since(past_time)
print(duration_since)  // Duration since Jan 1, 2024
```

### `datetime.until(time)`

Get the duration until a future time.

```js
let future_time = datetime.parse("2025-12-31 23:59:59", "2006-01-02 15:04:05")
let duration_until = datetime.until(future_time)
print(duration_until)  // Duration until Dec 31, 2025
```

### `datetime.isLeapYear(year)`

Check if a year is a leap year.

```js
print(datetime.isLeapYear(2024))  // true
print(datetime.isLeapYear(2023))  // false
```

### `datetime.daysInMonth(year, month)`

Get the number of days in a specific month.

```js
print(datetime.daysInMonth(2024, 2))  // 29 (February in leap year)
print(datetime.daysInMonth(2023, 2))  // 28 (February in regular year)
```

## Period Boundary Functions

### `datetime.startOfDay(time)`

Get the start of the day (00:00:00) for a given time.

```js
let current = datetime.now()
let start = datetime.startOfDay(current)
print(start)  // 00:00:00 26-09-2025
```

### `datetime.endOfDay(time)`

Get the end of the day (23:59:59) for a given time.

```js
let current = datetime.now()
let end = datetime.endOfDay(current)
print(end)  // 23:59:59 26-09-2025
```

### `datetime.startOfWeek(time)`

Get the start of the week (Sunday 00:00:00) for a given time.

```js
let start_week = datetime.startOfWeek(datetime.now())
```

### `datetime.endOfWeek(time)`

Get the end of the week (Saturday 23:59:59) for a given time.

```js
let end_week = datetime.endOfWeek(datetime.now())
```

### `datetime.startOfMonth(time)`

Get the start of the month for a given time.

```js
let start_month = datetime.startOfMonth(datetime.now())
```

### `datetime.endOfMonth(time)`

Get the end of the month for a given time.

```js
let end_month = datetime.endOfMonth(datetime.now())
```

### `datetime.startOfYear(time)`

Get the start of the year for a given time.

```js
let start_year = datetime.startOfYear(datetime.now())
```

### `datetime.endOfYear(time)`

Get the end of the year for a given time.

```js
let end_year = datetime.endOfYear(datetime.now())
```

## Time Object Methods

Time objects returned by datetime functions have many useful methods:

### Basic Properties
```js
let time = datetime.now()
print(time.year())      // 2025
print(time.month())     // 9
print(time.day())       // 26
print(time.hour())      // 10
print(time.minute())    // 15
print(time.second())    // 32
print(time.nanosecond()) // Nanosecond component
print(time.weekday())   // "Friday"
print(time.yearDay())   // Day of year (1-366)
```

### ISO Week
```js
let iso = time.isoWeek()
print(iso["year"])  // ISO week year
print(iso["week"])  // ISO week number
```

### Time Arithmetic
```js
let time = datetime.now()
let duration = datetime.duration(hours=2, minutes=30)

// Add/subtract durations
let future = time.add(duration)
let past = time.subtract(duration)

// Add/subtract specific units
let tomorrow = time.add(days=1)
let last_week = time.subtract(weeks=1)
```

### Time Comparisons
```js
let time1 = datetime.now()
let time2 = datetime.parse("2025-01-01 00:00:00", "2006-01-02 15:04:05")

print(time1.before(time2))  // true/false
print(time1.after(time2))   // true/false
print(time1.equal(time2))   // true/false
print(time1.compare(time2)) // -1, 0, or 1
```

### Timezone Operations
```js
let time = datetime.now()

// Get current timezone
print(time.timezone())  // "UTC" or local timezone name

// Convert to specific timezone
let ny_time = time.timezone("America/New_York")
let utc_time = time.utc()
let local_time = time.local()
```

### Other Methods
```js
let time = datetime.now()

// Get Unix timestamp
print(time.timestamp())  // Unix timestamp as integer

// Format the time
print(time.format("2006-01-02 15:04:05"))  // Custom formatting

// Truncate/round to duration
let truncated = time.truncate("1h")  // Truncate to hour boundary
let rounded = time.round("15m")      // Round to nearest 15 minutes
```

## Duration Object Methods

Duration objects have methods for accessing different time units:

```js
let duration = datetime.duration(hours=2, minutes=30, seconds=15)

print(duration.hours())        // 2.5041666666666664
print(duration.minutes())      // 150.25
print(duration.seconds())      // 9015
print(duration.milliseconds()) // 9015000
print(duration.nanoseconds())  // 9015000000000
print(duration.string())       // "2h30m15s"
```

### Duration Arithmetic
```js
let dur1 = datetime.duration("1h")
let dur2 = datetime.duration("30m")

let sum = dur1.add(dur2)          // 1h30m
let diff = dur1.subtract(dur2)    // 30m
let product = dur1.multiply(2)    // 2h
let quotient = dur1.divide(2)     // 30m
let ratio = dur1.divide(dur2)     // 2.0 (ratio as float)
```

## Timezone Support

The datetime module supports timezone-aware operations:

### Available Timezones
Common timezone identifiers include:
- `UTC`
- `America/New_York`
- `America/Los_Angeles`
- `Europe/London`
- `Europe/Paris`
- `Asia/Tokyo`
- `Asia/Shanghai`
- And many more standard IANA timezone identifiers

### Examples
```js
// Current time in different timezones
let utc = datetime.now("UTC")
let ny = datetime.now("America/New_York")
let tokyo = datetime.now("Asia/Tokyo")

// Convert between timezones
let local_time = datetime.now()
let ny_time = local_time.timezone("America/New_York")
```

## Practical Examples

### Age Calculator
```js
let calculate_age = func(birth_date_str) {
    let birth = datetime.parse(birth_date_str, "2006-01-02")
    let current = datetime.now()
    let age_duration = current.subtract(birth)
    let age_years = age_duration.hours() / (24 * 365.25)
    return age_years.floor()
}

let age = calculate_age("1990-05-15")
print("Age:", age, "years")
```

### Meeting Scheduler
```js
let schedule_meeting = func(date_str, duration_str, timezone) {
    let start_time = datetime.parse(date_str, "2006-01-02 15:04:05", timezone)
    let duration = datetime.duration(duration_str)
    let end_time = start_time.add(duration)
    
    print("Meeting scheduled:")
    print("Start:", start_time.format("2006-01-02 15:04:05 MST"))
    print("End:", end_time.format("2006-01-02 15:04:05 MST"))
    print("Duration:", duration)
}

schedule_meeting("2024-12-25 14:00:00", "1h30m", "America/New_York")
```

### Time Until Event
```js
let time_until_event = func(event_date_str) {
    let event = datetime.parse(event_date_str, "2006-01-02 15:04:05")
    let now = datetime.now()
    
    if (now.after(event)) {
        print("Event has already passed!")
        return
    }
    
    let duration = datetime.until(event)
    let days = duration.hours() / 24
    print("Time until event:", days.floor(), "days")
}

time_until_event("2024-12-31 23:59:59")
```

## Integration with Time Module

The datetime module works alongside the existing time module. You can use both:

```js
import time
import datetime

// Traditional time module
let time_now = time.now()
print("Time module:", time_now)

// Enhanced datetime module  
let datetime_now = datetime.now()
print("DateTime module:", datetime_now)

// They can work together
let formatted = time.format(time_now, "2006-01-02 15:04:05")
let parsed = datetime.parse(formatted, "2006-01-02 15:04:05")
```

The datetime module provides all the functionality of the time module and much more, making it the recommended choice for complex date and time operations.
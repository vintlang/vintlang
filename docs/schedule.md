# Schedule Module

The `schedule` module provides powerful scheduling capabilities for VintLang, similar to Go's ticker and NestJS's schedule decorators. It allows you to execute functions at regular intervals or at specific times using cron-like expressions.

## Functions

### `ticker(intervalSeconds, callback)`

Creates a ticker that executes a callback function at regular intervals.

**Parameters:**
- `intervalSeconds` (integer): The interval in seconds between executions
- `callback` (function): The function to execute at each interval

**Returns:** A ticker object that can be stopped with `stopTicker()`

**Example:**
```javascript
import schedule

// Execute every 5 seconds
let ticker = schedule.ticker(5, func() {
    print("Tick at", time.now())
})

// Stop the ticker after some time
time.sleep(20)
schedule.stopTicker(ticker)
```

### `stopTicker(tickerObj)`

Stops a running ticker.

**Parameters:**
- `tickerObj`: The ticker object returned by `ticker()`

**Returns:** Boolean indicating if the ticker was successfully stopped

### `schedule(cronExpr, callback)`

Schedules a function to execute at specific times using cron-like expressions.

**Parameters:**
- `cronExpr` (string): A cron expression in the format "second minute hour day month weekday"
- `callback` (function): The function to execute when the schedule triggers

**Returns:** A schedule object that can be stopped with `stopSchedule()`

**Cron Expression Format:**
```
second minute hour day month weekday
  |      |     |    |    |      |
  |      |     |    |    |      +-- Day of week (0-6, 0=Sunday)
  |      |     |    |    +--------- Month (1-12)
  |      |     |    +-------------- Day of month (1-31)
  |      |     +------------------- Hour (0-23)
  |      +------------------------- Minute (0-59)
  +-------------------------------- Second (0-59)
```

Use `*` for wildcards (any value).

**Examples:**
```javascript
import schedule

// Every minute at second 0
schedule.schedule("0 * * * * *", func() {
    print("Top of the minute!")
})

// Daily at 9:30 AM
schedule.schedule("0 30 9 * * *", func() {
    print("Good morning!")
})

// Every 30 seconds
schedule.schedule("*/30 * * * * *", func() {
    print("Every 30 seconds")
})

// Every Friday at 5:00 PM
schedule.schedule("0 0 17 * * 5", func() {
    print("TGIF!")
})
```

### `stopSchedule(scheduleObj)`

Stops a running scheduled task.

**Parameters:**
- `scheduleObj`: The schedule object returned by `schedule()`

**Returns:** Boolean indicating if the schedule was successfully stopped

## Helper Functions

The module provides convenient helper functions for common scheduling patterns:

### `everySecond(callback)`

Executes a callback every second. Equivalent to `ticker(1, callback)`.

**Example:**
```javascript
let job = schedule.everySecond(func() {
    print("Ping!")
})
```

### `everyMinute(callback)`

Executes a callback every minute at second 0. Equivalent to `schedule("0 * * * * *", callback)`.

**Example:**
```javascript
let job = schedule.everyMinute(func() {
    print("Another minute passed")
})
```

### `everyHour(callback)`

Executes a callback every hour at minute 0, second 0. Equivalent to `schedule("0 0 * * * *", callback)`.

**Example:**
```javascript
let job = schedule.everyHour(func() {
    print("It's a new hour!")
})
```

### `daily(hour, minute, callback)`

Executes a callback daily at the specified time.

**Parameters:**
- `hour` (integer): Hour of the day (0-23)
- `minute` (integer): Minute of the hour (0-59)
- `callback` (function): The function to execute

**Example:**
```javascript
// Daily reminder at 2:30 PM
let job = schedule.daily(14, 30, func() {
    print("Time for afternoon coffee!")
})
```

## Complete Example

```javascript
import schedule
import time

print("Starting scheduling demo...")

// Ticker example - every 3 seconds
let ticker = schedule.ticker(3, func() {
    print("[TICKER] Current time:", time.now())
})

// Schedule example - every 10 seconds
let job1 = schedule.schedule("*/10 * * * * *", func() {
    print("[SCHEDULE] Every 10 seconds")
})

// Helper function example
let job2 = schedule.everySecond(func() {
    print("[HELPER] Ping!")
})

// Daily example (would execute at specified time)
let job3 = schedule.daily(9, 0, func() {
    print("[DAILY] Good morning! Time to start the day.")
})

// Let everything run for 30 seconds
print("Running for 30 seconds...")
time.sleep(30)

// Clean up
print("Stopping all jobs...")
schedule.stopTicker(ticker)
schedule.stopSchedule(job1)
schedule.stopTicker(job2)
schedule.stopSchedule(job3)

print("Demo completed!")
```

## Notes

- All scheduling is done using Go's `time.Ticker` and `time.Timer` internally
- Callbacks currently log execution messages to demonstrate functionality
- The module uses goroutines for non-blocking execution
- Proper cleanup is important - always stop tickers and schedules when done
- Cron expressions support basic patterns; advanced features like ranges (1-5) or lists (1,3,5) are not yet implemented

## Error Handling

The module provides comprehensive error messages for common mistakes:

- Invalid argument types
- Invalid time ranges (e.g., hour > 23, minute > 59)
- Malformed cron expressions
- Negative intervals for tickers

All errors include usage examples to help with correct implementation. 
# schedule Module

The `schedule` module provides scheduling utilities similar to Go's `ticker` and NestJS's `schedule`. It is designed to allow repeated or scheduled execution of code at intervals or specific times.

**Note:** Callback execution is not yet supported in this build. The API is provided for future compatibility, but calling these functions will return an error until callback support is implemented.

## Functions

### ticker(intervalSeconds, callback)
Creates a ticker that would call the provided callback every `intervalSeconds` seconds.

- `intervalSeconds`: Integer. The interval in seconds between each tick.
- `callback`: Function. The function to call on each tick.

**Returns:** Ticker object (for future use with `stopTicker`).

**Current Limitation:** Returns an error: "ticker callback execution is not yet supported in this build".

#### Example
```vint
import schedule
let t = schedule.ticker(2, func() { print("Tick!") })
```

### stopTicker(tickerObj)
Stops a running ticker.

- `tickerObj`: The ticker object returned by `ticker()`.

**Returns:** Boolean indicating success.

### schedule(cronExpr, callback)
Schedules a callback to run at times specified by a cron-like expression.

- `cronExpr`: String. Format: "second minute hour day month weekday" (e.g., "0 30 14 * * *" for 14:30:00 every day).
- `callback`: Function. The function to call at the scheduled time.

**Returns:** Schedule object (for future use with `stopSchedule`).

**Current Limitation:** Returns an error: "schedule callback execution is not yet supported in this build".

#### Example
```vint
import schedule
let s = schedule.schedule("0 0 9 * * *", func() { print("Good morning!") })
```

### stopSchedule(scheduleObj)
Stops a scheduled job.

- `scheduleObj`: The schedule object returned by `schedule()`.

**Returns:** Boolean indicating success.

## Limitations
- Callback execution is not yet supported. The API is present for future compatibility.
- Cron expressions are basic and support wildcards (*).

## Future Work
- Implement a Vint event/callback queue to allow safe callback execution from Go routines.
- Support for more advanced cron syntax. 
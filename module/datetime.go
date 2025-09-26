package module

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vintlang/vintlang/object"
)

var DatetimeFunctions = map[string]object.ModuleFunction{}

func init() {
	DatetimeFunctions["now"] = datetimeNow
	DatetimeFunctions["parse"] = datetimeParse
	DatetimeFunctions["fromTimestamp"] = datetimeFromTimestamp
	DatetimeFunctions["utcNow"] = datetimeUtcNow
	DatetimeFunctions["duration"] = datetimeDuration
	DatetimeFunctions["timezone"] = datetimeTimezone
	DatetimeFunctions["sleep"] = datetimeSleep
	DatetimeFunctions["since"] = datetimeSince
	DatetimeFunctions["until"] = datetimeUntil
	DatetimeFunctions["isLeapYear"] = datetimeIsLeapYear
	DatetimeFunctions["daysInMonth"] = datetimeDaysInMonth
	DatetimeFunctions["startOfDay"] = datetimeStartOfDay
	DatetimeFunctions["endOfDay"] = datetimeEndOfDay
	DatetimeFunctions["startOfWeek"] = datetimeStartOfWeek
	DatetimeFunctions["endOfWeek"] = datetimeEndOfWeek
	DatetimeFunctions["startOfMonth"] = datetimeStartOfMonth
	DatetimeFunctions["endOfMonth"] = datetimeEndOfMonth
	DatetimeFunctions["startOfYear"] = datetimeStartOfYear
	DatetimeFunctions["endOfYear"] = datetimeEndOfYear
}

func datetimeNow(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 1 {
		return ErrorMessage(
			"datetime", "now",
			"0 or 1 argument (optional timezone)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.now() or datetime.now('America/New_York')",
		)
	}

	var location *time.Location = time.Local
	var err error

	// Handle timezone parameter
	if len(args) == 1 {
		timezoneStr := args[0].Inspect()
		location, err = time.LoadLocation(timezoneStr)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Invalid timezone: %s", timezoneStr)}
		}
	}

	now := time.Now().In(location)
	return &object.Time{TimeValue: now.Format("15:04:05 02-01-2006")}
}

func datetimeUtcNow(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 || len(defs) != 0 {
		return ErrorMessage(
			"datetime", "utcNow",
			"no arguments", 
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.utcNow() -> returns current UTC timestamp",
		)
	}

	utcNow := time.Now().UTC()
	return &object.Time{TimeValue: utcNow.Format("15:04:05 02-01-2006")}
}

func datetimeParse(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 3 {
		return ErrorMessage(
			"datetime", "parse",
			"1-3 arguments (datetime_string, format, timezone)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.parse('2024-01-15 10:30:00', '2006-01-02 15:04:05', 'America/New_York')",
		)
	}

	dateTimeStr := args[0].Inspect()
	format := "2006-01-02 15:04:05"
	var location *time.Location = time.Local

	// Handle format parameter
	if len(args) >= 2 {
		format = args[1].Inspect()
	}

	// Handle timezone parameter
	if len(args) == 3 {
		timezoneStr := args[2].Inspect()
		var err error
		location, err = time.LoadLocation(timezoneStr)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Invalid timezone: %s", timezoneStr)}
		}
	}

	t, err := time.ParseInLocation(format, dateTimeStr, location)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to parse datetime: %s", err.Error())}
	}

	return &object.Time{TimeValue: t.Format("15:04:05 02-01-2006")}
}

func datetimeFromTimestamp(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return ErrorMessage(
			"datetime", "fromTimestamp",
			"1-2 arguments (timestamp, timezone)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.fromTimestamp(1704063000, 'America/New_York')",
		)
	}

	timestampObj, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "First argument must be a timestamp (integer)"}
	}

	var location *time.Location = time.Local

	// Handle timezone parameter
	if len(args) == 2 {
		timezoneStr := args[1].Inspect()
		var err error
		location, err = time.LoadLocation(timezoneStr)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Invalid timezone: %s", timezoneStr)}
		}
	}

	t := time.Unix(timestampObj.Value, 0).In(location)
	return &object.Time{TimeValue: t.Format("15:04:05 02-01-2006")}
}

func datetimeDuration(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) > 0 {
		// Handle keyword arguments for duration creation
		var totalDuration time.Duration
		
		for key, value := range defs {
			valueStr := value.Inspect()
			valueInt, err := strconv.ParseInt(valueStr, 10, 64)
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Invalid duration value for %s: %s", key, valueStr)}
			}

			switch key {
			case "nanoseconds":
				totalDuration += time.Duration(valueInt) * time.Nanosecond
			case "microseconds":
				totalDuration += time.Duration(valueInt) * time.Microsecond
			case "milliseconds":
				totalDuration += time.Duration(valueInt) * time.Millisecond
			case "seconds":
				totalDuration += time.Duration(valueInt) * time.Second
			case "minutes":
				totalDuration += time.Duration(valueInt) * time.Minute
			case "hours":
				totalDuration += time.Duration(valueInt) * time.Hour
			case "days":
				totalDuration += time.Duration(valueInt) * time.Hour * 24
			case "weeks":
				totalDuration += time.Duration(valueInt) * time.Hour * 24 * 7
			default:
				return &object.Error{Message: fmt.Sprintf("Invalid duration unit: %s", key)}
			}
		}
		
		return &object.Duration{Value: totalDuration}
	}

	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "duration",
			"1 argument (duration string) or keyword arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.duration('2h30m') or datetime.duration(hours=2, minutes=30)",
		)
	}

	durationStr := args[0].Inspect()
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Invalid duration format: %s", err.Error())}
	}

	return &object.Duration{Value: duration}
}

func datetimeTimezone(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"datetime", "timezone",
			"2 arguments (time, timezone)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.timezone(time.now(), 'America/New_York')",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "First argument must be a time object"}
	}

	timezoneStr := args[1].Inspect()
	location, err := time.LoadLocation(timezoneStr)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Invalid timezone: %s", timezoneStr)}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	convertedTime := t.In(location)
	return &object.Time{TimeValue: convertedTime.Format("15:04:05 02-01-2006")}
}

func datetimeSleep(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "sleep",
			"1 argument (duration)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.sleep(datetime.duration('2s'))",
		)
	}

	switch arg := args[0].(type) {
	case *object.Duration:
		time.Sleep(arg.Value)
	case *object.Integer:
		time.Sleep(time.Duration(arg.Value) * time.Second)
	case *object.String:
		duration, err := time.ParseDuration(arg.Value)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Invalid duration format: %s", err.Error())}
		}
		time.Sleep(duration)
	default:
		return &object.Error{Message: "Argument must be a duration, integer (seconds), or duration string"}
	}

	return nil
}

func datetimeSince(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "since",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.since(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	duration := time.Since(t)
	return &object.Duration{Value: duration}
}

func datetimeUntil(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "until",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.until(future_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	duration := time.Until(t)
	return &object.Duration{Value: duration}
}

func datetimeIsLeapYear(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "isLeapYear",
			"1 argument (year)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.isLeapYear(2024)",
		)
	}

	yearObj, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "Argument must be an integer year"}
	}

	year := int(yearObj.Value)
	isLeap := time.Date(year, time.February, 29, 0, 0, 0, 0, time.UTC).Month() == time.February
	return &object.Boolean{Value: isLeap}
}

func datetimeDaysInMonth(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"datetime", "daysInMonth",
			"2 arguments (year, month)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.daysInMonth(2024, 2)",
		)
	}

	yearObj, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "First argument must be an integer year"}
	}

	monthObj, ok := args[1].(*object.Integer)
	if !ok {
		return &object.Error{Message: "Second argument must be an integer month"}
	}

	year := int(yearObj.Value)
	month := time.Month(monthObj.Value)

	if month < 1 || month > 12 {
		return &object.Error{Message: "Month must be between 1 and 12"}
	}

	firstOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastOfThisMonth := firstOfNextMonth.AddDate(0, 0, -1)
	
	return &object.Integer{Value: int64(lastOfThisMonth.Day())}
}

func datetimeStartOfDay(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "startOfDay",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.startOfDay(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return &object.Time{TimeValue: startOfDay.Format("15:04:05 02-01-2006")}
}

func datetimeEndOfDay(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "endOfDay",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.endOfDay(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	return &object.Time{TimeValue: endOfDay.Format("15:04:05 02-01-2006")}
}

func datetimeStartOfWeek(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "startOfWeek",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.startOfWeek(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	// Calculate days since Sunday (Go's weekday starts with Sunday = 0)
	daysFromSunday := int(t.Weekday())
	startOfWeek := t.AddDate(0, 0, -daysFromSunday)
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())
	
	return &object.Time{TimeValue: startOfWeek.Format("15:04:05 02-01-2006")}
}

func datetimeEndOfWeek(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "endOfWeek",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.endOfWeek(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	// Calculate days until Saturday
	daysUntilSaturday := 6 - int(t.Weekday())
	endOfWeek := t.AddDate(0, 0, daysUntilSaturday)
	endOfWeek = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 999999999, endOfWeek.Location())
	
	return &object.Time{TimeValue: endOfWeek.Format("15:04:05 02-01-2006")}
}

func datetimeStartOfMonth(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "startOfMonth",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.startOfMonth(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	startOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return &object.Time{TimeValue: startOfMonth.Format("15:04:05 02-01-2006")}
}

func datetimeEndOfMonth(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "endOfMonth",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.endOfMonth(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	// Get first day of next month, then subtract one day
	firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	endOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	endOfMonth = time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, endOfMonth.Location())
	
	return &object.Time{TimeValue: endOfMonth.Format("15:04:05 02-01-2006")}
}

func datetimeStartOfYear(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "startOfYear",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.startOfYear(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	startOfYear := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
	return &object.Time{TimeValue: startOfYear.Format("15:04:05 02-01-2006")}
}

func datetimeEndOfYear(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"datetime", "endOfYear",
			"1 argument (time)",
			fmt.Sprintf("%d arguments", len(args)),
			"datetime.endOfYear(some_time)",
		)
	}

	timeObj, ok := args[0].(*object.Time)
	if !ok {
		return &object.Error{Message: "Argument must be a time object"}
	}

	t, err := time.Parse("15:04:05 02-01-2006", timeObj.TimeValue)
	if err != nil {
		return &object.Error{Message: "Invalid time format"}
	}

	endOfYear := time.Date(t.Year(), time.December, 31, 23, 59, 59, 999999999, t.Location())
	return &object.Time{TimeValue: endOfYear.Format("15:04:05 02-01-2006")}
}
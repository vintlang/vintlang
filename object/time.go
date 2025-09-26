package object

import (
	"fmt"
	"strconv"
	"time"
)

type Time struct {
	TimeValue string
}

func (t *Time) Type() ObjectType { return TIME_OBJ }
func (t *Time) Inspect() string  { return t.TimeValue }
func (t *Time) Method(method string, args []Object, defs map[string]Object) Object {
	switch method {
	case "add":
		return t.add(args, defs)
	case "subtract":
		return t.subtract(args, defs)
	case "since":
		return t.since(args, defs)
	case "until":
		return t.until(args, defs)
	case "format":
		return t.format(args, defs)
	case "year":
		return t.year(args, defs)
	case "month":
		return t.month(args, defs)
	case "day":
		return t.day(args, defs)
	case "hour":
		return t.hour(args, defs)
	case "minute":
		return t.minute(args, defs)
	case "second":
		return t.second(args, defs)
	case "nanosecond":
		return t.nanosecond(args, defs)
	case "weekday":
		return t.weekday(args, defs)
	case "yearDay":
		return t.yearDay(args, defs)
	case "isoWeek":
		return t.isoWeek(args, defs)
	case "timezone":
		return t.timezone(args, defs)
	case "utc":
		return t.utc(args, defs)
	case "local":
		return t.local(args, defs)
	case "timestamp":
		return t.timestamp(args, defs)
	case "compare":
		return t.compare(args, defs)
	case "before":
		return t.before(args, defs)
	case "after":
		return t.after(args, defs)
	case "equal":
		return t.equal(args, defs)
	case "truncate":
		return t.truncate(args, defs)
	case "round":
		return t.round(args, defs)
	}
	return nil
}

func (t *Time) add(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		var sec, min, hr, d, m, y int
		for k, v := range defs {
			objvalue := v.Inspect()
			inttime, err := strconv.Atoi(objvalue)
			if err != nil {
				return newError("Only numbers are allowed as arguments")
			}
			switch k {
			case "seconds":
				sec = inttime
			case "minutes":
				min = inttime
			case "hours":
				hr = inttime
			case "days":
				d = inttime
			case "months":
				m = inttime
			case "years":
				y = inttime
			default:
				return newError("Invalid time key provided")
			}
		}
		curTime, _ := time.Parse("15:04:05 02-01-2006", t.Inspect())
		nextTime := curTime.
			Add(time.Duration(sec)*time.Second).
			Add(time.Duration(min)*time.Minute).
			Add(time.Duration(hr)*time.Hour).
			AddDate(y, m, d)
		return &Time{TimeValue: string(nextTime.Format("15:04:05 02-01-2006"))}
	}

	if len(args) != 1 {
		return newError("We require exactly 1 argument, but you provided %d", len(args))
	}

	curTime, _ := time.Parse("15:04:05 02-01-2006", t.Inspect())

	switch arg := args[0].(type) {
	case *Duration:
		nextTime := curTime.Add(arg.Value)
		return &Time{TimeValue: nextTime.Format("15:04:05 02-01-2006")}
	case *String:
		duration, err := time.ParseDuration(arg.Value)
		if err != nil {
			return newError("Invalid duration format: %s", err.Error())
		}
		nextTime := curTime.Add(duration)
		return &Time{TimeValue: nextTime.Format("15:04:05 02-01-2006")}
	default:
		// Legacy behavior: treat as hours
		objvalue := args[0].Inspect()
		inttime, err := strconv.Atoi(objvalue)
		if err != nil {
			return newError("Only numbers, durations, or duration strings are allowed as arguments")
		}
		nextTime := curTime.Add(time.Duration(inttime) * time.Hour)
		return &Time{TimeValue: string(nextTime.Format("15:04:05 02-01-2006"))}
	}
}

func (t *Time) since(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return &Error{Message: "This argument is not allowed"}
	}
	if len(args) != 1 {
		return &Error{Message: "We require exactly one argument"}
	}

	var (
		o   time.Time
		err error
	)

	switch m := args[0].(type) {
	case *Time:
		o, _ = time.Parse("15:04:05 02-01-2006", m.TimeValue)
	case *String:
		o, err = time.Parse("15:04:05 02-01-2006", m.Value)
		if err != nil {
			return &Error{Message: fmt.Sprintf("Invalid argument: %s", args[0].Inspect())}
		}
	default:
		return &Error{Message: fmt.Sprintf("Invalid argument: %s", args[0].Inspect())}
	}

	ct, _ := time.Parse("15:04:05 02-01-2006", t.TimeValue)

	diff := ct.Sub(o)
	durationInSeconds := diff.Seconds()

	return &Integer{Value: int64(durationInSeconds)}
}

func (t *Time) format(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("format() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("format() expects 1 argument, got %d", len(args))
	}
	
	pattern, ok := args[0].(*String)
	if !ok {
		return newError("Pattern must be a string")
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &String{Value: curTime.Format(pattern.Value)}
}

func (t *Time) year(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("year() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("year() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.Year())}
}

func (t *Time) month(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("month() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("month() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.Month())}
}

func (t *Time) day(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("day() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("day() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.Day())}
}

func (t *Time) hour(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("hour() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("hour() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.Hour())}
}

func (t *Time) minute(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("minute() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("minute() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.Minute())}
}

func (t *Time) second(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("second() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("second() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.Second())}
}

func (t *Time) weekday(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("weekday() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("weekday() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &String{Value: curTime.Weekday().String()}
}

func (t *Time) subtract(args []Object, defs map[string]Object) Object {
	if len(args) == 1 {
		// Subtracting duration from time
		switch arg := args[0].(type) {
		case *Duration:
			curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
			if err != nil {
				return newError("Invalid time format")
			}
			newTime := curTime.Add(-arg.Value)
			return &Time{TimeValue: newTime.Format("15:04:05 02-01-2006")}
		case *String:
			duration, err := time.ParseDuration(arg.Value)
			if err != nil {
				return newError("Invalid duration format: %s", err.Error())
			}
			curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
			if err != nil {
				return newError("Invalid time format")
			}
			newTime := curTime.Add(-duration)
			return &Time{TimeValue: newTime.Format("15:04:05 02-01-2006")}
		case *Time:
			// Subtracting time from time returns duration
			curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
			if err != nil {
				return newError("Invalid time format")
			}
			otherTime, err := time.Parse("15:04:05 02-01-2006", arg.TimeValue)
			if err != nil {
				return newError("Invalid time format")
			}
			diff := curTime.Sub(otherTime)
			return &Duration{Value: diff}
		default:
			return newError("subtract() argument must be a duration or time")
		}
	}

	// Handle keyword arguments for precise time subtraction
	if len(defs) != 0 {
		var sec, min, hr, d, m, y int
		for k, v := range defs {
			objvalue := v.Inspect()
			inttime, err := strconv.Atoi(objvalue)
			if err != nil {
				return newError("Only numbers are allowed as arguments")
			}
			switch k {
			case "seconds":
				sec = inttime
			case "minutes":
				min = inttime
			case "hours":
				hr = inttime
			case "days":
				d = inttime
			case "months":
				m = inttime
			case "years":
				y = inttime
			default:
				return newError("Invalid time key provided")
			}
		}
		curTime, _ := time.Parse("15:04:05 02-01-2006", t.TimeValue)
		nextTime := curTime.
			Add(-time.Duration(sec)*time.Second).
			Add(-time.Duration(min)*time.Minute).
			Add(-time.Duration(hr)*time.Hour).
			AddDate(-y, -m, -d)
		return &Time{TimeValue: nextTime.Format("15:04:05 02-01-2006")}
	}

	return newError("subtract() requires 1 argument or keyword arguments")
}

func (t *Time) until(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("until() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("until() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	duration := time.Until(curTime)
	return &Duration{Value: duration}
}

func (t *Time) nanosecond(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("nanosecond() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("nanosecond() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.Nanosecond())}
}

func (t *Time) yearDay(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("yearDay() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("yearDay() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: int64(curTime.YearDay())}
}

func (t *Time) isoWeek(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("isoWeek() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("isoWeek() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	year, week := curTime.ISOWeek()
	
	pairs := make(map[HashKey]DictPair)
	yearKey := &String{Value: "year"}
	weekKey := &String{Value: "week"}
	
	pairs[yearKey.HashKey()] = DictPair{Key: yearKey, Value: &Integer{Value: int64(year)}}
	pairs[weekKey.HashKey()] = DictPair{Key: weekKey, Value: &Integer{Value: int64(week)}}
	
	return &Dict{Pairs: pairs}
}

func (t *Time) timezone(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("timezone() does not accept keyword arguments")
	}
	if len(args) > 1 {
		return newError("timezone() expects 0 or 1 argument, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	if len(args) == 0 {
		// Return current timezone
		zone, _ := curTime.Zone()
		return &String{Value: zone}
	}
	
	// Convert to specified timezone
	timezoneStr := args[0].Inspect()
	location, err := time.LoadLocation(timezoneStr)
	if err != nil {
		return newError("Invalid timezone: %s", timezoneStr)
	}
	
	convertedTime := curTime.In(location)
	return &Time{TimeValue: convertedTime.Format("15:04:05 02-01-2006")}
}

func (t *Time) utc(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("utc() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("utc() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	utcTime := curTime.UTC()
	return &Time{TimeValue: utcTime.Format("15:04:05 02-01-2006")}
}

func (t *Time) local(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("local() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("local() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	localTime := curTime.Local()
	return &Time{TimeValue: localTime.Format("15:04:05 02-01-2006")}
}

func (t *Time) timestamp(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("timestamp() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("timestamp() expects 0 arguments, got %d", len(args))
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Integer{Value: curTime.Unix()}
}

func (t *Time) compare(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("compare() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("compare() expects 1 argument, got %d", len(args))
	}
	
	otherTime, ok := args[0].(*Time)
	if !ok {
		return newError("compare() argument must be a time object")
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	otherTimeParsed, err := time.Parse("15:04:05 02-01-2006", otherTime.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	if curTime.Before(otherTimeParsed) {
		return &Integer{Value: -1}
	} else if curTime.After(otherTimeParsed) {
		return &Integer{Value: 1}
	} else {
		return &Integer{Value: 0}
	}
}

func (t *Time) before(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("before() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("before() expects 1 argument, got %d", len(args))
	}
	
	otherTime, ok := args[0].(*Time)
	if !ok {
		return newError("before() argument must be a time object")
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	otherTimeParsed, err := time.Parse("15:04:05 02-01-2006", otherTime.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Boolean{Value: curTime.Before(otherTimeParsed)}
}

func (t *Time) after(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("after() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("after() expects 1 argument, got %d", len(args))
	}
	
	otherTime, ok := args[0].(*Time)
	if !ok {
		return newError("after() argument must be a time object")
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	otherTimeParsed, err := time.Parse("15:04:05 02-01-2006", otherTime.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Boolean{Value: curTime.After(otherTimeParsed)}
}

func (t *Time) equal(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("equal() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("equal() expects 1 argument, got %d", len(args))
	}
	
	otherTime, ok := args[0].(*Time)
	if !ok {
		return newError("equal() argument must be a time object")
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	otherTimeParsed, err := time.Parse("15:04:05 02-01-2006", otherTime.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	return &Boolean{Value: curTime.Equal(otherTimeParsed)}
}

func (t *Time) truncate(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("truncate() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("truncate() expects 1 argument, got %d", len(args))
	}
	
	var duration time.Duration
	var err error
	
	switch arg := args[0].(type) {
	case *Duration:
		duration = arg.Value
	case *String:
		duration, err = time.ParseDuration(arg.Value)
		if err != nil {
			return newError("Invalid duration format: %s", err.Error())
		}
	default:
		return newError("truncate() argument must be a duration")
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	truncatedTime := curTime.Truncate(duration)
	return &Time{TimeValue: truncatedTime.Format("15:04:05 02-01-2006")}
}

func (t *Time) round(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return newError("round() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("round() expects 1 argument, got %d", len(args))
	}
	
	var duration time.Duration
	var err error
	
	switch arg := args[0].(type) {
	case *Duration:
		duration = arg.Value
	case *String:
		duration, err = time.ParseDuration(arg.Value)
		if err != nil {
			return newError("Invalid duration format: %s", err.Error())
		}
	default:
		return newError("round() argument must be a duration")
	}
	
	curTime, err := time.Parse("15:04:05 02-01-2006", t.TimeValue)
	if err != nil {
		return newError("Invalid time format")
	}
	
	roundedTime := curTime.Round(duration)
	return &Time{TimeValue: roundedTime.Format("15:04:05 02-01-2006")}
}

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
	case "since":
		return t.since(args, defs)
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

	objvalue := args[0].Inspect()
	inttime, err := strconv.Atoi(objvalue)

	if err != nil {
		return newError("Only numbers are allowed as arguments")
	}

	nextTime := curTime.Add(time.Duration(inttime) * time.Hour)
	return &Time{TimeValue: string(nextTime.Format("15:04:05 02-01-2006"))}
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

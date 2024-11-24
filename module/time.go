package module

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ekilie/vint-lang/object"
)

var TimeFunctions = map[string]object.ModuleFunction{}

func init() {
	TimeFunctions["now"] = now
	TimeFunctions["sleep"] = sleep
	TimeFunctions["since"] = since
}

func now(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 || len(defs) != 0 {
		return &object.Error{Message: "No arguments required here"}
	}

	tn := time.Now()
	time_string := tn.Format("15:04:05 02-01-2006")

	return &object.Time{TimeValue: time_string}
}

func sleep(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "This argument is not allowed"}
	}
	if len(args) != 1 {
		return &object.Error{Message: "We only need one argument"}
	}

	objvalue := args[0].Inspect()
	inttime, err := strconv.Atoi(objvalue)

	if err != nil {
		return &object.Error{Message: "Only numbers are allowed as arguments"}
	}

	time.Sleep(time.Duration(inttime) * time.Second)

	return nil
}

func since(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "This argument is not allowed"}
	}
	if len(args) != 1 {
		return &object.Error{Message: "We only need one argument"}
	}

	var (
		t   time.Time
		err error
	)

	switch m := args[0].(type) {
	case *object.Time:
		t, _ = time.Parse("15:04:05 02-01-2006", m.TimeValue)
	case *object.String:
		t, err = time.Parse("15:04:05 02-01-2006", m.Value)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Argument %s is not valid", args[0].Inspect())}
		}
	default:
		return &object.Error{Message: fmt.Sprintf("Argument %s is not valid", args[0].Inspect())}
	}

	current_time := time.Now().Format("15:04:05 02-01-2006")
	ct, _ := time.Parse("15:04:05 02-01-2006", current_time)

	diff := ct.Sub(t)
	durationInSeconds := diff.Seconds()

	return &object.Integer{Value: int64(durationInSeconds)}
}

func format(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: time and format string"}
	}

	// Parse the time argument
	var t time.Time
	switch m := args[0].(type) {
	case *object.Time:
		var err error
		t, err = time.Parse("15:04:05 02-01-2006", m.TimeValue)
		if err != nil {
			return &object.Error{Message: "Invalid time format"}
		}
	case *object.String:
		var err error
		t, err = time.Parse("15:04:05 02-01-2006", m.Value)
		if err != nil {
			return &object.Error{Message: "Invalid time format"}
		}
	default:
		return &object.Error{Message: "Invalid time argument"}
	}

	formatStr := args[1].Inspect()
	formattedTime := t.Format(formatStr)
	return &object.String{Value: formattedTime}
}


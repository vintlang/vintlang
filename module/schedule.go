package module

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/vintlang/vintlang/object"
)

var ScheduleFunctions = map[string]object.ModuleFunction{}

var (
	tickerStore     = make(map[*time.Ticker]struct{})
	tickerStoreMu   sync.Mutex
	scheduleStore   = make(map[*time.Timer]struct{})
	scheduleStoreMu sync.Mutex
)

func init() {
	ScheduleFunctions["ticker"] = tickerFunc
	ScheduleFunctions["stopTicker"] = stopTickerFunc
	ScheduleFunctions["schedule"] = scheduleFunc
	ScheduleFunctions["stopSchedule"] = stopScheduleFunc
}

// ticker(intervalSeconds, callback)
func tickerFunc(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "ticker(intervalSeconds, callback) requires 2 arguments"}
	}
	// interval, ok := args[0].(*object.Integer)
	// if !ok {
	// 	return &object.Error{Message: "intervalSeconds must be an integer"}
	// }
	// callback, ok := args[1].(*object.Function)
	// if !ok {
	// 	return &object.Error{Message: "callback must be a function"}
	// }
	return &object.Error{Message: "ticker callback execution is not yet supported in this build"}
}

// stopTicker(tickerObj)
func stopTickerFunc(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "stopTicker(tickerObj) requires 1 argument"}
	}
	native, ok := args[0].(*object.NativeObject)
	if !ok {
		return &object.Error{Message: "Argument must be a ticker object"}
	}
	ticker, ok := native.Value.(*time.Ticker)
	if !ok {
		return &object.Error{Message: "Not a valid ticker object"}
	}
	ticker.Stop()
	tickerStoreMu.Lock()
	delete(tickerStore, ticker)
	tickerStoreMu.Unlock()
	return &object.Boolean{Value: true}
}

// schedule(cronExpr, callback) - cronExpr: "second minute hour day month weekday" (basic, not full cron)
func scheduleFunc(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "schedule(cronExpr, callback) requires 2 arguments"}
	}
	// expr, ok := args[0].(*object.String)
	// if !ok {
	// 	return &object.Error{Message: "cronExpr must be a string"}
	// }
	// callback, ok := args[1].(*object.Function)
	// if !ok {
	// 	return &object.Error{Message: "callback must be a function"}
	// }
	return &object.Error{Message: "schedule callback execution is not yet supported in this build"}
}

// stopSchedule(scheduleObj)
func stopScheduleFunc(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "stopSchedule(scheduleObj) requires 1 argument"}
	}
	native, ok := args[0].(*object.NativeObject)
	if !ok {
		return &object.Error{Message: "Argument must be a schedule object"}
	}
	timer, ok := native.Value.(*time.Timer)
	if !ok {
		return &object.Error{Message: "Not a valid schedule object"}
	}
	stopped := timer.Stop()
	scheduleStoreMu.Lock()
	delete(scheduleStore, timer)
	scheduleStoreMu.Unlock()
	return &object.Boolean{Value: stopped}
}

// nextSchedule parses a basic cron string and returns the next time.Time
func nextSchedule(expr string) time.Time {
	// Only supports: "second minute hour * * *" (wildcards allowed)
	// Example: "0 30 14 * * *" = 14:30:00 every day
	fields := [6]int{-1, -1, -1, -1, -1, -1}
	parts := splitAndTrim(expr)
	if len(parts) != 6 {
		return time.Time{}
	}
	for i, p := range parts {
		if p == "*" {
			fields[i] = -1
		} else {
			v, err := parseInt(p)
			if err != nil {
				return time.Time{}
			}
			fields[i] = v
		}
	}
	now := time.Now()
	for i := 0; i < 366; i++ {
		cand := now.Add(time.Duration(i) * 24 * time.Hour)
		for h := 0; h < 24; h++ {
			for m := 0; m < 60; m++ {
				for s := 0; s < 60; s++ {
					if (fields[2] == -1 || h == fields[2]) && (fields[1] == -1 || m == fields[1]) && (fields[0] == -1 || s == fields[0]) {
						if (fields[3] == -1 || cand.Day() == fields[3]) && (fields[4] == -1 || int(cand.Month()) == fields[4]) && (fields[5] == -1 || int(cand.Weekday()) == fields[5]) {
							t := time.Date(cand.Year(), cand.Month(), cand.Day(), h, m, s, 0, cand.Location())
							if t.After(now) {
								return t
							}
						}
					}
				}
			}
		}
	}
	return time.Time{}
}

func splitAndTrim(s string) []string {
	return strings.Fields(s)
}

func parseInt(s string) (int, error) {
	var v int
	_, err := fmt.Sscanf(s, "%d", &v)
	return v, err
}

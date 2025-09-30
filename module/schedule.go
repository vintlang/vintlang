// Experimental module for scheduling tasks using ticker and cron-like expressions
// This module provides functions to create periodic tasks and one-time scheduled tasks
// using a simple cron-like syntax. It is not fully implemented yet, but provides a structure
// for future development.
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

// TickerControl represents a ticker instance that can be controlled
type TickerControl struct {
	Ticker   *time.Ticker
	StopChan chan bool
	ID       string
}

func init() {
	ScheduleFunctions["ticker"] = tickerFunc
	ScheduleFunctions["stopTicker"] = stopTickerFunc
	ScheduleFunctions["schedule"] = scheduleFunc
	ScheduleFunctions["stopSchedule"] = stopScheduleFunc
	ScheduleFunctions["everySecond"] = everySecondFunc
	ScheduleFunctions["everyMinute"] = everyMinuteFunc
	ScheduleFunctions["everyHour"] = everyHourFunc
	ScheduleFunctions["daily"] = dailyFunc
}

// ticker(intervalSeconds, callback)
func tickerFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"schedule", "ticker",
			"2 arguments: intervalSeconds (integer) and callback (function)",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.ticker(5, func() { print(\"tick\") })",
		)
	}

	// Validate interval argument
	interval, ok := args[0].(*object.Integer)
	if !ok {
		return ErrorMessage(
			"schedule", "ticker",
			"intervalSeconds to be an integer",
			fmt.Sprintf("'%s' (%T)", args[0].Inspect(), args[0]),
			"schedule.ticker(5, func() { print(\"tick\") })",
		)
	}

	// Validate callback argument
	callback, ok := args[1].(*object.Function)
	if !ok {
		return ErrorMessage(
			"schedule", "ticker",
			"callback to be a function",
			fmt.Sprintf("'%s' (%T)", args[1].Inspect(), args[1]),
			"schedule.ticker(5, func() { print(\"tick\") })",
		)
	}

	if interval.Value <= 0 {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> schedule.ticker()\033[0m:\n"+
				"  Interval must be positive, got %d seconds.\n"+
				"  Usage: schedule.ticker(5, func() { print(\"tick\") })\n",
				interval.Value),
		}
	}

	// Create ticker
	ticker := time.NewTicker(time.Duration(interval.Value) * time.Second)
	stopChan := make(chan bool)

	// Store ticker for management
	tickerStoreMu.Lock()
	tickerStore[ticker] = struct{}{}
	tickerStoreMu.Unlock()

	// Start ticker goroutine
	go func() {
		defer func() {
			ticker.Stop()
			tickerStoreMu.Lock()
			delete(tickerStore, ticker)
			tickerStoreMu.Unlock()
		}()

		for {
			select {
			case <-ticker.C:
				// For now, we'll return a success message to demonstrate the ticker is working
				// A full implementation would need access to the evaluator to execute the callback
				fmt.Printf("[schedule.ticker] Tick at %s (callback: %s)\n",
					time.Now().Format("15:04:05"), callback.Inspect())
			case <-stopChan:
				return
			}
		}
	}()

	// Create and return a control object
	control := &TickerControl{
		Ticker:   ticker,
		StopChan: stopChan,
		ID:       fmt.Sprintf("ticker_%p", ticker),
	}

	return &object.NativeObject{Value: control}
}

// stopTicker(tickerObj)
func stopTickerFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"schedule", "stopTicker",
			"1 argument: ticker object",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.stopTicker(ticker)",
		)
	}

	native, ok := args[0].(*object.NativeObject)
	if !ok {
		return ErrorMessage(
			"schedule", "stopTicker",
			"ticker object",
			fmt.Sprintf("'%s' (%T)", args[0].Inspect(), args[0]),
			"schedule.stopTicker(ticker)",
		)
	}

	control, ok := native.Value.(*TickerControl)
	if !ok {
		return ErrorMessage(
			"schedule", "stopTicker",
			"valid ticker object",
			"invalid ticker object",
			"schedule.stopTicker(ticker)",
		)
	}

	// Stop the ticker
	select {
	case control.StopChan <- true:
	default:
		// Channel might be closed already
	}

	control.Ticker.Stop()
	tickerStoreMu.Lock()
	delete(tickerStore, control.Ticker)
	tickerStoreMu.Unlock()

	return &object.Boolean{Value: true}
}

// ScheduleControl represents a schedule instance that can be controlled
type ScheduleControl struct {
	Timer    *time.Timer
	StopChan chan bool
	ID       string
}

// schedule(cronExpr, callback) - cronExpr: "second minute hour day month weekday" (basic, not full cron)
func scheduleFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"schedule", "schedule",
			"2 arguments: cronExpr (string) and callback (function)",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.schedule(\"0 30 14 * * *\", func() { print(\"Good afternoon!\") })",
		)
	}

	// Validate cron expression argument
	expr, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"schedule", "schedule",
			"cronExpr to be a string",
			fmt.Sprintf("'%s' (%T)", args[0].Inspect(), args[0]),
			"schedule.schedule(\"0 30 14 * * *\", func() { print(\"Good afternoon!\") })",
		)
	}

	// Validate callback argument
	callback, ok := args[1].(*object.Function)
	if !ok {
		return ErrorMessage(
			"schedule", "schedule",
			"callback to be a function",
			fmt.Sprintf("'%s' (%T)", args[1].Inspect(), args[1]),
			"schedule.schedule(\"0 30 14 * * *\", func() { print(\"Good afternoon!\") })",
		)
	}

	// Parse the cron expression to get the next execution time
	nextTime := nextSchedule(expr.Value)
	if nextTime.IsZero() {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> schedule.schedule()\033[0m:\n"+
				"  Invalid cron expression: '%s'\n"+
				"  Format: \"second minute hour day month weekday\"\n"+
				"  Example: \"0 30 14 * * *\" (daily at 14:30:00)\n"+
				"  Use '*' for wildcards\n",
				expr.Value),
		}
	}

	// Calculate duration until next execution
	duration := time.Until(nextTime)

	// Create timer
	timer := time.NewTimer(duration)
	stopChan := make(chan bool)

	// Store timer for management
	scheduleStoreMu.Lock()
	scheduleStore[timer] = struct{}{}
	scheduleStoreMu.Unlock()

	// Start schedule goroutine
	go func() {
		defer func() {
			timer.Stop()
			scheduleStoreMu.Lock()
			delete(scheduleStore, timer)
			scheduleStoreMu.Unlock()
		}()

		select {
		case <-timer.C:
			// For now, we'll log the execution to demonstrate the schedule is working
			fmt.Printf("[schedule.schedule] Executing at %s (callback: %s)\n",
				time.Now().Format("15:04:05"), callback.Inspect())
		case <-stopChan:
			return
		}
	}()

	// Create and return a control object
	control := &ScheduleControl{
		Timer:    timer,
		StopChan: stopChan,
		ID:       fmt.Sprintf("schedule_%p", timer),
	}

	return &object.NativeObject{Value: control}
}

// stopSchedule(scheduleObj)
func stopScheduleFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"schedule", "stopSchedule",
			"1 argument: schedule object",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.stopSchedule(scheduleObj)",
		)
	}

	native, ok := args[0].(*object.NativeObject)
	if !ok {
		return ErrorMessage(
			"schedule", "stopSchedule",
			"schedule object",
			fmt.Sprintf("'%s' (%T)", args[0].Inspect(), args[0]),
			"schedule.stopSchedule(scheduleObj)",
		)
	}

	control, ok := native.Value.(*ScheduleControl)
	if !ok {
		return ErrorMessage(
			"schedule", "stopSchedule",
			"valid schedule object",
			"invalid schedule object",
			"schedule.stopSchedule(scheduleObj)",
		)
	}

	// Stop the schedule
	select {
	case control.StopChan <- true:
	default:
		// Channel might be closed already
	}

	stopped := control.Timer.Stop()
	scheduleStoreMu.Lock()
	delete(scheduleStore, control.Timer)
	scheduleStoreMu.Unlock()

	return &object.Boolean{Value: stopped}
}

// nextSchedule parses a basic cron string and returns the next time.Time
func nextSchedule(expr string) time.Time {
	// Supports: "second minute hour day month weekday" (wildcards allowed)
	// Example: "0 30 14 * * *" = 14:30:00 every day
	// Example: "*/10 * * * * *" = every 10 seconds (basic step values)
	// Weekday: 0=Sunday, 1=Monday, ..., 6=Saturday

	fields := [6]int{-1, -1, -1, -1, -1, -1}
	stepValues := [6]int{1, 1, 1, 1, 1, 1} // Step values for */n patterns
	parts := splitAndTrim(expr)
	if len(parts) != 6 {
		return time.Time{}
	}

	for i, p := range parts {
		if p == "*" {
			fields[i] = -1
		} else if strings.HasPrefix(p, "*/") {
			// Handle step values like */5, */10, etc.
			stepStr := strings.TrimPrefix(p, "*/")
			step, err := parseInt(stepStr)
			if err != nil || step <= 0 {
				return time.Time{}
			}
			fields[i] = -1 // Wildcard with step
			stepValues[i] = step
		} else {
			v, err := parseInt(p)
			if err != nil {
				return time.Time{}
			}
			fields[i] = v
		}
	}

	now := time.Now()

	// Look for the next valid time within the next year
	for days := 0; days < 366; days++ {
		cand := now.Add(time.Duration(days) * 24 * time.Hour)

		// Check if day matches (if specified)
		if fields[3] != -1 && cand.Day() != fields[3] {
			continue
		}

		// Check if month matches (if specified)
		if fields[4] != -1 && int(cand.Month()) != fields[4] {
			continue
		}

		// Check if weekday matches (if specified)
		if fields[5] != -1 && int(cand.Weekday()) != fields[5] {
			continue
		}

		// For the current day, start from current time; for future days, start from 00:00:00
		startHour := 0
		if days == 0 {
			startHour = now.Hour()
		}

		for h := startHour; h < 24; h++ {
			// Check if hour matches (if specified) or step value
			if fields[2] != -1 && h != fields[2] {
				continue
			}
			if fields[2] == -1 && stepValues[2] > 1 && h%stepValues[2] != 0 {
				continue
			}

			startMinute := 0
			if days == 0 && h == now.Hour() {
				startMinute = now.Minute()
			}

			for m := startMinute; m < 60; m++ {
				// Check if minute matches (if specified) or step value
				if fields[1] != -1 && m != fields[1] {
					continue
				}
				if fields[1] == -1 && stepValues[1] > 1 && m%stepValues[1] != 0 {
					continue
				}

				startSecond := 0
				if days == 0 && h == now.Hour() && m == now.Minute() {
					startSecond = now.Second() + 1 // Next second
				}

				for s := startSecond; s < 60; s++ {
					// Check if second matches (if specified) or step value
					if fields[0] != -1 && s != fields[0] {
						continue
					}
					if fields[0] == -1 && stepValues[0] > 1 && s%stepValues[0] != 0 {
						continue
					}

					t := time.Date(cand.Year(), cand.Month(), cand.Day(), h, m, s, 0, cand.Location())
					if t.After(now) {
						return t
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

// Helper functions for common scheduling patterns

// everySecond(callback) - executes callback every second
func everySecondFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"schedule", "everySecond",
			"1 argument: callback (function)",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.everySecond(func() { print(\"tick\") })",
		)
	}

	// Call ticker with 1 second interval
	tickerArgs := []object.VintObject{&object.Integer{Value: 1}, args[0]}
	return tickerFunc(tickerArgs, defs)
}

// everyMinute(callback) - executes callback every minute at second 0
func everyMinuteFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"schedule", "everyMinute",
			"1 argument: callback (function)",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.everyMinute(func() { print(\"Every minute!\") })",
		)
	}

	// Call schedule with "0 * * * * *" (every minute at second 0)
	scheduleArgs := []object.VintObject{&object.String{Value: "0 * * * * *"}, args[0]}
	return scheduleFunc(scheduleArgs, defs)
}

// everyHour(callback) - executes callback every hour at minute 0, second 0
func everyHourFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"schedule", "everyHour",
			"1 argument: callback (function)",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.everyHour(func() { print(\"Every hour!\") })",
		)
	}

	// Call schedule with "0 0 * * * *" (every hour at minute 0, second 0)
	scheduleArgs := []object.VintObject{&object.String{Value: "0 0 * * * *"}, args[0]}
	return scheduleFunc(scheduleArgs, defs)
}

// daily(hour, minute, callback) - executes callback daily at specified time
func dailyFunc(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 {
		return ErrorMessage(
			"schedule", "daily",
			"3 arguments: hour (integer), minute (integer), and callback (function)",
			fmt.Sprintf("%d arguments", len(args)),
			"schedule.daily(9, 30, func() { print(\"Good morning!\") })",
		)
	}

	// Validate hour argument
	hour, ok := args[0].(*object.Integer)
	if !ok {
		return ErrorMessage(
			"schedule", "daily",
			"hour to be an integer",
			fmt.Sprintf("'%s' (%T)", args[0].Inspect(), args[0]),
			"schedule.daily(9, 30, func() { print(\"Good morning!\") })",
		)
	}

	// Validate minute argument
	minute, ok := args[1].(*object.Integer)
	if !ok {
		return ErrorMessage(
			"schedule", "daily",
			"minute to be an integer",
			fmt.Sprintf("'%s' (%T)", args[1].Inspect(), args[1]),
			"schedule.daily(9, 30, func() { print(\"Good morning!\") })",
		)
	}

	// Validate time ranges
	if hour.Value < 0 || hour.Value > 23 {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> schedule.daily()\033[0m:\n"+
				"  Hour must be between 0 and 23, got %d.\n"+
				"  Usage: schedule.daily(9, 30, func() { print(\"Good morning!\") })\n",
				hour.Value),
		}
	}

	if minute.Value < 0 || minute.Value > 59 {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> schedule.daily()\033[0m:\n"+
				"  Minute must be between 0 and 59, got %d.\n"+
				"  Usage: schedule.daily(9, 30, func() { print(\"Good morning!\") })\n",
				minute.Value),
		}
	}

	// Create cron expression: "0 minute hour * * *"
	cronExpr := fmt.Sprintf("0 %d %d * * *", minute.Value, hour.Value)
	scheduleArgs := []object.VintObject{&object.String{Value: cronExpr}, args[2]}
	return scheduleFunc(scheduleArgs, defs)
}

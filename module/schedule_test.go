package module

import (
	"testing"

	"github.com/vintlang/vintlang/object"
)

func TestTickerValidation(t *testing.T) {
	// Test with no arguments
	result := tickerFunc([]object.Object{}, map[string]object.Object{})
	if err, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for no arguments, got %T", result)
	} else if err.Message == "" {
		t.Errorf("Expected error message for no arguments")
	}

	// Test with wrong argument types
	result = tickerFunc([]object.Object{
		&object.String{Value: "not_a_number"},
		&object.Function{},
	}, map[string]object.Object{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for invalid interval type, got %T", result)
	}

	// Test with negative interval
	result = tickerFunc([]object.Object{
		&object.Integer{Value: -1},
		&object.Function{},
	}, map[string]object.Object{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for negative interval, got %T", result)
	}
}

func TestTickerCreation(t *testing.T) {
	// Test valid ticker creation
	result := tickerFunc([]object.Object{
		&object.Integer{Value: 1},
		&object.Function{},
	}, map[string]object.Object{})

	if native, ok := result.(*object.NativeObject); !ok {
		t.Errorf("Expected NativeObject, got %T", result)
	} else {
		if control, ok := native.Value.(*TickerControl); !ok {
			t.Errorf("Expected TickerControl, got %T", native.Value)
		} else {
			// Stop the ticker to clean up
			select {
			case control.StopChan <- true:
			default:
			}
			control.Ticker.Stop()
		}
	}
}

func TestScheduleValidation(t *testing.T) {
	// Test with invalid cron expression
	result := scheduleFunc([]object.Object{
		&object.String{Value: "invalid cron"},
		&object.Function{},
	}, map[string]object.Object{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for invalid cron expression, got %T", result)
	}

	// Test with wrong argument types
	result = scheduleFunc([]object.Object{
		&object.Integer{Value: 123},
		&object.Function{},
	}, map[string]object.Object{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for invalid cron expression type, got %T", result)
	}
}

func TestScheduleCreation(t *testing.T) {
	// Test valid schedule creation
	result := scheduleFunc([]object.Object{
		&object.String{Value: "0 0 * * * *"}, // Every hour
		&object.Function{},
	}, map[string]object.Object{})

	if native, ok := result.(*object.NativeObject); !ok {
		t.Errorf("Expected NativeObject, got %T", result)
	} else {
		if control, ok := native.Value.(*ScheduleControl); !ok {
			t.Errorf("Expected ScheduleControl, got %T", native.Value)
		} else {
			// Stop the schedule to clean up
			select {
			case control.StopChan <- true:
			default:
			}
			control.Timer.Stop()
		}
	}
}

func TestCronParsing(t *testing.T) {
	// Test valid cron expressions
	testCases := []struct {
		expr     string
		shouldWork bool
	}{
		{"* * * * * *", true},     // Every second
		{"0 * * * * *", true},     // Every minute
		{"0 0 * * * *", true},     // Every hour
		{"0 30 14 * * *", true},   // Daily at 14:30
		{"0 0 0 1 * *", true},     // First day of month
		{"*/5 * * * * *", true},   // Every 5 seconds
		{"*/10 * * * * *", true},  // Every 10 seconds
		{"0 */15 * * * *", true},  // Every 15 minutes
		{"invalid", false},        // Invalid format
		{"", false},               // Empty string
		{"* * *", false},          // Too few fields
		{"*/0 * * * * *", false},  // Invalid step value
		{"*/-1 * * * * *", false}, // Negative step value
	}

	for _, tc := range testCases {
		result := nextSchedule(tc.expr)
		if tc.shouldWork && result.IsZero() {
			t.Errorf("Expected valid time for cron expression '%s', got zero time", tc.expr)
		}
		if !tc.shouldWork && !result.IsZero() {
			t.Errorf("Expected zero time for invalid cron expression '%s', got %v", tc.expr, result)
		}
	}
}

func TestHelperFunctions(t *testing.T) {
	// Test everySecond
	result := everySecondFunc([]object.Object{&object.Function{}}, map[string]object.Object{})
	if native, ok := result.(*object.NativeObject); !ok {
		t.Errorf("Expected NativeObject from everySecond, got %T", result)
	} else {
		if control, ok := native.Value.(*TickerControl); ok {
			control.Ticker.Stop()
		}
	}

	// Test everyMinute
	result = everyMinuteFunc([]object.Object{&object.Function{}}, map[string]object.Object{})
	if native, ok := result.(*object.NativeObject); !ok {
		t.Errorf("Expected NativeObject from everyMinute, got %T", result)
	} else {
		if control, ok := native.Value.(*ScheduleControl); ok {
			control.Timer.Stop()
		}
	}

	// Test daily with valid time
	result = dailyFunc([]object.Object{
		&object.Integer{Value: 9},  // 9 AM
		&object.Integer{Value: 30}, // 30 minutes
		&object.Function{},
	}, map[string]object.Object{})
	if native, ok := result.(*object.NativeObject); !ok {
		t.Errorf("Expected NativeObject from daily, got %T", result)
	} else {
		if control, ok := native.Value.(*ScheduleControl); ok {
			control.Timer.Stop()
		}
	}

	// Test daily with invalid hour
	result = dailyFunc([]object.Object{
		&object.Integer{Value: 25}, // Invalid hour
		&object.Integer{Value: 30},
		&object.Function{},
	}, map[string]object.Object{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for invalid hour, got %T", result)
	}

	// Test daily with invalid minute
	result = dailyFunc([]object.Object{
		&object.Integer{Value: 9},
		&object.Integer{Value: 70}, // Invalid minute
		&object.Function{},
	}, map[string]object.Object{})
	if _, ok := result.(*object.Error); !ok {
		t.Errorf("Expected error for invalid minute, got %T", result)
	}
}

func TestStopFunctions(t *testing.T) {
	// Create a ticker and test stopping it
	tickerResult := tickerFunc([]object.Object{
		&object.Integer{Value: 1},
		&object.Function{},
	}, map[string]object.Object{})

	if native, ok := tickerResult.(*object.NativeObject); ok {
		stopResult := stopTickerFunc([]object.Object{native}, map[string]object.Object{})
		if boolean, ok := stopResult.(*object.Boolean); !ok {
			t.Errorf("Expected Boolean from stopTicker, got %T", stopResult)
		} else if !boolean.Value {
			t.Errorf("Expected true from stopTicker, got false")
		}
	}

	// Create a schedule and test stopping it
	scheduleResult := scheduleFunc([]object.Object{
		&object.String{Value: "0 0 * * * *"},
		&object.Function{},
	}, map[string]object.Object{})

	if native, ok := scheduleResult.(*object.NativeObject); ok {
		stopResult := stopScheduleFunc([]object.Object{native}, map[string]object.Object{})
		if boolean, ok := stopResult.(*object.Boolean); !ok {
			t.Errorf("Expected Boolean from stopSchedule, got %T", stopResult)
		} else if !boolean.Value {
			t.Errorf("Expected true from stopSchedule, got false")
		}
	}
}
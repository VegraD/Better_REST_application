package utils

import (
	"time"
)

// StartTime is the time when the server starts
var StartTime = time.Now()

/*
Uptime returns the uptime of the server
*/
func Uptime() int {
	// Get the time since the server started
	runtime := int(time.Since(StartTime).Seconds())

	// Convert the time to a string with 2 decimals
	//uptime := strconv.FormatFloat(runtime, 'f', 2, 64)

	// Convert the time to hours, minutes and seconds. E.g., 1h 2m 3s and round to 2 decimals
	//uptime := time.Duration(runtime * float64(time.Second)).Round(time.Second).String()

	return runtime
}

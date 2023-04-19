package res

import (
	"strconv"
	"time"
)

// StartTime is the time when the server starts
var StartTime = time.Now()

/*
Uptime returns the uptime of the server
*/
func Uptime() string {
	// Get the time since the server started
	runtime := time.Since(StartTime).Seconds()

	// Convert the time to a string with 2 decimals
	uptime := strconv.FormatFloat(runtime, 'f', 2, 64)

	return uptime
}

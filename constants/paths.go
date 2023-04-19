/*
Package constants contains all constants used in the application. This includes paths, ports, error messages, etc.
*/
package constants

// Ports
const (
	// DefaultPort is the default port used by the application if the environment variable PORT is not set.
	DefaultPort string = "8080"
)

// TODO: ?maybe? add comments to explain what each endpoint is used for as it's done for the default port
// Endpoint paths
const (
	DefaultEndpoint       string = "/"
	CurrentEndpoint       string = renewablesEndpoint + "/current"
	HistoryEndpoint       string = renewablesEndpoint + "/history"
	NotificationsEndpoint string = baseEndpoint + "/notifications"
	StatusEndpoint        string = baseEndpoint + "/status"
)

// ################################# Unexported constants below this line #################################

// Base endpoint paths
const (
	baseEndpoint       string = "/energy/v1"
	renewablesEndpoint string = baseEndpoint + "/renewables"
)

/*
Package constants contains all constants used in the application. This includes paths, ports, error messages, etc.
*/
package constants

// Ports
const (
	// DefaultPort is the default port used by the application if the environment variable PORT is not set.
	DefaultPort string = "8080"
)

// Endpoint paths
const (
	DefaultEndpoint       string = "/"
	CurrentEndpoint       string = renewablesEndpoint + "/current"
	HistoryEndpoint       string = renewablesEndpoint + "/history"
	NotificationsEndpoint string = baseEndpoint + "/notifications"
	StatusEndpoint        string = baseEndpoint + "/status"
)

// External API paths
const (
	CountryApi    = "http://129.241.150.113:8080/v3.1"
	RenewablesApi = "https://drive.google.com/file/d/18G470pU2NRniDfAYJ27XgHyrWOThP__p/view"
)

// ################################# Unexported constants below this line #################################

// Base endpoint paths
const (
	baseEndpoint       string = "/energy/v1"
	renewablesEndpoint string = baseEndpoint + "/renewables"
)

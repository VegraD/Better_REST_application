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
	// DefaultEP is the default endpoint path.
	DefaultEP string = "/"

	// CurrentEP is the endpoint path for retrieving current renewable energy production.
	CurrentEP = baseEP + "/renewables/current/"

	// HistoryEP is the endpoint path for retrieving historical renewable energy production.
	HistoryEP = baseEP + "/renewables/history/"

	// NotificationsEP is the endpoint path for retrieving notifications.
	NotificationsEP = baseEP + "/notifications/"

	// StatusEP is the endpoint path for retrieving the status of the application.
	StatusEP = baseEP + "/status/"
)

// External API paths
const (
	// CountryApi is the endpoint path for retrieving country information.
	CountryApi = "http://129.241.150.113:8080/v3.1"

	// CountryFullText is the query parameter to limit country results to full text matches.
	CountryFullText = "?fullText=true"

	// CountryFullTextName is the query parameter to search for country based on full text name.
	CountryFullTextName = "/name/"

	// CountryAlpha it the subdirectory to search for country based on alpha code
	CountryAlpha = "/alpha/"

	// RenewablesApi is the endpoint path for retrieving renewable energy production.
	RenewablesApi = "https://drive.google.com/file/d/18G470pU2NRniDfAYJ27XgHyrWOThP__p/view"
)

// ################################# Unexported constants below this line #################################

// Base endpoint paths
const (
	// baseEP is the base path for all endpoints.
	baseEP string = "/energy/v1"
)

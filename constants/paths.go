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

	// StaticEP is the endpoint path for static files.
	StaticEP string = "/static/"

	// CurrentEP is the endpoint path for retrieving current renewable energy production.
	CurrentEP = BasePath + "/renewables/current/"

	// HistoryEP is the endpoint path for retrieving historical renewable energy production.
	HistoryEP = BasePath + "/renewables/history/"

	// NotificationsEP is the endpoint path for retrieving notifications.
	NotificationsEP = BasePath + "/notifications/"

	// StatusEP is the endpoint path for retrieving the status of the application.
	StatusEP = BasePath + "/status/"

	// ReadmeEP is the endpoint path for retrieving the readme file.
	ReadmeEP = BasePath + "/readme/"
)

// External API paths
const (
	// CountryApi is the endpoint path for retrieving country information.
	CountryApi = "http://129.241.150.113:8080/"

	// countyreApiVersion is the version of the country API.
	CountryApiVersion = "/v3.1/"

	// CountryAlpha it the subdirectory to search for country based on alpha code
	CountryAlpha = "/alpha/"

	// RenewablesApi is the endpoint path for retrieving renewable energy production.
	RenewablesApi = "https://drive.google.com/file/d/18G470pU2NRniDfAYJ27XgHyrWOThP__p/view"
)

// ################################# Unexported constants below this line #################################

// Base endpoint paths
const (
	// BasePath is the base path for all endpoints.
	BasePath string = "/energy/v1"
)

// ServiceAccountLocation Firestore service account location
const (
	ServiceAccountLocation string = "./assignment-2-key.json"
)

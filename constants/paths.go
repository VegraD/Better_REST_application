/*
Package constants contains all constants used in the application. This includes paths, ports, error messages, etc.
*/
package constants

// TODO: ?maybe? add comments to explain what each endpoint is used for as it's done for the default port
// TODO: Sånn her du tenker? -Nicolai
// Endpoint paths
const (
	// DefaultEndpoint is the default endpoint for the application.
	DefaultEndpoint string = "/"

	// currentEndpoint is the endpoint for the current renewable energy production for the given country
	CurrentEndpoint = DEFAULT_PATH + RENEWABLES_PATH + CURRENT_PATH

	// HistoryEndpoint is the endpoint for the history of renewable energy production for the given country
	HistoryEndpoint = DEFAULT_PATH + RENEWABLES_PATH + HISTORY_PATH

	// NotificationsEndpoint is the endpoint for the notifications for the given country
	NotificationsEndpoint = DEFAULT_PATH + NOTIFICATION_PATH

	// StatusEndpoint is the endpoint for the status of the api
	StatusEndpoint = DEFAULT_PATH + STATUS_PATH
)

/*
Package constants contains all constants used in the application. This includes paths, ports, error messages, etc.
*/
package constants

// TODO: ?maybe? add comments to explain what each endpoint is used for as it's done for the default port
// Endpoint paths
const (
	DefaultEndpoint       string = "/"
	CurrentEndpoint              = DEFAULT_PATH + RENEWABLES_PATH + CURRENT_PATH
	HistoryEndpoint              = DEFAULT_PATH + RENEWABLES_PATH + HISTORY_PATH
	NotificationsEndpoint        = DEFAULT_PATH + NOTIFICATION_PATH
	StatusEndpoint               = DEFAULT_PATH + STATUS_PATH
)

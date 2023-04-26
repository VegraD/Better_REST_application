package constants

const Version string = "v1"
const NullString string = "null"

// Constants for choosing the correct data.
const (
	Current string = "current"
	History string = "history"

	CurrentYear string = "-1"
)

// Common error messages
const (
	CheckURLErr    string = ": Check that you entered the correct URL!"
	MarshallingErr string = "Error during marshalling: "
	PrettyPrintErr string = "Error during pretty printing: "
	WebhookNotFound string = "webhook was not found"
)

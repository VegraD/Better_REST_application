package constants

const Version string = "v1"

// MarkdownConverter related constants
const (
	ReadmeHtml       string = "<head><meta charset=\"UTF-8\"><link rel=\"stylesheet\" type=\"text/css\" href=\"%s\"></head><body>%s</body>"
	MdConvertPostReq string = `{"text": "Status"}`
)

// Constants for URL parameters
const (
	NullString        string = "null"
	CountryString     string = "country"
	BeginString       string = "begin"
	EndString         string = "end"
	TrueString        string = "true"
	SortByValueString string = "sortByValue"
	NeighboursString  string = "neighbours"
)

// Constants for choosing the correct data.
const (
	Current string = "current"
	History string = "history"

	CurrentYear string = "-1"
)

// Common error messages
const (
	CheckURLErr     string = ": Check that you entered the correct URL!"
	MarshallingErr  string = "Error during marshalling: "
	PrettyPrintErr  string = "Error during pretty printing: "
	WebhookNotFound string = "webhook was not found"
	EmptyDatabase   string = "database is empty"
)

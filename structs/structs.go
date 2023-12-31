package structs

// TODO: We should consider separating the structs into different files if we get many structs.

type Status struct {
	CountriesApi    string `json:"countries_api"`
	MarkdownHtmlApi string `json:"markdown_html_api"`
	NotificationDB  string `json:"notification_db"`
	Webhooks        int    `json:"webhooks"`
	Version         string `json:"version"`
	Uptime          int    `json:"uptime"`
}

type URLParams struct {
	Country     string
	BeginYear   string
	EndYear     string
	SortByValue bool
	Neighbours  bool
	EndPoint    string
}

// ################################################## Country structs ##################################################

type CountryInfo struct {
	Country    string  `json:"name"`
	IsoCode    string  `json:"isoCode"`
	Year       int     `json:"year,omitempty"`
	Percentage float32 `json:"percentage"`
}

type Border struct {
	Borders []string `json:"borders"`
}

// ################################################## Webhook structs ##################################################

type RegisteredWebhook struct {
	WebHookID string `json:"webhook_id"`
	URL       string `json:"url"`
	Country   string `json:"country"`
	Calls     int    `json:"calls"`
	Count     int    `json:"count"`
}

type DisplayWebhook struct {
	WebHookID string `json:"webhook_id"`
	URL       string `json:"url"`
	Country   string `json:"country"`
	Calls     int    `json:"calls"`
}

type WebHookRequest struct {
	URL     string `json:"url"`
	Country string `json:"country"`
	Calls   int    `json:"calls"`
}

type WebHookIDResponse struct {
	WebhookID string `json:"webhook_id"`
}

type WebHookInvocationResponse struct {
	WebhookID string `json:"webhook_id"`
	Country   string `json:"country"`
	Calls     int    `json:"calls"`
}

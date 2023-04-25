package structs

// TODO: We should consider separating the structs into different files if we get many structs.

type CountryInfo struct {
	Country    string  `json:"name"`
	IsoCode    string  `json:"isoCode"`
	Year       int     `json:"year,omitempty"`
	Percentage float32 `json:"percentage"`
}

type Status struct {
	CountriesApi   string `json:"countries_api"`
	RenewablesApi  string `json:"renewable_api"`
	NotificationDB string `json:"notification_db"`
	Webhooks       string `json:"webhooks"`
	Version        string `json:"version"`
	Uptime         string `json:"uptime"`
}

type Renewables struct {
	Entity     string
	Code       string
	Year       int
	Renewables float64
}

type URLParams struct {
	Country     string
	BeginYear   string
	EndYear     string
	SortByValue bool
	Neighbours  bool
	EndPoint    string
}

type Border struct {
	Cca3    string   `json:"cca3"`
	Borders []string `json:"borders"`
}

// ################################################## Webhook structs ##################################################

type RegisteredWebHook struct {
	WebHookID string `json:"webhook_id"`
	Url       string `json:"url"`
	Country   string `json:"country"`
	CallS     int    `json:"calls"`
	Count     int    `json:"count"`
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

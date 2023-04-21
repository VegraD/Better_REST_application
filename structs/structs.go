package structs

// TODO: We should consider separating the structs into different files if we get many structs.

type CountryInfo struct {
	Country    string  `json:"name"`
	Cca3       string  `gorm:"isoCode" json:"cca3"`
	Year       int     `json:"year"`
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

// ################################################## Webhook structs ##################################################

type RegisteredWebHook struct {
	WebHookID string `json:"webhook_id"`
	Url       string `json:"url"`
	Country   string `json:"country"`
	CallS     int    `json:"calls"`
}

type WebHookRequest struct {
	URL     string `json:"url"`
	Country string `json:"country"`
	Calls   int    `json:"calls"`
}

type WebHookIDResponse struct {
	WebhookID string `json:"webhook_id"`
}

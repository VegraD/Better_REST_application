package structs

type CountryInfo struct {
	Country    string  `json:"name"`
	IsoCode    string  `json:"isoCode"`
	Year       int     `json:"year"`
	Percentage float32 `json:"percentage"`
}

type WebHook struct {
	WebHookID string `json:"webhook_id"`
	Url       string `json:"url"`
	Country   string `json:"country"`
	CallS     int    `json:"calls"`
}

type Status struct {
	CountriesApi   string `json:"countries_api"`
	RenewablesApi  string `json:"renewable_api"`
	NotificationDB string `json:"notification_db"`
	Webhooks       string `json:"webhooks"`
	Version        string `json:"version"`
	Uptime         string `json:"uptime"`
}

package statusHandler

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"bytes"
	"io"
	"log"
	"net/http"
)

// StatusHandler is for redirecting the http request. it currently only supports get requests.
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		status := Status()
		json_coder.PrettyPrint(w, status)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
	}
}

// getApiStatus is a method to check the status for the apis that is used
func getApiStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error in get request to %s: %s", url, err)
	}
	return resp.Status
}

// getWebhookAmount is a method to get the amount of webhooks in the database.
func getWebhookAmount() int {
	webhookAmount, err := database.GetWebhookAmount()
	if err != nil {
		log.Printf("Error in getting webhook amount: %s", err)
	}
	return webhookAmount
}

// getMarkdownApiStatus The markdown api does not accept any get requests, so this function will send a post request with a markdown
// string to the api, and return the status of the response.
func getMarkdownApiStatus(url string) string {

	payload := []byte(constants.MdConvertPostReq)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Error in post request to %s: %s", url, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error in closing body: %s", err)
		}
	}(resp.Body)

	return resp.Status
}

// handleStatusRequest is a method to check the status for the apis that is used
// in the program, and "pretty print" it to the user.
func handleStatusRequest() (string, string, string) {

	countryResp := getApiStatus(constants.CountryApi)

	markdownResp := getMarkdownApiStatus(constants.MarkdownToHTMLApi)

	//TODO: Add correct url
	notificationResp := getApiStatus(constants.FireStoreApi)

	return countryResp, markdownResp, notificationResp
}

// Status is a method to create a status struct
func Status() structs.Status {
	country, markdown, notification := handleStatusRequest()

	status := structs.Status{
		CountriesApi:    country,
		MarkdownHtmlApi: markdown,
		NotificationDB:  notification,
		Webhooks:        getWebhookAmount(),
		Version:         constants.Version,
		Uptime:          utils.Uptime(),
	}

	return status
}

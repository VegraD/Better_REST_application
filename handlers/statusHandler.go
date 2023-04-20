package handlers

import (
	"assignment-2/constants"
	"assignment-2/json_coder"
	"assignment-2/structs"
	"assignment-2/utils"
	"log"
	"net/http"
)

/*
DiagnosticHandler is for redirecting the http request. it curently only support get requests.
*/
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		handleStatusRequest(w, r)
	} else {
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

/*
handleStatusRequest is a method to check the status for the apis that is used
in the program, and "pretty print" it to the user.
*/
func handleStatusRequest(w http.ResponseWriter, r *http.Request) {

	countryResp, err := http.Get(constants.CountryApi)
	if err != nil {
		log.Fatal(err)
	}
	renewablesResp, err := http.Get(constants.RenewablesApi)
	if err != nil {
		log.Fatal(err)
	}
	//TODO: Add correct url
	notificationResp, err := http.Get("https://restcountries.com/")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Add correct url
	webhookResp, err := http.Get("https://restcountries.com/")
	if err != nil {
		log.Fatal(err)
	}

	json_coder.PrettyPrint(w, Status(countryResp, renewablesResp, notificationResp, webhookResp))
}

func Status(country *http.Response, renewables *http.Response, notification *http.Response, webhook *http.Response) structs.Status {
	return structs.Status{
		CountriesApi:   country.Status,
		RenewablesApi:  renewables.Status,
		NotificationDB: notification.Status,
		Webhooks:       webhook.Status,
		Version:        constants.Version,
		Uptime:         utils.Uptime(),
	}
}

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

func getApiStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Status
}

/*
handleStatusRequest is a method to check the status for the apis that is used
in the program, and "pretty print" it to the user.
*/
func handleStatusRequest(w http.ResponseWriter, r *http.Request) {

	countryResp := getApiStatus(constants.CountryApi)
	renewablesResp := getApiStatus(constants.RenewablesApi)

	//TODO: Add correct url
	notificationResp := getApiStatus("https://restcountries.com/")

	//TODO: Add correct url
	webhookResp := getApiStatus("https://restcountries.com/")

	json_coder.PrettyPrint(w, Status(countryResp, renewablesResp, notificationResp, webhookResp))
}

func Status(country string, renewables string, notification string, webhook string) structs.Status {
	return structs.Status{
		CountriesApi:   country,
		RenewablesApi:  renewables,
		NotificationDB: notification,
		Webhooks:       webhook,
		Version:        constants.Version,
		Uptime:         utils.Uptime(),
	}
}

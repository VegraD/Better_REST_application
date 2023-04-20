package handlers

import "net/http"

/*
RenewablesCurrentHandler is the handler to get current information about countries renewable energy
*/
func RenewablesCurrentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		handleRenewablesCurrentGetRequest(w, r)
	} else {
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

/*
handleRenewablesCurrentGetRequest handles the get request for the current renewable energy information
*/
func handleRenewablesCurrentGetRequest(w http.ResponseWriter, r *http.Request) {

}

package defaultHandler

import (
	"assignment-2/constants"
	"assignment-2/utils"
	"net/http"
)

// DefaultHandler handles requests to the default endpoint.
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	// Switch for the different methods
	switch r.Method {
	case http.MethodGet:
		getDefault(w, r)
	default:
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+"!", http.StatusMethodNotAllowed)
	}
}

func getDefault(w http.ResponseWriter, r *http.Request) {

	// switch for endpoints
	switch r.URL.Path {
	case constants.DefaultEP, constants.BasePath, constants.BasePath + constants.DefaultEP:
		// Custom func for displaying the HTML file in the browser that handles errors.
		utils.DisplayDefaultPage(w, constants.DefaultHtml)
	default:
		// http.Error(w, "Endpoint not found!", http.StatusNotFound)
	}
}

package handlers

import (
	"assignment-2/constants"
	"assignment-2/utils"
	"net/http"
)

// DefaultHandler handles requests to the default endpoint.
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: Temporary! Find a better way.
	if r.URL.Path != constants.DefaultEP {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Custom func for displaying the HTML file in the browser that handles errors.
	utils.DisplayHTML(w, constants.DefaultHtml)
}

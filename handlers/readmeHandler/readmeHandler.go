package readmeHandler

import (
	"assignment-2/constants"
	"assignment-2/utils"
	"net/http"
)

func ReadmeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getReadme(w, r)
	default:
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+"!", http.StatusMethodNotAllowed)
	}
}

func getReadme(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == constants.ReadmeEP {
		err := utils.DisplayReadme(w, constants.Readme)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

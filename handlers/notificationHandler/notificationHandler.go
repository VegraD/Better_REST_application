package notificationHandler

import (
	"assignment-2/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

var db = []structs.RegisteredWebHook{}

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleNotificationGetRequest(w, r)
		fmt.Println("Get request")
	case http.MethodPost:
		handleNotificationPostRequest(w, r)
		fmt.Println("Post request")
	case http.MethodDelete:
		handleNotificationDeleteRequest(w, r)
		fmt.Println("Delete request")
	default:
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+", "+http.MethodPost+" or "+http.MethodDelete+"!", http.StatusMethodNotAllowed)
	}
}

func handleNotificationGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(db)
	if err != nil {
		http.Error(w, "Error during database encoding", http.StatusInternalServerError)
	}

}

func handleNotificationPostRequest(w http.ResponseWriter, r *http.Request) {

}

func handleNotificationDeleteRequest(w http.ResponseWriter, r *http.Request) {

}
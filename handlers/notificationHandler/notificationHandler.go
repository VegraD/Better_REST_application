package notificationHandler

import (
	"assignment-2/structs"
	"encoding/json"
	"fmt"
	"math/rand"
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
		return
	}

	// No content if no action is taken above this point.
	http.Error(w, "", http.StatusNoContent)

}

func handleNotificationPostRequest(w http.ResponseWriter, r *http.Request) {

}

func handleNotificationDeleteRequest(w http.ResponseWriter, r *http.Request) {

}

// TODO: error handling
func requestToResponse(request structs.WebHookRequest, id string) (structs.RegisteredWebHook, error) {

	return structs.RegisteredWebHook{
		WebHookID: fmt.Sprintf(id),
		Url:       fmt.Sprintf(request.URL),
		Country:   fmt.Sprintf(request.Country),
		CallS:     request.Calls,
	}, nil
}

func getRandomId() string {

	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randID := ""

	//TODO: change to check on firebase db
	for i, v := range db {
		l := make([]byte, 13)
		for j := range l {
			l[j] = letters[rand.Intn(len(letters))]
		}
		randID = string(l)

		//this should run for loop on same element again
		if v.WebHookID == randID {
			i--
			continue
		}
	}
	return randID
}

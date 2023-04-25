package notificationHandler

import (
	"assignment-2/structs"
	"assignment-2/webhooks"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	webhooks.WebhookHandler(w, r, db)
	w.Header().Add("content-type", "application/json")
	keyword := ""

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) >= 3 {
		keyword = parts[4]
	}

	//TODO: Check for blank spaces! unicode.IsSpace (need to check each byte)
	if len(keyword) == 0 || keyword == "" {
		err := json.NewEncoder(w).Encode(db)
		if err != nil {
			http.Error(w, "error during encoding", http.StatusInternalServerError)
			return
		}
		return
	}

	//TODO: Implement check in firebase
	for _, v := range db {
		if keyword == v.WebHookID {
			err := json.NewEncoder(w).Encode(v)
			if err != nil {
				http.Error(w, "error during database encoding", http.StatusInternalServerError)
				return
			}
		}
	}

	// No content if no action is taken above this point.
	http.Error(w, "", http.StatusNoContent)

}

func handleNotificationPostRequest(w http.ResponseWriter, r *http.Request) {
	//webhooks.WebhookHandler(w, r, db)
	webhook := structs.WebHookRequest{}

	w.Header().Add("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&webhook)

	if err != nil {
		http.Error(w, "cannot decode request", http.StatusBadRequest)
		return
	}

	webhookR, err := requestToRegistered(webhook, validateAndSetID())

	if err != nil {
		http.Error(w, "error during JSON request translation", http.StatusInternalServerError)
		return
	}

	db = append(db, webhookR)

	//TODO: do this smoother, what if encoder fails??
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(structs.WebHookIDResponse{WebhookID: webhookR.WebHookID})

	if err != nil {
		http.Error(w, "error during response decoding", http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}

func handleNotificationDeleteRequest(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
}

// TODO: error handling if input is empty
func requestToRegistered(request structs.WebHookRequest, id string) (structs.RegisteredWebHook, error) {

	return structs.RegisteredWebHook{
		WebHookID: fmt.Sprintf(id),
		Url:       fmt.Sprintf(request.URL),
		Country:   fmt.Sprintf(request.Country),
		CallS:     request.Calls,
	}, nil
}

func validateAndSetID() string {

	randID := ""

	if len(db) == 0 {
		return idGen()
	}
	//TODO: change to check in firebase db
	for i, v := range db {
		//generate random ID
		randID = idGen()

		//this should run for loop on same element again
		if v.WebHookID == randID {
			i--
			continue
		}
	}
	return randID
}

func idGen() string {

	//possible letters to have in ID
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Seed so the random integer is not the same each time.
	rand.Seed(time.Now().UnixNano())
	l := make([]byte, 13)
	for j := range l {
		l[j] = letters[rand.Intn(len(letters))]
	}
	return string(l)
}

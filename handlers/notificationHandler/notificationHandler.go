package notificationHandler

import (
	"assignment-2/database"
	"assignment-2/structs"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var Db = []structs.RegisteredWebHook{}

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleNotificationGetRequest(w, r)

	case http.MethodPost:
		handleNotificationPostRequest(w, r)

	case http.MethodDelete:
		handleNotificationDeleteRequest(w, r)

	default:
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+", "+http.MethodPost+" or "+http.MethodDelete+"!", http.StatusMethodNotAllowed)
	}
}

func handleNotificationGetRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	keyword := ""

	// Split URL paths into parts
	parts := strings.Split(r.URL.Path, "/")

	// Check if there are more than 3 parts. If not, dont give keyword a value
	if len(parts) >= 3 {
		keyword = parts[4]
	}

	//TODO: Check for blank spaces! unicode.IsSpace (need to check each byte)

	// Check if keyword is empty, if so; return all webhooks
	if len(keyword) == 0 || keyword == "" {
		// Display all elements in database
		err := json.NewEncoder(w).Encode(Db)
		if err != nil {
			http.Error(w, "error during encoding", http.StatusInternalServerError)
			return
		}
		return
	}

	//TODO: Implement check in firebase

	// Only relevant if keyword is set; checks if one of the elements in database has the relevant
	for _, v := range Db {
		if keyword == v.WebHookID {
			webhook, err := database.GetAndDisplayWebhook(v.WebHookID)
			if err != nil {
				http.Error(w, "error fetching webhook", http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(w).Encode(webhook)
			if err != nil {
				http.Error(w, "error during database encoding", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	// No content if no action is taken above this point.
	http.Error(w, "", http.StatusNoContent)

}

func handleNotificationPostRequest(w http.ResponseWriter, r *http.Request) {

	// Allocate empty struct
	webhook := structs.WebHookRequest{}

	w.Header().Add("content-type", "application/json")

	// Decode POST request
	err := json.NewDecoder(r.Body).Decode(&webhook)

	if err != nil {
		http.Error(w, "cannot decode request", http.StatusBadRequest)
		return
	}

	// Check if id is a valid cca3 code
	if len(webhook.Country) != 3 {
		http.Error(w, "invalid cca3 code!", http.StatusBadRequest)
		return
	}

	// Change struct into RegisteredWebHook and set ID
	webhookR, err := requestToRegistered(webhook, validateAndSetID())

	if err != nil {
		http.Error(w, "error during JSON request translation", http.StatusInternalServerError)
		return
	}

	/*
		id, err := firestore.WebhookAddition(webhookR.Url, webhookR.Country, webhookR.CallS)

		if err != nil {
			http.Error(w, "couldnt add webhook to server", http.StatusInternalServerError)
		}

	*/

	// Append webhook to database
	Db = append(Db, webhookR)

	//TODO: do this smoother, what if encoder fails??

	// Set header to display "201 - created"
	w.WriteHeader(http.StatusCreated)

	// encode response into JSON format
	err = json.NewEncoder(w).Encode(structs.WebHookIDResponse{WebhookID: webhookR.WebHookID /*id*/})

	if err != nil {
		http.Error(w, "error during response decoding", http.StatusInternalServerError)
		return
	}

	// No content if no action is taken above this point.
	http.Error(w, "", http.StatusNoContent)
}

func handleNotificationDeleteRequest(w http.ResponseWriter, r *http.Request) {
	//TODO: implement

	w.Header().Add("content-type", "application/json")
	keyword := ""

	// split URL into parts
	parts := strings.Split(r.URL.Path, "/")

	// Check the amount of parts in the URL, if less or equal 3, dont set keyword
	if len(parts) >= 3 {
		keyword = parts[4]
	}

	//TODO: Check for blank spaces! unicode.IsSpace (need to check each byte)
	if len(keyword) == 0 || keyword == "" {
		http.Error(w, "please enter a valid ID to delete!", http.StatusNoContent)
		return
	}

	//TODO: Implement check in firebase

	// Check for ID to delete in database (append deletion if found)
	for i, v := range Db {
		if keyword == v.WebHookID {
			Db = append(Db[:i], Db[i+1:]...)
			http.Error(w, "webhook successfully deleted", http.StatusOK)
			return
		}
	}

	// No content if no action is taken above this point.
	http.Error(w, "no valid webhook found", http.StatusNotModified)
}

// TODO: error handling if input is empty
func requestToRegistered(request structs.WebHookRequest, id string) (structs.RegisteredWebHook, error) {

	// Return new struct, set count to 0 as it is yet to be called
	return structs.RegisteredWebHook{
		WebHookID: fmt.Sprintf(id),
		Url:       fmt.Sprintf(request.URL),
		Country:   fmt.Sprintf(request.Country),
		CallS:     request.Calls,
		Count:     0,
	}, nil
}

func validateAndSetID() string {

	randID := ""

	// Dont validate if database so far is empty; no reason to check against other webhooks if no other webhooks exist.
	if len(Db) == 0 {
		return idGen()
	}

	//TODO: change to check in firebase Db
	for i, v := range Db {
		// Generate random ID
		randID = idGen()

		// This will run for loop on same element again
		if v.WebHookID == randID {
			i--
			continue
		}
	}
	return randID
}

func idGen() string {

	// Possible letters to have in ID
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Seed so the random integer is not the same each time.
	rand.Seed(time.Now().UnixNano())

	// Allocate empty byte array with 13 bytes
	id := make([]byte, 13)
	for j := range id {
		id[j] = letters[rand.Intn(len(letters))]
	}
	return string(id)
}

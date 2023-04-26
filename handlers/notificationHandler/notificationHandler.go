package notificationHandler

import (
	"assignment-2/database"
	"assignment-2/structs"
	"encoding/json"
	"net/http"
	"strings"
)

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

	// Get webhooks from firestore
	db, err := database.GetAllWebhooks()

	if err != nil {
		http.Error(w, "database is empty", http.StatusNoContent)
	}

	// Split URL paths into parts
	parts := strings.Split(r.URL.Path, "/")

	// Check if there are more than 3 parts. If not, dont give keyword a value
	if len(parts) >= 3 {
		keyword = parts[4]
	}

	//TODO: Check for blank spaces! unicode.IsSpace (need to check each byte)

	// Check if keyword is empty, if so; return all webhooks
	if len(keyword) == 0 || keyword == "" {
		// encode and display database
		err = json.NewEncoder(w).Encode(db)
		if err != nil {
			http.Error(w, "error during encoding", http.StatusInternalServerError)
			return
		}
		return
	}

	//TODO: Implement check in firebase

	// Only relevant if keyword is set; checks if one of the elements in database has the relevant
	for _, v := range db {
		if keyword == v.WebHookID {
			// Get webhook from database
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
	var webhook structs.WebHookRequest

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

	// Add webhook to firestore database
	id, err := database.AddWebhook(webhook.URL, webhook.Country, webhook.Calls)

	// throw error if webhook cannot be added to database
	if err != nil {
		http.Error(w, "couldnt add webhook to server", http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	// Set header to display "201 - created"
	w.WriteHeader(http.StatusCreated)

	//TODO: do this smoother, what if encoder fails?? ^

	// encode response into JSON format
	err = json.NewEncoder(w).Encode(structs.WebHookIDResponse{WebhookID: id})

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
	// Get webhooks from firestore
	db, err := database.GetAllWebhooks()

	if err != nil {
		http.Error(w, "database is empty", http.StatusNoContent)
	}

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
	//TODO: how to delete from firebase? IMPLEMENT
	for _, v := range db {
		if keyword == v.WebHookID {
			err = database.DeletionOfWebhook(keyword)
			if err != nil {
				http.Error(w, "deletion of webhook failed", http.StatusInternalServerError)
			}
			http.Error(w, "webhook successfully deleted", http.StatusOK)
			return
		}
	}

	// No content if no action is taken above this point.
	http.Error(w, "no valid webhook found", http.StatusNotModified)
}

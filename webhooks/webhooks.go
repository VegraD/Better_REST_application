package webhooks

import (
	"assignment-2/handlers/notificationHandler"
	"assignment-2/structs"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

/*
func WebhookHandler(w http.ResponseWriter, r *http.Request, db []structs.RegisteredWebHook) {
	str, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request body", http.StatusInternalServerError)
	}
	switch r.Method {
	case http.MethodGet:
		log.Println("Received " + http.MethodGet + " request...")
		// Iterate through registered webhooks and invoke based on registered URL, method, and with received content
		for _, v := range db {
			log.Println("Trigger event: Call to service endpoint with method " + http.MethodGet +
				" and content '" + string(str) + "'.")

			go callURL(v.Url, http.MethodGet, string(str))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	case http.MethodPost:
		log.Println("Received " + http.MethodPost + " request...")
		// Iterate through registered webhooks and invoke based on registered URL, method, and with received content
		for _, v := range db {
			log.Println("Trigger event: Call to service endpoint with method " + http.MethodPost +
				" and content '" + string(str) + "'.")
			go callURL(v.Url, http.MethodPost, string(str))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	case http.MethodDelete:
		log.Println("Received " + http.MethodDelete + " request...")
		// Iterate through registered webhooks and invoke based on registered URL, method, and with received content
		for _, v := range db {
			log.Println("Trigger event: Call to service endpoint with method " + http.MethodDelete +
				" and content '" + string(str) + "'.")
			go callURL(v.Url, http.MethodDelete, string(str))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	default:
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+", "+http.MethodPost+" or "+http.MethodDelete+"!", http.StatusMethodNotAllowed)
	}

}

*/

// TODO: implement with persistent storage
func InvokeWebhook(w http.ResponseWriter, country string) {
	webhooks := notificationHandler.Db

	for _, v := range webhooks {
		if v.Country == country {
			v.Count++
			if v.CallS == v.Count {
				v.Count = 0

				go callURL(http.MethodPost, v)
			}
		}
	}
}

// TODO: add functionality for incrementing calls
func callURL(method string, webhook structs.RegisteredWebHook) error {

	response := registeredToResponse(webhook)

	body, err := json.Marshal(response)

	if err != nil {
		return errors.New("unable to marshal content")
	}

	req, err := http.NewRequest(method, webhook.Url, bytes.NewReader(body))
	if err != nil {
		return errors.New("unable to create request")
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return errors.New("unable to send request")
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("error fetching response body")
	}

	log.Println("Webhook " + webhook.Url + " invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(resp))

	return nil
}

/*
// TODO: add functionality for incrementing calls
func callURL(url string, method string, content string) error {
	log.Println("Attempting invocation of url " + url + " with content '" + content + "'.")

	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(content)))
	if err != nil {
		return errors.New("unable to create request")
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return errors.New("unable to send request")
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("error fetching response body")
	}

	log.Println("Webhook " + url + " invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(resp))

	return nil
}
*/

func registeredToResponse(registered structs.RegisteredWebHook) structs.WebHookInvocationResponse {

	return structs.WebHookInvocationResponse{
		WebhookID: fmt.Sprintf(registered.WebHookID),
		Country:   fmt.Sprintf(registered.Country),
		Calls:     registered.CallS,
	}
}

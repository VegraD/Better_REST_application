package webhooks

import (
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
// TODO: implement with persistent storage
func InvokeWebhook(w http.ResponseWriter, country string) {
	webhooks := notificationHandler.Db

	for _, v := range webhooks {
		if v.Country == country {
			v.Count = v.Count + 1
			if v.CallS == v.Count {
				v.Count = 0

				go callURL(http.MethodPost, v)
			}
		}
	}
}

*/

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

func registeredToResponse(registered structs.RegisteredWebHook) structs.WebHookInvocationResponse {

	return structs.WebHookInvocationResponse{
		WebhookID: fmt.Sprintf(registered.WebHookID),
		Country:   fmt.Sprintf(registered.Country),
		Calls:     registered.CallS,
	}
}

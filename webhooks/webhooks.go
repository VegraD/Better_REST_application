package webhooks

import (
	"assignment-2/database"
	"assignment-2/structs"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TODO: implement with persistent storage
func InvokeWebhook(country string) error {
	webhooks, err := database.GetAllWebhooks()

	if err != nil {
		return errors.New("Webhooks is empty")
	}
	if country == "" {
		for _, v := range webhooks {
			v.Count = v.Count + 1
			if v.Calls <= v.Count {
				v.Count = 0

				go callURL(http.MethodPost, v)
			}
			_, err = database.UpdateWebhooks(v.URL, v.Country, v.Calls, v.Count)
		}
	}
	for _, v := range webhooks {
		if strings.EqualFold(v.Country, country) {
			v.Count = v.Count + 1
			if v.Calls <= v.Count {
				v.Count = 0

				go callURL(http.MethodPost, v)
			}
			_, err = database.UpdateWebhooks(v.URL, v.Country, v.Calls, v.Count)
		}
	}
	return nil
}

// TODO: add functionality for incrementing calls
func callURL(method string, webhook structs.RegisteredWebhook) error {

	response := registeredToResponse(webhook)

	body, err := json.Marshal(response)

	if err != nil {
		return errors.New("unable to marshal content")
	}

	req, err := http.NewRequest(method, webhook.URL, bytes.NewReader(body))
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

	log.Println("Webhook " + webhook.URL + " invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(resp))

	return nil
}

func registeredToResponse(registered structs.RegisteredWebhook) structs.WebHookInvocationResponse {

	return structs.WebHookInvocationResponse{
		WebhookID: fmt.Sprintf(registered.WebHookID),
		Country:   fmt.Sprintf(registered.Country),
		Calls:     registered.Calls,
	}
}

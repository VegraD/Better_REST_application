package firestore

import (
	"assignment-2/structs"
	"assignment-2/utils"
	"errors"
)

// Collection name in Firestore
const collection = "webhooks"

var ct = 0

/*
Reads a string from the body in plain-text and sends it to Firestore to be registered as a document
*/
func WebhookAddition(url string, country string, no_calls int) (string, error) {

	webhookId := utils.HashingTheWebhook(url, country, no_calls)

	response := client.Collection(collection).Doc(webhookId)
	_, err := response.Get(ctx)

	if err != nil {
		return "", errors.New("webhooks already exist")
	}
	_, err = response.Set(ctx, map[string]interface{}{
		"webhookId": webhookId,
		"url":       url,
		"country":   country,
		"calls":     no_calls,
		"count":     0,
	})
	if err != nil {
		return "", err
	} else {
		return webhookId, nil
	}
}
func GetAndDisplayWebhook(webhookID string) (structs.RegisteredWebHook, error) {
	getResponse := client.Collection(collection).Doc(webhookID)
	doc, err := getResponse.Get(ctx)
	if err != nil {
		return structs.RegisteredWebHook{}, errors.New("webhook not found")
	}

	var webhookToBeDisplayed structs.RegisteredWebHook
	err = doc.DataTo(&webhookToBeDisplayed)
	if err != nil {
		return structs.RegisteredWebHook{}, err
	}
	return structs.RegisteredWebHook{}, nil
}

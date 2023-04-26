package firestore

import (
	"assignment-2/constants"
	"assignment-2/structs"
	"assignment-2/utils/hashing-utility"
	"errors"
)

// Collection name in Firestore
const collection = "webhooks"

var ct = 0

/*
Reads a string from the body in plain-text and sends it to Firestore to be registered as a document
*/
func addWebhook(url string, country string, noCalls int) (string, error) {

	webhookId := hashing_utility.HashingTheWebhook(url, country, noCalls)

	response := client.Collection(collection).Doc(webhookId)
	_, err := response.Get(ctx)

	if err != nil {
		return "", errors.New("webhooks already exist")
	}
	_, err = response.Set(ctx, map[string]interface{}{
		"webhookId": webhookId,
		"url":       url,
		"country":   country,
		"calls":     noCalls,
		"count":     0,
	})
	if err != nil {
		return "", err
	} else {
		return webhookId, nil
	}
}
func getAndDisplayWebhook(webhookID string) (structs.RegisteredWebHook, error) {
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

func deletionOfWebhook(webhookID string) error {
	getResponse := client.Collection(collection).Doc(webhookID)
	_, err := getResponse.Get(ctx)
	if err != nil {
		return errors.New(constants.WebhookNotFound)
	}

	_, err = getResponse.Delete(ctx)
	if err != nil {
		return err
	} else {
		return nil

	}
}

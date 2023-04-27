package database

import (
	"assignment-2/constants"
	"assignment-2/structs"
	"assignment-2/utils/hashing-utility"
	"errors"
	"google.golang.org/api/iterator"
	"log"
)

// Collection name in Firestore
var collection = "webhooks"

/*
Function that adds webhook to firestore-database.
*/
func AddWebhook(url string, country string, noCalls int) (string, error) {

	//Hashes webhook to create a unique ID
	webhookId := hashing_utility.HashingTheWebhook(url, country, noCalls)

	response := client.Collection(collection).Doc(webhookId)
	_, err := response.Get(ctx)

	//Adds values to the webhook.
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

/*
Function to get and display webhook. Uses the webhookID as paramter, and returns a webhook-struct and/or error.
*/
func GetAndDisplayWebhook(webhookID string) (structs.RegisteredWebhook, error) {
	getResponse := client.Collection(collection).Doc(webhookID)
	doc, err := getResponse.Get(ctx)
	if err != nil {
		return structs.RegisteredWebhook{}, errors.New(constants.WebhookNotFound)
	}

	var webhookToBeDisplayed structs.RegisteredWebhook
	err = doc.DataTo(&webhookToBeDisplayed)
	if err != nil {
		return structs.RegisteredWebhook{}, err
	}
	return structs.RegisteredWebhook{}, nil
}

/*
Function that deletes a webhook from the database. Takes in the webhookID to be deleted.
*/
func DeletionOfWebhook(webhookID string) error {
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

func GetAllWebhooks() ([]structs.RegisteredWebhook, error) {

	var webhooks []structs.RegisteredWebhook

	collection := GetClient().Collection(collection).Documents(GetContext())

	for {

		wh, err := collection.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		var webhookToAdd structs.RegisteredWebhook
		err = wh.DataTo(&webhookToAdd)

		if err != nil {
			return nil, err
		}

		webhooks = append(webhooks, webhookToAdd)

	}

	if len(webhooks) == 0 {
		return []structs.RegisteredWebhook{}, errors.New(constants.EmptyDatabase)
	}

	return webhooks, nil
}

func UpdateWebhooks(url string, country string, noCalls int, count int) (string, error) {

	webhookId := hashing_utility.HashingTheWebhook(url, country, noCalls)

	response := client.Collection(collection).Doc(webhookId)
	_, err := response.Get(ctx)

	_, err = response.Set(ctx, map[string]interface{}{
		"webhookId": webhookId,
		"url":       url,
		"country":   country,
		"calls":     noCalls,
		"count":     count,
	})
	if err != nil {
		return "", err
	} else {
		return webhookId, nil
	}
}

/*
A method for getting the total number of webhooks in the database
*/
func GetWebhookAmount() int {

	// Get all webhooks from database
	webhooks, err := GetClient().Collection(collection).Documents(GetContext()).GetAll()

	if err != nil {
		log.Print("No documents found in the specified collection" + err.Error())
		return 0
	}
	if len(webhooks) == 0 {
		log.Print("No documents found in the specified collection")
		return 0
	}

	return len(webhooks)
}

/*
Function for clearing the database of potential webhooks
*/
func ClearDB() {
	db, _ := GetAllWebhooks()

	for _, v := range db {
		_ = DeletionOfWebhook(v.WebHookID)
	}
}

/*
A function for setting up a test database
*/
func TestDatabaseSetup() []string {
	collection = "testing"
	ClearDB()

	id1, _ := AddWebhook("https://testurl.com", "Norway", 2)
	id2, _ := AddWebhook("https://testurl2.com", "Finland", 2)
	id3, _ := AddWebhook("https://testurl3.com", "Sweden", 2)

	return []string{id1, id2, id3}
}

package database

import (
	"assignment-2/constants"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

var ctx context.Context
var client *firestore.Client

func InitFirestore() {
	ctx = context.Background()

	sa := option.WithCredentialsFile(constants.ServiceAccountLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func CloseDB() error {
	err := client.Close()
	if err != nil {
		return err
	}
	return nil
}

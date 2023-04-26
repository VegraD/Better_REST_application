package firestore

import (
	"assignment-2/constants"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var ctx context.Context
var client *firestore.Client

func InitFirestore() error {
	ctx = context.Background()

	sa := option.WithCredentialsFile(constants.ServiceAccountLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		//Bedre ERRORHENALDING HER
		return err
	}

	_, err = app.Firestore(ctx)
	if err != nil {
		return err
	}
	return nil
}

func closeDB() error {
	err := client.Close()
	if err != nil {
		return err
	}
	return nil
}

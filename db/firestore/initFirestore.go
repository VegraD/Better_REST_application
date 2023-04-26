package firestore

import (
	"assignment-2/constants"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
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

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	//Close down client
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error: ", err)
		}
	}()
	return nil
}

/*
func handleMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		(w, r)
	case http.MethodGet:
		displayDocuments(w, r)
	default:
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
		return
	}
}

*/

func initFirestoreDatabase() {
	// Make it Heroku-compatible
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	//http.HandleFunc("/messages", handleMessage) // Be forgiving in case people forget the trailing slash ('/')
	//http.HandleFunc("/messages/", handleMessage)
	log.Printf("Firestore REST service listening on %s ...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}

}

func GetClient() *firestore.Client {
	return client
}

func GetContext() context.Context {
	return ctx
}

package utils

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"io"
	"log"
	"net/http"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

// Collection name in Firestore
const collection = "messages"

var ct = 0

/*
Reads a string from the body in plain-text and sends it to Firestore to be registered as a document
*/
func addDocument(w http.ResponseWriter, r *http.Request) {
	text, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Reading payload from body failed")
		http.Error(w, "Reading payload failed", http.StatusInternalServerError)
		return
	}
	log.Println("Recieved request to add document for content ", string(text))
	if len(string(text)) == 0 {
		log.Println("Content appears to be empty")
		http.Error(w, "Your payload (to be stored as document) appears to be empty. Ensure to terminate URI with /.", http.StatusBadRequest)
	} else {
		// Add element in embedded structure.
		// Note: this structure is defined by the client, not the server!; it exemplifies the use of a complex structure
		// and illustrates how you can use Firestore features such as Firestore timestamps.
		id, _, err := client.Collection(collection).Add(ctx,
			map[string]interface{}{
				"text": string(text),
				"ct":   ct,
				"time": firestore.ServerTimestamp,
			})
		ct++
		if err != nil {
			// Error handling
			log.Println("Error when adding document " + string(text) + ", Error: " + err.Error())
			http.Error(w, "Error when adding document "+string(text)+", Error: "+err.Error(), http.StatusBadRequest)
			return
		} else {
			// Returns document ID in body
			log.Println("Document added to collection. Identifier of returned document: " + id.ID)
			http.Error(w, id.ID, http.StatusCreated)
			return
		}
	}
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		//addMessage(w,r)
	case http.MethodGet:
		//displayMessage(w,r)
	default:
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
		return
	}
}

func firebaseAndClientInit() {
	// Firebase initialisation
	ctx = context.Background()

	opt := option.WithCredentialsFile("./assignment-2-key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	//Instantiate client
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	//Close down client
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error: ", err)
		}
	}()

}

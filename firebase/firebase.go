package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile("firebase-adminsdk.json") // letakkan file ini di root project
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase App: %v", err)
	}
	App = app
}

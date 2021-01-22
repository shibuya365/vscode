package fs

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// CTX is ConTeXt
var CTX context.Context

// App is Firebase App
var App *firebase.App

// FirebaseInit initialize
func FirebaseInit() {
	var err error

	CTX = context.Background()
	sa := option.WithCredentialsFile("/Users/masashishibuya/firebase/vscode-72dc9-firebase-adminsdk-4bhgc-59cdd11e2e.json")
	App, err = firebase.NewApp(CTX, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
}

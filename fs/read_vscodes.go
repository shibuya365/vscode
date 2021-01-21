package fs

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// VSCode is ShortCut for VSCode
type VSCode struct {
	ID          string `firestore:"id"`
	Title       string `firestore:"title"`
	Shortcut    string `firestore:"shortcut"`
	Description string `firestore:"description"`
	Visiable    bool
	// `json:"visiable"`
}

// VSCodes is VSCode's shortcuts
var VSCodes []VSCode

// ReadVSCodes read VSCodes
func ReadVSCodes() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("/Users/masashishibuya/firebase/vscode-72dc9-firebase-adminsdk-4bhgc-59cdd11e2e.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// 全てのショートカットを取得
	iter := client.Collection("vscodes").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		var vscode VSCode
		doc.DataTo(&vscode)
		VSCodes = append(VSCodes, vscode)
	}
}

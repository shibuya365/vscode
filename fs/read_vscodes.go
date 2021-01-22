package fs

import (
	"log"

	"google.golang.org/api/iterator"
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

	// get client
	client, err := App.Firestore(CTX)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// 全てのショートカットを取得
	iter := client.Collection("vscodes").Documents(CTX)
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

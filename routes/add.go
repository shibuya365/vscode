package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shibuya365/VSCode.git/fs"
)

// Add delete inivisi from arary
func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("root: Add")

		// PathからIDを取得
		id := c.Param("id")
		fmt.Println("id: ", id)

		// get client
		client, err := fs.App.Firestore(fs.CTX)
		if err != nil {
			log.Fatalln(err)
		}
		defer client.Close()

		// クッキー読み込み
		cookie, err := c.Cookie("vscode_scs")

		// userの初期化
		var user User

		// Login?
		if err != nil {
			// ログインしてなかった場合
			fmt.Println("Not Login")
			return
		} else {
			// ログインしている場合
			fmt.Println("Login")
		}

		// strs := make([]string, 0)

		// userを取得
		dsnap, err := client.Collection("users").Doc(cookie).Get(fs.CTX)
		if err != nil {
			fmt.Println("Firebase have no data: ", err)
		}
		dsnap.DataTo(&user)

		// strsへ表示しないIDの配列を入れる
		strs := user.Invisis

		// delete id
		strs = remove(strs, id)

		// userへ戻す
		user.Invisis = strs

		// データベースへ書き込み
		_, err = client.Collection("users").Doc(cookie).Set(fs.CTX, user)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}

		// Redirect to /showall
		c.Redirect(302, "/showall")
	}
}

// スライスの中身削除
func remove(strings []string, search string) []string {
	for i, v := range strings {
		if v == search {
			return append(strings[:i], strings[i+1:]...)
		}
	}
	fmt.Println("not found: ", search)
	return strings
}

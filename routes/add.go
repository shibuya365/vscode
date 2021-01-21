package routes

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"google.golang.org/api/option"
)

// Add delete inivisi from arary
func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("root: Add")

		// 表示を戻すShortcutを取得
		id := c.Param("id")
		fmt.Println("id: ", id)

		// userの初期化
		var user User

		// get client
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

		// クッキー読み込み
		cookie, err := c.Cookie("vscode_scs")

		// Login?
		if err != nil {
			// ログインしてなかった場合
			fmt.Println("Not Login")
			// idを生成
			guid := xid.New()
			cookie = guid.String()
			c.SetCookie("vscode_scs", cookie, 60*60*24*31*12*2, "/", "localhost", false, true)

			// forestore追加
			_, err := client.Collection("users").Doc(cookie).Set(ctx, user)
			if err != nil {
				log.Printf("An error has occurred: %s", err)
			}
		} else {
			// ログインしている場合
			fmt.Println("Login")
		}

		// strs := make([]string, 0)
		// userを取得
		dsnap, err := client.Collection("users").Doc(cookie).Get(ctx)
		if err != nil {
			fmt.Println("Firebase have no data: ", err)
		}
		dsnap.DataTo(&user)

		strs := user.Invisis

		// delete id
		strs = remove(strs, id)

		// userへ戻す
		user.Invisis = strs

		// データベースへ書き込み
		_, err = client.Collection("users").Doc(cookie).Set(ctx, user)
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

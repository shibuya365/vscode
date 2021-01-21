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

// Delete add invisi to array
func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("R Delete")

		// 表示しないShortcutを取得
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

		// userを取得
		dsnap, err := client.Collection("users").Doc(cookie).Get(ctx)
		if err != nil {
			fmt.Println("Firebase have no data: ", err)
		}
		dsnap.DataTo(&user)

		strs := user.Invisis

		// IDと追加
		strs = append(strs, id)

		// userへ戻す
		user.Invisis = strs

		// データベースへ書き込み
		_, err = client.Collection("users").Doc(cookie).Set(ctx, user)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}

		// Redirect to root
		c.Redirect(302, "/")
	}
}

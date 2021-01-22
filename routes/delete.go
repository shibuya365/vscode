package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/shibuya365/VSCode.git/fs"
)

// Delete add invisi to array
func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Root: Delete")

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

		// ログインしてない場合に備えてからの配列を用意
		strs := make([]string, 0)

		// Login?
		if err != nil {
			// ログインしてなかった場合
			fmt.Println("Not Login")
			// idを生成
			guid := xid.New()
			cookie = guid.String()
			c.SetCookie("vscode_scs", cookie, 60*60*24*31*12*2, "/", "localhost", false, true)

			// forestore追加
			// _, err := client.Collection("users").Doc(cookie).Set(ctx, user)
			// if err != nil {
			// 	log.Printf("An error has occurred: %s", err)
			// }
		} else {
			// ログインしている場合
			fmt.Println("Login")

			// userを取得
			dsnap, err := client.Collection("users").Doc(cookie).Get(fs.CTX)
			if err != nil {
				fmt.Println("Firebase have no data: ", err)
			}
			dsnap.DataTo(&user)

			// strsへ表示しないIDの配列を入れる
			strs = user.Invisis
		}

		// IDと追加
		strs = append(strs, id)

		// userへ戻す
		user.Invisis = strs

		// データベースへ書き込み
		_, err = client.Collection("users").Doc(cookie).Set(fs.CTX, user)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}

		// Redirect to root
		c.Redirect(302, "/")
	}
}

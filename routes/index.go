package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/shibuya365/VSCode.git/fs"
	"google.golang.org/api/option"
)

type User struct {
	Invisis []string `firestore:"vscodes"`
}

// Index ルートを表示する
func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("root: /")

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
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
		} else {
			// ログインしている場合
			fmt.Println("Login")

			// VSCodesのVisiableの初期化
			for i := 0; i < len(fs.VSCodes); i++ {
				fs.VSCodes[i].Visiable = true
			}

			// userを取得
			dsnap, err := client.Collection("users").Doc(cookie).Get(ctx)
			if err != nil {
				fmt.Println("Firebase have no data: ", err)
			}
			dsnap.DataTo(&user)

			// fmt.Println("user: ", user)
			strs := user.Invisis
			// 表示しないものにfalseを代入
			for _, str := range strs {
				for j, vs := range fs.VSCodes {
					if vs.ID == str {
						fs.VSCodes[j].Visiable = false
					}
				}
			}
		}

		// fmt.Println("fs.VSCodes", fs.VSCodes)
		// クッキーをログに表示
		fmt.Printf("Cookie value: %s \n", cookie)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"vscs": fs.VSCodes,
		})
	}
}

// NoRoute 当てはまるものがない場合はルートへ
// func NoRoute(c *gin.Context) {
// 	c.Redirect(http.StatusMovedPermanently, "/")
// }

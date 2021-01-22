package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shibuya365/VSCode.git/fs"
)

type User struct {
	Invisis []string `firestore:"vscodes"`
}

// Index ルートを表示する
func Index(temp string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("root: /")

		// userの初期化
		var user User

		// get client
		client, err := fs.App.Firestore(fs.CTX)
		if err != nil {
			log.Fatalln(err)
		}
		defer client.Close()

		// クッキー読み込み
		cookie, err := c.Cookie("vscode_scs")

		// ログインしてない場合に備えてからの配列を用意
		strs := make([]string, 0)

		// Login?
		if err != nil {
			// ログインしてなかった場合
			fmt.Println("Not Login")
		} else {
			// ログインしている場合
			fmt.Println("Login")

			// クッキーをログに表示
			fmt.Printf("Cookie value: %s \n", cookie)

			// VSCodesのVisiableの初期化
			for i := 0; i < len(fs.VSCodes); i++ {
				fs.VSCodes[i].Visiable = true
			}

			// userを取得
			dsnap, err := client.Collection("users").Doc(cookie).Get(fs.CTX)
			if err != nil {
				fmt.Println("Firebase have no data: ", err)
			}
			dsnap.DataTo(&user)

			strs = user.Invisis

			// 表示しないものにfalseを代入
			for _, str := range strs {
				for j, vs := range fs.VSCodes {
					if vs.ID == str {
						fs.VSCodes[j].Visiable = false
					}
				}
			}
		}

		c.HTML(http.StatusOK, temp, gin.H{
			"vscs": fs.VSCodes,
		})
	}
}

// NoRoute 当てはまるものがない場合はルートへ
// func NoRoute(c *gin.Context) {
// 	c.Redirect(http.StatusMovedPermanently, "/")
// }

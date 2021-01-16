package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/shibuya365/VSCode.git/dbscs"
	"github.com/shibuya365/VSCode.git/dbus"
)

// ShowAll 全てを表示
func ShowAll(scs dbscs.Shortcuts) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("root: /showall")
		// User 読み込み
		users := dbus.ReadUsersDB()

		// クッキー読み込み
		cookie, err := c.Cookie("vscode_scs")

		if err != nil {
			// ログインしてなかった場合
			fmt.Println("Not Login")
			// idを生成
			guid := xid.New()
			cookie = guid.String()
			c.SetCookie("vscode_scs", cookie, 60*60*24*31*12*2, "/", "localhost", false, true)
		} else {
			// ログインしている場合
			fmt.Println("Login")

			// scsの初期化
			for i := 0; i < len(scs); i++ {
				scs[i].Visiable = true
			}

			// 表示しないものにfalseを代入
			strs := users[cookie]
			fmt.Println("/ users: ", users)
			for _, str := range strs {
				for j, sc := range scs {
					if sc.Shortcut == str {
						scs[j].Visiable = false
					}
				}
			}
		}

		c.HTML(http.StatusOK, "showall.html", gin.H{
			"scs": scs,
		})
	}

}

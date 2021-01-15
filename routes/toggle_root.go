package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/shibuya365/VSCode.git/dbscs"
	"github.com/shibuya365/VSCode.git/dbus"
)

// Toggle is toggle users flag
func Toggle(scs dbscs.Shortcuts) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("R Delete")
		// IDを取得
		id, _ := strconv.Atoi(c.Param("id"))
		fmt.Println("id: ", id)

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
		}
		// nums := users[cookie]
		// usersにIDを追加
		// users[cookie] = append(users[cookie], id)
		users[cookie][id] = !users[cookie][id]
		fmt.Println("delete users: ", users)
		dbus.WriteUsersDB(users)

		c.Redirect(302, "/")
	}
}

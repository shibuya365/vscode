package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/shibuya365/VSCode.git/dbscs"
	"github.com/shibuya365/VSCode.git/dbus"
)

// Toggle is toggle users flag
func Add(scs dbscs.Shortcuts) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("root: Add")
		// 表示を戻すShortcutを取得
		id := c.Param("id")
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
		// usersからSCを削除
		fmt.Println("add users: ", users)
		users[cookie] = remove(users[cookie], id)

		// 新しいusersデータ書き込み
		dbus.WriteUsersDB(users)

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

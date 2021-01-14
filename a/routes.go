package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/shibuya365/VSCode.git/dbscs"
	"github.com/shibuya365/VSCode.git/dbus"
)

// Index ルートを表示する
func Index(scs dbscs.Shortcuts) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("R /")
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
			nums := users[cookie]
			fmt.Println("index users: ", users)
			for _, num := range nums {
				scs[num].Visiable = false
			}
		}

		// クッキーをログに表示
		fmt.Printf("Cookie value: %s \n", cookie)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"scs": scs,
		})
	}
}

// 消す
func Delete(scs dbscs.Shortcuts) gin.HandlerFunc {
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
		users[cookie] = append(users[cookie], id)
		fmt.Println("delete users: ", users)
		dbus.WriteUsersDB(users)

		c.Redirect(302, "/")
	}
}

// 表示させる
func Add(scs dbscs.Shortcuts) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("R Add")
		// IDを取得
		id, _ := strconv.Atoi(c.Param("id"))
		fmt.Println("add id: ", id)

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
		users[cookie] = append(users[cookie], id)
		fmt.Println("add users: ", users)
		dbus.WriteUsersDB(users)

		c.Redirect(302, "/showall")
	}
}

// 全てを表示
func ShowAll(scs dbscs.Shortcuts) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("R ShowAll")
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
			nums := users[cookie]
			fmt.Println("showall users: ", users)
			for _, num := range nums {
				scs[num].Visiable = false
			}
		}

		c.HTML(http.StatusOK, "showall.html", gin.H{
			"scs": scs,
		})
	}

}

// NoRoute 当てはまるものがない場合はルートへ
// func NoRoute(c *gin.Context) {
// 	c.Redirect(http.StatusMovedPermanently, "/")
// }

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shibuya365/VSCode.git/dbscs"
	"github.com/shibuya365/VSCode.git/routes"
)

func main() {

	// ショートカットのファイルを読み込む
	scs := dbscs.ReadShortcutsDB()

	// ルーター
	r := gin.Default()

	// 事前にテンプレートをロード 相対パス
	// router.LoadHTMLGlob("templates/*/**") などもいけるらしい
	r.LoadHTMLGlob("templates/*.html")

	// 静的ファイルのパスを指定
	r.Static("/assets", "./assets")

	// ハンドラの指定
	r.GET("/", routes.Index(scs))
	r.GET("/toggle/:id", routes.Toggle(scs))
	r.GET("/toggleall/:id", routes.ToggleAll(scs))
	r.GET("/showall", routes.ShowAll(scs))

	// どのルーティングにも当てはまらなかった場合に処理
	// r.NoRoute(routes.NoRoute)

	r.Run(":3000")
}

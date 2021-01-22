package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shibuya365/VSCode.git/fs"
	"github.com/shibuya365/VSCode.git/routes"
)

// Shortcut はDB設定の構造体
type Shortcut struct {
	ID          string `firestore:"id"`
	Title       string `firestore:"title"`
	Shortcut    string `firestore:"shortcut"`
	Description string `firestore:"description"`
	Visiable    bool
}

func main() {

	// Firebaseを初期化
	fs.FirebaseInit()

	// Firestoreから読み込み
	fs.ReadVSCodes()

	// ルーター
	r := gin.Default()

	// 事前にテンプレートをロード 相対パス
	// router.LoadHTMLGlob("templates/*/**") などもいけるらしい
	r.LoadHTMLGlob("templates/*.html")

	// 静的ファイルのパスを指定
	r.Static("/assets", "./assets")

	// ハンドラの指定
	r.GET("/", routes.Index("index.html"))
	r.GET("/delete/:id", routes.Delete())
	r.GET("/add/:id", routes.Add())
	r.GET("/showall", routes.Index("showall.html"))

	// どのルーティングにも当てはまらなかった場合に処理
	// r.NoRoute(routes.NoRoute)

	r.Run(":3000")
}

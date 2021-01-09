package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shibuya365/VSCode.git/routes"
)

func main() {
	router := gin.Default()

	// 事前にテンプレートをロード 相対パス
	// router.LoadHTMLGlob("templates/*/**") などもいけるらしい
	router.LoadHTMLGlob("templates/*.html")

	// 静的ファイルのパスを指定
	router.Static("/assets", "./assets")

	// ハンドラの指定
	router.GET("/hello", routes.Hello)

	// グルーピング
	user := router.Group("/api")
	{
		user.GET("/hello", routes.HelloJson)
	}

	router.NoRoute(routes.NoRoute) // どのルーティングにも当てはまらなかった場合に処理
	router.Run(":8080")
}

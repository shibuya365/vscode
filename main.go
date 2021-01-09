package main

import (
	"fmt"

	"github.com/gin-gonic/conf"
	"github.com/gin-gonic/gin"
	"github.com/shibuya365/VSCode.git/routes"
)

func main() {
	// 設定ファイルを読み込む
	lines, err := conf.ReadConfDB()
	// もしファイルがなかったら
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	// ルーター
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

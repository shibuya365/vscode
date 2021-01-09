package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"name": "Taro",
	})
}

func NoRoute(c *gin.Context) {
	// helloに飛ばす（ログインしていない場合に、ログイン画面に飛ばす
	c.Redirect(http.StatusMovedPermanently, "/hello")
}

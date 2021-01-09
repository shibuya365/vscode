package routes

import (
	"github.com/gin-gonic/gin"
)

func HelloJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "taro",
	})
}

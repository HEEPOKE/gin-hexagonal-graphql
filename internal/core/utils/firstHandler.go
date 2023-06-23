package utils

import "github.com/gin-gonic/gin"

func HandleFirst(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello",
	})
}

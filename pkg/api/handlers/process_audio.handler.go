package handlers

import "github.com/gin-gonic/gin"

func ProcessAudioHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ProcessAudioHandler",
	})
}

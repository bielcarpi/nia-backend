package handlers

import (
	"github.com/gin-gonic/gin"
	"nia-backend/pkg/openai"
)

func ProcessAudioHandler(c *gin.Context) {
	client := openai.OpenAIClientProvider()
	// TODO handle 429: exceeded current plan limits

	text, err := client.SpeechToText(c, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	generatedText, err := client.SendToGPT3(c, text)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": generatedText,
	})
}

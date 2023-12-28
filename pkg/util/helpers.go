package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nia-backend/config"
	"nia-backend/pkg/api"
)

func main() {
	gin.SetMode(gin.DebugMode) //TODO Change to gin.ReleaseMode
	server := api.NewServer()

	// Start the server
	err := server.Start(fmt.Sprintf(":%s", config.GetConfig().Port))
	if err != nil {
		panic(err)
	}
}

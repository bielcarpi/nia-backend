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
	port := config.GetConfig().Port
	if port == "" {
		port = "8080"
	}
	err := server.Start(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

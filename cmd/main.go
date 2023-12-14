package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nia-backend/config"
	"nia-backend/pkg/api"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode) //TODO Change to gin.ReleaseMode
	server := api.NewServer()

	// Start the server
	err = server.Start(fmt.Sprintf(":%s", configs.Port))
	if err != nil {
		panic(err)
	}
}

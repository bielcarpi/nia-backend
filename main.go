package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nia-backend/api"
	"nia-backend/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode) //TODO Change to gin.ReleaseMode
	server := api.NewServer()

	// Start the server
	err = server.Start(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		panic(err)
	}
}

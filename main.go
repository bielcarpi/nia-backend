package main

import (
	"github.com/gin-gonic/gin"
	"nia-backend/api"
)

func main() {
	/*config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs", err)
	}
	*/

	gin.SetMode(gin.DebugMode) //TODO Change to gin.ReleaseMode
	server := api.NewServer(store)

	// Start the server
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}

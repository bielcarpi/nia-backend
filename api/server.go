package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	// Set up routing
	// We can add middleware to the routes here too
	//router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Start starts the HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse Returns an error response in JSON format to the client
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
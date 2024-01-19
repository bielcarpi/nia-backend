package api

import (
	"github.com/gin-gonic/gin"
	"nia-backend/pkg/api/handlers"
	"nia-backend/pkg/api/middleware"
)

type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	// Set up /api routing
	apiRouter := router.Group("/api")
	apiRouter.Use(middleware.AuthMiddleware())
	apiRouter.GET("/audio", handlers.ProcessAudioHandler)

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

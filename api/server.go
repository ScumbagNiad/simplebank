package api

import (
	"net/http"
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/v1/accounts", server.createAccount)
	router.GET("/v1/accounts/:id", server.getAccount)
	router.GET("/v1/accounts", server.listAccounts)
	router.NoRoute(gin.WrapH(http.FileServer(http.Dir("static"))))

	server.router = router
	return server
}

// Start server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

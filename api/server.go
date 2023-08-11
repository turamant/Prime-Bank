package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/turamant/simplebank/db/sqlc"
)


type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	server.router = router
	return server
}
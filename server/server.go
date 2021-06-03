package server

import (
	"go-ginapi-base/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func StartServer() Server {
	return Server{
		port:   ":8000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.Run(s.port)
}

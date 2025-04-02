package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"websocket-high-tps-chat/repository"
	"websocket-high-tps-chat/service"
)

type Server struct {
	engine *gin.Engine

	service    *service.Service
	repository *repository.Repository

	port string
	ip   string
}

func NewServer(service *service.Service, repository *repository.Repository, port string) *Server {
	s := &Server{
		engine:     gin.New(),
		service:    service,
		repository: repository,
		port:       port,
	}

	// api가 들어오는 것에 대한 로그
	s.engine.Use(gin.Logger())
	// panic에 의한 서버가 죽어버릴 경우 다시 서버를 재기동
	s.engine.Use(gin.Recovery())
	s.engine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	return s
}

func (s *Server) StartServer() error {
	log.Println("Starting Server")
	return s.engine.Run(s.port)
}

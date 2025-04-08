package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang-chat-controller/service"
	"log"
)

type Server struct {
	engine *gin.Engine

	service *service.Service

	port          string
	avgServerList map[string]bool
}

func NewNetwork(service *service.Service, port string) *Server {
	s := &Server{
		engine:        gin.New(),
		service:       service,
		port:          port,
		avgServerList: make(map[string]bool),
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

	s.setServerInfo()

	return s
}

func (s *Server) setServerInfo() {
	if serverList, err := s.service.GetAvailableServerList(); err != nil {
		panic(err)
	} else {
		for _, server := range serverList {
			s.avgServerList[server.IP] = true
		}
	}
}

func (s *Server) Start() error {
	log.Println("Starting Server")
	return s.engine.Run(s.port)
}

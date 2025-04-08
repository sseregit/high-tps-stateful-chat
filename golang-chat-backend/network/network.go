package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang-chat-backend/service"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	engine *gin.Engine

	service *service.Service

	port string
	ip   string
}

func NewNetwork(service *service.Service, port string) *Server {
	s := &Server{
		engine:  gin.New(),
		service: service,
		port:    port,
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

	registerServer(s)

	return s
}

func (s *Server) setServerInfo() {
	if addrs, err := net.InterfaceAddrs(); err != nil {
		panic(err.Error())
	} else {
		var ip net.IP

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					ip = ipnet.IP
					break
				}
			}
		}

		if ip == nil {
			panic("no ip address found")
		} else {
			if err = s.service.ServerSet(ip.String()+s.port, true); err != nil {
				panic(err)
			} else {
				s.ip = ip.String()
			}

			s.service.PublishServerStatusEvent(s.ip+s.port, true)
		}
	}

}

func (s *Server) StartServer() error {
	s.setServerInfo()

	// 서버가 죽는것을 캐치할 수 있는법
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)

	go func() {
		<-channel

		if err := s.service.ServerSet(s.ip+s.port, false); err != nil {
			log.Println("Failed To Set Server Info When Close", "err", err)
		}

		// Kafka에 이벤트 전송

		s.service.PublishServerStatusEvent(s.ip+s.port, false)

		os.Exit(1)
	}()

	log.Println("Starting Server")
	return s.engine.Run(s.port)
}

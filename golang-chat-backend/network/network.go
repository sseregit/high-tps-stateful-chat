package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Network struct {
	engin *gin.Engine
}

func NewServer() *Network {
	n := &Network{
		engin: gin.New(),
	}

	// api가 들어오는 것에 대한 로그
	n.engin.Use(gin.Logger())
	// panic에 의한 서버가 죽어버릴 경우 다시 서버를 재기동
	n.engin.Use(gin.Recovery())
	n.engin.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	r := NewRoom()
	go r.RunInit()

	n.engin.GET("/room", r.SocketServe)

	return n
}

func (n *Network) StartServer() error {
	log.Println("Starting Server")
	return n.engin.Run(":8080")
}

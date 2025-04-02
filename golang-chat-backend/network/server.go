package network

import "github.com/gin-gonic/gin"

type data struct {
}

func registerServer(engine *gin.Engine) *data {
	d := &data{}

	r := NewRoom()
	go r.RunInit()

	engine.GET("/room", r.SocketServe)

	return d
}

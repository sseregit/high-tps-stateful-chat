package network

import "github.com/gin-gonic/gin"

type api struct {
	server *Server
}

func registerServer(server *Server) {
	a := &api{server: server}

	server.engine.GET("/room-list", a.roomList)
	server.engine.POST("/make-room", a.makeRoom)
	server.engine.GET("/room", a.room)
	server.engine.GET("/enter-room", a.enterRoom)

	//r := NewRoom()
	//go r.RunInit()
	//
	//server.engine.GET("/room", r.SocketServe)
}

func (a *api) roomList(c *gin.Context) {

}
func (a *api) makeRoom(c *gin.Context) {

}
func (a *api) room(c *gin.Context) {

}
func (a *api) enterRoom(c *gin.Context) {

}

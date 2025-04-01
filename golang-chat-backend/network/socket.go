package network

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"websocket-high-tps-chat/types"
)

var upgrader = &websocket.Upgrader{ReadBufferSize: types.SocketBufferSize, WriteBufferSize: types.MessageBufferSize, CheckOrigin: func(r *http.Request) bool {
	return true
}}

type message struct {
	Name    string
	Message string
	Time    int64
}

type Room struct {
	Forward chan *message // 수신되는 메시지 보관, 들어오는 메시지를 다른 클라이언트들에게 전송
	Join    chan *client  // Socket이 연결되는 경우에 작동
	Leave   chan *client  // Socket이 끊어지는 경우에 대해서 작동

	clients map[*client]bool // 현재 방에 있는 client 정보를 저장
}

type client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *client),
		Leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (c *client) Read() {
	defer c.Socket.Close()
	for {
		var msg *message
		err := c.Socket.ReadJSON(&msg)
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				break
			} else {
				panic(err)
			}
		} else {
			log.Println("READ : ", msg, "client", c.Name)
			log.Println()
			msg.Time = time.Now().Unix()
			msg.Name = c.Name

			c.Room.Forward <- msg
		}
	}
}

func (c *client) Write() {
	defer c.Socket.Close()

	for msg := range c.Send {
		log.Println("WRITE : ", msg, "client", c.Name)
		log.Println()
		err := c.Socket.WriteJSON(msg)
		if err != nil {
			panic(err)
		}
	}
}

func (r *Room) RunInit() {
	for {
		select {
		case client := <-r.Join:
			r.clients[client] = true
		case client := <-r.Leave:
			r.clients[client] = false
			delete(r.clients, client)
			close(client.Send)
		case msg := <-r.Forward:
			for client := range r.clients {
				client.Send <- msg
			}
		}
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	userCookie, err := c.Request.Cookie("auth")
	if err != nil {
		panic(err)
	}

	client := &client{
		Socket: socket,
		Send:   make(chan *message, types.MessageBufferSize),
		Room:   r,
		Name:   userCookie.Value,
	}

	r.Join <- client

	defer func() {
		r.Leave <- client
	}()

	go client.Write()

	client.Read()
}

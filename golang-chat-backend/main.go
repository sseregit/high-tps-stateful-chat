package main

import (
	"flag"
	"websocket-high-tps-chat/config"
	"websocket-high-tps-chat/network"
	"websocket-high-tps-chat/repository"
	"websocket-high-tps-chat/service"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)

	if rep, err := repository.NewRepository(c); err != nil {
		panic(err)
	} else {
		s := network.NewServer(service.NewService(rep), rep, *port)
		s.StartServer()
	}
}

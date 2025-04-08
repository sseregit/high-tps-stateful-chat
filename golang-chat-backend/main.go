package main

import (
	"flag"
	"golang-chat-backend/config"
	"golang-chat-backend/network"
	"golang-chat-backend/repository"
	"golang-chat-backend/service"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)

	if rep, err := repository.NewRepository(c); err != nil {
		panic(err)
	} else {
		s := network.NewNetwork(service.NewService(rep), *port)
		s.StartServer()
	}
}

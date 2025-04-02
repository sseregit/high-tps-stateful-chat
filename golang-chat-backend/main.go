package main

import (
	"flag"
	"websocket-high-tps-chat/config"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)

	//n := network.NewServer()
	//n.StartServer()
}

package main

import (
	"flag"
	"fmt"
	"websocket-high-tps-chat/config"
	"websocket-high-tps-chat/repository"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)

	if rep, err := repository.NewRepository(c); err != nil {
		panic(err)
	} else {
		fmt.Println(*rep)
	}

	fmt.Println(c)
	//n := network.NewServer()
	//n.StartServer()
}

package main

import "websocket-high-tps-chat/network"

func main() {
	n := network.NewServer()
	n.StartServer()
}

package main

import (
	"flag"

	"github.com/kafy11/gosocket/message"
	"github.com/kafy11/gosocket/server"
)

func getPort() int {
	var port int
	flag.IntVar(&port, "p", 8080, "port to be used by the websocket")
	flag.Parse()
	return port
}

func main() {
	port := getPort()

	server.Start(message.Handler, port)
}

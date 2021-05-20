package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/kafy11/gosocket/log"
	"github.com/kafy11/gosocket/server/client"

	"github.com/gobwas/ws"
)

type MessageHandlerFunction func(self *client.Data)

var clients map[int]*client.Data

func Start(messageHandler MessageHandlerFunction, port int) {
	clients = make(map[int]*client.Data)

	websocketPort := fmt.Sprint(":", port)
	http.ListenAndServe(websocketPort, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := getId(r)
		if id == 0 {
			http.NotFound(w, r)
			return
		}

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Error(err)
			return
		}

		go handleConnection(id, conn, messageHandler)
	}))
}

func handleConnection(id int, conn net.Conn, messageHandler MessageHandlerFunction) {
	client := client.New(id, conn)
	clients[id] = client
	defer client.Close()

	for {
		if err := client.WaitMessage(); err != nil {
			log.Error(err)
			return
		}

		messageHandler(client)
	}
}

func getId(r *http.Request) (id int) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		return 0
	}
	return id
}

func SendMessage(clientId int, message interface{}) (bool, error) {
	client, ok := clients[clientId]
	if ok {
		if err := client.SendMessage(message); err != nil {
			return false, err
		}
	}
	return ok, nil
}

package message

import (
	"github.com/kafy11/gosocket/log"
	"github.com/kafy11/gosocket/server"
	"github.com/kafy11/gosocket/server/client"
)

type Received struct {
	Action string `json:"action"`
	To     int    `json:"to"`
	Msg    string `json:"msg"`
}

type ToSend struct {
	Action string `json:"action"`
	From   int    `json:"from"`
	Msg    string `json:"msg"`
}

func Handler(self *client.Data) {
	var message Received
	self.DecodeMessageReceived(&message)
	log.Info(message.Action)

	sent, err := server.SendMessage(message.To, &ToSend{
		Action: message.Action,
		From:   self.Id,
		Msg:    message.Msg,
	})

	if err != nil {
		log.Error(err)
	} else if !sent {
		self.SendMessage(&ToSend{
			Action: message.Action,
			From:   self.Id,
		})
	}
}

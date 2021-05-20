package message

import (
	"github.com/kafy11/gosocket/log"
	"github.com/kafy11/gosocket/server"
	"github.com/kafy11/gosocket/server/client"
)

type MessageReceived struct {
	Action string `json:"action"`
	To     int    `json:"to"`
	Msg    string `json:"msg"`
}

type MessageToSent struct {
	Action string `json:"action"`
	From   int    `json:"from"`
	Msg    string `json:"msg"`
}

func Handler(self *client.Data) {
	var message MessageReceived
	self.DecodeMessageReceived(&message)
	log.Info(message.Action)

	sent, err := server.SendMessage(message.To, &MessageToSent{
		Action: message.Action,
		From:   self.Id,
		Msg:    message.Msg,
	})

	if err != nil {
		log.Error(err)
	} else if !sent {
		self.SendMessage(&MessageToSent{
			Action: message.Action,
			From:   self.Id,
		})
	}
}

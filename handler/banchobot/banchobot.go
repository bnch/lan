package banchobot

import (
	"fmt"
	"time"

	"github.com/bnch/lan/handler"
	"github.com/bnch/lan/packets"
)

var self *handler.Session

func init() {
	self = &handler.Session{
		Username: "BanchoBot",
		UserID:   1,
		Token:    handler.GenerateGUID(),
		Admin:    true,
		// packets should be handled rather quickly, so we need a small buffer
		Packets:  make(chan packets.Packet, 50),
		LastSeen: time.Now().Add(time.Hour * 24 * 365 * 10),
		State: packets.OsuSendUserState{
			Action: 2,
			Text:   "with itself",
		},
	}
	handler.Sessions.Add(self)
	go packetHandler()
}

func packetHandler() {
	for packet := range self.Packets {
		switch p := packet.(type) {
		case *packets.BanchoSendMessage:
			// message received
			fmt.Printf("B!%s: %s\n", p.SenderName, p.Content)
			sess := handler.Sessions.GetByID(p.SenderID)
			handleMessage(sess, p.Content)
		}
	}
}

func handleMessage(user *handler.Session, msg string) {
	newMsg := &packets.BanchoSendMessage{
		SenderName: self.Username,
		SenderID:   self.UserID,
		Channel:    user.Username,
	}
	if msg == "!ping" {
		newMsg.Content = "Pong, motherfucker."
	}
	if newMsg.Content != "" {
		user.Send(newMsg)
	}
}

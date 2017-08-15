package banchobot

import (
	"fmt"
	"time"

	"github.com/bnch/lan/handler"
	"github.com/bnch/lan/packets"
)

var self handler.Session

// Start initialises banchobot.
func Start() {
	prevSession := handler.GetSession(handler.Sessions.TokenFromUsername("BanchoBot"))
	if prevSession != nil {
		self = *prevSession
		go packetHandler()
		return
	}
	self = handler.Session{
		Username: "BanchoBot",
		UserID:   1,
		Token:    handler.GenerateGUID(),
		Admin:    true,
		LastSeen: time.Now().Add(time.Hour * 24 * 365 * 10),
		State: packets.OsuSendUserState{
			Action: 2,
			Text:   "with itself",
		},
	}
	handler.Sessions.Add(self)
	handler.SaveSession(self)
	go packetHandler()
}

func packetHandler() {
	for {
		msg := handler.Redis.BLPop(time.Second*20, "lan/queues/"+self.Token).Val()
		if len(msg) < 2 {
			continue
		}

		pks, err := packets.Depacketify([]byte(msg[1]))
		if err != nil {
			fmt.Println("banchobot error", err)
			continue
		}

		for _, p := range pks {
			switch p := p.(type) {
			case *packets.BanchoSendMessage:
				// message received
				fmt.Printf("B!%s: %s\n", p.SenderName, p.Content)
				token := handler.Sessions.TokenFromID(p.SenderID)
				handleMessage(handler.GetSession(token), p.Content)
			}
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

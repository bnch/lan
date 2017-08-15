package banchobot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bnch/lan/handler"
	"github.com/bnch/lan/packets"
)

var self handler.Session

// Start initialises banchobot.
func Start() {
	prevSession := handler.GetSession(handler.Sessions.TokenFromUsername("BanchoBot"))
	guid := handler.GenerateGUID()
	if prevSession != nil {
		guid = prevSession.Token
	}
	self = handler.Session{
		Username: "BanchoBot",
		UserID:   1,
		Token:    guid,
		Admin:    true,
		LastSeen: time.Now().Add(time.Hour * 24 * 365 * 10),
		State: packets.OsuSendUserState{
			Action: 0,
		},
	}
	if prevSession != nil {
		handler.Sessions.Add(self)
	}
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
	parts := strings.Split(msg, " ")
	switch parts[0] {
	case "ping":
		newMsg.Content = "What? You're sending me a ping and you are " +
			"expecting me to respond?!? You're so-- wait, I guess this kind " +
			"of counts as a response. Oh well."
	case "admin":
		if user.Admin {
			newMsg.Content = "yes"
		} else {
			newMsg.Content = "no"
		}
	case "create_channel":
		if len(parts) < 2 || !user.Admin {
			return
		}
		handler.AddChannel(handler.Channel{
			Name:        parts[1],
			Description: strings.Join(parts[2:], " "),
		})
		newMsg.Content = parts[1] + " created!"
	case "remove_channel":
		if len(parts) != 2 || !user.Admin {
			return
		}
		handler.RemoveChannel(parts[1])
		newMsg.Content = "channel deleted!"
	}
	if newMsg.Content != "" {
		user.Send(newMsg)
	}
}

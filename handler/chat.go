package handler

import (
	"github.com/bnch/lan/packets"
)

// Handler functions related to the osu! chat.

func sendMessage(p *packets.OsuSendMessage, s Session) {
	if !s.In("chan/" + p.Channel) {
		s.Send(&packets.BanchoSendMessage{
			SenderName: "BanchoBot",
			SenderID:   1,
			Content:    "You haven't joined that channel yet!",
			Channel:    "BanchoBot",
		})
		return
	}
	p.SenderName = s.Username
	p.SenderID = s.UserID
	converted := packets.BanchoSendMessage(*p)
	SendMessageToChannel(&converted)
}

func sendPrivateMessage(p *packets.OsuSendPrivateMessage, s Session) {
	u := GetSession(Sessions.TokenFromUsername(p.Channel))
	if u == nil {
		return
	}
	u.Send(&packets.BanchoSendMessage{
		SenderID:   s.UserID,
		SenderName: s.Username,
		Content:    p.Content,
		Channel:    s.Username,
	})
}

func init() {
	RegisterHandlers(sendMessage, sendPrivateMessage)
}

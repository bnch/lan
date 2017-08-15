package handler

import (
	"github.com/bnch/lan/packets"
)

// Handler functions related to the osu! chat.

func banchobotMessage(s Session, message string) {
	s.Send(&packets.BanchoSendMessage{
		SenderName: "BanchoBot",
		SenderID:   1,
		Content:    message,
		Channel:    "BanchoBot",
	})
}

func sendMessage(p *packets.OsuSendMessage, s Session) {
	if !s.In("chan/" + p.Channel) {
		banchobotMessage(s, "You haven't joined that channel yet!")
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

func joinChannel(p *packets.OsuChannelJoin, s Session) {
	if !ChannelExists(p.Channel) {
		banchobotMessage(s, "That channel does not exist.")
		s.Send(&packets.BanchoChannelRevoked{
			Channel: p.Channel,
		})
		return
	}
	s.SubscribeChannel(p.Channel)
	s.Send(&packets.BanchoChannelJoinSuccess{
		Channel: p.Channel,
	})
}

func leaveChannel(p *packets.OsuChannelLeave, s Session) {
	if !s.In("chan/" + p.Channel) {
		banchobotMessage(s, "You are not even in that channel in the first place!")
		return
	}
	s.UnsubscribeChannel(p.Channel)
	s.Send(&packets.BanchoChannelRevoked{Channel: p.Channel})
}

func init() {
	RegisterHandlers(
		sendMessage,
		sendPrivateMessage,
		joinChannel,
		leaveChannel,
	)
}

package handler

import "github.com/bnch/lan/packets"

// Stuff that deals with channels

// A Channel is a group of users that can chat together.
type Channel struct {
	Name        string
	Description string
}

// Users gets the number of users in a channel.
func (c Channel) Users() uint16 {
	return uint16(Redis.SCard("lan/session_list/chan/" + c.Name).Val())
}

// ToChannelAvailable converts a Channel to a BanchoChannelAvailable
func (c Channel) ToChannelAvailable() *packets.BanchoChannelAvailable {
	return &packets.BanchoChannelAvailable{
		Channel:     c.Name,
		Description: c.Description,
		Users:       c.Users(),
	}
}

// AddChannel initialises a channel, creating it on redis.
func AddChannel(c Channel) {
	Redis.HSet("lan/channels", c.Name, c.Description)
}

// GetChannels makes a copy of channels.
func GetChannels() []Channel {
	chMap := Redis.HGetAll("lan/channels").Val()

	chans := make([]Channel, len(chMap))
	i := 0
	for k, v := range chMap {
		chans[i] = Channel{Name: k, Description: v}
	}

	return chans
}

// RemoveChannel removes a channel.
func RemoveChannel(name string) {
	Redis.HDel("lan/channels", name)
	// TODO: Send BanchoChannelRevoked to all members, remove it from their
	// collections, etc.
}

// ChannelExists checks whether a channel exists
func ChannelExists(ch string) bool {
	return Redis.HExists("lan/channels", ch).Val()
}

// SubscribeChannel subscribes the user to a channel.
func (s Session) SubscribeChannel(ch string) {
	s.Subscribe("chan/" + ch)
}

// UnsubscribeChannel unsubscribes the user from a channel.
func (s Session) UnsubscribeChannel(ch string) {
	s.Unsubscribe("chan/" + ch)
}

// SendMessageToChannel broadcasts a BanchoSendMessage to all the members of a
// channel.
func SendMessageToChannel(m *packets.BanchoSendMessage) {
	if !Redis.HExists("lan/channels", m.Channel).Val() {
		return
	}
	SessionCollection("chan/"+m.Channel).SendExcept([]int32{m.SenderID}, m)
}

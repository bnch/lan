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
	streamsMutex.RLock()
	defer streamsMutex.RUnlock()
	return uint16(streams["chan/"+c.Name].Len())
}

// ToChannelAvailable converts a Channel to a BanchoChannelAvailable
func (c Channel) ToChannelAvailable() *packets.BanchoChannelAvailable {
	return &packets.BanchoChannelAvailable{
		Channel:     c.Name,
		Description: c.Description,
		Users:       c.Users(),
	}
}

// It's not thread-safe, but it's only meant to be used rarely when it comes to
// writing it so fuck it.
var channels []Channel

// AddChannel initialises a channel, creating the stream for it and adding it to
// the channels slice.
func AddChannel(c Channel) {
	streamsMutex.Lock()
	defer streamsMutex.Unlock()
	streams["chan/"+c.Name] = new(SessionCollection)
	channels = append(channels, c)
}

// GetChannels makes a copy of channels.
func GetChannels() []Channel {
	ch := make([]Channel, len(channels))
	copy(ch, channels)
	return ch
}

// RemoveChannel removes a channel.
func RemoveChannel(name string) {
	for i, x := range channels {
		if x.Name == name {
			channels[i] = channels[len(channels)-1]
			channels = channels[:len(channels)-1]
			streamsMutex.Lock()
			defer streamsMutex.Unlock()
			delete(streams, "chan/"+name)
			return
		}
	}
}

// SubscribeChannel subscribes the user to a channel.
func (s *Session) SubscribeChannel(ch string) {
	s.Subscribe("chan/" + ch)
	s.Send(&packets.BanchoChannelJoinSuccess{ch})
}

// UnsubscribeChannel unsubscribes the user from a channel.
func (s *Session) UnsubscribeChannel(ch string) {
	s.Unsubscribe("chan/" + ch)
}

// SendMessageToChannel broadcasts a BanchoSendMessage to all the members of a
// channel.
func SendMessageToChannel(m *packets.BanchoSendMessage) {
	s, ok := streams["chan/"+m.Channel]
	if !ok {
		return
	}
	s.SendExcept([]int32{m.SenderID}, m)
}

func init() {
	AddChannel(Channel{
		"#osu",
		"General discussion about everything",
	})
	AddChannel(Channel{
		"#announce",
		"Probably not so important announcements.",
	})
}

package handler

import (
	"fmt"

	"github.com/bnch/lan/packets"
)

// SendUsers sends to this session all the online users.
func (s *Session) SendUsers() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("critical:", err)
		}
	}()
	c := Sessions.Copy()
	chunk := make([]int32, 0, 512)
	for _, e := range c {
		chunk = append(chunk, e.UserID)
		if len(chunk) == 512 {
			s.Send(&packets.BanchoUserPresenceBundle{chunk})
			chunk = make([]int32, 0, 512)
		}
	}
	s.Send(&packets.BanchoUserPresenceBundle{chunk})
}

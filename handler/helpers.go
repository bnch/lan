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

// ToUserPresence converts a session to a packets.BanchoUserPresence
func (s *Session) ToUserPresence() *packets.BanchoUserPresence {
	return &packets.BanchoUserPresence{
		ID:         s.UserID,
		Name:       s.Username,
		UTCOffset:  24,
		Country:    0,
		Privileges: uint8(s.Permissions()),
		Rank:       s.UserID,
	}
}

// ToHandleUserUpdate converts a session to a packets.BanchoHandleUserUpdate
func (s *Session) ToHandleUserUpdate() *packets.BanchoHandleUserUpdate {
	return &packets.BanchoHandleUserUpdate{
		ID:          s.UserID,
		Action:      packets.StatusIdle,
		RankedScore: 0,
		TotalScore:  0,
		Rank:        s.UserID,
	}
}

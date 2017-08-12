package handler

import (
	"fmt"
	"time"

	"github.com/bnch/lan/packets"
)

// ProtocolVersion is the version of the Bancho protocol.
const ProtocolVersion = 19

// Handle takes a set of packets, handles them and then pushes the results
func (s *Session) Handle(pks []packets.Packet) {
	for _, p := range pks {
		go s.handle(p)
	}
}

func (s *Session) handle(p packets.Packet) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Error while handling a %T: %v\n%#v\n", p, err, p)
		}
	}()
	fmt.Printf("< %#v\n", p)

	// Update last seen
	s.Mutex.Lock()
	s.LastSeen = time.Now()
	s.Mutex.Unlock()

	switch p := p.(type) {
	case *packets.OsuSendUserState:
		// Set the user's state to that requested to have.
		s.State = *p
	case *packets.OsuSendMessage:
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
	case *packets.OsuExit:
		// Log out the user
		s.Mutex.RLock()
		fmt.Printf("> %s has logged out\n", s.Username)
		s.Mutex.RUnlock()
		s.Dispose()
	case *packets.OsuUserStatsRequest:
		// Send to osu! information about the users it requests.
		for _, i := range p.IDs {
			u := Sessions.GetByID(i)
			if u == nil {
				continue
			}
			s.Send(u.ToHandleUserUpdate())
		}
	case *packets.OsuUserPresenceRequest:
		// Send to osu! information about the users it requests.
		for _, i := range p.IDs {
			u := Sessions.GetByID(i)
			if u == nil {
				continue
			}
			s.Send(u.ToUserPresence())
		}
	case *packets.OsuPong:
		// Do nothing
	case *packets.OsuRequestStatusUpdate:
		s.Send(s.ToHandleUserUpdate())
		// Should also be sent to spectators if any
	case *packets.OsuSendPrivateMessage:
		u := Sessions.GetByUsername(p.Channel)
		if u == nil {
			return
		}
		u.Send(&packets.BanchoSendMessage{
			SenderID:   s.UserID,
			SenderName: s.Username,
			Content:    p.Content,
			Channel:    s.Username,
		})
	default:
		fmt.Printf("> got %T\n", p)
	}
}

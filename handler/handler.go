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
	switch p := p.(type) {
	case *packets.OsuSendUserState:
		// Set the user's state to that requested to have.
		s.State = *p
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
			s.Send(u.ToUserPresence())
		}
	case *packets.OsuPong:
		// Update the Last Seen. That's it.
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		s.LastSeen = time.Now()
	default:
		fmt.Printf("> got %T\n", p)
	}
}

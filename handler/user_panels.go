package handler

import (
	"fmt"

	"github.com/bnch/lan/packets"
)

// This files contains all the handlers related to changing the user panels in
// f9, even indirectly (such as logout etc.).

// Set the user's state to that requested to have.
func sendUserState(p *packets.OsuSendUserState, s Session) {
	s.State = *p
	SaveSession(s)
}

// Log out the user
func osuExit(p *packets.OsuExit, s Session) {
	fmt.Printf("> %s has logged out\n", s.Username)
	s.Dispose()
}

func userStatsRequest(p *packets.OsuUserStatsRequest, s Session) {
	// Send to osu! information about the users it requests.
	for _, i := range p.IDs {
		u := GetSession(Sessions.TokenFromID(i))
		if u == nil {
			continue
		}
		s.Send(u.ToHandleUserUpdate())
	}
}

func userPresenceRequest(p *packets.OsuUserPresenceRequest, s Session) {
	// Send to osu! information about the users it requests.
	for _, i := range p.IDs {
		u := GetSession(Sessions.TokenFromID(i))
		if u == nil {
			continue
		}
		s.Send(u.ToUserPresence())
	}
}

func updateUserStats(p *packets.OsuRequestStatusUpdate, s Session) {
	s.Send(s.ToHandleUserUpdate())
	// Should also be sent to spectators if any
}

func init() {
	RegisterHandlers(
		sendUserState,
		osuExit,
		userStatsRequest,
		userPresenceRequest,
		updateUserStats,
	)
}

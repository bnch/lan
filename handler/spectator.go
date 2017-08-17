package handler

import (
	"github.com/bnch/lan/packets"
)

func startSpectating(p *packets.OsuStartSpectating, s Session) {
	// If we are spectating someone, stop that.
	if s.Spectating != "" {
		stopSpectating(nil, s)
	}

	// Remove spectators of user.
	mySpectators := SessionCollection("spec/" + s.Token)
	if mySpectators.Len() > 0 {
		mySpectators.Send(&packets.BanchoChannelRevoked{Channel: "#spectator"})
		mySpectators.Destroy()
	}

	// Get who we want to spectate.
	specWant := p.User
	tok := Sessions.TokenFromID(specWant)
	if tok == "" {
		return
	}
	target := GetSession(tok)

	// user does not exist.
	if target == nil {
		return
	}

	specColl := SessionCollection("spec/" + target.Token)
	if specColl.Len() == 0 {
		target.Send(
			&packets.BanchoChannelAvailableAutojoin{Channel: "#spectator"},
			&packets.BanchoChannelJoinSuccess{Channel: "#spectator"},
		)
		specColl.Add(*target)
	}
	specColl.Add(s)
	s.Spectating = target.Token
	SaveSession(s)
	s.Send(
		&packets.BanchoChannelAvailableAutojoin{Channel: "#spectator"},
		&packets.BanchoChannelJoinSuccess{Channel: "#spectator"},
	)
	specColl.SendExcept(
		[]int32{target.UserID},
		&packets.BanchoFellowSpectatorJoined{User: s.UserID},
	)
	target.Send(&packets.BanchoSpectatorJoined{User: s.UserID})
	specColl.Send(&packets.BanchoSendMessage{
		SenderName: "BanchoBot",
		SenderID:   1,
		Channel:    "#spectator",
		Content:    s.Username + " joined. Welcome, and have fun!",
	})
}

func stopSpectating(p *packets.OsuStopSpectating, s Session) {
	// check that we are spectating someone
	if s.Spectating == "" {
		return
	}

	host := GetSession(s.Spectating)

	// unsubscribe from stream and revoke #spectator channel
	s.Unsubscribe("spec/" + s.Spectating)
	s.Send(&packets.BanchoChannelRevoked{Channel: "#spectator"})
	s.Spectating = ""
	SaveSession(s)

	// send the necessary packets to host and all the fellow spectators
	if host != nil {
		specCollection := SessionCollection("spec/" + s.Spectating)
		specCollection.SendExcept([]int32{s.UserID, host.UserID}, &packets.BanchoFellowSpectatorLeft{User: s.UserID})
		host.Send(&packets.BanchoSpectatorLeft{User: s.UserID})

		// if there are no other users left, we can destroy the collection.
		if specCollection.Len() == 1 && specCollection.AllTokens()[0] == host.Token {
			specCollection.Destroy()
		}
	}
}

func spectateFrames(p *packets.OsuSpectateFrames, s Session) {
	converted := packets.BanchoSpectateFrames(*p)
	// Even if nobody is actually spectating this user, come what may.
	SessionCollection("spec/"+s.Token).SendExcept([]int32{s.UserID}, &converted)
}

func init() {
	RegisterHandlers(
		startSpectating,
		stopSpectating,
		spectateFrames,
	)
}

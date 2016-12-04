package handler

import "github.com/bnch/lan/packets"

// special user id to tell the osu! client the authentication failed
const (
	authFailed = iota*-1 - 1
	authClientTooOld
	authBanned
	authBannedAgain
	authFailedRetryLater
	authNeedSupporter
	authWeakPassword
	authRequireVerification
)

// Authenticate logs in into lan.
func Authenticate(username string, password string) *Session {
	sess := NewSession(username, password)

	// if userid is negative, it's one of the special user IDs
	if sess.UserID < 0 {
		sess.Send(&packets.BanchoLoginReply{sess.UserID})
		return sess
	}

	sess.Send(
		&packets.BanchoBanInfo{0},
		&packets.BanchoLoginReply{sess.UserID},
		&packets.BanchoProtocolVersion{ProtocolVersion},
		&packets.BanchoLoginPermissions{sess.Permissions()},
		&packets.BanchoFriendList{},
		sess.ToUserPresence(),
		sess.ToHandleUserUpdate(),
		&packets.BanchoChannelJoinSuccess{"#osu"},
		&packets.BanchoChannelJoinSuccess{"#announce"},
		&packets.BanchoChannelAvailable{"#osu", "osu", 1},
		&packets.BanchoChannelAvailable{"#announce", "announce", 1},
		&packets.BanchoChannelListingComplete{},
	)

	go sess.SendUsers()

	return sess
}

// LogoutTokenNotFound returns a Session for attempts to connect when there's
// no such user stored with such token.
func LogoutTokenNotFound() *Session {
	sess := &Session{
		Token:   GenerateGUID(),
		Packets: make(chan packets.Packet, 5),
	}
	sess.Send(&packets.BanchoLoginReply{authFailedRetryLater})
	sess.Send(&packets.BanchoAnnounce{"Sorry, but we had an issue logging you in. Trying again..."})
	return sess
}

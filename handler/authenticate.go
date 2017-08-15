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

// ProtocolVersion is the version of the Bancho protocol.
const ProtocolVersion = 19

// Authenticate logs in into lan.
func Authenticate(username string, password string) Session {
	sess := NewSession(username, password)

	// if userid is negative, it's one of the special user IDs
	if sess.UserID < 0 {
		sess.Send(&packets.BanchoLoginReply{UserID: sess.UserID})
		return sess
	}

	sess.Send(
		&packets.BanchoBanInfo{Seconds: 0},
		&packets.BanchoLoginReply{UserID: sess.UserID},
		&packets.BanchoProtocolVersion{Version: ProtocolVersion},
		&packets.BanchoLoginPermissions{Permissions: sess.Permissions()},
		&packets.BanchoFriendList{},
		sess.ToUserPresence(),
		sess.ToHandleUserUpdate(),
	)

	sess.SubscribeChannel("#osu")
	sess.SubscribeChannel("#announce")
	sess.Send(&packets.BanchoChannelJoinSuccess{Channel: "#osu"},
		&packets.BanchoChannelJoinSuccess{Channel: "#announce"})

	go sess.SendUsers()

	ch := GetChannels()
	for _, c := range ch {
		sess.Send(c.ToChannelAvailable())
	}
	sess.Send(&packets.BanchoChannelListingComplete{})

	Sessions.Send(&packets.BanchoUserPresenceSingle{ID: sess.UserID})

	return sess
}

// LogoutTokenNotFound returns a Session for attempts to connect when there's
// no such user stored with such token.
func LogoutTokenNotFound() Session {
	sess := Session{
		Token: GenerateGUID(),
	}
	sess.Send(&packets.BanchoLoginReply{UserID: authFailedRetryLater})
	sess.Send(&packets.BanchoAnnounce{Message: "Sorry, but we had an issue logging you in. Trying again..."})
	return sess
}

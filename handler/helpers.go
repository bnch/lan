package handler

import (
	"fmt"

	"github.com/bnch/lan/packets"
)

// SendUsers sends to this session all the online users.
func (s Session) SendUsers() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("critical:", err)
		}
	}()

	// Simpler demonstration: https://play.golang.org/p/QKp3emL58a

	const chunkSize = 512

	c := Sessions.AllUserIDs()
	for i := 0; i < (len(c)/chunkSize + 1); i++ {
		start := chunkSize * i
		end := start + chunkSize
		if end > len(c) {
			end = len(c)
		}

		if len(c[start:end]) > 0 {
			s.Send(&packets.BanchoUserPresenceBundle{IDs: c[start:end]})
		}
	}
}

// ToUserPresence converts a session to a packets.BanchoUserPresence
func (s Session) ToUserPresence() *packets.BanchoUserPresence {
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
func (s Session) ToHandleUserUpdate() *packets.BanchoHandleUserUpdate {
	return &packets.BanchoHandleUserUpdate{
		ID:              s.UserID,
		Action:          s.State.Action,
		ActionText:      s.State.Text,
		ActionMapMD5:    s.State.MapMD5,
		ActionMods:      s.State.Mods,
		ActionGameMode:  s.State.GameMode,
		ActionBeatmapID: s.State.BeatmapID,
		RankedScore:     0,
		TotalScore:      0,
		Rank:            s.UserID,
	}
}

// Start runs all initialisations required for this package.
func Start() {
	AddChannel(Channel{
		"#osu",
		"General discussion about everything",
	})
	AddChannel(Channel{
		"#announce",
		"Probably not so important announcements.",
	})
}

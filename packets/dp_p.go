// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	"errors"
	"fmt"
	"io"
	
	"github.com/bnch/lan/osubinary"
)

// Packetify transforms a slice of packets into a slice of bytes, to be
// transmitted to the osu! client.
func Packetify(packets []Packetifier) ([]byte, error) {
	w := osubinary.NewWriter()
	for _, packet := range packets {
		err := packetify(w, packet)
		if err != nil {
			return nil, err
		}
	}
	_, err := w.End()
	return w.Bytes(), err
}

func packetify(w *osubinary.OsuWriteChain, p Packetifier) error {
	data, err := p.Packetify()
	if err != nil {
		return err
	}
	var id uint16
	switch p.(type) {
	case *OsuSendUserState: id = 0
	case *OsuSendMessage: id = 1
	case *OsuExit: id = 2
	case *OsuRequestStatusUpdate: id = 3
	case *OsuPong: id = 4
	case *BanchoLoginReply: id = 5
	case *BanchoSendMessage: id = 7
	case *BanchoHandleUserUpdate: id = 11
	case *BanchoUserQuit: id = 12
	case *BanchoAnnounce: id = 24
	case *OsuSendPrivateMessage: id = 25
	case *OsuChannelJoin: id = 63
	case *BanchoChannelJoinSuccess: id = 64
	case *BanchoChannelAvailable: id = 65
	case *BanchoChannelRevoked: id = 66
	case *BanchoChannelAvailableAutojoin: id = 67
	case *BanchoLoginPermissions: id = 71
	case *BanchoFriendList: id = 72
	case *BanchoProtocolVersion: id = 75
	case *OsuChannelLeave: id = 78
	case *BanchoUserPresence: id = 83
	case *OsuUserStatsRequest: id = 85
	case *BanchoChannelListingComplete: id = 89
	case *BanchoBanInfo: id = 92
	case *BanchoUserSilenced: id = 94
	case *BanchoUserPresenceSingle: id = 95
	case *BanchoUserPresenceBundle: id = 96
	case *OsuUserPresenceRequest: id = 97

	default:
		return errors.New("invalid packet")
	}
	w.Packet(id, data)
	return nil
}

// Depacketify decodes a byte slice received from the osu! client into a
// packet slice.
func Depacketify(b []byte) ([]Packet, error) {
	r := osubinary.NewReaderFromBytes(b)
	var packets []Packet
	for {
		id, pack, err := r.Packet()
		if err != nil {
			return nil, err
		}
		_, err = r.End()
		if err == io.EOF {
			return packets, nil
		}
		if err != nil {
			return nil, err
		}
		packet, err := depacketify(id, pack)
		if err != nil {
			return nil, err
		}
		if packet != nil {	
			packets = append(packets, packet)
		}
	}
}

func depacketify(id uint16, packet []byte) (Packet, error) {
	var p Packet
	switch id {
	case 0: p = &OsuSendUserState{}
	case 1: p = &OsuSendMessage{}
	case 2: p = &OsuExit{}
	case 3: p = &OsuRequestStatusUpdate{}
	case 4: p = &OsuPong{}
	case 5: p = &BanchoLoginReply{}
	case 7: p = &BanchoSendMessage{}
	case 11: p = &BanchoHandleUserUpdate{}
	case 12: p = &BanchoUserQuit{}
	case 24: p = &BanchoAnnounce{}
	case 25: p = &OsuSendPrivateMessage{}
	case 63: p = &OsuChannelJoin{}
	case 64: p = &BanchoChannelJoinSuccess{}
	case 65: p = &BanchoChannelAvailable{}
	case 66: p = &BanchoChannelRevoked{}
	case 67: p = &BanchoChannelAvailableAutojoin{}
	case 71: p = &BanchoLoginPermissions{}
	case 72: p = &BanchoFriendList{}
	case 75: p = &BanchoProtocolVersion{}
	case 78: p = &OsuChannelLeave{}
	case 83: p = &BanchoUserPresence{}
	case 85: p = &OsuUserStatsRequest{}
	case 89: p = &BanchoChannelListingComplete{}
	case 92: p = &BanchoBanInfo{}
	case 94: p = &BanchoUserSilenced{}
	case 95: p = &BanchoUserPresenceSingle{}
	case 96: p = &BanchoUserPresenceBundle{}
	case 97: p = &OsuUserPresenceRequest{}

	default:
		fmt.Printf("Asked to depacketify an unknown packet: %d\n", int(id))
		return nil, nil // errors.New("invalid packet ID (" + strconv.Itoa(int(id)) + ")")
	}
	err := p.Depacketify(packet)
	return p, err
}

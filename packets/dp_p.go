// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	"errors"
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
	case *OsuExit: id = 2
	case *BanchoLoginReply: id = 5
	case *BanchoHandleUserUpdate: id = 11
	case *BanchoAnnounce: id = 24
	case *OsuChannelJoin: id = 63
	case *BanchoChannelJoinSuccess: id = 64
	case *BanchoChannelAvailable: id = 65
	case *BanchoChannelRevoked: id = 66
	case *BanchoChannelAvailableAutojoin: id = 67
	case *BanchoLoginPermissions: id = 71
	case *BanchoFriendList: id = 72
	case *BanchoProtocolVersion: id = 75
	case *BanchoUserPresence: id = 83
	case *BanchoChannelListingComplete: id = 89
	case *BanchoBanInfo: id = 92
	case *BanchoUserSilenced: id = 94
	case *BanchoUserPresenceSingle: id = 95
	case *BanchoUserPresenceBundle: id = 96

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
		packets = append(packets, packet)
	}
}

func depacketify(id uint16, packet []byte) (Packet, error) {
	var p Packet
	switch id {
	case 2: p = &OsuExit{}
	case 5: p = &BanchoLoginReply{}
	case 11: p = &BanchoHandleUserUpdate{}
	case 24: p = &BanchoAnnounce{}
	case 63: p = &OsuChannelJoin{}
	case 64: p = &BanchoChannelJoinSuccess{}
	case 65: p = &BanchoChannelAvailable{}
	case 66: p = &BanchoChannelRevoked{}
	case 67: p = &BanchoChannelAvailableAutojoin{}
	case 71: p = &BanchoLoginPermissions{}
	case 72: p = &BanchoFriendList{}
	case 75: p = &BanchoProtocolVersion{}
	case 83: p = &BanchoUserPresence{}
	case 89: p = &BanchoChannelListingComplete{}
	case 92: p = &BanchoBanInfo{}
	case 94: p = &BanchoUserSilenced{}
	case 95: p = &BanchoUserPresenceSingle{}
	case 96: p = &BanchoUserPresenceBundle{}

	default:
		return nil, errors.New("invalid packet ID")
	}
	err := p.Depacketify(packet)
	return p, err
}

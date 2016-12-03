// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoChannelAvailable is sent by the server to inform the osu! client about
// the existance of a channel. For channels where the user auto-joins,
// BanchoChannelAvailableAutojoin is used instead.
type BanchoChannelAvailable struct {
	Channel string
	Description string
	Users uint16
}

// Packetify encodes a BanchoChannelAvailable into
// a byte slice.
func (p BanchoChannelAvailable) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.Channel)
	w.BanchoString(p.Description)
	w.Uint16(p.Users)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoChannelAvailable.
func (p *BanchoChannelAvailable) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Channel = r.BanchoString()
	p.Description = r.BanchoString()
	p.Users = r.Uint16()

	_, err := r.End()
	return err
}

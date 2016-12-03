// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoChannelRevoked is sent by the server to the osu! client to get the fuck
// out of a channel.
type BanchoChannelRevoked struct {
	Channel string
}

// Packetify encodes a BanchoChannelRevoked into
// a byte slice.
func (p BanchoChannelRevoked) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.Channel)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoChannelRevoked.
func (p *BanchoChannelRevoked) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Channel = r.BanchoString()

	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuChannelJoin is sent by the osu! client whenever it wishes to join a channel.
type OsuChannelJoin struct {
	Channel string
}

// Packetify encodes a OsuChannelJoin into
// a byte slice.
func (p OsuChannelJoin) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.Channel)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuChannelJoin.
func (p *OsuChannelJoin) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Channel = r.BanchoString()

	_, err := r.End()
	return err
}

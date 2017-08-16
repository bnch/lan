// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuStartSpectating is sent by the osu! client when it wishes to start spectating
// somebody.
type OsuStartSpectating struct {
	User int32
}

// Packetify encodes a OsuStartSpectating into
// a byte slice.
func (p OsuStartSpectating) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.User)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuStartSpectating.
func (p *OsuStartSpectating) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.User = r.Int32()

	_, err := r.End()
	return err
}

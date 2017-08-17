// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoSpectatorCantSpectate is sent when an user who is spectating you doesn't
// have the beatmap you are playing
type BanchoSpectatorCantSpectate struct {
	User int32
}

// Packetify encodes a BanchoSpectatorCantSpectate into
// a byte slice.
func (p BanchoSpectatorCantSpectate) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.User)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoSpectatorCantSpectate.
func (p *BanchoSpectatorCantSpectate) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.User = r.Int32()

	_, err := r.End()
	return err
}

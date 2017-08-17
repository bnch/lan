// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoFellowSpectatorLeft is sent to all other spectators of an user when they
// aren't spectating the user anymore.
type BanchoFellowSpectatorLeft struct {
	User int32
}

// Packetify encodes a BanchoFellowSpectatorLeft into
// a byte slice.
func (p BanchoFellowSpectatorLeft) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.User)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoFellowSpectatorLeft.
func (p *BanchoFellowSpectatorLeft) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.User = r.Int32()

	_, err := r.End()
	return err
}

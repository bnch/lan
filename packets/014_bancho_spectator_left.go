// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoSpectatorLeft tells the client that a spectator has left.
type BanchoSpectatorLeft struct {
	User int32
}

// Packetify encodes a BanchoSpectatorLeft into
// a byte slice.
func (p BanchoSpectatorLeft) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.User)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoSpectatorLeft.
func (p *BanchoSpectatorLeft) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.User = r.Int32()

	_, err := r.End()
	return err
}

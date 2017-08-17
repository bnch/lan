// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoSpectatorJoined tells the client that there is a new spectator.
type BanchoSpectatorJoined struct {
	User int32
}

// Packetify encodes a BanchoSpectatorJoined into
// a byte slice.
func (p BanchoSpectatorJoined) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.User)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoSpectatorJoined.
func (p *BanchoSpectatorJoined) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.User = r.Int32()

	_, err := r.End()
	return err
}

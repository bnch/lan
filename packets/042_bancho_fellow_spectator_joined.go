// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoFellowSpectatorJoined is sent to all other spectators of an user when they
// connect (except to the new spectator.)
type BanchoFellowSpectatorJoined struct {
	User int32
}

// Packetify encodes a BanchoFellowSpectatorJoined into
// a byte slice.
func (p BanchoFellowSpectatorJoined) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.User)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoFellowSpectatorJoined.
func (p *BanchoFellowSpectatorJoined) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.User = r.Int32()

	_, err := r.End()
	return err
}

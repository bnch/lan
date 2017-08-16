// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoUserQuit tells clients that an user has exited the game.
type BanchoUserQuit struct {
	ID int32
	State byte // Gone = 0, OsuRemaining = 1, IRCRemaining = 2
}

// Packetify encodes a BanchoUserQuit into
// a byte slice.
func (p BanchoUserQuit) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.ID)
	w.Byte(p.State)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoUserQuit.
func (p *BanchoUserQuit) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.ID = r.Int32()
	p.State = r.Byte()

	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoUserSilenced is sent when an user (specified by its ID) is silenced,
// so that their messages are cleared.
type BanchoUserSilenced struct {
	ID int32
}

// Packetify encodes a BanchoUserSilenced into
// a byte slice.
func (p BanchoUserSilenced) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.ID)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoUserSilenced.
func (p *BanchoUserSilenced) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.ID = r.Int32()

	_, err := r.End()
	return err
}

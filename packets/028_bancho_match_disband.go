// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoMatchDisband tells a user of the match that the match has been disbanded.
// This is also sent in response to an user requesting to leave the match.
type BanchoMatchDisband struct {
	Match int32
}

// Packetify encodes a BanchoMatchDisband into
// a byte slice.
func (p BanchoMatchDisband) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.Match)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoMatchDisband.
func (p *BanchoMatchDisband) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Match = r.Int32()

	_, err := r.End()
	return err
}

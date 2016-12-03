// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoUserPresenceSingle is used to broadcasts to users about when an user
// comes online.
type BanchoUserPresenceSingle struct {
	ID int32
}

// Packetify encodes a BanchoUserPresenceSingle into
// a byte slice.
func (p BanchoUserPresenceSingle) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.ID)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoUserPresenceSingle.
func (p *BanchoUserPresenceSingle) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.ID = r.Int32()

	_, err := r.End()
	return err
}

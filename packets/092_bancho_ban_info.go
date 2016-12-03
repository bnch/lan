// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoBanInfo tells the client information about the lenght of the silence for
// the user.
type BanchoBanInfo struct {
	Seconds int32
}

// Packetify encodes a BanchoBanInfo into
// a byte slice.
func (p BanchoBanInfo) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.Seconds)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoBanInfo.
func (p *BanchoBanInfo) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Seconds = r.Int32()

	_, err := r.End()
	return err
}

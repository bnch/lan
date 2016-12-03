// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoAnnounce is sent by bancho to notify the osu! clients of something.
type BanchoAnnounce struct {
	Message string
}

// Packetify encodes a BanchoAnnounce into
// a byte slice.
func (p BanchoAnnounce) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.Message)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoAnnounce.
func (p *BanchoAnnounce) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Message = r.BanchoString()

	_, err := r.End()
	return err
}

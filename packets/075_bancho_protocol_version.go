// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoProtocolVersion is the version of the bancho protocol, which is sent to
// the client at login.
type BanchoProtocolVersion struct {
	Version int32
}

// Packetify encodes a BanchoProtocolVersion into
// a byte slice.
func (p BanchoProtocolVersion) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.Version)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoProtocolVersion.
func (p *BanchoProtocolVersion) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Version = r.Int32()

	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuPong is received from the client to notify it's still alive (sent when
// there would otherwise be nothing in the POST body)
type OsuPong struct {
}

// Packetify encodes a OsuPong into
// a byte slice.
func (p OsuPong) Packetify() ([]byte, error) {
	w := ob.NewWriter()


	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuPong.
func (p *OsuPong) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)


	_, err := r.End()
	return err
}

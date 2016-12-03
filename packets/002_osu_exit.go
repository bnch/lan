// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuExit is sent by the osu! client whenever the client closes
type OsuExit struct {
	Reason int32
}

// Packetify encodes a OsuExit into
// a byte slice.
func (p OsuExit) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.Reason)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuExit.
func (p *OsuExit) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Reason = r.Int32()

	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuStopSpectating is sent by the client when it wishes to stop spectating the
// user it is spectating.
type OsuStopSpectating struct {
}

// Packetify encodes a OsuStopSpectating into
// a byte slice.
func (p OsuStopSpectating) Packetify() ([]byte, error) {
	w := ob.NewWriter()


	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuStopSpectating.
func (p *OsuStopSpectating) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)


	_, err := r.End()
	return err
}

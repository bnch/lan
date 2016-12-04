// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuRequestStatusUpdate is a request to update the status of the current user. If
// it changed, then it is also a request to distribute it to all those who need it
// (spectators)
type OsuRequestStatusUpdate struct {
}

// Packetify encodes a OsuRequestStatusUpdate into
// a byte slice.
func (p OsuRequestStatusUpdate) Packetify() ([]byte, error) {
	w := ob.NewWriter()


	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuRequestStatusUpdate.
func (p *OsuRequestStatusUpdate) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)


	_, err := r.End()
	return err
}

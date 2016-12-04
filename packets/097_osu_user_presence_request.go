// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuUserPresenceRequest is a request to have a BanchoUserPresence of some
// users.
type OsuUserPresenceRequest struct {
	IDs []int32 // []
}

// Packetify encodes a OsuUserPresenceRequest into
// a byte slice.
func (p OsuUserPresenceRequest) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32Slice(p.IDs)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuUserPresenceRequest.
func (p *OsuUserPresenceRequest) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.IDs = r.Int32Slice()

	_, err := r.End()
	return err
}

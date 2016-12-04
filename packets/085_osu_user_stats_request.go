// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuUserStatsRequest is a request to have a BanchoHandleUserUpdate of some users.
type OsuUserStatsRequest struct {
	IDs []int32 // []
}

// Packetify encodes a OsuUserStatsRequest into
// a byte slice.
func (p OsuUserStatsRequest) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32Slice(p.IDs)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuUserStatsRequest.
func (p *OsuUserStatsRequest) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.IDs = r.Int32Slice()

	_, err := r.End()
	return err
}

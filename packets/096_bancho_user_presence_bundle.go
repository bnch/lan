// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoUserPresenceBundle is used to initially broadcast to an user the users
// currently online on the server.
type BanchoUserPresenceBundle struct {
	IDs []int32
}

// Packetify encodes a BanchoUserPresenceBundle into
// a byte slice.
func (p BanchoUserPresenceBundle) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32Slice(p.IDs)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoUserPresenceBundle.
func (p *BanchoUserPresenceBundle) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.IDs = r.Int32Slice()

	_, err := r.End()
	return err
}

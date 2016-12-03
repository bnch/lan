// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoLoginReply is a packet containing either an user ID or an authentication
// error (see handler/authenticate.go).
type BanchoLoginReply struct {
	UserID int32
}

// Packetify encodes a BanchoLoginReply into
// a byte slice.
func (p BanchoLoginReply) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.UserID)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoLoginReply.
func (p *BanchoLoginReply) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.UserID = r.Int32()

	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoLoginPermissions are the privileges assigned to the user.
type BanchoLoginPermissions struct {
	Permissions int32
}

// Packetify encodes a BanchoLoginPermissions into
// a byte slice.
func (p BanchoLoginPermissions) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.Permissions)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoLoginPermissions.
func (p *BanchoLoginPermissions) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Permissions = r.Int32()

	_, err := r.End()
	return err
}

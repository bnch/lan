// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoFriendList is a list of the friend of the user.
type BanchoFriendList struct {
	Friends []int32 // []
}

// Packetify encodes a BanchoFriendList into
// a byte slice.
func (p BanchoFriendList) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32Slice(p.Friends)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoFriendList.
func (p *BanchoFriendList) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Friends = r.Int32Slice()

	_, err := r.End()
	return err
}

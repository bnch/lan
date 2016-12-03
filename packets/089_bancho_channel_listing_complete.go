// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoChannelListingComplete says the channel listing has been completed.
type BanchoChannelListingComplete struct {
}

// Packetify encodes a BanchoChannelListingComplete into
// a byte slice.
func (p BanchoChannelListingComplete) Packetify() ([]byte, error) {
	w := ob.NewWriter()


	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoChannelListingComplete.
func (p *BanchoChannelListingComplete) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)


	_, err := r.End()
	return err
}

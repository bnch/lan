// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoChannelAvailableAutojoin tells the user about a channel, straight before
// adding them to it. I don't quite understand why this is a thing, but whatever.
type BanchoChannelAvailableAutojoin struct {
	Channel string
}

// Packetify encodes a BanchoChannelAvailableAutojoin into
// a byte slice.
func (p BanchoChannelAvailableAutojoin) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.Channel)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoChannelAvailableAutojoin.
func (p *BanchoChannelAvailableAutojoin) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Channel = r.BanchoString()

	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoChannelJoinSuccess is sent by the server the client has successfully
// joined a channel.
type BanchoChannelJoinSuccess struct {
	Channel string
}

// Packetify encodes a BanchoChannelJoinSuccess into
// a byte slice.
func (p BanchoChannelJoinSuccess) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.Channel)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoChannelJoinSuccess.
func (p *BanchoChannelJoinSuccess) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Channel = r.BanchoString()

	_, err := r.End()
	return err
}

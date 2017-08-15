// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuChannelLeave requests to part a channel.
type OsuChannelLeave struct {
	Channel string
}

// Packetify encodes a OsuChannelLeave into
// a byte slice.
func (p OsuChannelLeave) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.Channel)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuChannelLeave.
func (p *OsuChannelLeave) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Channel = r.BanchoString()

	_, err := r.End()
	return err
}

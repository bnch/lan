// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuCantSpectate is sent by the client when they don't have the beatmap the host
// is playing.
type OsuCantSpectate struct {
}

// Packetify encodes a OsuCantSpectate into
// a byte slice.
func (p OsuCantSpectate) Packetify() ([]byte, error) {
	w := ob.NewWriter()


	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuCantSpectate.
func (p *OsuCantSpectate) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)


	_, err := r.End()
	return err
}

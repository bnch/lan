// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuLobbyJoin requests to join the multiplayer lobby.
type OsuLobbyJoin struct {
}

// Packetify encodes a OsuLobbyJoin into
// a byte slice.
func (p OsuLobbyJoin) Packetify() ([]byte, error) {
	w := ob.NewWriter()


	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuLobbyJoin.
func (p *OsuLobbyJoin) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)


	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuLobbyPart requests to leave the multiplayer lobby and not get any more
// updates about the lobby.
type OsuLobbyPart struct {
}

// Packetify encodes a OsuLobbyPart into
// a byte slice.
func (p OsuLobbyPart) Packetify() ([]byte, error) {
	w := ob.NewWriter()


	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuLobbyPart.
func (p *OsuLobbyPart) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)


	_, err := r.End()
	return err
}

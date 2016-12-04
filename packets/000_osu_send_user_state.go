// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuSendUserState informs bancho about what's the user currently doing (playing,
// listening, multiplaying what beatmap with what mods on what mode)
type OsuSendUserState struct {
	Action uint8
	Text string
	MapMD5 string
	Mods int32
	GameMode uint8
	BeatmapID uint32
}

// Packetify encodes a OsuSendUserState into
// a byte slice.
func (p OsuSendUserState) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Uint8(p.Action)
	w.BanchoString(p.Text)
	w.BanchoString(p.MapMD5)
	w.Int32(p.Mods)
	w.Uint8(p.GameMode)
	w.Uint32(p.BeatmapID)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuSendUserState.
func (p *OsuSendUserState) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Action = r.Uint8()
	p.Text = r.BanchoString()
	p.MapMD5 = r.BanchoString()
	p.Mods = r.Int32()
	p.GameMode = r.Uint8()
	p.BeatmapID = r.Uint32()

	_, err := r.End()
	return err
}

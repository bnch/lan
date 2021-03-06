// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoHandleUserUpdate contains all the information about an user you will ever
// need.
type BanchoHandleUserUpdate struct {
	ID int32
	Action uint8 // afk, modding, playing, osudirect...
	ActionText string
	ActionMapMD5 string
	ActionMods int32
	ActionGameMode uint8
	ActionBeatmapID uint32
	RankedScore uint64
	Accuracy float32 // 0.1337 = 13.37%
	Playcount uint32
	TotalScore uint64
	Rank int32
	PP uint16
}

// Packetify encodes a BanchoHandleUserUpdate into
// a byte slice.
func (p BanchoHandleUserUpdate) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.ID)
	w.Uint8(p.Action)
	w.BanchoString(p.ActionText)
	w.BanchoString(p.ActionMapMD5)
	w.Int32(p.ActionMods)
	w.Uint8(p.ActionGameMode)
	w.Uint32(p.ActionBeatmapID)
	w.Uint64(p.RankedScore)
	w.Float32(p.Accuracy)
	w.Uint32(p.Playcount)
	w.Uint64(p.TotalScore)
	w.Int32(p.Rank)
	w.Uint16(p.PP)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoHandleUserUpdate.
func (p *BanchoHandleUserUpdate) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.ID = r.Int32()
	p.Action = r.Uint8()
	p.ActionText = r.BanchoString()
	p.ActionMapMD5 = r.BanchoString()
	p.ActionMods = r.Int32()
	p.ActionGameMode = r.Uint8()
	p.ActionBeatmapID = r.Uint32()
	p.RankedScore = r.Uint64()
	p.Accuracy = r.Float32()
	p.Playcount = r.Uint32()
	p.TotalScore = r.Uint64()
	p.Rank = r.Int32()
	p.PP = r.Uint16()

	_, err := r.End()
	return err
}

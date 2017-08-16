// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuSpectateFrames is sent by the client, and contains the frames that must be
// sent to all the spectators of the user.
type OsuSpectateFrames struct {
	Extra int32
	Frames []ob.SpectateFrame
	Action byte
	Time int32
	ID byte
	Count300 uint16
	Count100 uint16
	Count50 uint16
	Geki uint16
	Katu uint16
	Miss uint16
	TotalScore int32
	MaxCombo uint16
	CurrentCombo uint16
	Perfect bool
	HP byte
	TagByte byte // ???
	UsingScorev2 bool
	ComboPortion float64 // only if UsingScorev2 is true
	BonusPortion float64 // only if UsingScorev2 is true
}

// Packetify encodes a OsuSpectateFrames into
// a byte slice.
func (p OsuSpectateFrames) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.Extra)
	w.SpectateFrameSlice(p.Frames)
	w.Byte(p.Action)
	w.Int32(p.Time)
	w.Byte(p.ID)
	w.Uint16(p.Count300)
	w.Uint16(p.Count100)
	w.Uint16(p.Count50)
	w.Uint16(p.Geki)
	w.Uint16(p.Katu)
	w.Uint16(p.Miss)
	w.Int32(p.TotalScore)
	w.Uint16(p.MaxCombo)
	w.Uint16(p.CurrentCombo)
	w.Bool(p.Perfect)
	w.Byte(p.HP)
	w.Byte(p.TagByte)
	w.Bool(p.UsingScorev2)
	w.Float64(p.ComboPortion)
	w.Float64(p.BonusPortion)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuSpectateFrames.
func (p *OsuSpectateFrames) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.Extra = r.Int32()
	p.Frames = r.SpectateFrameSlice()
	p.Action = r.Byte()
	p.Time = r.Int32()
	p.ID = r.Byte()
	p.Count300 = r.Uint16()
	p.Count100 = r.Uint16()
	p.Count50 = r.Uint16()
	p.Geki = r.Uint16()
	p.Katu = r.Uint16()
	p.Miss = r.Uint16()
	p.TotalScore = r.Int32()
	p.MaxCombo = r.Uint16()
	p.CurrentCombo = r.Uint16()
	p.Perfect = r.Bool()
	p.HP = r.Byte()
	p.TagByte = r.Byte()
	p.UsingScorev2 = r.Bool()
	p.ComboPortion = r.Float64()
	p.BonusPortion = r.Float64()

	_, err := r.End()
	return err
}

// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoSpectateFrames sends the client frames of the user they are spectating.
type BanchoSpectateFrames struct {
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
	Scorev2Portion ob.Scorev2Portion
	SomeUint16 uint16
}

// Packetify encodes a BanchoSpectateFrames into
// a byte slice.
func (p BanchoSpectateFrames) Packetify() ([]byte, error) {
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
	w.Scorev2Portion(p.Scorev2Portion)
	w.Uint16(p.SomeUint16)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoSpectateFrames.
func (p *BanchoSpectateFrames) Depacketify(b []byte) error {
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
	p.Scorev2Portion = r.Scorev2Portion()
	p.SomeUint16 = r.Uint16()

	_, err := r.End()
	return err
}

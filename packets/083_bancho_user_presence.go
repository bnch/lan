// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoUserPresence tells the client basic information about a certain user.
type BanchoUserPresence struct {
	ID int32
	Name string
	UTCOffset uint8
	Country uint8
	Privileges uint8
	Longitude float32
	Latitude float32
	Rank int32
}

// Packetify encodes a BanchoUserPresence into
// a byte slice.
func (p BanchoUserPresence) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.ID)
	w.BanchoString(p.Name)
	w.Uint8(p.UTCOffset)
	w.Uint8(p.Country)
	w.Uint8(p.Privileges)
	w.Float32(p.Longitude)
	w.Float32(p.Latitude)
	w.Int32(p.Rank)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoUserPresence.
func (p *BanchoUserPresence) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.ID = r.Int32()
	p.Name = r.BanchoString()
	p.UTCOffset = r.Uint8()
	p.Country = r.Uint8()
	p.Privileges = r.Uint8()
	p.Longitude = r.Float32()
	p.Latitude = r.Float32()
	p.Rank = r.Int32()

	_, err := r.End()
	return err
}

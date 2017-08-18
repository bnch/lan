// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// BanchoMatchNew is sent when an user joins a multiplayer lobby, for every match
// currently being played.
type BanchoMatchNew struct {
	ID int32
	InProgress bool
	MatchType byte
	ActiveMods uint32
	Name string
	Password string
	BeatmapName string
	BeatmapID int32
	BeatmapChecksum string
	SlotStatuses [16]byte
	SlotTeams [16]byte
	SlotUsers [16]int32
	Host int32
	PlayMode byte
	MatchScoringType byte
	MatchTeamType byte
	Freemod ob.Freemod
	Seed int32
}

// Packetify encodes a BanchoMatchNew into
// a byte slice.
func (p BanchoMatchNew) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.Int32(p.ID)
	w.Bool(p.InProgress)
	w.Byte(p.MatchType)
	w.Uint32(p.ActiveMods)
	w.BanchoString(p.Name)
	w.BanchoString(p.Password)
	w.BanchoString(p.BeatmapName)
	w.Int32(p.BeatmapID)
	w.BanchoString(p.BeatmapChecksum)
	w.ByteSlice(p.SlotStatuses[:])
	w.ByteSlice(p.SlotTeams[:])
	w.Int32Slice(p.SlotUsers[:])
	w.Int32(p.Host)
	w.Byte(p.PlayMode)
	w.Byte(p.MatchScoringType)
	w.Byte(p.MatchTeamType)
	w.Freemod(p.Freemod)
	w.Int32(p.Seed)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a BanchoMatchNew.
func (p *BanchoMatchNew) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.ID = r.Int32()
	p.InProgress = r.Bool()
	p.MatchType = r.Byte()
	p.ActiveMods = r.Uint32()
	p.Name = r.BanchoString()
	p.Password = r.BanchoString()
	p.BeatmapName = r.BanchoString()
	p.BeatmapID = r.Int32()
	p.BeatmapChecksum = r.BanchoString()
	copy(p.SlotStatuses[:], r.Reader.ByteSlice(16))
	copy(p.SlotTeams[:], r.Reader.ByteSlice(16))
	copy(p.SlotUsers[:], r.Reader.Int32Slice(16))
	p.Host = r.Int32()
	p.PlayMode = r.Byte()
	p.MatchScoringType = r.Byte()
	p.MatchTeamType = r.Byte()
	p.Freemod = r.Freemod()
	p.Seed = r.Int32()

	_, err := r.End()
	return err
}

// Package osubinary provides fast binary reading and writing for Bancho
// packets/osu! database files. Its main advantage from bnch/osubinary is
// speed: this package uses github.com/thehowl/binary, which is a much
// faster binary encoder and does not make use of reflection.
package osubinary

import (
	"bytes"
	"errors"
	"io"

	"github.com/bnch/lan/osubinary/uleb128"
	"github.com/thehowl/binary"
)

// OsuReader extends binary.Reader, providing methods related to the osu!
// binary format.
type OsuReader struct {
	binary.Reader
}

// NewReader creates a new reader.
func NewReader(r io.Reader) *OsuReader {
	return &OsuReader{
		binary.Reader{
			Reader:    r,
			ByteOrder: binary.LittleEndian,
		},
	}
}

// NewReaderFromBytes calls NewReader automatically creating an io.Reader from
// a []byte
func NewReaderFromBytes(b []byte) *OsuReader {
	return NewReader(bytes.NewReader(b))
}

// Uleb128 reads an uleb128 int from the Reader.
func (o *OsuReader) Uleb128() uint64 {
	return uleb128.UnmarshalReader(o.Reader.Reader)
}

// BanchoString reads an osu!db-string from the Reader. Specifically, an osu!db
// string is an uleb128 int, specifying the string's length, and then the
// actual raw string.
func (o *OsuReader) BanchoString() string {
	o.Byte() // ignored byte \x11
	len := o.Uleb128()
	return o.String(int(len))
}

// Int32Slice gets a slice of int32s, in the osu! format.
func (o *OsuReader) Int32Slice() []int32 {
	length := o.Uint16()
	return o.Reader.Int32Slice(int(length))
}

// Uint32Slice gets a slice of uint32s, in the osu! format.
func (o *OsuReader) Uint32Slice() []uint32 {
	length := o.Uint16()
	return o.Reader.Uint32Slice(int(length))
}

// SpectateFrame represents a single frame of a replay, or in this case, of the
// information sent to users through bancho.
type SpectateFrame struct {
	ButtonState byte
	MouseX      float32
	MouseY      float32
	Time        int32
}

// SpectateFrameSlice reads multiple frames broadcasted to spectators.
func (o *OsuReader) SpectateFrameSlice() []SpectateFrame {
	length := o.Uint16()
	sl := make([]SpectateFrame, length)
	for i := uint16(0); i < length; i++ {
		var frame SpectateFrame
		frame.ButtonState = o.Byte()
		o.Byte()
		frame.MouseX = o.Float32()
		frame.MouseY = o.Float32()
		frame.Time = o.Int32()
		sl[i] = frame
	}
	return sl
}

// Scorev2Portion is a part of a replay frame that has parts sent conditionally.
type Scorev2Portion struct {
	UsingScorev2 bool
	ComboPortion float64
	BonusPortion float64
}

// Scorev2Portion decodes a Scorev2Portion.
func (o *OsuReader) Scorev2Portion() Scorev2Portion {
	var p Scorev2Portion
	p.UsingScorev2 = o.Bool()
	if p.UsingScorev2 {
		p.ComboPortion = o.Float64()
		p.BonusPortion = o.Float64()
	}
	return p
}

// SanityLimit is the maximum size of a packet.
const SanityLimit = 1024 * 1024 * 5

// Packet reads a packet from the Reader.
func (o *OsuReader) Packet() (uint16, []byte, error) {
	id := o.Uint16()
	o.Byte() // Skip one empty byte
	length := o.Uint32()
	if length > SanityLimit {
		// if it's insane, skip it
		return 0, nil, errors.New("osubinary: insane packet size (more than 5 MB)")
	}
	if length == 0 {
		return id, nil, nil
	}
	return id, o.ByteSlice(int(length)), nil
}

// OsuWriteChain extends binary.WriteChain, providing methods related to the osu!
// binary format.
type OsuWriteChain struct {
	binary.WriteChain
}

// NewWriter creates a new OsuWriteChain.
func NewWriter() *OsuWriteChain {
	buf := &bytes.Buffer{}
	return &OsuWriteChain{
		WriteChain: binary.WriteChain{
			Writer:    buf,
			ByteOrder: binary.LittleEndian,
		},
	}
}

// Bytes retrieves the bytes written. This function may panic if the Writer does not
// satisfy the method `Bytes() []byte`.
func (o *OsuWriteChain) Bytes() []byte {
	fixed, can := o.Writer.(interface {
		Bytes() []byte
	})
	if !can {
		panic("can't call Bytes() on OsuWriteChain.Writer")
	}
	return fixed.Bytes()
}

// Uleb128 writes an uleb128 int into the reader.
func (o *OsuWriteChain) Uleb128(i uint64) *OsuWriteChain {
	b := uleb128.Marshal(i)
	o.ByteSlice(b)
	return o
}

// BanchoString writes a string in the osu!db-format, as in an uleb128 int
// containing the length of the string and the string itself.
func (o *OsuWriteChain) BanchoString(s string) *OsuWriteChain {
	o.Byte('\x11')
	o.Uleb128(uint64(len(s)))
	o.String(s)
	return o
}

// Int32Slice writes a slice of int32s, in the osu! format.
func (o *OsuWriteChain) Int32Slice(s []int32) *OsuWriteChain {
	o.Uint16(uint16(len(s)))
	o.WriteChain.Int32Slice(s)
	return o
}

// Uint32Slice writes a slice of uint32s, in the osu! format.
func (o *OsuWriteChain) Uint32Slice(s []uint32) *OsuWriteChain {
	o.Uint16(uint16(len(s)))
	o.WriteChain.Uint32Slice(s)
	return o
}

// SpectateFrameSlice writes multiple SpectateFrames.
func (o *OsuWriteChain) SpectateFrameSlice(fs []SpectateFrame) *OsuWriteChain {
	o.Uint16(uint16(len(fs)))
	for _, f := range fs {
		o.
			Byte(f.ButtonState).
			Byte(0).
			Float32(f.MouseX).
			Float32(f.MouseY).
			Int32(f.Time)
	}
	return o
}

// Scorev2Portion writes a Scorev2Portion. I'm not in the mood for descriptive
// comments today.
func (o *OsuWriteChain) Scorev2Portion(p Scorev2Portion) *OsuWriteChain {
	o.Bool(p.UsingScorev2)
	if p.UsingScorev2 {
		o.Float64(p.ComboPortion).Float64(p.BonusPortion)
	}
	return o
}

// Packet writes a packet, in the standard form of
//
//   <uint16 PacketID>\x00<uint32 Length><[]byte Data>
func (o *OsuWriteChain) Packet(id uint16, data []byte) *OsuWriteChain {
	o.Uint16(id)
	o.Byte(0)
	o.Uint32(uint32(len(data)))
	o.ByteSlice(data)
	return o
}

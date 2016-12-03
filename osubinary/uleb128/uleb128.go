package uleb128

import "io"

// Marshal converts an int into a uleb128-encoded byte array.
func Marshal(i uint64) (r []byte) {
	var len int
	if i == 0 {
		r = []byte{0}
		return
	}

	for i > 0 {
		r = append(r, 0)
		r[len] = byte(i & 0x7F)
		i >>= 7
		if i != 0 {
			r[len] |= 0x80
		}
		len++
	}

	return
}

// Unmarshal converts a uleb128-encoded byte array into an int.
func Unmarshal(r []byte) (total uint64, len int) {
	var shift uint

	for {
		b := r[len]
		len++
		total |= (uint64(b&0x7F) << shift)
		if b&0x80 == 0 {
			break
		}
		shift += 7
	}

	return
}

// UnmarshalReader unmarshals something with an io.Reader
func UnmarshalReader(r io.Reader) (total uint64) {
	var shift uint
	var lastByte byte

	for {
		b := make([]byte, 1)
		r.Read(b)
		lastByte = b[0]
		total |= (uint64(lastByte&0x7F) << shift)
		if lastByte&0x80 == 0 {
			break
		}
		shift += 7
	}

	return
}

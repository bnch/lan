// Package packets implements reading and writing of packets from bancho.
package packets

// Packet is a generic interface that can be transformed into a Bancho packet
// through the function Packetify.
type Packet interface {
	Packetifier
	Depacketifier
}

// Packetifier wraps around Packetify
type Packetifier interface {
	Packetify() ([]byte, error)
}

// Depacketifier wraps around Depacketify
type Depacketifier interface {
	Depacketify([]byte) error
}

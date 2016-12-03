package handler

import "github.com/bnch/lan/packets"

// ProtocolVersion is the version of the Bancho protocol.
const ProtocolVersion = 19

// Handle takes a set of packets, handles them and then pushes the results
func (s *Session) Handle(pks []packets.Packet) {

}

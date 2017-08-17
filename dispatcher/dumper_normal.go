// +build !banchoreader

package dispatcher

// DumpPackets in normal circumstances does nothing. If you want to enable
// logging of packets using banchoreader, you should build using the tag
// banchoreader.
func DumpPackets(data []byte) {}

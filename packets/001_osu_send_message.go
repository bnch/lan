// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuSendMessage sends a message through the osu! chat.
type OsuSendMessage struct {
	SenderName string // Sender's username, must be replaced server-side
	Content string
	Channel string
	SenderID int32 // Sender's ID, should be replaced server-side by actual user ID
}

// Packetify encodes a OsuSendMessage into
// a byte slice.
func (p OsuSendMessage) Packetify() ([]byte, error) {
	w := ob.NewWriter()

	w.BanchoString(p.SenderName)
	w.BanchoString(p.Content)
	w.BanchoString(p.Channel)
	w.Int32(p.SenderID)

	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a OsuSendMessage.
func (p *OsuSendMessage) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.SenderName = r.BanchoString()
	p.Content = r.BanchoString()
	p.Channel = r.BanchoString()
	p.SenderID = r.Int32()

	_, err := r.End()
	return err
}

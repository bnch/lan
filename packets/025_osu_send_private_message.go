// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

// OsuSendPrivateMessage is a request to send a private message to an user.
// The target is specified in the field "Channel".
type OsuSendPrivateMessage struct {
	SenderName string // Sender's username
	Content string
	Channel string
	SenderID int32 // Sender's ID
}

// Packetify encodes a OsuSendPrivateMessage into
// a byte slice.
func (p OsuSendPrivateMessage) Packetify() ([]byte, error) {
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

// Depacketify decodes a OsuSendPrivateMessage.
func (p *OsuSendPrivateMessage) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

	p.SenderName = r.BanchoString()
	p.Content = r.BanchoString()
	p.Channel = r.BanchoString()
	p.SenderID = r.Int32()

	_, err := r.End()
	return err
}

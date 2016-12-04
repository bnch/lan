// generator generates packet structs and methods based on the packets.txt file.
//
// packets.txt structure:
//
//   - A packet is separated by another with a line of 4 more dashes.
//   - The first line of a package declaration must have the Packet ID and the
//     name, like this:
//   ID: 0, Name: OsuSendUserStatus
//   - The next lines before the structure declaration are considered as the
//     description of the packet, and will be written in the final go file as
//     the documentation of the struct.
//   - The following lines must contain the structure of the Packet, in the
//     following format (description is within brackets because optional):
//   - StructField Type[, Description]
//
// For instance:
//
//   ID: 24, Name: BanchoAnnounce
//
//   BanchoAnnounce is sent by bancho to notify the osu! clients of something.
//
//   - Message string
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	packetSeparator = regexp.MustCompile("^-{4,}$")
	// https://regex101.com/r/sX3eH6/1
	fieldDeclarationReg = regexp.MustCompile(`^- ([a-zA-Z_][a-zA-Z_0-9]*) ((:?\[\])?[a-zA-Z_][a-zA-Z_0-9]*)(?:, (.*))?$`)
)

func main() {
	b := time.Now()
	f, err := os.Open("packets.txt")
	if err != nil {
		fail(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var (
		nextIsDeclaration   = true
		currentIsFieldBlock bool
		currentPacket       *packetDeclaration
	)
	files, err := filepath.Glob("[0-9][0-9][0-9]\\_*.go")
	if err != nil {
		fail(err)
	}
	for _, f := range files {
		os.Remove(f)
	}
	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fail(err)
		}
		if err == io.EOF {
			currentPacket.Export()
			break
		}
		l = strings.TrimSpace(l)
		switch {
		case packetSeparator.MatchString(l):
			nextIsDeclaration = true
			currentIsFieldBlock = false
			currentPacket.Export()
			currentPacket = nil
			continue
		case nextIsDeclaration:
			newPacket := declaration(l)
			if newPacket == nil {
				continue
			}
			currentPacket = newPacket
			nextIsDeclaration = false
		case !currentIsFieldBlock:
			if len(l) > 2 && l[:2] == "- " {
				currentIsFieldBlock = true
				currentPacket.AddField(l)
				continue
			}
			currentPacket.Description += l + "\n"
		default:
			currentPacket.AddField(l)
		}
	}
	buildPacketifyDepacketify()
	fmt.Println("Time:", time.Since(b))
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func declaration(s string) *packetDeclaration {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	var ret packetDeclaration
	parts := strings.Split(s, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		subparts := strings.Split(part, ":")
		if len(subparts) != 2 {
			fail(errors.New("packet declaration must be key: value[, key: value...]"))
		}
		switch strings.ToLower(strings.TrimSpace(subparts[0])) {
		case "id":
			ret.ID, _ = strconv.Atoi(strings.TrimSpace(subparts[1]))
		case "name":
			ret.Name = subparts[1]
		}
	}
	return &ret
}

type packetDeclaration struct {
	ID          int
	Name        string
	Description string
	Fields      []fieldDeclaration
}

type fieldDeclaration struct {
	FieldName   string
	Type        string
	Description string
}

func (p *packetDeclaration) AddField(s string) {
	res := fieldDeclarationReg.FindStringSubmatch(s)
	if len(res) == 0 {
		return
	}
	p.Fields = append(p.Fields, fieldDeclaration{
		FieldName:   strings.TrimSpace(res[1]),
		Type:        strings.TrimSpace(res[2]),
		Description: strings.TrimSpace(res[3]),
	})
}

const baseFileFormat = `// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	ob "github.com/bnch/lan/osubinary"
)

%s
type %s struct {
%s}

// Packetify encodes a %[2]s into
// a byte slice.
func (p %[2]s) Packetify() ([]byte, error) {
	w := ob.NewWriter()

%[4]s
	data := w.Bytes()
	_, err := w.End()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Depacketify decodes a %[2]s.
func (p *%[2]s) Depacketify(b []byte) error {
	r := ob.NewReaderFromBytes(b)

%[5]s
	_, err := r.End()
	return err
}
`

func (p *packetDeclaration) Export() {
	if p == nil {
		return
	}
	p.Name = strings.TrimSpace(p.Name)
	fmt.Printf("Exporting %s... ", p.Name)
	begin := time.Now()
	p.Description = strings.TrimSpace(p.Description)
	filename := fmt.Sprintf("%03d_%s.go", p.ID, toSnake(p.Name))
	content := fmt.Sprintf(
		baseFileFormat,
		commentify(p.Description),
		p.Name,
		fieldsInStruct(p.Fields),
		fieldsInWriter(p.Fields),
		fieldsInReader(p.Fields),
	)
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Now().Sub(begin))
	pairs = append(pairs, idNamePair{
		ID:   p.ID,
		Name: p.Name,
	})
}

func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

var commentifyReplacer = strings.NewReplacer(
	"\n", "\n// ",
	"\r", "",
)

func commentify(s string) string {
	s = strings.TrimSpace(s)
	s = "// " + commentifyReplacer.Replace(s)
	return s
}

func fieldsInStruct(fs []fieldDeclaration) (x string) {
	for _, f := range fs {
		x += fmt.Sprintf("\t%s %s", f.FieldName, f.Type)
		if f.Description != "" {
			x += " " + commentify(f.Description)
		}
		x += "\n"
	}
	return
}
func fieldsInWriter(fs []fieldDeclaration) (x string) {
	for _, f := range fs {
		x += fmt.Sprintf("\tw.%s(p.%s)\n", typeNameToOsuBinaryName(f.Type), f.FieldName)
	}
	return
}
func fieldsInReader(fs []fieldDeclaration) (x string) {
	for _, f := range fs {
		x += fmt.Sprintf("\tp.%s = r.%s()\n", f.FieldName, typeNameToOsuBinaryName(f.Type))
	}
	return x
}

var specialCases = map[string]string{
	"string": "BanchoString",
}

func typeNameToOsuBinaryName(s string) string {
	if c, exist := specialCases[s]; exist {
		return c
	}
	var isSlice bool
	if len(s) > 2 && s[:2] == "[]" {
		s = s[2:]
		isSlice = true
	}
	newS := []byte(s)
	newS[0] = byte(unicode.ToUpper(rune(newS[0])))
	s = string(newS)
	if isSlice {
		s += "Slice"
	}
	return s
}

const basePacketifierDepacketifierFormat = `// THIS FILE HAS BEEN AUTOMATICALLY GENERATED. DO NOT EDIT.
// (modify packets.txt to make changes, run go generate to build)

package packets

import (
	"errors"
	"fmt"
	"io"
	
	"github.com/bnch/lan/osubinary"
)

// Packetify transforms a slice of packets into a slice of bytes, to be
// transmitted to the osu! client.
func Packetify(packets []Packetifier) ([]byte, error) {
	w := osubinary.NewWriter()
	for _, packet := range packets {
		err := packetify(w, packet)
		if err != nil {
			return nil, err
		}
	}
	_, err := w.End()
	return w.Bytes(), err
}

func packetify(w *osubinary.OsuWriteChain, p Packetifier) error {
	data, err := p.Packetify()
	if err != nil {
		return err
	}
	var id uint16
	switch p.(type) {
%s
	default:
		return errors.New("invalid packet")
	}
	w.Packet(id, data)
	return nil
}

// Depacketify decodes a byte slice received from the osu! client into a
// packet slice.
func Depacketify(b []byte) ([]Packet, error) {
	r := osubinary.NewReaderFromBytes(b)
	var packets []Packet
	for {
		id, pack, err := r.Packet()
		if err != nil {
			return nil, err
		}
		_, err = r.End()
		if err == io.EOF {
			return packets, nil
		}
		if err != nil {
			return nil, err
		}
		packet, err := depacketify(id, pack)
		if err != nil {
			return nil, err
		}
		if packet != nil {	
			packets = append(packets, packet)
		}
	}
}

func depacketify(id uint16, packet []byte) (Packet, error) {
	var p Packet
	switch id {
%s
	default:
		fmt.Printf("Asked to depacketify an unknown packet: %%d\n", int(id))
		return nil, nil // errors.New("invalid packet ID (" + strconv.Itoa(int(id)) + ")")
	}
	err := p.Depacketify(packet)
	return p, err
}
`

type idNamePair struct {
	ID   int
	Name string
}

func (p idNamePair) forPacketify() string {
	return fmt.Sprintf("\tcase *%s: id = %d\n", p.Name, p.ID)
}
func (p idNamePair) forDepacketify() string {
	return fmt.Sprintf("\tcase %d: p = &%s{}\n", p.ID, p.Name)
}

var pairs []idNamePair

func buildPacketifyDepacketify() {
	var (
		_p string
		dp string
	)
	for _, p := range pairs {
		_p += p.forPacketify()
		dp += p.forDepacketify()
	}
	ioutil.WriteFile("dp_p.go", []byte(fmt.Sprintf(basePacketifierDepacketifierFormat, _p, dp)), 0644)
}

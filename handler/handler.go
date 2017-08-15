package handler

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/bnch/lan/packets"
	"gopkg.in/redis.v5"
)

// Redis client that will be used to fetch and retrieve information
var Redis *redis.Client

// Handle takes a set of packets, handles them and then pushes the results
func (s Session) Handle(pks []packets.Packet) {
	s.LastSeen = time.Now()
	SaveSession(s)

	for _, p := range pks {
		go s.handle(p)
	}
}

func (s Session) handle(p packets.Packet) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Error while handling a %T: %v\n%#v\n", p, err, p)
		}
	}()
	fmt.Printf("< %#v\n", p)

	f := handlers[reflect.TypeOf(p).Elem().Name()]

	// function not found
	if !f.IsValid() {
		fmt.Printf("> couldn't handle packet of type %T\n", p)
		return
	}

	f.Call([]reflect.Value{
		reflect.ValueOf(p),
		reflect.ValueOf(s),
	})
}

// At the time of writing this, there are 109 known osu! packets.
var handlers = make(map[string]reflect.Value, 109)

// apparently, this is the way you do it. Don't ask me why.
// https://golang.org/pkg/reflect/#TypeOf
var packetType = reflect.TypeOf((*packets.Packet)(nil)).Elem()
var sessionType = reflect.TypeOf(Session{})

// RegisterHandler registers a handler function.
//
// Handler functions are simply in the signature func (p *packets.<something>).
func RegisterHandler(f interface{}) {
	val := reflect.ValueOf(f)
	t := val.Type()

	// if the function doesn't have exactly two input parameter, then we can't
	// register it.
	if t.NumIn() != 2 {
		panic(errors.New("handler.RegisterHandler: function doesn't have exactly 1 input argument"))
	}

	// make sure the packet argument is a valid Packet
	argPacket := t.In(0)
	if !argPacket.Implements(packetType) {
		panic(errors.New("handler.RegisterHandler: function argument is not a packets.Packet"))
	}

	if !t.In(1).AssignableTo(sessionType) {
		panic(errors.New("handler.RegisterHandler: second argument is not a Session"))
	}

	// register the function in the handlers
	handlers[argPacket.Elem().Name()] = val
}

// RegisterHandlers simply executes RegisterHandler for many handlers.
func RegisterHandlers(handlers ...interface{}) {
	for _, h := range handlers {
		RegisterHandler(h)
	}
}

func init() {
	RegisterHandler(func(*packets.OsuPong, Session) {})
}

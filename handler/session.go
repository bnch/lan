package handler

import (
	"sync"
	"sync/atomic"

	"github.com/bnch/lan/packets"
)

const packetChanBufSize = 500

// AdminPassword is md5 hash of the password.
var AdminPassword string

var userID int32 = 49

// Session is an user's session on Bancho.
type Session struct {
	Username string
	UserID   int32
	Token    string
	Admin    bool
	Packets  chan packets.Packet
	Mutex    sync.RWMutex
}

// Send sends a packet down the session.
func (s *Session) Send(ps ...packets.Packet) {
	for _, p := range ps {
		s.send(p)
	}
}

// Permissions converts the bool Admin to osu!'s permissions.
func (s *Session) Permissions() int32 {
	if s.Admin {
		// all permissions
		return 63
	}
	// supporter
	return 5
}

func (s *Session) send(p packets.Packet) {
	// We want this function to return immediately.
	// So if s.Packets is full, we'll fire a goroutine to wait for it.
	select {
	case s.Packets <- p:
		return
	default:
		go func() {
			s.Packets <- p
		}()
	}
}

// NewSession creates a new session.
func NewSession(username string, passMD5 string) (s *Session) {
	Sessions.Lock()
	for _, el := range Sessions {
		if el.Username == username {
			return &Session{
				UserID:  authFailed,
				Token:   GenerateGUID(),
				Packets: make(chan packets.Packet, 5),
			}
		}
	}
	Sessions.Unlock()

	// user not found, all good

	// get userid
	uid := atomic.AddInt32(&userID, 1)
	newSess := &Session{
		Username: username,
		UserID:   uid,
		Token:    GenerateGUID(),
		Admin:    passMD5 == AdminPassword,
		Packets:  make(chan packets.Packet, packetChanBufSize),
	}
	Sessions.Add(newSess)
	return newSess
}

// SessionCollection is simply a []*Session, with a few helper methods.
type SessionCollection []*Session

// Lock gains exclusive write access to the collection.
func (s SessionCollection) Lock() {
	sessionMutex.Lock()
}

// Unlock removes the exclusive write access to the collection.
func (s SessionCollection) Unlock() {
	sessionMutex.Unlock()
}

// RLock gains exclusive write access to the collection.
func (s SessionCollection) RLock() {
	sessionMutex.RLock()
}

// RUnlock removes the exclusive write access to the collection.
func (s SessionCollection) RUnlock() {
	sessionMutex.RUnlock()
}

// Copy gains a copy of a SessionCollection, unreferenced from the master
// SessionCollection.
func (s SessionCollection) Copy() SessionCollection {
	s.RLock()
	defer s.RUnlock()
	x := make(SessionCollection, len(s))
	copy(x, s)
	return x
}

// GetByUsername retrieves a Session by its user's username.
func (s SessionCollection) GetByUsername(u string) (sess *Session) {
	s.RLock()
	defer s.RUnlock()
	for _, el := range s {
		el.Mutex.RLock()
		if el.Username == u {
			sess = el
		}
		el.Mutex.RUnlock()
		if sess != nil {
			return sess
		}
	}
	return sess
}

// GetByID retrieves a Session by its user's ID.
func (s SessionCollection) GetByID(u int32) (sess *Session) {
	s.RLock()
	defer s.RUnlock()
	for _, el := range s {
		el.Mutex.RLock()
		if el.UserID == u {
			sess = el
		}
		el.Mutex.RUnlock()
		if sess != nil {
			return sess
		}
	}
	return sess
}

// GetByToken retrieves a Session by its user's token.
func (s SessionCollection) GetByToken(u string) (sess *Session) {
	s.RLock()
	defer s.RUnlock()
	for _, el := range s {
		el.Mutex.RLock()
		if el.Token == u {
			sess = el
		}
		el.Mutex.RUnlock()
		if sess != nil {
			return
		}
	}
	return
}

// Add adds a Session to the collection.
func (s SessionCollection) Add(sess *Session) {
	s.Lock()
	defer s.Unlock()
	// check an user with same username/userid/token exists
	for _, el := range s {
		var alreadyFound bool
		el.Mutex.RLock()
		if el.Username == sess.Username ||
			el.UserID == sess.UserID ||
			el.Token == sess.Token {
			alreadyFound = true
		}
		el.Mutex.RUnlock()
		if alreadyFound {
			return
		}
	}
	// not found, add it
	s = append(s, sess)
}

// Delete removes a Session from the collection.
func (s SessionCollection) Delete(token string) {
	s.Lock()
	defer s.Unlock()
	for i, el := range s {
		var isthis bool
		el.Mutex.RLock()
		if el.Token == token {
			isthis = true
		}
		el.Mutex.RUnlock()
		if isthis {
			s[i] = s[len(s)-1]
			s[len(s)-1] = nil
			s = s[:len(s)-1]
			close(el.Packets)
			return
		}
	}
}

// Sessions is a slice containing all the sessions of the online users.
var Sessions SessionCollection
var sessionMutex = &sync.RWMutex{}

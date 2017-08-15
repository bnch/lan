package handler

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bnch/lan/packets"
)

const packetChanBufSize = 500

// Sessions is a slice containing all the sessions of the online users.
var Sessions = SessionCollection("")

// AdminPassword is md5 hash of the password.
var AdminPassword string

// Session is an user's session on Bancho.
type Session struct {
	Username string
	UserID   int32
	Token    string
	Admin    bool
	LastSeen time.Time
	State    packets.OsuSendUserState
}

// Send sends a packets to the session.
func (s *Session) Send(ps ...packets.Packet) {
	packetifiers := make([]packets.Packetifier, len(ps))
	for i, p := range ps {
		packetifiers[i] = packets.Packetifier(p)
	}
	b, err := packets.Packetify(packetifiers)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.SendBytes(b)
}

// SendBytes sends bytes to the session to be send to the client. These are
// usually packetified packets.
func (s *Session) SendBytes(b []byte) {
	Redis.RPush("lan/queues/"+s.Token, string(b))
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

// Subscribe adds the Session to a stream in streams.
func (s Session) Subscribe(stream string) bool {
	if s.In(stream) {
		return false
	}
	SessionCollection(stream).Add(s)
	Redis.SAdd("lan/my_collections/"+s.Token, stream)
	return true
}

// Unsubscribe removes the user from a stream.
func (s Session) Unsubscribe(stream string) {
	SessionCollection(stream).Delete(s)
	Redis.SRem("lan/my_collections/"+s.Token, stream)
}

// In checks whether a Session is in a stream.
func (s Session) In(stream string) bool {
	return Redis.SIsMember("lan/my_collections/"+s.Token, stream).Val()
}

// Dispose finalizes the session removing all references to it
// (including Sessions) and setting s.disposed to true.
func (s Session) Dispose() {
	Redis.Del("lan/queues/" + s.Token)
	Redis.Del("lan/sessions/" + s.Token)
	Sessions.Delete(s)
	for _, x := range Redis.SMembers("lan/my_collections/" + s.Token).Val() {
		SessionCollection(x).Delete(s)
	}
	Redis.Del("lan/my_collections/" + s.Token)
}

// NewSession creates a new session.
func NewSession(username string, passMD5 string) Session {
	if Sessions.TokenFromUsername(username) != "" {
		return Session{
			UserID: authFailed,
			Token:  GenerateGUID(),
		}
	}

	// user not found, all good

	if Redis.Get("lan/user_id").Val() == "" {
		Redis.Set("lan/user_id", 49, 0)
	}

	newSess := Session{
		Username: username,
		UserID:   int32(Redis.Incr("lan/user_id").Val()),
		Token:    GenerateGUID(),
		Admin:    passMD5 == AdminPassword,
	}
	Sessions.Add(newSess)
	SaveSession(newSess)
	return newSess
}

// SessionCollection is simply a string, and it mainly indicates a redis list
// of tokens.
type SessionCollection string

// AllTokens retrieves all the tokens in the session collection.
func (s SessionCollection) AllTokens() []string {
	return Redis.SMembers("lan/session_list/" + string(s)).Val()
}

// TokenFromUsername retrieves a token of a session using its user's username.
func (s SessionCollection) TokenFromUsername(u string) string {
	return Redis.HGet("lan/session_lookup_tables/username/"+string(s), u).Val()
}

// TokenFromID retrieves a Session by its user's ID.
func (s SessionCollection) TokenFromID(u int32) string {
	return Redis.HGet("lan/session_lookup_tables/id/"+string(s), strconv.Itoa(int(u))).Val()
}

// Add adds a Session to the collection.
func (s SessionCollection) Add(sess Session) {
	if Redis.SAdd("lan/session_list/"+string(s), sess.Token).Val() == 0 {
		// Value already present
		return
	}
	Redis.HSet("lan/session_lookup_tables/username/"+string(s), sess.Username, sess.Token)
	Redis.HSet("lan/session_lookup_tables/id/"+string(s), strconv.Itoa(int(sess.UserID)), sess.Token)
}

// Delete removes a Session from the collection.
func (s SessionCollection) Delete(sess Session) {
	Redis.SRem("lan/session_list/"+string(s), sess.Token)
	Redis.HDel("lan/session_lookup_tables/username/"+string(s), sess.Username)
	Redis.HDel("lan/session_lookup_tables/id/"+string(s), strconv.Itoa(int(sess.UserID)))
}

// AllUserIDs fetches all the user IDs in the SessionCollection.
func (s SessionCollection) AllUserIDs() []int32 {
	m := Redis.HGetAll("lan/session_lookup_tables/id/" + string(s)).Val()

	ids := make([]int32, len(m))
	i := 0
	for v := range m {
		id, _ := strconv.ParseInt(v, 10, 32)
		ids[i] = int32(id)
		i++
	}

	return ids
}

// GetSession retrieves a Session by knowing its token.
func GetSession(token string) (sess *Session) {
	val := Redis.Get("lan/sessions/" + token).Val()
	if val == "" {
		return nil
	}
	sess = new(Session)
	gob.NewDecoder(strings.NewReader(val)).Decode(sess)
	return sess
}

// SaveSession saves the session currently being handled.
func SaveSession(sess Session) {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(sess)
	Redis.Set("lan/sessions/"+sess.Token, buf.String(), 0)
}

// Send sends packets to all members of the SessionCollection
func (s *SessionCollection) Send(ps ...packets.Packet) {
	packetifiers := make([]packets.Packetifier, len(ps))
	for i, p := range ps {
		packetifiers[i] = packets.Packetifier(p)
	}
	b, err := packets.Packetify(packetifiers)
	if err != nil {
		return
	}

	c := s.AllTokens()
	for _, e := range c {
		Redis.RPush("lan/queues/"+e, string(b))
	}
}

// SendExcept sends packets to all members of the SessionCollection, except for
// those with the ID specified.
func (s SessionCollection) SendExcept(except []int32, ps ...packets.Packet) {
	packetifiers := make([]packets.Packetifier, len(ps))
	for i, p := range ps {
		packetifiers[i] = packets.Packetifier(p)
	}
	b, err := packets.Packetify(packetifiers)
	if err != nil {
		return
	}

	m := Redis.HGetAll("lan/session_lookup_tables/id/" + string(s)).Val()

SessionLooper:
	for k, v := range m {
		for _, x := range except {
			if strconv.Itoa(int(x)) == k {
				continue SessionLooper
			}
		}

		Redis.RPush("lan/queues/"+v, string(b))
	}
}

// Len returns the length of the SessionCollection's slice.
func (s SessionCollection) Len() int {
	return int(Redis.SCard("lan/session_list/" + string(s)).Val())
}

// disposer automatically disposes sessions older than 120 seconds.
func disposer() {
	for {
		time.Sleep(time.Second * 10)
		cp := Sessions.AllTokens()
		for _, c := range cp {
			sess := GetSession(c)
			if sess == nil {
				continue
			}
			if time.Now().Sub(sess.LastSeen) > time.Second*120 {
				sess.Dispose()
				Sessions.Send(&packets.BanchoUserQuit{ID: sess.UserID, State: 0})
			}
		}
	}
}

func init() {
	go disposer()
	gob.Register(Session{})
}

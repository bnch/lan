package dispatcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bnch/lan/handler"
	"github.com/bnch/lan/packets"
)

func panicRecover() {
	err := recover()
	if err != nil {
		fmt.Printf("Critical error! %v\n", err)
	}
}

// exportQueue takes from the queue the packets it can fetch in 50 milliseconds.
func exportQueue(token string) []byte {
	parts := make([]byte, 0, 1024*64)

	p := handler.Redis.BLPop(time.Millisecond*100, "lan/queues/"+token).Val()
	if len(p) < 2 {
		return nil
	}
	parts = append(parts, []byte(p[1])...)

	for {
		p := handler.Redis.LPop("lan/queues/" + token).Val()
		if p == "" {
			break
		}
		parts = append(parts, []byte(p)...)
	}
	return parts
}

func (s Server) bancho(w http.ResponseWriter, r *http.Request) {
	defer panicRecover()

	if r.Method != "POST" || r.UserAgent() != "osu!" {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("lan\ngithub.com/bnch/lan\n"))
		return
	}

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header()["cho-protocol"] = []string{strconv.Itoa(handler.ProtocolVersion)}

	tok := r.Header.Get("osu-token")
	var sess handler.Session
	if tok == "" {
		br := bufio.NewReader(r.Body)
		u, _ := br.ReadString('\n')
		p, _ := br.ReadString('\n')
		sess = handler.Authenticate(strings.TrimSpace(u), strings.TrimSpace(p))
	} else {
		sess = s.banchoHandle(r)
	}

	if tok == "" {
		w.Header()["cho-token"] = []string{sess.Token}
	}

	w.Write(exportQueue(sess.Token))
}

func (s Server) banchoHandle(r *http.Request) handler.Session {
	sessPtr := handler.GetSession(r.Header.Get("osu-token"))
	if sessPtr == nil {
		return handler.LogoutTokenNotFound()
	}
	sess := *sessPtr
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading post body: %v\n", err)
		return handler.LogoutTokenNotFound()
	}
	pks, err := packets.Depacketify(d)
	if err != nil {
		fmt.Printf("Error depacketifying: %v\n", err)
		return handler.LogoutTokenNotFound()
	}
	sess.Handle(pks)
	return sess
}

package dispatcher

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"fmt"

	"github.com/bnch/banchoreader/lib"
	"github.com/bnch/lan/handler"
	"github.com/bnch/lan/packets"
)

func (s Server) bancho(w http.ResponseWriter, r *http.Request) {
	//begin := time.Now()

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

	parts := make([]byte, 0, 1024*64)
	for {
		p := handler.Redis.LPop("lan/queues/" + sess.Token).Val()
		if p == "" {
			break
		}
		parts = append(parts, []byte(p)...)
	}

	w.Write(parts)

	//os.Stdout.WriteString("=> request done in " + time.Now().Sub(begin).String() + "\n")
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

var dumper = banchoreader.New()

func panicRecover() {
	err := recover()
	if err != nil {
		fmt.Printf("Critical error! %v\n", err)
	}
}

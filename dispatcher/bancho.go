package dispatcher

import (
	"bufio"
	"net/http"
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/bnch/banchoreader/lib"
	"github.com/bnch/lan/handler"
	"github.com/bnch/lan/packets"
)

func (s Server) bancho(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()

	if r.Method != "POST" || r.UserAgent() != "osu!" {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("lan - a private bancho server designed for playing in the local network\n"))
		return
	}

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header()["cho-protocol"] = []string{strconv.Itoa(handler.ProtocolVersion)}

	tok := r.Header.Get("osu-token")
	var sess *handler.Session
	if tok == "" {
		br := bufio.NewReader(r.Body)
		u, _ := br.ReadString('\n')
		p, _ := br.ReadString('\n')
		sess = handler.Authenticate(u, p)
	} else {
		sess = handler.Sessions.GetByToken(tok)
		if sess == nil {
			sess = handler.LogoutTokenNotFound()
		}
	}

	if tok == "" {
		sess.Mutex.RLock()
		w.Header()["cho-token"] = []string{sess.Token}
		sess.Mutex.RUnlock()
	}

	pks := make([]packets.Packetifier, 0, 10)
Looper:
	for {
		select {
		case x := <-sess.Packets:
			pks = append(pks, packets.Packetifier(x))
		default:
			break Looper
		}
	}

	encoded, err := packets.Packetify(pks)
	if err != nil {
		fmt.Println(err)
	}
	dumper.Dump(os.Stdout, encoded)
	w.Write(encoded)

	os.Stdout.WriteString("=> request done in " + time.Now().Sub(begin).String() + "\n")
}

var dumper = banchoreader.New()

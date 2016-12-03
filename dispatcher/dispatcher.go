// Package dispatcher handles interfacing with the osu! client, and sending
// enqueued packets to it.
package dispatcher

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Server is the HTTP server that handles the Bancho requests.
type Server struct{}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// ServeHTTP handles the requests.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IBMD")
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		s.switcher(w, r)
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	s.switcher(gzr, r)
}

func (s Server) switcher(w http.ResponseWriter, r *http.Request) {
	switch r.Host {
	case "c.ppy.sh", "c1.ppy.sh", "c2.ppy.sh":
		s.bancho(w, r)
	case "a.ppy.sh":
		// temporary
		w.WriteHeader(404)
	case "osu.ppy.sh":
		scoreServer.ServeHTTP(w, r)
	default:
		w.Write([]byte("lan"))
	}
}

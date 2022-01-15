//
// HTTP Server
//

package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/eserilev/migration.winc.services/corporate"
)

type Handler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

func Start() {
}

type HTTPServer struct{}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: normalize path. ends in a slash, strip the slash.
	//
	switch p := r.URL.Path; p {
	case "/status/run":
		s.statusRunHandler(w, r)
	case "/status":
		s.statusHandler(w, r)
	default:
		if r.Method == "POST" {
			if strings.HasPrefix(p, "/__s/v1/orders/corporate") {
				corporate.ServeHTTP(w, r)
			} else {
				s.defaultHandler(w, r)
			}
		}
	}
}

func RegisterPaths(paths []string, h *Handler) {
	fmt.Println("Register paths", paths, h)
}

//
// TODO: Move these to config
//

func AddDefaultResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Server", "Winc")
	w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	w.Header().Set("Date", "Mon, 01 Jan 1990 00:00:00 GMT")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Ua-Compatible", "IE=Edge")
	w.Header().Set("X-Xss-Protection", "1; mode=block")
}

func (s *HTTPServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Status: ", r.Method, r.URL.Path, r.Header["X-Forwarded-Host"])
	fmt.Println("Status 2: ", r.Header)
	// fmt.Println("Current Hostname: ", config.CurrentHostname(r))
	w.WriteHeader(200)
	AddDefaultResponseHeaders(w)
	io.WriteString(w, r.URL.Path)
}

func (s *HTTPServer) statusRunHandler(w http.ResponseWriter, r *http.Request) {
	// run := config.RunId()
	// fmt.Println("Run Id: ", run)
	w.WriteHeader(200)
	AddDefaultResponseHeaders(w)
	// io.WriteString(w, run)
}

func (s *HTTPServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	io.WriteString(w, "Bad Request")
}

//
// Corporate Orders
//

package corporate

import (
	"io"
	"net/http"
	"strings"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//
	// TODO: normalize path. ends in a slash, strip the slash.
	//
	p := r.URL.Path

	if !strings.HasPrefix(p, "/__s/v1/orders/corporate/") {
		handleBadRequest(w, r)
		return
	}

	ProcessCorporateOrders(r)

	handleBadRequest(w, r)
}

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

func handleBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	io.WriteString(w, "Bad Request")
}

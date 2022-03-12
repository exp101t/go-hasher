package middlewares

import (
	"net/http"
	"strings"
)

func getContentType(r *http.Request) string {
	header := r.Header.Get("Content-Type")

	if i := strings.IndexRune(header, ';'); i != -1 {
		header = header[0:i]
	}

	return header
}

func JsonHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "PUT" || r.Method == "POST" || r.Method == "PATCH") &&
			!(getContentType(r) == "application/json") {
			http.Error(w, "Unsupported Content Type", http.StatusUnsupportedMediaType)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

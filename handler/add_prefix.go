package handler

import (
	"net/http"
	"net/url"
)

// AddPrefix creates a new Handler with a prefix appended to all requests
//
// AddPrefix is symmetrical to http.StripPrefix
func AddPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = prefix + r.URL.Path
		h.ServeHTTP(w, r2)
	})
}

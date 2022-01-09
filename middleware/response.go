package middleware

import (
	"net/http"
	"os"
)

type ResponseHeader struct {
	http.Handler
}

func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		w.Header()[k] = v
	}
	version := os.Getenv("VERSION")
	w.Header().Add("X-Server-Version", version)
	w.Header().Add("Content-Type", "application/json")
	rh.Handler.ServeHTTP(w, req)
}

func NewResponseHeader(handler http.Handler) *ResponseHeader {
	return &ResponseHeader{Handler: handler}
}

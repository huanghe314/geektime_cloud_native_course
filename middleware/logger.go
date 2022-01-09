package middleware

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type EnhancedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (ew *EnhancedWriter) WriteHeader(code int) {
	ew.statusCode = code
	ew.ResponseWriter.WriteHeader(code)
}

func (ew *EnhancedWriter) Code() int {
	return ew.statusCode
}

func NewEnhancedWriter(rw http.ResponseWriter) *EnhancedWriter {
	return &EnhancedWriter{
		ResponseWriter: rw,
	}
}

type Logger struct {
	http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ew := NewEnhancedWriter(w)
	l.serveHTTP(ew, r)
}

func (l *Logger) serveHTTP(w *EnhancedWriter, r *http.Request) {
	// log the client ip
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	log.Printf("remoteAddr: %s\n", r.RemoteAddr)
	if err != nil {
		_, err = fmt.Fprintf(w, "userip: %s is not IP:port", r.RemoteAddr)
		if err != nil {
			log.Printf("write to responseWrite has err: %e", err)
			return
		}
	}
	userIP := net.ParseIP(ip)
	if userIP == nil { // not valid IP Address
		_, err = fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
		if err != nil {
			log.Printf("write to responseWrite has err: %e", err)
			return
		}
		return
	}

	forward := r.Header.Get("X-Forwarded-For")

	log.Printf("IP: %s\n", ip)
	log.Printf("Port: %s\n", port)
	log.Printf("Forwared: %s\n", forward)

	l.Handler.ServeHTTP(w, r)

	// log response code
	log.Printf("code: %d\n", w.Code())
}

func NewLogger(h http.Handler) *Logger {
	return &Logger{h}
}

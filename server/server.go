package server

import (
	"github.com/huanghe314/geektime_cloud_native_course/middleware"
	"log"
	"net/http"
)

// homework 1: implement a http server in go

const (
	localAddr = ":8081"
)

func Serve() {
	// initialize serveMux
	mux := http.NewServeMux()

	// handler registration
	mux.HandleFunc("/a", RootHandler)
	mux.HandleFunc("/healthz", HealthHandler)

	// chain middlewares
	wrappedMux := middleware.NewLogger(middleware.NewResponseHeader(mux))

	err := http.ListenAndServe(localAddr, wrappedMux)
	if err != nil {
		log.Fatalf("fatal err from ListenAndServe: %e", err)
	}
	log.Printf("Http Server is Listening on: %s", localAddr)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

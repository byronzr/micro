package main

import (
	"net/http"
	"time"

	. "micro/fundational"
	"micro/handlers"
)

func Start() {
	Inf("service start.")
	mux := http.NewServeMux()
	mux.Handle("/", handlers.ROUTER{})
	server := &http.Server{
		Addr:         ":8000",
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}
	server.ListenAndServe()
}

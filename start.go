package micro

import (
	"net/http"
	"time"

	. "github.com/byronzr/micro/fundational"
	"github.com/byronzr/micro/handlers"
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

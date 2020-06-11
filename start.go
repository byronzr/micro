package micro

import (
	"fmt"
	"net/http"
	"time"

	"github.com/byronzr/micro/helper"
)

type SERVICE struct {
	Mux *http.ServeMux
}

func Register(hands ...interface{}) *SERVICE {
	s := &SERVICE{}
	if len(hands) == 0 {
		helper.Inf("not handler register. service shutdown.")
	}
	for _, h := range hands {
		helper.RegisterHandler(h)
	}
	helper.Inf("service start.")
	s.Mux = http.NewServeMux()
	s.Mux.Handle("/", helper.ROUTER{})
	return s
}

func (s *SERVICE) Start(port, timeout int) {
	pstr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         pstr,
		WriteTimeout: time.Second * time.Duration(timeout),
		Handler:      s.Mux,
	}
	server.ListenAndServe()
}

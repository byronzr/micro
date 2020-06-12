package micro

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/byronzr/micro/helper"
)

type SERVICE struct {
	Mux        *http.ServeMux
	PrefixPath string
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

func (s *SERVICE) Prefix(p string) *SERVICE {
	prefix := []byte{}
	if []byte(p)[0] != '/' {
		prefix = []byte{'/'}
	}
	prefix = append(prefix, []byte(p)...)
	s.PrefixPath = strings.ToLower(string(prefix))
	return s
}

func (s *SERVICE) Start(port, timeout int) {
	for uri, _ := range helper.ActionFuncMap {
		helper.Inf(">> registered >> ", uri)
	}
	pstr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         pstr,
		WriteTimeout: time.Second * time.Duration(timeout),
		Handler:      s.Mux,
	}
	log.Fatal(server.ListenAndServe())

}

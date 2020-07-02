package micro

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/byronzr/micro/helper"
)

type SERVICE struct {
	Mux *http.ServeMux
}

// register handlers
func Register(hands ...interface{}) *SERVICE {
	s := &SERVICE{}
	if len(hands) == 0 {
		helper.Inf("not handler register. service shutdown.")
	}
	for _, h := range hands {
		helper.RegisterHandler(h)
	}
	s.Mux = http.NewServeMux()
	s.Mux.Handle("/", helper.ROUTER{})
	return s
}

// global before call
func (s *SERVICE) Before(f func(*http.Request) (interface{}, bool)) *SERVICE {
	helper.MiddleFuncMap["GLOBAL_BEFORE"] = f
	helper.Wrn(">> MIDDLE BEFORE >> GLOBAL ")
	return s
}

// global after call
func (s *SERVICE) After(f func(*http.Request) (interface{}, bool)) *SERVICE {
	helper.MiddleFuncMap["GLOBAL_AFTER"] = f
	helper.Wrn(">> MIDDLE AFTER >> GLOBAL ")
	return s
}

// service start
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
	helper.Inf(":::::: service start ::::::")
	log.Fatal(server.ListenAndServe())

}

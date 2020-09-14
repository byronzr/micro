package micro

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/modood/table"
)

type SERVICE struct {
	Mux *http.ServeMux
}

type TableInfo struct {
	Type     string
	Register string
}

// register handlers
func Register(hands ...interface{}) *SERVICE {
	s := &SERVICE{}
	if len(hands) == 0 {
		Inf("not handler register. service shutdown.")
	}
	for _, h := range hands {
		RegisterHandler(h)
	}
	s.Mux = http.NewServeMux()
	s.Mux.Handle("/", ROUTER{})
	return s
}

// default router "/"
func (s *SERVICE) Default(f func(*MicroRequest) int) *SERVICE {
	ActionFuncMap["_DEFAULT_"] = f
	return s
}

// global before call
func (s *SERVICE) Before(f func(*MicroRequest) bool) *SERVICE {
	MiddleFuncMap["GLOBAL before"] = f
	return s
}

// global after call
func (s *SERVICE) After(f func(*MicroRequest) bool) *SERVICE {
	MiddleFuncMap["GLOBAL after"] = f
	return s
}

// service start
func (s *SERVICE) Start(port, timeout int) {

	inf := []TableInfo{}

	for mid, _ := range MiddleFuncMap {
		inf = append(inf, TableInfo{"MIDDLE", mid})
	}

	for uri, _ := range ActionFuncMap {
		inf = append(inf, TableInfo{"ACTION", uri})
	}

	table.Output(inf)

	pstr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         pstr,
		WriteTimeout: time.Second * time.Duration(timeout),
		Handler:      s.Mux,
	}

	msg := fmt.Sprintf(":::::: service [ %d ] start ::::::", port)
	Inf(msg)
	log.Fatal(server.ListenAndServe())

}

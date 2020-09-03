package micro

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/modood/table"
	"github.com/byronzr/micro/helper"
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
func (s *SERVICE) Before(f func(*helper.MicroRequest) bool) *SERVICE {
	helper.MiddleFuncMap["GLB before"] = f
	return s
}

// global after call
func (s *SERVICE) After(f func(*helper.MicroRequest) bool) *SERVICE {
	helper.MiddleFuncMap["GLB after"] = f
	return s
}

// service start
func (s *SERVICE) Start(port, timeout int) {

	inf := []TableInfo{}

	for mid, _ := range helper.MiddleFuncMap {
		inf = append(inf, TableInfo{"MIDDLE", mid})
	}

	for uri, _ := range helper.ActionFuncMap {
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
	helper.Inf(msg)
	log.Fatal(server.ListenAndServe())

}

package micro

import (
	"fmt"
	"net/http"
	"time"

	"github.com/byronzr/micro/helper"
)

// type GET struct{}
// type HEAD struct{}
// type POST struct{}
// type PUT struct{}
// type PATCH struct{}
// type DELETE struct{}
// type TRACE struct{}
// type OPTIONS struct{}

// type Response struct {
// 	Result interface{} `json:"result"`
// }

// func main() {
// 	Start(GET{}, POST{})
// }

// func (GET) Check(r *http.Request) (response []byte, err error) {
// 	msg := "GET.Check"
// 	return json.Marshal(&Response{Result: msg})
// }

// func (POST) Check(r *http.Request) (response []byte, err error) {
// 	msg := "POST.Check"
// 	return json.Marshal(&Response{Result: msg})
// }
type SERVICE struct {
	Mux *http.ServeMux
}

var (
	S = &SERVICE{}
)

func (s *SERVICE) Register(hands ...interface{}) *SERVICE {
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

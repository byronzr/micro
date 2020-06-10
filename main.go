package micro

import (
	"net/http"
	"time"

	"micro/helper"
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

func Start(hands ...interface{}) {
	if len(hands) == 0 {
		helper.Inf("not handler register. service shutdown.")
	}
	for _, h := range hands {
		helper.RegisterHandler(h)
	}
	helper.Inf("service start.")
	mux := http.NewServeMux()
	mux.Handle("/", helper.ROUTER{})
	server := &http.Server{
		Addr:         ":8000",
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	server.ListenAndServe()
}

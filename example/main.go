package main

import (
	"test/handlers"

	"github.com/byronzr/micro"
)

// request method set
//type GET struct{}
// type HEAD struct{}
// type PUT struct{}
// type PATCH struct{}
// type DELETE struct{}
// type TRACE struct{}
// type OPTIONS struct{}

func main() {
	micro.S.Register(handlers.POST{}, handlers.OPTIONS{}).Start(8000)
}

// func (GET) Check(r *http.Request) (response []byte, err error) {
// 	msg := "GET.Check"
// 	return json.Marshal(msg)
// }

// func (GET) ReportThisweek(r *http.Request) (response []byte, err error) {
// 	msg := "GET.ReprotThisweek."
// 	return json.Marshal(&Response{Result: msg})
// }

// func (POST) ReportThisweek(r *http.Request) (response []byte, err error) {
// 	msg := "POST.ReprotThisweek."
// 	return json.Marshal(&Response{Result: msg})
// }

// func (OPTIONS) ReportThisweek(r *http.Request) (response []byte, err error) {
// 	msg := "OPTIONS.ReprotThisweek."
// 	return json.Marshal(&Response{Result: msg})
// }

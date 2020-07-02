package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/handlers"

	"github.com/byronzr/micro"
)

// request method set
type GET struct{}

// type HEAD struct{}
// type PUT struct{}
// type PATCH struct{}
// type DELETE struct{}
// type TRACE struct{}
// type OPTIONS struct{}

func main() {
	// example 1
	// no prefix start
	// micro.Register(handlers.POST{}, handlers.OPTIONS{}).Start(8000, 10)

	// example 2
	// has prefix start
	// micro.Register(handlers.POST{}, handlers.OPTIONS{}).Prefix("byron").Start(8000, 10)

	// example 3
	// not chan call
	service := micro.Register(GET{}, "byron", handlers.POST{}, "wl", handlers.OPTIONS{}) // must be first

	// add middle hook on before
	before := func(r *http.Request) (interface{}, bool) {
		fmt.Println("i'm running before <global>")
		return nil, true
	}
	service.Before(before)

	// add middle hook on before
	after := func(r *http.Request) (interface{}, bool) {
		fmt.Println("i'm running after <global>")
		return nil, true
	}
	service.After(after)

	service.Start(8000, 10)
	// INF 2020/06/12 14:38:04 service start.
	// INF 2020/06/12 14:38:04 >> registered >> POST /byron/report/thisweek
	// INF 2020/06/12 14:38:04 >> registered >> OPTIONS /wl/report/this/week
}

func (GET) Check(r *http.Request) (response []byte, err error) {
	msg := "GET.Check"
	return json.Marshal(msg)
}

func (GET) Before(r *http.Request) (interface{}, bool) {
	fmt.Println("i'm running before <get>")
	return nil, true
}

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

// func Test_main(t *testing.T) {
// 	str := "REPORTThisWEEK"
// 	fixNameRe := regexp.MustCompile(`([A-Z]+?)`)
// 	fixName := fixNameRe.FindStringSubmatch(str)
// 	fmt.Println(fixName)

// }

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type POST struct{}
type OPTIONS struct{}

func (POST) ReportThisweek(r *http.Request) (response []byte, err error) {
	r.Response = &http.Response{}
	r.Response.Header = http.Header{}
	r.Response.Header.Add("cool", "damn")
	fmt.Printf("%#v\n", r)
	fmt.Println(len(r.Response.Header))
	msg := "POST.ReprotThisweek."
	return json.Marshal(msg)
}

func (OPTIONS) REPORTThisWEEK(r *http.Request) (response []byte, err error) {
	fmt.Printf("%#v", r)
	msg := "POST.ReprotThisweek."
	panic(msg)
	return json.Marshal(msg)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/byronzr/micro/helper"
)

type POST struct{}
type OPTIONS struct{}

func (POST) ReportThisweek(r *http.Request) (response []byte, err error) {

	// like http.Header.Add
	helper.HeaderAdd(r, "foo", "bar")
	fmt.Printf("%#v\n", r)

	// like http.Header.Set
	helper.HeaderSet(r, "foo", "bar1")
	fmt.Printf("%#v\n", r)

	msg := "POST.ReprotThisweek."
	return json.Marshal(msg)
}

func (OPTIONS) REPORTThisWEEK(r *http.Request) (response []byte, err error) {
	fmt.Printf("%#v", r)
	msg := "POST.ReprotThisweek."
	panic(msg)
	return json.Marshal(msg)
}

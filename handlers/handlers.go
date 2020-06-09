package handlers

import (
	"encoding/json"
	"net/http"
)

func (GET) ListAll(r *http.Request) (response []byte, err error) {
	msg := "GET.ListAll"
	return json.Marshal(&Response{Result: msg})
}

func (GET) ReportThisweek(r *http.Request) (response []byte, err error) {
	msg := "GET.ReprotThisweek."
	return json.Marshal(&Response{Result: msg})
}

func (POST) ReportThisweek(r *http.Request) (response []byte, err error) {
	msg := "POST.ReprotThisweek."
	return json.Marshal(&Response{Result: msg})
}

func (OPTIONS) ReportThisweek(r *http.Request) (response []byte, err error) {
	msg := "OPTIONS.ReprotThisweek."
	return json.Marshal(&Response{Result: msg})
}

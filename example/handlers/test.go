package handlers

import (
	"gopkg.in/byronzr/micro.v2/helper"
)

type POST struct{}
type OPTIONS struct{}

func (POST) ReportThisweek(m *helper.MicroRequest) int {
	msg := "POST.ReprotThisweek."
	l, err := m.W.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	return l
}

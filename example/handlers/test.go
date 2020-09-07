package handlers

import (
	"github.com/byronzr/micro"
)

type POST struct{}
type OPTIONS struct{}

func (POST) ReportThisweek(m *micro.MicroRequest) int {
	msg := "POST.ReprotThisweek."
	l, err := m.W.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	return l
}

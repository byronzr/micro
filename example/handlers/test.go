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

func (POST) ResponseJsonFail(m *micro.MicroRequest) int {
	micro.Err("check this longfile flag.")
	return m.Fail("fail", -1)
}

func (POST) ResponseJsonSuccess(m *micro.MicroRequest) int {
	result := "this is result."
	return m.Success("success", result)
}

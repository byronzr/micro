package micro

import (
	"encoding/json"
	"net/http"
)

type MicroRequest struct {
	R *http.Request
	W http.ResponseWriter
	M *Middle
}

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func (m *MicroRequest) Success(msg string, result interface{}) int {
	ret := Response{}
	ret.Result = result
	ret.Msg = msg
	v, err := json.Marshal(ret)
	if err != nil {
		Err(err)
		return 0
	}
	if r, err := m.W.Write(v); err != nil {
		Err(err)
		return 0
	} else {
		return r
	}
}

func (m *MicroRequest) Fail(msg string, code int) int {
	ret := Response{}
	ret.Code = code
	ret.Msg = msg
	v, err := json.Marshal(ret)

	if err != nil {
		Err(err)
		return 0
	}

	if r, err := m.W.Write(v); err != nil {
		Err(err)
		return 0
	} else {
		return r
	}
}

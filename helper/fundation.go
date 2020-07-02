package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DD(o ...interface{}) {
	buf := make([]byte, 0)
	for idx, s := range o {
		out, err := json.MarshalIndent(s, "", "\t")
		if err != nil {
			panic(err)
		}
		buf = append(buf, []byte(fmt.Sprintf("\n\033[33m----%02d----------------------------------------------\033[0m\n%s", idx, string(out)))...)
	}
	fmt.Println(string(buf))
}

// header init
func HeaderInit(r *http.Request) {
	if r.Response == nil {
		r.Response = &http.Response{}
	}
	if len(r.Response.Header) == 0 {
		r.Response.Header = http.Header{}
	}
}

// for add
func HeaderAdd(r *http.Request, key, value string) {
	HeaderInit(r)
	r.Response.Header.Add(key, value)
}

// for set
func HeaderSet(r *http.Request, key, value string) {
	HeaderInit(r)
	r.Response.Header.Set(key, value)
}

package helper

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ROUTER struct{}

func (ROUTER) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	uri := strings.ToLower(r.URL.Path)
	method := strings.ToUpper(r.Method)
	furi := r.URL.RequestURI()

	target := fmt.Sprintf("%s %s", method, uri)
	if fn, ok := ActionFuncMap[target]; ok {
		if response, err := fn(r); err == nil {
			// write header
			resp := r.Response
			if resp != nil && len(resp.Header) > 0 {
				h := w.Header()
				for k, vs := range resp.Header {
					for _, v := range vs {
						h.Add(k, v)
					}
				}
			}

			// write response
			ret, err := w.Write(response)
			if err != nil {
				panic(err)
			}
			s := time.Since(t)
			Inf(method, " ", furi, " write:", ret, " t:", s)
			return
		} else {
			panic(err)
		}
	}

	Err(furi, " ", method, " ", " NOT FOUND ")
	http.NotFound(w, r)

}

func (ROUTER) CrossHeader(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Accept,Origin,XRequestedWith,Content-Type,LastModified,DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
}

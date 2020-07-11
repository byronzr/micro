package helper

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	data = sync.Map{}
)

type ROUTER struct{}

// middle function call on before
type BeforeCall interface {
	Before(*http.Request) (interface{}, bool)
}

// middle function call
type AfterCall interface {
	After(*http.Request) (interface{}, bool)
}

func (ROUTER) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	uri := strings.ToLower(r.URL.Path)
	method := strings.ToUpper(r.Method)
	bf := fmt.Sprintf("before.%s", method)
	af := fmt.Sprintf("after.%s", method)
	furi := r.URL.RequestURI()

	target := fmt.Sprintf("%s %s", method, uri)
	if fn, ok := ActionFuncMap[target]; ok {
		// MIDDLE BEFORE
		if f, ok := MiddleFuncMap["GLOBAL_BEFORE"]; ok {
			if rs, ok := f(r); !ok {
				Wrn("halt on Middle Global before")
				return
			} else {
				data.Store(r.RemoteAddr, rs)
			}
		}
		if f, ok := MiddleFuncMap[bf]; ok {
			if rs, ok := f(r); !ok {
				Wrn("halt on Middle before ", method)
				return
			} else {
				data.Store(r.RemoteAddr, rs)
			}
		}
		// clear sync.Map
		defer func() {
			if _, ok := data.Load(r.RemoteAddr); ok {
				data.Delete(r.RemoteAddr)
			}
		}()

		// run serve http
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
		} else {
			panic(err)
		}

		// MIDDLE AFTER
		if f, ok := MiddleFuncMap["GLOBAL_AFTER"]; ok {
			f(r)
		}
		if f, ok := MiddleFuncMap[af]; ok {
			f(r)
		}
		return
	}

	Err(furi, " ", method, " ", " NOT FOUND ")
	http.NotFound(w, r)

}

func (ROUTER) CrossHeader(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Accept,Origin,XRequestedWith,Content-Type,LastModified,DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
}

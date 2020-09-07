package micro

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ROUTER struct{}

// middle function call on before
type BeforeCall interface {
	Before(*MicroRequest) bool
}

// middle function call
type AfterCall interface {
	After(*MicroRequest) bool
}

func (ROUTER) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	// middle store
	mr := &MicroRequest{r, w, &Middle{}}

	uri := strings.TrimRight(strings.ToLower(r.URL.Path), "/")
	method := strings.ToUpper(r.Method)

	// middle name
	bf := fmt.Sprintf("%s before", method)
	af := fmt.Sprintf("%s after", method)

	furi := r.URL.RequestURI()

	target := fmt.Sprintf("%s %s", method, uri)
	if fn, ok := ActionFuncMap[target]; ok {
		// GLOBAL MIDDLE BEFORE
		if f, ok := MiddleFuncMap["GLOBAL before"]; ok {
			if ok := f(mr); !ok {
				// Wrn("halt on global Middle before")
				return
			}
		}
		// PARTIAL BEFORE
		if f, ok := MiddleFuncMap[bf]; ok {
			if ok := f(mr); !ok {
				// Wrn("halt on partial Middle before ", method)
				return
			}
		}

		// run serve http
		lenOfContents := fn(mr)

		s := time.Since(t)
		Inf(method, " ", furi, " write:", PackSize(lenOfContents, "b"), " t:", s)

		// MIDDLE AFTER
		if f, ok := MiddleFuncMap["GLOBAL after"]; ok {
			f(mr)
		}
		if f, ok := MiddleFuncMap[af]; ok {
			f(mr)
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

// auto covert content length B/KB/MB/GB/TB
func PackSize(v int, uint string) string {
	us := map[string]string{"b": "kb", "kb": "mb", "mb": "gb", "gb": "tb"}

	if v < 1024 {
		return fmt.Sprintf("%d%s", v, uint)
	}
	nv := v / 1024
	if nv >= 1024 {
		return PackSize(nv, us[uint])
	} else {
		return fmt.Sprintf("%d%s", nv, us[uint])
	}
}

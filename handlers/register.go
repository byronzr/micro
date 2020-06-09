package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	. "github.com/byronzr/micro/fundational"
)

type Response struct {
	Result interface{} `json:"result"`
}

var (
	ActionFuncMap = make(map[string]func(*http.Request) ([]byte, error), 0)
)

func RegisterHandler(h interface{}) {

	re := regexp.MustCompile(`<func\(handlers.([A-Z]+?), \*http.Request\) \(\[]uint8, error\) Value>`)

	t := reflect.TypeOf(h)
	methodCount := t.NumMethod()
	for i := 0; i < methodCount; i++ {
		uriKey := ""
		method := t.Method(i)

		tys := re.FindStringSubmatch(method.Func.String())
		if len(tys) < 2 {
			Inf(">> continue >> ", method.Name)
			continue
		}
		action := tys[1]

		// 转换驼峰函数名为URI路径名
		rawName := []byte(method.Name)
		uriName := []byte{}
		for _, b := range rawName {
			if b >= 'A' && b <= 'Z' {
				uriName = append(uriName, '/')
			}
			uriName = append(uriName, b)
		}
		uriKey = fmt.Sprintf("%s %s", action, strings.ToLower(string(uriName)))

		// TODO: 未来泛型优化
		switch action {
		case "GET":
			fn := method.Func.Interface().(func(GET, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(GET{}, r) }
		case "POST":
			fn := method.Func.Interface().(func(POST, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(POST{}, r) }
		case "OPTIONS":
			fn := method.Func.Interface().(func(OPTIONS, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(OPTIONS{}, r) }
		case "HEAD":
			fn := method.Func.Interface().(func(HEAD, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(HEAD{}, r) }
		case "PUT":
			fn := method.Func.Interface().(func(PUT, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(PUT{}, r) }
		case "PATCH":
			fn := method.Func.Interface().(func(PATCH, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(PATCH{}, r) }
		case "DELETE":
			fn := method.Func.Interface().(func(DELETE, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(DELETE{}, r) }
		case "TRACE":
			fn := method.Func.Interface().(func(TRACE, *http.Request) ([]uint8, error))
			ActionFuncMap[uriKey] = func(r *http.Request) ([]byte, error) { return fn(TRACE{}, r) }
		}
		Inf(">> registered >> ", uriKey)
	}
}

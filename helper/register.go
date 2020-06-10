package helper

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

// request method set

var (
	ActionFuncMap = make(map[string]func(*http.Request) ([]byte, error), 0)
)

func RegisterHandler(h interface{}) {

	re := regexp.MustCompile(`<func\((\S+?)\.([A-Z]+?), \*http.Request\) \(\[]uint8, error\) Value>`)

	t := reflect.TypeOf(h)
	v := reflect.ValueOf(h)
	methodCount := t.NumMethod()
	for i := 0; i < methodCount; i++ {
		uriKey := ""
		method := t.Method(i)

		tys := re.FindStringSubmatch(method.Func.String())
		if len(tys) < 2 {
			Inf(">> continue >> ", method.Func.String(), " >> ", method.Name)
			continue
		}
		action := tys[2]

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
		ActionFuncMap[uriKey] = func(r *http.Request) (result []byte, err error) {
			rs := method.Func.Call([]reflect.Value{v, reflect.ValueOf(r)})
			if rs[0].Interface() != nil {
				result = rs[0].Interface().([]byte)
			}
			if rs[1].Interface() != nil {
				err = rs[1].Interface().(error)
			}
			return
		}
		Inf(">> registered >> ", uriKey)
	}
}

// func RegisterHandler(h Action) {
// 	h.Register()
// }
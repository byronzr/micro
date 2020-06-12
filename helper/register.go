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
	patchPrefix   = ""
)

func RegisterHandler(h interface{}) {

	t := reflect.TypeOf(h)
	v := reflect.ValueOf(h)

	// 分组边界前缀
	if t.Kind() == reflect.String {
		str := v.Interface().(string)
		if !strings.HasPrefix(str, "/") {
			patchPrefix = fmt.Sprintf("/%s", str)
		} else {
			patchPrefix = str
		}
		return
	}
	re := regexp.MustCompile(`<func\((\S+?)\.([A-Z]+?), \*http.Request\) \(\[]uint8, error\) Value>`)

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
		// 转换驼峰函数名为URI路径名
		rawName := []byte(method.Name)
		lenName := len(rawName)
		uriName := []byte{}
		for idx, b := range rawName {
			prefix := byte('a')
			suffix := byte('A')
			if idx > 0 {
				prefix = rawName[idx-1]
			}
			if idx != lenName-1 {
				suffix = rawName[idx+1]
			}
			if (up(b) && !up(prefix)) || (up(b) && !up(suffix)) {
				uriName = append(uriName, '/')
			}
			uriName = append(uriName, b)
		}
		uriKey = fmt.Sprintf("%s %s%s", action, patchPrefix, strings.ToLower(string(uriName)))

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

	}
}

func up(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

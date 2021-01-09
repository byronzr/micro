package micro

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// request method set
var (
	// V2
	ActionFuncMap = make(map[string]func(*MicroRequest) int, 0)
	MiddleFuncMap = make(map[string]func(*MicroRequest) bool, 0)
	patchPrefix   = ""
)

func RegisterHandler(h interface{}) {

	// init middle register
	if mf, ok := h.(BeforeCall); ok {
		structname := fmt.Sprintf("%#v", mf)
		names := strings.Split(strings.Trim(structname, "{}"), ".")
		method := fmt.Sprintf("%s before", names[len(names)-1])
		MiddleFuncMap[method] = mf.Before
	}

	if mf, ok := h.(AfterCall); ok {
		structname := fmt.Sprintf("%#v", mf)
		names := strings.Split(strings.Trim(structname, "{}"), ".")
		method := fmt.Sprintf("%s after", names[len(names)-1])
		MiddleFuncMap[method] = mf.After
	}

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

	// V2
	re := regexp.MustCompile(`<func\((\S+?)\.([A-Z]+?), \*micro.MicroRequest\) int Value>`)

	methodCount := t.NumMethod()
	for i := 0; i < methodCount; i++ {
		uriKey := ""
		method := t.Method(i)

		// PPP(method.Name, "==", method.Func.String())

		tys := re.FindStringSubmatch(method.Func.String())
		if len(tys) < 2 {
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
		ActionFuncMap[uriKey] = func(m *MicroRequest) int {
			rs := method.Func.Call([]reflect.Value{v, reflect.ValueOf(m)})
			if rs[0].Interface() != nil {
				return rs[0].Interface().(int)
			}
			return 0
		}

	}
}

func up(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

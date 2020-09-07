package micro

import "fmt"

type Middle struct {
	prefix string                 // global or partial(get/post/...)
	suffix string                 // before or after
	values map[string]interface{} // global before
}

func (m *Middle) init() {
	if m.values == nil {
		m.values = make(map[string]interface{}, 0)
	}
}

func (m *Middle) Global() *Middle {
	m.prefix = "global"
	return m
}

func (m *Middle) Before() *Middle {
	m.suffix = "before"
	return m
}

func (m *Middle) After() *Middle {
	m.suffix = "after"
	return m
}

func (m *Middle) Set(v interface{}) (update bool, err error) {
	m.init()
	if m.suffix == "" {
		err = fmt.Errorf("suffix(.Before / .After) not choice on set.")
		return
	}
	target := fmt.Sprintf("%s%s", m.prefix, m.suffix)
	_, update = m.values[target]
	m.values[target] = v
	// reset
	m.suffix = ""
	m.prefix = ""
	return
}

func (m *Middle) Value() (v interface{}, has bool) {
	m.init()
	if m.suffix == "" {
		return
	}
	target := fmt.Sprintf("%s%s", m.prefix, m.suffix)
	v, has = m.values[target]
	// reset
	m.suffix = ""
	m.prefix = ""
	return
}

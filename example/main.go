package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/handlers"

	"github.com/byronzr/micro"
	"github.com/byronzr/micro/helper"
)

// request method set
type GET struct{}

func main() {
	service := micro.Register(GET{}, "byron", handlers.POST{}) // must be first
	bf := func(m *helper.MicroRequest) bool {
		fmt.Println("i'm running before <GLB>")
		return true
	}
	af := func(m *helper.MicroRequest) bool {
		fmt.Println("i'm running after <GLB>")
		return true
	}

	service.Before(bf)
	service.After(af)

	service.Start(8000, 10)
}

// (无效的) handler V1
func (GET) Check(r *http.Request) (response []byte, err error) {
	msg := "GET.Check"
	return json.Marshal(msg)
}

// 局部中间件
func (GET) Before(m *helper.MicroRequest) bool {
	// 设置与传递值
	mid := m.M.Before()
	mid.Set("value from partial get.")
	// false, 中断执行
	return true
}

func (GET) After(m *helper.MicroRequest) bool {
	// 获取中间件的传递值
	mid := m.M.Before()
	v, ok := mid.Value()
	if ok {
		fmt.Printf("after: %s\n", v)
	}
	return true
}

// 常规 handler
func (GET) FullCheck(m *helper.MicroRequest) int {
	content := "fulll check."
	l, err := m.W.Write([]byte(content))
	if err != nil {
		panic(err)
	}
	// 返回写入长度
	return l
}

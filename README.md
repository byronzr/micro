# micro
simple base RESTFUL framework

# Quick Start

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/handlers"

	"github.com/byronzr/micro.v2"
	"github.com/byronzr/micro.v2/helper"/'/'
)

// 定义一个请求方式
type GET struct{}

func main() {
  // 注册 handler
  // 注册过程中可进行分组 
	service := micro.Register(GET{}, "byron", handlers.POST{}) // must be first
  
  // 注册全局中间件 before / after
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

  // 启动服务（端口，超时）
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

```

### runtime

```shell
➜  example git:(V2) ✗ go run main.go
┌────────┬─────────────────────────────┐
│ Type   │ Register                    │
├────────┼─────────────────────────────┤
│ MIDDLE │ GET before                  │
│ MIDDLE │ GET after                   │
│ MIDDLE │ GLB before                  │
│ MIDDLE │ GLB after                   │
│ ACTION │ GET /full/check             │
│ ACTION │ POST /byron/report/thisweek │
└────────┴─────────────────────────────┘
INF 2020/09/03 12:04:18 :::::: service [ 8000 ] start ::::::
i'm running before <GLB>
INF 2020/09/03 12:04:23 GET /full/check write:12b t:45.292µs
i'm running after <GLB>
after: value from partial get.
```


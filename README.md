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

	"github.com/byronzr/micro"
)

// request method set
type GET struct{}

func main() {
	service := micro.Register(GET{}, "byron", handlers.POST{}) // must be first
	bf := func(m *micro.MicroRequest) bool {
		fmt.Println("i'm running before <GLB>")
		return true
	}
	af := func(m *micro.MicroRequest) bool {
		fmt.Println("i'm running after <GLB>")
		return true
	}
    
    // 添加一个默认路由
	// 返回 -1 则是 404
	service.Default(func(m *micro.MicroRequest) int {
		if len(m.R.URL.Query().Get("test")) == 0 {
			return -1
		}
		return GET{}.FullCheck(m)
	})

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
func (GET) Before(m *micro.MicroRequest) bool {
	// 设置与传递值
	mid := m.M.Before()
	mid.Set("value from partial get.")
	// false, 中断执行
	return true
}

func (GET) After(m *micro.MicroRequest) bool {
	// 获取中间件的传递值
	mid := m.M.Before()
	v, ok := mid.Value()
	if ok {
		fmt.Printf("after: %s\n", v)
	}
	return true
}

// 常规 handler
func (GET) FullCheck(m *micro.MicroRequest) int {
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
➜  example git:(master) ✗ go run main.go
┌────────┬─────────────────────────────┐
│ Type   │ Register                    │
├────────┼─────────────────────────────┤
│ MIDDLE │ GLOBAL after                │
│ MIDDLE │ GET before                  │
│ MIDDLE │ GET after                   │
│ MIDDLE │ GLOBAL before               │
│ ACTION │ GET /full/check             │
│ ACTION │ POST /byron/report/thisweek │
└────────┴─────────────────────────────┘
INF 2020/09/07 10:24:03 :::::: service [ 8000 ] start ::::::
```

# changelog
> 2020-01-09 add Fail / Success / Standard Response
```go

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func (POST) ResponseJsonFail(m *micro.MicroRequest) int {
	micro.Err("check this longfile flag.")
	return m.Fail("fail", -1)
}

//{
//  "code": -1,
//  "msg": "fail",
//  "result": null
//}

func (POST) ResponseJsonSuccess(m *micro.MicroRequest) int {
	result := "this is result."
	return m.Success("success", result)
}

//{
//  "code": 0,
//  "msg": "success",
//  "result": "this is result."
//}
```

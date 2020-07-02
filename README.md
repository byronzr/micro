# micro
simple base RESTFUL framework
简单粗暴的快速构建微服务脚手架 micro。别跟我说什么性能，并发，优化，老夫业务流程，数据分析从来不考虑这些重构的事。再说，性能也不差。

# 粗暴的启动

```go
package main

import (
	"test/handlers"

	"github.com/byronzr/micro"
)
func main() {
	// example 1
	// no prefix start
	// micro.Register(handlers.POST{}, handlers.OPTIONS{}).Start(8000, 10)
    
	// example 2
	// not chan call
	service := micro.Register("byron", handlers.POST{}, "wl", handlers.OPTIONS{}) // must be first
	service.Start(8000, 10)
	// INF 2020/06/12 14:38:04 service start.
	// INF 2020/06/12 14:38:04 >> registered >> POST /byron/report/thisweek
	// INF 2020/06/12 14:38:04 >> registered >> OPTIONS /wl/report/this/week
}
```
# 粗暴的分组混合注册前缀
```go
func main() {
    // Register 方法提供任何时候进行分组
    // 以下用例中 "byron" 注册后 handlers.POST中的所有路由自动加上 /byron 前缀
    // 当遇到 wl 后，后续前缀为 /wl 前缀
	service := micro.Register("byron", handlers.POST{}, "wl", handlers.OPTIONS{}) // must be first
	service.Start(8000, 10)
}
```

# 粗暴的路由设置

```bash
.
├── go.mod
├── go.sum
├── handlers
│   └── test.go // 你可以将路由处理程序集中规划在一个独立的目录中
└── main.go     // 也可以将路由直接写在 package main 里，更可以混合归纳
```

# 粗暴的一致性

以全大写定义一个 RESTFUL 的方法结构，也可以自定义自已的方法。

```go
package handlers
import(

	"encoding/json"
	"net/http"

)

type POST struct{}
type OPTIONS struct{}

func (POST) ReportThisweek(r *http.Request) (response []byte, err error) {
	msg := "POST.ReprotThisweek."
	return json.Marshal(msg)
}

func (OPTIONS) ReportThisweek(r *http.Request) (response []byte, err error) {
	msg := "POST.ReprotThisweek."
	return json.Marshal(msg)
}
```

# 粗暴的路由转换
将驼峰法则自动转换成路由 URL。
> REPORTThisweek => /report/thisweek
> ReportThisweek => /report/thisweek
```bash
INF 2020/06/10 21:01:46 >> registered >> POST /report/thisweek     # 自动将 ReportThisweek 首字母大写位置添加左竖线
INF 2020/06/10 21:01:46 >> registered >> OPTIONS /report/thisweek
INF 2020/06/10 21:01:46 service start.
```

# response header 返回头部的管理
业务返回时，需要对header进行精确控制，需要调用 helper 的相关方法
```go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/byronzr/micro/helper"
)

type POST struct{}
type OPTIONS struct{}

func (POST) ReportThisweek(r *http.Request) (response []byte, err error) {

	// like http.Header.Add
	helper.HeaderAdd(r, "foo", "bar")
	fmt.Printf("%#v\n", r)

	// like http.Header.Set
	helper.HeaderSet(r, "foo", "bar1")
	fmt.Printf("%#v\n", r)

	msg := "POST.ReprotThisweek."
	return json.Marshal(msg)
}

func (OPTIONS) REPORTThisWEEK(r *http.Request) (response []byte, err error) {
	fmt.Printf("%#v", r)
	msg := "POST.ReprotThisweek."
	panic(msg)
	return json.Marshal(msg)
}

// 中间件接口
func (OPTIONS) Before(r *http.Request) (interface{}, bool) {
	fmt.Println("i'm halt on before <options>")
    // 返回为 false 可以用在鉴权失败，中断业务流的调用
    // 返回 interface{} 可在业务调用中，再次计算中间件，以获得可用值（解析后的用户信息）
	return nil, false
}
// bash out:
// i'm running before <global>
// i'm halt on before <options>
// WRN 2020/07/02 15:25:27 halt on Middle before OPTIONS
```

# 中间件的注册
`全局中间件`永远优先于`方法中间件`调用，全局中间件需要显示调用注册，方法中间件只需要实现接口即可。

```go
// middle function call on before
type BeforeCall interface {
	Before(*http.Request) (interface{}, bool)
}

// middle function call
type AfterCall interface {
	After(*http.Request) (interface{}, bool)
}
```

全局与方法绑定中间件的用例
```go
// 全局中间件的注册范例
func main() {
	service := micro.Register(GET{}, "byron", handlers.POST{}, "wl", handlers.OPTIONS{}) // must be first

	// add middle hook on before
	before := func(r *http.Request) (interface{}, bool) {
		fmt.Println("i'm running before <global>")
		return nil, true
	}
	service.Before(before)

	// add middle hook on before
	after := func(r *http.Request) (interface{}, bool) {
		fmt.Println("i'm running after <global>")
		return nil, true
	}
	service.After(after)
    
    // start
	service.Start(8000, 10)
}

func (GET) Check(r *http.Request) (response []byte, err error) {
	msg := "GET.Check"
	return json.Marshal(msg)
}

// 局部中间件的注册范例
func (GET) Before(r *http.Request) (interface{}, bool) {
	fmt.Println("i'm running before <get>")
	return nil, true
}
```
调用顺序可见
```bash
WRN 2020/07/02 15:18:52 >> MIDDLE BEFORE >> before.GET
WRN 2020/07/02 15:18:52 >> MIDDLE BEFORE >> before.OPTIONS
WRN 2020/07/02 15:18:52 >> MIDDLE BEFORE >> GLOBAL
WRN 2020/07/02 15:18:52 >> MIDDLE AFTER >> GLOBAL
INF 2020/07/02 15:18:52 >> registered >> GET /check
INF 2020/07/02 15:18:52 >> registered >> POST /byron/report/thisweek
INF 2020/07/02 15:18:52 >> registered >> OPTIONS /wl/report/this/week
INF 2020/07/02 15:18:52 :::::: service start ::::::
i'm running before <global>
i'm running before <get>
INF 2020/07/02 15:18:59 GET /check write:11 t:107.625µs
i'm running after <global>
```

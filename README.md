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

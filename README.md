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

    // Register 注册需要处理的路由
    // Start 启动服务，监听端口与服务链接超时设置
    // micro.Register(handlers.POST{}, handlers.OPTIONS{}).Start(8000, 10)
  
    // example 1
    // no prefix start
    // micro.Register(handlers.POST{}, handlers.OPTIONS{}).Start(8000, 10)

	// example 2
	// has prefix start
	// micro.Register(handlers.POST{}, handlers.OPTIONS{}).Prefix("byron").Start(8000, 10)

	// example 3
	// not chan call
	service := micro.Register(handlers.POST{}, handlers.OPTIONS{}) // must be first
	service.Prefix("byron")                                        // optional method
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

  注： AbcdEfg => /abcd/efg 所以要注意 ABCDEFG => /a/b/c/e/d/f/g
  
  
```bash
INF 2020/06/10 21:01:46 >> registered >> POST /report/thisweek     # 自动将 ReportThisweek 首字母大写位置添加左竖线
INF 2020/06/10 21:01:46 >> registered >> OPTIONS /report/thisweek
INF 2020/06/10 21:01:46 service start.
```

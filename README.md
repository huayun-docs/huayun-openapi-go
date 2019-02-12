# huayun-openapi-go
- 提供调用华云公有云开放接口的Golang SDK代码
- 调用如下：<br />
执行go get -u github.com/huayun-docs/huayun-openapi-go<br />
代码：<br />
```
import (
    "fmt"

    "github.com/huayun-docs/huayun-openapi-go/chinacgateway/core"
    "github.com/huayun-docs/huayun-openapi-go/chinacgateway/ecs"
)

g := chinacgateway.NewGateWayApi("region", "ak", "sk") //分别传入机房标识、ak和sk

descVol := chinacgatewayecs.DescribeInstances()
r := descVol.GetRequest()

g, err := g.SetRequest(r)
if err != nil {
    fmt.Println(err)
    return
}
s, err := g.Request()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(g.GetStatusCode())
fmt.Println(s)

```

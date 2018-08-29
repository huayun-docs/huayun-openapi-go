# huayun-openapi-go
- 提供调用华云公有云开放接口的Golang SDK代码
- 调用如下：<br />
执行go get -u github.com/huayun-docs/huayun-openapi-go<br />
代码：<br />
```
import (
    "fmt"

    "github.com/huayun-docs/huayun-openapi-go/chinacgateway"
)

g := chinacgateway.NewGateWayApi("region", "ak", "sk") //分别传入机房标识、ak和sk
r := map[string]string{
    "Action": "DescribeRouters",
    "Id.0": "r-pp16mi6rgfo12c",
}
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

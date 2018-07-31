# huayun-openapi-go
- 提供调用华云公有云开放接口的Golang SDK代码
- 调用如下：<br />
执行go get github.com/huayun-docs/huayun-openapi-go<br />
代码：<br />
import (<br />
  "fmt"<br />
  "github.com/huayun-docs/huayun-openapi-go/chinacgateway"<br />
)<br /><br />
g := chinacgateway.NewGateWayApi("region", "ak", "sk")&nbsp;//分别传入机房标识、ak和sk<br >
r := map[string]string{<br >
&nbsp;&nbsp;&nbsp;&nbsp;"Action": "DescribeRouters",<br >
&nbsp;&nbsp;&nbsp;&nbsp;"Id.0":   "r-pp16mi6rgfo12c",<br >
}<br >
g, err := g.SetRequest(r)<br >
if err != nil {<br >
&nbsp;&nbsp;&nbsp;&nbsp;fmt.Println(err)<br >
&nbsp;&nbsp;&nbsp;&nbsp;return<br >
}<br >
s, err := g.Request()<br >
if err != nil {<br >
&nbsp;&nbsp;&nbsp;&nbsp;fmt.Println(err)<br >
&nbsp;&nbsp;&nbsp;&nbsp;return<br >
}<br >
fmt.Println(s)

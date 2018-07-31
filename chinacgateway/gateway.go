package chinacgateway

import (
	"fmt"
	"strings"
	"time"
)

type GateWayApi struct {
	Url             string
	AccessKeyId     string
	AccessKeySecret string
	Region          string
	Version         string
	Date            string

	Method string
	Header map[string]string
	Req    map[string]string
	Ourl   string
	Oreq   string
	Res    string
}

func NewGateWayApi(tag, ak, sk string) *GateWayApi {
	g := &GateWayApi{
		Url:             "https://api.chinac.com",
		AccessKeyId:     ak,
		AccessKeySecret: sk,
		Region:          tag,
		Version:         "1.0",
		Date:            time.Now().Format("2006-01-02T15:04:05 +0800"),
		Method:          "GET",
	}
	header := map[string]string{
		"Content-Type": `application/json;charset=UTF-8`,
		"Date":         g.Date}
	g.Header = header

	return g
}

func (g *GateWayApi) SetMethod(m string) *GateWayApi {
	g.Method = m
	return g
}

func (g *GateWayApi) SetUrl(nurl string) *GateWayApi {
	g.Url = nurl
	return g
}

func (g *GateWayApi) Clear() *GateWayApi {
	g = &GateWayApi{}
	return g
}

func (g *GateWayApi) SetRequest(r map[string]string) (*GateWayApi, error) {
	if _, ok := r["Action"]; !ok {
		return nil, fmt.Errorf("The params has no 'Action'")
	}
	g.Req = r
	return g, nil
}

func (g *GateWayApi) Request() (string, error) {
	if len(g.Req) <= 0 {
		return "", fmt.Errorf("No params")
	}
	request := g.Req
	request["Region"] = g.Region
	request["AccessKeyId"] = g.AccessKeyId
	request["Version"] = g.Version
	request["Date"] = g.Date

	var q []string
	for k, v := range request {
		q = append(q, percentEncode(k))
		q = append(q, "=")
		q = append(q, percentEncode(v))
		q = append(q, "&")
	}
	q = q[:(len(q) - 1)]
	quest := strings.Join(q, "")
	signature := getSignature(g, quest)
	quest += "&Signature=" + signature

	rurl := g.Url
	if g.Method == "GET" {
		rurl += "?" + quest
	}
	g.Ourl = rurl
	g.Oreq = quest

	return callServerRes(g)
}

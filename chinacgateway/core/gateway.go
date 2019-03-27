package chinacgateway

import (
	"fmt"
	"strings"
	"time"
)

type GateWayApi struct {
	url             string
	accessKeyId     string
	accessKeySecret string
	region          string
	version         string
	date            string

	method     string
	header     map[string]string
	req        map[string]string
	ourl       string
	oreq       map[string]string
	Res        string
	statusCode int
}

func NewGateWayApi(tag, ak, sk string) *GateWayApi {
	g := &GateWayApi{
		url:             "https://api.chinac.com",
		accessKeyId:     ak,
		accessKeySecret: sk,
		region:          tag,
		version:         "1.0",
		date:            time.Now().Format("2006-01-02T15:04:05 +0800"),
		method:          "GET",
	}
	header := map[string]string{
		"Content-Type":    `application/json;charset=UTF-8`,
		"Accept-Encoding": "*",
		"Date":            g.date}
	g.header = header

	return g
}

func (g *GateWayApi) SetMethod(m string) *GateWayApi {
	g.method = m
	return g
}

func (g *GateWayApi) SetUrl(nurl string) *GateWayApi {
	g.url = nurl
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
	g.req = r
	return g, nil
}

func (g *GateWayApi) Request() (string, error) {
	if len(g.req) <= 0 {
		return "", fmt.Errorf("No params")
	}
	request := g.req
	request["Region"] = g.region
	request["AccessKeyId"] = g.accessKeyId
	request["Version"] = g.version
	request["Date"] = g.date
	g.oreq = request
	if g.method != "GET" {
		request = map[string]string{
			"Region":      g.region,
			"AccessKeyId": g.accessKeyId,
			"Date":        g.date,
			"Action":      g.req["Action"],
			"Version":     g.version,
		}
	}

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

	rurl := g.url + "?" + quest
	g.ourl = rurl
	return callServerRes(g)
}

func (g *GateWayApi) GetStatusCode() int {
	return g.statusCode
}

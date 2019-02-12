package chinacgateway

import (
	"fmt"
	"strings"
	"time"
)

type gateWayApi struct {
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

func NewGateWayApi(tag, ak, sk string) *gateWayApi {
	g := &gateWayApi{
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

func (g *gateWayApi) SetMethod(m string) *gateWayApi {
	g.method = m
	return g
}

func (g *gateWayApi) SetUrl(nurl string) *gateWayApi {
	g.url = nurl
	return g
}

func (g *gateWayApi) Clear() *gateWayApi {
	g = &gateWayApi{}
	return g
}

func (g *gateWayApi) SetRequest(r map[string]string) (*gateWayApi, error) {
	if _, ok := r["Action"]; !ok {
		return nil, fmt.Errorf("The params has no 'Action'")
	}
	g.req = r
	return g, nil
}

func (g *gateWayApi) Request() (string, error) {
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

func (g *gateWayApi) GetStatusCode() int {
	return g.statusCode
}

package chinacgateway

import (
	//"fmt"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func percentEncode(str string) string {
	str = url.QueryEscape(str)
	str = strings.Replace(str, `+`, `%20`, -1)
	str = strings.Replace(str, `*`, `%2A`, -1)
	str = strings.Replace(str, `%7E`, `~`, -1)
	return str
}

func getSignature(g *GateWayApi, request string) string {
	header := g.header
	var t []string
	t = append(t, g.method)
	t = append(t, "\n")
	h := md5.New()
	h.Write([]byte(request))
	cipherStr := h.Sum(nil)
	request = hex.EncodeToString(cipherStr)
	t = append(t, request)
	t = append(t, "\n")
	t = append(t, header["Content-Type"])
	t = append(t, "\n")
	t = append(t, percentEncode(header["Date"]))
	t = append(t, "\n")
	str := strings.Join(t, "")

	return computeHmac256(str, g.accessKeySecret)
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return percentEncode(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}

func callServerRes(g *GateWayApi) (string, error) {
	//请求
	client := &http.Client{}
	var req *http.Request
	var err error
	if g.method == "GET" {
		req, err = http.NewRequest(g.method, g.ourl, nil)
	} else {
		body, _ := json.Marshal(g.oreq)
		req, err = http.NewRequest(g.method, g.ourl, strings.NewReader(string(body)))
	}

	if err != nil {
		return "", err
	}
	for k, v := range g.header {
		req.Header.Set(k, v)
	}

	//处理返回结果
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	g.statusCode = res.StatusCode

	defer res.Body.Close()
	by, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(by), nil
}

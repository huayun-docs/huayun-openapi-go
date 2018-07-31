package chinacgateway

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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
	header := g.Header
	var t []string
	t = append(t, g.Method)
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

	return computeHmac256(str, g.AccessKeySecret)
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
	if g.Method == "GET" {
		req, err = http.NewRequest(g.Method, g.Ourl, nil)
	} else {
		req, err = http.NewRequest(g.Method, g.Ourl, strings.NewReader(g.Oreq))
	}

	if err != nil {
		return "", err
	}
	for k, v := range g.Header {
		req.Header.Set(k, v)
	}

	//处理返回结果
	res, _ := client.Do(req)

	defer res.Body.Close()
	by, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(by), nil
}

package chinacgatewayvpc

import (
    "strconv"
)

type changeEipBandwidth struct {
    req        	map[string]string
}

func ChangeEipBandwidth() *changeEipBandwidth {
    c := &changeEipBandwidth {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ChangeEipBandwidth"
    return c
}

func (c *changeEipBandwidth) GetRequest() map[string]string {
    return c.req
}

func (c *changeEipBandwidth) GetAction() string {
    return c.req["Action"]
}

func (c *changeEipBandwidth) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *changeEipBandwidth) GetRegion() string {
    return c.req["Region"]
}

func (c *changeEipBandwidth) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *changeEipBandwidth) GetEip() string {
    return c.req["Eip"]
}

func (c *changeEipBandwidth) SetEip(varEip string) {
    c.req["Eip"] = varEip
}

func (c *changeEipBandwidth) GetBandwidth() int {
    nBandwidth, _ := strconv.Atoi(c.req["Bandwidth"])
    return nBandwidth
}

func (c *changeEipBandwidth) SetBandwidth(varBandwidth int) {
    c.req["Bandwidth"] = strconv.Itoa(varBandwidth)
}


